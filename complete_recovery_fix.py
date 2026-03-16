#!/usr/bin/env python3
"""
Comprehensive API package recovery fix script.
Fixes all syntax errors, type references, and import issues in one pass.
"""

import re
from pathlib import Path
from typing import Dict, List, Tuple

# Type mappings: unqualified type -> (package, qualified type)
TYPE_MAPPINGS = {
    'TagFilter': ('tagfilter', 'tagfilter.TagFilter'),
    'CustomPayloadField': ('types', 'types.CustomPayloadField'),
    'BoundaryScope': ('types', 'types.BoundaryScope'),
    'AccessRule': ('types', 'types.AccessRule'),
    'Aggregation': ('types', 'types.Aggregation'),
    'Granularity': ('types', 'types.Granularity'),
    'AlertSeverity': ('types', 'types.AlertSeverity'),
    'ThresholdRule': ('types', 'types.ThresholdRule'),
    'RuleWithThreshold': ('types', 'types.RuleWithThreshold'),
    'Threshold': ('types', 'types.Threshold'),
    'WebsiteTimeThreshold': ('types', 'types.WebsiteTimeThreshold'),
    'ExpressionOperator': ('types', 'types.ExpressionOperator'),
    'LogLevel': ('types', 'types.LogLevel'),
    'WebsiteImpactMeasurementMethod': ('types', 'types.WebsiteImpactMeasurementMethod'),
    'APIMember': ('types', 'types.APIMember'),
    'IncludedApplication': ('types', 'types.IncludedApplication'),
    'Scheduling': ('types', 'types.Scheduling'),
}

# Packages that need automationaction import
AUTOMATION_CROSS_REFS = {
    'AutomationAction': 'automationaction.AutomationAction',
    'Parameter': 'automationaction.Parameter',
}

def read_file(filepath: Path) -> str:
    """Read file content."""
    with open(filepath, 'r') as f:
        return f.read()

def write_file(filepath: Path, content: str):
    """Write file content."""
    with open(filepath, 'w') as f:
        f.write(content)

def fix_package_file(filepath: Path) -> bool:
    """Fix a single package file."""
    package_name = filepath.parent.name
    print(f"\nProcessing {package_name}/{filepath.name}...")
    
    content = read_file(filepath)
    original_content = content
    
    # Step 1: Fix malformed structures
    content = fix_syntax_errors(content, package_name)
    
    # Step 2: Determine needed imports
    needed_imports = determine_imports(content, package_name)
    
    # Step 3: Replace type references
    content = replace_type_references(content, package_name)
    
    # Step 4: Fix import block
    content = fix_import_block(content, needed_imports)
    
    # Step 5: Remove duplicate constants
    content = remove_duplicate_constants(content)
    
    if content != original_content:
        write_file(filepath, content)
        print(f"  ✅ Fixed")
        return True
    else:
        print(f"  ℹ️  No changes needed")
        return False

def fix_syntax_errors(content: str, package_name: str) -> str:
    """Fix syntax errors specific to certain packages."""
    
    # Fix hostagent - move imports before const
    if package_name == 'hostagent':
        # Extract package declaration
        lines = content.split('\n')
        package_line = lines[0]
        rest = '\n'.join(lines[1:])
        
        # Add proper import
        content = f"{package_line}\n\nimport (\n\t\"encoding/json\"\n)\n\n{rest}"
        # Remove any malformed import later in file
        content = re.sub(r'\n\nimport\s*\(\s*"encoding/json"\s*\)\s*\n(?=\ntype)', '\n', content)
    
    # Fix syntheticlocation - const before imports
    if package_name == 'syntheticlocation':
        content = re.sub(
            r'(package syntheticlocation\n\n)// SyntheticLocationResourcePath.*?\nconst SyntheticLocationResourcePath = "[^"]+"\n\n',
            r'\1',
            content
        )
    
    # Fix customeventspec - literal in import
    if package_name == 'customeventspec':
        content = re.sub(
            r'import \(\s*\n\s*"[^"]+"\s*\n\s*"/api/events/settings/event-specifications"[^\)]*\)',
            'import (\n\t"github.com/instana/instana-go-client/shared/tagfilter"\n\t"github.com/instana/instana-go-client/shared/types"\n)',
            content
        )
    
    return content

