#!/usr/bin/env python3
"""
Script to recover API package files from Terraform provider source files.
This script reads the source files from internal/restapi/ and creates
properly formatted API package files in ../instana-go-client/api/
"""

import os
import re
from pathlib import Path

# Base paths
TERRAFORM_RESTAPI = Path("../terraform-provider-instana/internal/restapi")
CLIENT_API = Path("api")

# Mapping of API packages to their source files
PACKAGE_MAPPINGS = {
    "apitoken": ["api-tokens-api.go"],
    "alertingconfig": ["alerts-api.go", "alert-event-type.go"],
    "automationaction": ["automation-action-api.go"],
    "automationpolicy": ["automation-policy-api.go"],
    "builtineventspec": ["builtin-event-specification-api.go"],
    "customdashboard": ["custom-dashboard.go"],
    "customeventspec": ["custom-event-specficiations-api.go"],
    "group": ["groups-api.go"],
    "hostagent": ["host-agents-api.go", "host-agent-json-unmarshaller.go"],
    "infraalertconfig": ["infra-alert-configs.go", "infra-alert-rule.go", "infra-alert-evaluation-type.go", "infra-time-threshold.go"],
    "logalertconfig": ["log-alert-config-api.go", "log-level.go"],
    "maintenancewindow": ["maintenance-window-config-api.go"],
    "mobilealertconfig": ["mobile-alert-config.go"],
    "role": ["roles-api.go"],
    "sliconfig": ["sli-config-api.go"],
    "sloalertconfig": ["slo-alert-config-api.go"],
    "sloconfig": ["slo-config-api.go"],
    "slocorrection": ["slo-correction-config-api.go"],
    "syntheticalertconfig": ["synthetic-alert-config.go"],
    "syntheticlocation": ["synthetic-location.go"],
    "synthetictest": ["synthetic-test.go", "synthetic-test-rest-resource.go"],
    "team": ["teams-api.go"],
    "user": ["users-api.go"],
    "websitealertconfig": ["website-alert-config.go", "website-alert-rule.go", "website-impact-measurement-method.go"],
    "websitemonitoring": ["website-monitoring-config-api.go", "website-monitoring-config-rest-resource.go", "website-time-threshold.go"],
    "applicationalertconfig": ["application-alert-config.go", "application-alert-config-types.go", "application-alert-rule.go", "application-alert-evaluation-type.go"],
    "applicationconfig": ["application-configs-api.go", "application-config-scope.go", "included-application.go"],
}

def read_source_file(filepath):
    """Read a source file and return its content."""
    try:
        with open(filepath, 'r') as f:
            return f.read()
    except Exception as e:
        print(f"Error reading {filepath}: {e}")
        return None

def extract_content_after_package(content):
    """Extract content after the package declaration."""
    lines = content.split('\n')
    result_lines = []
    found_package = False
    
    for line in lines:
        if line.strip().startswith('package '):
            found_package = True
            continue
        if found_package:
            result_lines.append(line)
    
    return '\n'.join(result_lines).strip()

def process_package(package_name, source_files):
    """Process a single API package."""
    print(f"Processing {package_name}...")
    
    # Collect all content from source files
    all_content = []
    
    for source_file in source_files:
        source_path = TERRAFORM_RESTAPI / source_file
        if not source_path.exists():
            print(f"  WARNING: Source file not found: {source_path}")
            continue
        
        content = read_source_file(source_path)
        if content:
            # Extract content after package declaration
            extracted = extract_content_after_package(content)
            if extracted:
                all_content.append(extracted)
    
    if not all_content:
        print(f"  ERROR: No content found for {package_name}")
        return False
    
    # Combine all content
    combined_content = '\n\n'.join(all_content)
    
    # Create the output file
    output_dir = CLIENT_API / package_name
    output_file = output_dir / f"{package_name}.go"
    
    # Ensure directory exists
    output_dir.mkdir(parents=True, exist_ok=True)
    
    # Write the file with package declaration
    try:
        with open(output_file, 'w') as f:
            f.write(f"package {package_name}\n\n")
            f.write(combined_content)
            f.write("\n")
        
        # Count lines
        line_count = len(open(output_file).readlines())
        print(f"  ✅ Created {output_file} ({line_count} lines)")
        return True
    except Exception as e:
        print(f"  ERROR writing {output_file}: {e}")
        return False

def main():
    """Main recovery function."""
    print("=" * 60)
    print("API Package Recovery Script")
    print("=" * 60)
    print()
    
    success_count = 0
    fail_count = 0
    
    for package_name, source_files in sorted(PACKAGE_MAPPINGS.items()):
        if process_package(package_name, source_files):
            success_count += 1
        else:
            fail_count += 1
        print()
    
    print("=" * 60)
    print(f"Recovery Complete!")
    print(f"  ✅ Success: {success_count} packages")
    print(f"  ❌ Failed: {fail_count} packages")
    print("=" * 60)

if __name__ == "__main__":
    main()

# Made with Bob
