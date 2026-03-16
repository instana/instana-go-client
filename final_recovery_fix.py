#!/usr/bin/env python3
"""
Final recovery fix script - handles all remaining compilation issues.
"""

import re
from pathlib import Path

def fix_file(filepath: Path) -> bool:
    """Fix a single file."""
    with open(filepath, 'r') as f:
        content = f.read()
    
    original = content
    package_name = filepath.parent.name
    
    # Fix 1: hostagent - const before imports
    if package_name == 'hostagent':
        # Completely rebuild the file structure
        lines = content.split('\n')
        # Find where actual types start
        type_start = -1
        for i, line in enumerate(lines):
            if line.strip().startswith('type HostAgent'):
                type_start = i
                break
        
        if type_start > 0:
            # Rebuild: package -> import -> const -> types
            new_content = 'package hostagent\n\nimport (\n\t"encoding/json"\n)\n\n'
            new_content += '// ResourcePath is the path to the Host Agents resource in the Instana API\n'
            new_content += 'const ResourcePath = "/api/host-agent"\n\n'
            new_content += '\n'.join(lines[type_start:])
            content = new_content
    
    # Fix 2: syntheticlocation - const before package end
    if package_name == 'syntheticlocation':
        lines = content.split('\n')
        # Find where types start
        type_start = -1
        for i, line in enumerate(lines):
            if line.strip().startswith('type SyntheticLocation'):
                type_start = i
                break
        
        if type_start > 0:
            new_content = 'package syntheticlocation\n\n'
            new_content += '// ResourcePath is the path to the Synthetic Locations resource in the Instana API\n'
            new_content += 'const ResourcePath = "/api/synthetics/settings/locations"\n\n'
            new_content += '\n'.join(lines[type_start:])
            content = new_content
    
    # Fix 3: customdashboard, sloalertconfig, slocorrection - imports after other declarations
    if package_name in ['customdashboard', 'sloalertconfig', 'slocorrection']:
        # Find and move import block to after package
        lines = content.split('\n')
        package_idx = 0
        import_start = -1
        import_end = -1
        
        for i, line in enumerate(lines):
            if line.strip().startswith('package '):
                package_idx = i
            elif line.strip().startswith('import ('):
                import_start = i
            elif import_start > 0 and line.strip() == ')':
                import_end = i
                break
        
        if import_start > package_idx + 1:
            # Extract import block
            import_block = lines[import_start:import_end+1]
            # Remove from current position
            del lines[import_start:import_end+1]
            # Insert after package
            for j, imp_line in enumerate(reversed(import_block)):
                lines.insert(package_idx + 1, imp_line)
            content = '\n'.join(lines)
    
    # Fix 4: customeventspec - literal in import
    if package_name == 'customeventspec':
        # Remove malformed import lines
        lines = content.split('\n')
        fixed_lines = []
        skip_next = False
        for line in lines:
            if '"/api/events/settings/event-specifications"' in line:
                skip_next = True
                continue
            if not skip_next:
                fixed_lines.append(line)
            else:
                skip_next = False
        content = '\n'.join(fixed_lines)
    
    # Fix 5: synthetictest - missing json import and old interfaces
    if package_name == 'synthetictest':
        # Add encoding/json import if using json
        if 'json.Unmarshal' in content and '"encoding/json"' not in content:
            content = re.sub(
                r'(package synthetictest\n)',
                r'\1\nimport (\n\t"encoding/json"\n)\n',
                content
            )
        # Remove old REST resource implementations
        content = re.sub(
            r'// NewSyntheticTestRestResource.*?(?=\n\ntype|\nconst|\Z)',
            '',
            content,
            flags=re.DOTALL
        )
    
    # Fix 6: websitemonitoring - old interfaces
    if package_name == 'websitemonitoring':
        content = re.sub(
            r'// NewWebsiteMonitoringConfigRestResource.*?(?=\n\ntype|\nconst|\Z)',
            '',
            content,
            flags=re.DOTALL
        )
    
    # Fix 7: Replace unqualified AlertSeverity with types.AlertSeverity
    content = re.sub(r'\bAlertSeverity\b(\s+`json:)', r'types.AlertSeverity\1', content)
    content = re.sub(r'\*AlertSeverity\b', r'*types.AlertSeverity', content)
    
    # Fix 8: LogLevel constants - need to be qualified
    if package_name == 'logalertconfig':
        content = re.sub(r'\bLogLevelError\b', 'types.LogLevelError', content)
        content = re.sub(r'\bLogLevelAny\b', 'types.LogLevelAny', content)
        content = re.sub(r'type LogLevel\b', 'type LogLevel = types.LogLevel', content)
    
    # Fix 9: WebsiteImpactMeasurementMethod constants
    if package_name == 'websitealertconfig':
        content = re.sub(r'\bWebsiteImpactMeasurementMethodPerWindow\b', 'types.WebsiteImpactMeasurementMethodPerWindow', content)
        content = re.sub(r'\bWebsiteImpactMeasurementMethodAggregated\b', 'types.WebsiteImpactMeasurementMethodAggregated', content)
        content = re.sub(r'type WebsiteImpactMeasurementMethod\b', 'type WebsiteImpactMeasurementMethod = types.WebsiteImpactMeasurementMethod', content)
    
    # Fix 10: sloconfig - missing fmt and os imports
    if package_name == 'sloconfig' and ('fmt.' in content or 'os.' in content):
        if '"fmt"' not in content or '"os"' not in content:
            # Add to imports
            content = re.sub(
                r'(import \(\n)',
                r'\1\t"fmt"\n\t"os"\n',
                content
            )
    
    # Fix 11: Remove unused imports
    if '"github.com/instana/instana-go-client/shared/types"' in content:
        if not re.search(r'\btypes\.[A-Z]', content):
            content = re.sub(r'\t"github://instana/instana-go-client/shared/types"\n', '', content)
    
    if content != original:
        with open(filepath, 'w') as f:
            f.write(content)
        return True
    return False

def main():
    """Main function."""
    print("=" * 70)
    print("Final Recovery Fix")
    print("=" * 70)
    print()
    
    api_dir = Path("api")
    fixed = 0
    
    for pkg_dir in sorted(api_dir.iterdir()):
        if not pkg_dir.is_dir():
            continue
        
        pkg_file = pkg_dir / f"{pkg_dir.name}.go"
        if not pkg_file.exists():
            continue
        
        print(f"Processing {pkg_dir.name}...")
        if fix_file(pkg_file):
            print(f"  ✅ Fixed")
            fixed += 1
        else:
            print(f"  ℹ️  No changes")
    
    print()
    print("=" * 70)
    print(f"Fixed {fixed} files")
    print("=" * 70)

if __name__ == "__main__":
    main()

# Made with Bob