def determine_imports(content: str, package_name: str) -> Dict[str, str]:
    """Determine which imports are needed."""
    imports = {}
    
    # Check for type references
    for type_name, (pkg, _) in TYPE_MAPPINGS.items():
        if re.search(r'\b' + type_name + r'\b', content):
            if pkg == 'tagfilter':
                imports['tagfilter'] = 'github.com/instana/instana-go-client/shared/tagfilter'
            elif pkg == 'types':
                imports['types'] = 'github.com/instana/instana-go-client/shared/types'
    
    # Check for automation cross-references
    if package_name == 'automationpolicy':
        for type_name in AUTOMATION_CROSS_REFS:
            if re.search(r'\b' + type_name + r'\b', content):
                imports['automationaction'] = 'github.com/instana/instana-go-client/api/automationaction'
                imports['types'] = 'github.com/instana/instana-go-client/shared/types'
                break
    
    return imports

def replace_type_references(content: str, package_name: str) -> str:
    """Replace unqualified type references with qualified ones."""
    
    # Replace standard type mappings
    for type_name, (_, qualified_type) in TYPE_MAPPINGS.items():
        # Match type name as a standalone word (not part of another word)
        # In struct fields, function parameters, return types, etc.
        patterns = [
            # Struct field: Name Type `json:"..."`
            (r'\b' + type_name + r'\b(\s+`json:)', qualified_type + r'\1'),
            # Function parameter/return: (name Type)
            (r'\(([^)]*\s)' + type_name + r'\b', r'(\1' + qualified_type),
            # Array/slice: []Type
            (r'\[\]' + type_name + r'\b', '[]' + qualified_type),
            # Pointer: *Type
            (r'\*' + type_name + r'\b', '*' + qualified_type),
            # Map value: map[string]Type
            (r'map\[string\]' + type_name + r'\b', 'map[string]' + qualified_type),
        ]
        
        for pattern, replacement in patterns:
            content = re.sub(pattern, replacement, content)
    
    # Handle automation cross-references
    if package_name == 'automationpolicy':
        for type_name, qualified_type in AUTOMATION_CROSS_REFS.items():
            content = re.sub(r'\b' + type_name + r'\b(\s+`json:)', qualified_type + r'\1', content)
            content = re.sub(r'\[\]' + type_name + r'\b', '[]' + qualified_type, content)
    
    return content

def fix_import_block(content: str, needed_imports: Dict[str, str]) -> str:
    """Fix the import block."""
    lines = content.split('\n')
    
    # Find package declaration
    package_idx = -1
    for i, line in enumerate(lines):
        if line.strip().startswith('package '):
            package_idx = i
            break
    
    if package_idx == -1:
        return content
    
    # Remove existing import block
    new_lines = [lines[package_idx]]
    i = package_idx + 1
    
    # Skip empty lines and old import block
    while i < len(lines):
        line = lines[i].strip()
        if line.startswith('import'):
            # Skip entire import block
            if '(' in line:
                # Multi-line import
                while i < len(lines) and ')' not in lines[i]:
                    i += 1
                i += 1  # Skip the closing )
            else:
                # Single line import
                i += 1
        elif line == '':
            i += 1
        else:
            break
    
    # Add new import block if needed
    if needed_imports:
        new_lines.append('')
        new_lines.append('import (')
        for import_path in sorted(needed_imports.values()):
            new_lines.append(f'\t"{import_path}"')
        new_lines.append(')')
    
    # Add rest of file
    new_lines.extend(lines[i:])
    
    return '\n'.join(new_lines)

def remove_duplicate_constants(content: str) -> str:
    """Remove duplicate ResourcePath constants."""
    # Keep only the first ResourcePath constant
    parts = content.split('// ResourcePath')
    if len(parts) > 2:
        # Keep first occurrence, remove others
        result = parts[0] + '// ResourcePath' + parts[1]
        for part in parts[2:]:
            # Skip the const line and rejoin
            lines = part.split('\n')
            if lines and lines[0].strip().startswith('const ResourcePath'):
                result += '\n'.join(lines[1:])
            else:
                result += part
        content = result
    
    return content

def main():
    """Main function."""
    print("=" * 70)
    print("Comprehensive API Package Recovery Fix")
    print("=" * 70)
    
    api_dir = Path("api")
    if not api_dir.exists():
        print("ERROR: api/ directory not found")
        return
    
    fixed_count = 0
    total_count = 0
    
    # Process all API packages
    for package_dir in sorted(api_dir.iterdir()):
        if not package_dir.is_dir():
            continue
        
        package_file = package_dir / f"{package_dir.name}.go"
        if not package_file.exists():
            continue
        
        total_count += 1
        if fix_package_file(package_file):
            fixed_count += 1
    
    print("\n" + "=" * 70)
    print(f"Recovery Fix Complete!")
    print(f"  Total packages: {total_count}")
    print(f"  Fixed: {fixed_count}")
    print(f"  Unchanged: {total_count - fixed_count}")
    print("=" * 70)

if __name__ == "__main__":
    main()

# Made with Bob
