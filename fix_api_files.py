#!/usr/bin/env python3
import os
import re
from pathlib import Path

api_dir = Path("api")

for pkg_dir in api_dir.iterdir():
    if not pkg_dir.is_dir():
        continue
    
    pkg_name = pkg_dir.name
    pkg_file = pkg_dir / f"{pkg_name}.go"
    
    if not pkg_file.exists():
        continue
    
    print(f"Fixing {pkg_name}...")
    
    # Read the file
    with open(pkg_file, 'r') as f:
        content = f.read()
    
    # Remove duplicate package declarations
    lines = content.split('\n')
    new_lines = []
    seen_package = False
    skip_next_empty = False
    
    for line in lines:
        # Skip duplicate package declarations
        if line.startswith('package '):
            if not seen_package:
                new_lines.append(line)
                seen_package = True
                skip_next_empty = True
            continue
        
        # Skip comments between package declarations
        if skip_next_empty and (line.strip() == '' or line.strip().startswith('//')):
            continue
        
        skip_next_empty = False
        new_lines.append(line)
    
    # Write back
    with open(pkg_file, 'w') as f:
        f.write('\n'.join(new_lines))
    
    print(f"  ✓ Fixed {pkg_name}")

print("\n✅ All files fixed!")

# Made with Bob
