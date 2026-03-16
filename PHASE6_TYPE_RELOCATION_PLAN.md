# Phase 6b: Type Relocation Plan

## Overview
After removing 88 old API implementation files from instana/, we need to relocate the types that were defined there. This document provides a comprehensive plan for where each type should be moved.

## Missing Types Analysis

### Category 1: Common Types (Already in shared/types/)
These types are already available in shared/types/ package:
- ✅ `Threshold` - shared/types/threshold.go
- ✅ `Severity` - shared/types/severity.go
- ✅ `Granularity` - shared/types/granularity.go
- ✅ `Operator` - shared/types/operator.go
- ✅ `Aggregation` - shared/types/aggregation.go

**Action**: Update imports in API packages to use `shared/types`

### Category 2: Tag Filtering (Already in shared/tagfilter/)
- ✅ `TagFilter` - shared/tagfilter/filter.go

**Action**: Update imports in API packages to use `shared/tagfilter`

### Category 3: Complex Shared Models (Need new shared/models/ package)
These are complex types used by multiple APIs:

1. **BoundaryScope** - Used by: ApplicationConfig, ApplicationAlertConfig
2. **AccessRule** - Used by: CustomDashboard, ApplicationConfig  
3. **CustomPayloadField** - Used by: ApplicationAlertConfig, AlertingConfig, InfraAlertConfig
4. **IncludedApplication** - Used by: ApplicationAlertConfig

**Action**: Create new `shared/models/` package with these types

### Category 4: Application Alert Specific (Move to api/applicationalertconfig/)
These are specific to application alerts:

1. **ApplicationAlertEvaluationType** - Enum for evaluation types
2. **ApplicationAlertRule** - Alert rule structure
3. **ApplicationAlertRuleWithThresholds** - Rule with thresholds
4. **ApplicationAlertTimeThreshold** - Time-based threshold

**Action**: Create new files in `api/applicationalertconfig/` package

### Category 5: Application Config Specific (Move to api/applicationconfig/)
1. **ApplicationConfigScope** - Scope configuration

**Action**: Add to `api/applicationconfig/` package

### Category 6: Alerting Channel Specific (Move to api/alertingchannel/)
1. **AlertingChannelType** - Enum for channel types

**Action**: Add to `api/alertingchannel/` package

### Category 7: Alerting Config Specific (Move to api/alertingconfig/)
1. **EventFilteringConfiguration** - Event filtering config

**Action**: Add to `api/alertingconfig/` package

### Category 8: Automation Specific (Move to api/automationaction/ and api/automationpolicy/)
1. **Field** - Field definition (automationaction)
2. **Parameter** - Parameter definition (automationaction)
3. **Trigger** - Trigger definition (automationpolicy)
4. **TypeConfiguration** - Type config (automationpolicy)

**Action**: Add to respective API packages

### Category 9: Custom Event Spec Specific (Move to api/customeventspec/)
1. **RuleSpecification** - Rule specification

**Action**: Add to `api/customeventspec/` package

### Category 10: Group/RBAC Specific (Move to api/group/)
1. **APIMember** - Member definition
2. **APIPermissionSetWithRoles** - Permission set

**Action**: Add to `api/group/` package

### Category 11: Infrastructure Alert Specific (Move to api/infraalertconfig/)
1. **InfraTimeThreshold** - Time threshold for infra alerts
2. **RuleWithThreshold** - Rule with threshold
3. **InfrastructureAlertConfig** - Main config (might be typo in model.go)

**Action**: Add to `api/infraalertconfig/` package

### Category 12: JSON Import Issue
1. **json** - Missing import in customdashboard

**Action**: Add `import "encoding/json"` to customdashboard/model.go

## Implementation Strategy

### Step 1: Create shared/models/ Package
Create new package for complex shared models:
```
shared/models/
├── boundary-scope.go
├── access-rule.go
├── custom-payload-field.go
├── included-application.go
└── doc.go
```

### Step 2: Add API-Specific Types
For each API package, create additional files for types:
- `types.go` - For enums and simple types
- `rules.go` - For rule-related types
- `config.go` - For configuration types

### Step 3: Update Imports
Update all model.go files to import from correct locations:
- `shared/types` for common types
- `shared/tagfilter` for tag filtering
- `shared/models` for complex shared models
- Local package for API-specific types

### Step 4: Fix Special Cases
- Fix json import in customdashboard
- Fix InfrastructureAlertConfig naming issue

## File Creation Summary

**New files to create:**
1. shared/models/boundary-scope.go
2. shared/models/access-rule.go
3. shared/models/custom-payload-field.go
4. shared/models/included-application.go
5. shared/models/doc.go
6. api/applicationalertconfig/types.go
7. api/applicationalertconfig/rules.go
8. api/applicationconfig/types.go
9. api/alertingchannel/types.go
10. api/alertingconfig/types.go
11. api/automationaction/types.go
12. api/automationpolicy/types.go
13. api/customeventspec/types.go
14. api/group/types.go
15. api/infraalertconfig/types.go
16. api/infraalertconfig/rules.go

**Total**: 16 new files

## Next Steps

1. Extract type definitions from old instana/ package files (if backed up)
2. Create shared/models/ package with 5 files
3. Create API-specific type files (11 files)
4. Update all imports in model.go files
5. Fix json import issue
6. Verify compilation

## Estimated Effort

- Type extraction and creation: 2-3 hours
- Import updates: 1 hour
- Testing and verification: 1 hour
- **Total**: 4-5 hours

This is a significant refactoring task that requires careful attention to ensure all types are correctly relocated and all imports are updated.