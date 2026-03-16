#!/usr/bin/env python3
"""
Script to fix imports and references in recovered API packages.
This adds proper imports for shared types and fixes path constants.
"""

import os
import re
from pathlib import Path

# Base path
CLIENT_API = Path("api")

# Type mappings: old reference -> new import path
TYPE_IMPORTS = {
    'TagFilter': 'github.com/instana/instana-go-client/shared/tagfilter',
    'CustomPayloadField': 'github.com/instana/instana-go-client/shared/types',
    'BoundaryScope': 'github.com/instana/instana-go-client/shared/types',
    'AccessRule': 'github.com/instana/instana-go-client/shared/types',
    'Aggregation': 'github.com/instana/instana-go-client/shared/types',
    'Granularity': 'github.com/instana/instana-go-client/shared/types',
    'AlertSeverity': 'github.com/instana/instana-go-client/shared/types',
    'ThresholdRule': 'github.com/instana/instana-go-client/shared/types',
    'RuleWithThreshold': 'github.com/instana/instana-go-client/shared/types',
    'Threshold': 'github.com/instana/instana-go-client/shared/types',
    'WebsiteTimeThreshold': 'github.com/instana/instana-go-client/shared/types',
    'ExpressionOperator': 'github.com/instana/instana-go-client/shared/types',
    'LogLevel': 'github.com/instana/instana-go-client/shared/types',
    'WebsiteImpactMeasurementMethod': 'github.com/instana/instana-go-client/shared/types',
    'APIMember': 'github.com/instana/instana-go-client/shared/types',
    'IncludedApplication': 'github.com/instana/instana-go-client/shared/types',
}

# Path constant replacements
PATH_CONSTANTS = {
    'EventSettingsBasePath': '"/api/events/settings"',
    'SettingsBasePath': '"/api/settings"',
    'AutomationBasePath': '"/api/automation"',
    'InstanaAPIBasePath': '"/api"',
    'RBACSettingsBasePath': '"/api/settings/rbac"',
    'EventSpecificationBasePath': '"/api/events/settings/event-specifications"',
    'WebsiteMonitoringResourcePath': '"/api/website-monitoring"',
    'SyntheticTestResourcePath': '"/api/synthetics/settings/tests"',
    'settingsPathElement': '"settings"',
}

# Special handling for cross-package references
CROSS_PACKAGE_REFS = {
    'AutomationAction': 'github.com/instana/instana-go-client/api/automationaction',
    'Parameter': 'github.com/instana/instana-go-client/api/automationaction',
    'Scheduling': 'github.com/instana/instana-go-client/shared/types',
}

def fix_package_file(package_name, filepath):
    """Fix imports and references in a single package file."""
    print(f"  Fixing {filepath.name}...")
    
    with open(filepath, 'r') as f:
        content = f.read()
    
    original_content = content
    lines = content.split('\n')
    
    # Find package declaration line
    package_line_idx = -1
    for i, line in enumerate(lines):
        if line.strip().startswith('package '):
            package_line_idx = i
            break
    
    if package_line_idx == -1:
        print(f"    ERROR: No package declaration found")
        return False
    
    # Collect needed imports
    needed_imports = set()
    
    # Check for type references
    for type_name, import_path in TYPE_IMPORTS.items():
        if re.search(r'\b' + type_name + r'\b', content):
            needed_imports.add(import_path)
    
    # Check for cross-package references
    for type_name, import_path in CROSS_PACKAGE_REFS.items():
        if re.search(r'\b' + type_name + r'\b', content):
            needed_imports.add(import_path)
    
    # Replace path constants
    for const_name, const_value in PATH_CONSTANTS.items():
        content = re.sub(
            r'\b' + const_name + r'\b',
            const_value,
            content
        )
    
    # Add import block if needed
    if needed_imports:
        import_lines = ['', 'import (']
        for imp in sorted(needed_imports):
            import_lines.append(f'\t"{imp}"')
        import_lines.append(')')
        import_lines.append('')
        
        # Insert after package declaration
        lines = content.split('\n')
        lines.insert(package_line_idx + 1, '\n'.join(import_lines))
        content = '\n'.join(lines)
    
    # Fix specific issues
    
    # Fix hostagent import issue (imports in wrong place)
    if package_name == 'hostagent':
        content = re.sub(
            r'(package hostagent\n)',
            r'\1\nimport (\n\t"encoding/json"\n)\n',
            content
        )
        # Remove any imports that appear later
        content = re.sub(r'\nimport\s*\(\s*"encoding/json"\s*\)\s*\n(?!package)', '\n', content)
    
    # Fix websitemonitoring and synthetictest - they reference old interfaces
    if package_name in ['websitemonitoring', 'synthetictest']:
        # These need special handling - remove the old resource implementations
        # Keep only the model structs
        content = re.sub(
            r'// NewWebsiteMonitoringConfigRestResource.*?(?=\n\n|\ntype|\nconst|\Z)',
            '',
            content,
            flags=re.DOTALL
        )
        content = re.sub(
            r'// NewSyntheticTestRestResource.*?(?=\n\n|\ntype|\nconst|\Z)',
            '',
            content,
            flags=re.DOTALL
        )
    
    # Write back if changed
    if content != original_content:
        with open(filepath, 'w') as f:
            f.write(content)
        print(f"    ✅ Fixed")
        return True
    else:
        print(f"    ℹ️  No changes needed")
        return True

def main():
    """Main function."""
    print("=" * 60)
    print("Fixing Recovered API Packages")
    print("=" * 60)
    print()
    
    success_count = 0
    fail_count = 0
    
    # Process each API package
    for package_dir in sorted(CLIENT_API.iterdir()):
        if not package_dir.is_dir():
            continue
        
        package_name = package_dir.name
        package_file = package_dir / f"{package_name}.go"
        
        if not package_file.exists():
            print(f"⚠️  {package_name}: No package file found")
            continue
        
        print(f"Processing {package_name}...")
        if fix_package_file(package_name, package_file):
            success_count += 1
        else:
            fail_count += 1
        print()
    
    print("=" * 60)
    print(f"Fix Complete!")
    print(f"  ✅ Success: {success_count} packages")
    print(f"  ❌ Failed: {fail_count} packages")
    print("=" * 60)

if __name__ == "__main__":
    main()

# Made with Bob
