# Instana Package Migration - File Classification

## Files to KEEP (Infrastructure - 11 files + 9 tests)

### Core Infrastructure Files:
1. `config.go` - Client configuration
2. `config_builder.go` - Configuration builder
3. `config_loader.go` - Configuration loader
4. `config_validator.go` - Configuration validator
5. `rest-client.go` - REST client implementation
6. `errors.go` - Error types
7. `logger.go` - Logger implementation
8. `log-level.go` - Log level types
9. `rate_limiter.go` - Rate limiter
10. `retry.go` - Retry logic
11. `Instana-api.go` - Initialization methods (updated)

### Infrastructure Test Files:
1. `config_test.go`
2. `config_builder_test.go`
3. `config_loader_test.go`
4. `config_validator_test.go`
5. `rest-client_test.go`
6. `errors_test.go`
7. `log-level_test.go`
8. `rate_limiter_test.go`
9. `retry_test.go`

### Utility Files:
1. `dummy.go` - Package marker

## Files to REMOVE (Old API Implementations - 90+ files)

### API Implementation Files (to be removed):
- All `*-api.go` files (alerting-channels-api.go, api-tokens-api.go, etc.)
- All `*-config.go` files (application-alert-config.go, mobile-alert-config.go, etc.)
- All model files (access-rule.go, synthetic-test.go, custom-dashboard.go, etc.)
- All type files (aggregation.go, severity.go, threshold.go, operator.go, granularity.go, etc.)
- All custom resource implementations (default-rest-resource.go, read-only-rest-resource.go, etc.)
- All unmarshaller files (default-json-unmarshaller.go, custom-payload-fields-unmarshaller-adapter.go, etc.)
- All special resource implementations (synthetic-test-rest-resource.go, website-monitoring-config-rest-resource.go)
- Interface file (instana-rest-resource.go)
- Provider models (provider-models.go)
- All corresponding test files

### Complete List of Files to Remove (90 files):
1. access-rule.go
2. access-type.go
3. access-type_test.go
4. aggregation.go
5. aggregation_test.go
6. alert-event-type.go
7. alerting-channel-type.go
8. alerting-channels-api.go
9. alerts-api.go
10. api-tokens-api.go
11. application-alert-config.go
12. application-alert-config-types.go
13. application-alert-evaluation-type.go
14. application-alert-evaluation-type_test.go
15. application-alert-rule.go
16. application-config-scope.go
17. application-config-scope_test.go
18. application-configs-api.go
19. automation-action-api.go
20. automation-policy-api.go
21. boundary-scope.go
22. boundary-scope_test.go
23. builtin-event-specification-api.go
24. custom-dashboard.go
25. custom-event-specficiations-api.go
26. custom-payload-field.go
27. custom-payload-field_test.go
28. custom-payload-fields-unmarshaller-adapter.go
29. custom-payload-fields-unmarshaller-adapter_test.go
30. default-json-unmarshaller.go
31. default-json-unmarshaller_test.go
32. default-rest-resource.go
33. default-rest-resource_test.go
34. default-rest-resource_create-post-update-not-supported_test.go
35. default-rest-resource_create-post-update-post_test.go
36. default-rest-resource_create-post-update-put_test.go
37. default-rest-resource_create-put-update-not-supported_test.go
38. default-rest-resource_create-put-update-put_test.go
39. granularity.go
40. granularity_test.go
41. groups-api.go
42. groups-api_test.go
43. host-agent-json-unmarshaller.go
44. host-agents-api.go
45. included-application.go
46. infra-alert-configs.go
47. infra-alert-evaluation-type.go
48. infra-alert-rule.go
49. infra-time-threshold.go
50. Instana-api_test.go
51. instana-rest-resource.go
52. log-alert-config-api.go
53. maintenance-window-config-api.go
54. mobile-alert-config.go
55. operator.go
56. operator_test.go
57. provider-models.go
58. read-only-rest-resource.go
59. read-only-rest-resource_test.go
60. relation-type.go
61. relation-type_test.go
62. roles-api.go
63. rule-with-threshold.go
64. severity.go
65. severity_test.go
66. sli-config-api.go
67. slo-alert-config-api.go
68. slo-config-api.go
69. slo-correction-config-api.go
70. synthetic-alert-config.go
71. synthetic-location.go
72. synthetic-test.go
73. synthetic-test-rest-resource.go
74. synthetic-test-rest-resource_test.go
75. tag-filter.go
76. tag-filter_test.go
77. teams-api.go
78. threshold.go
79. threshold_test.go
80. users-api.go
81. website-alert-config.go
82. website-alert-rule.go
83. website-impact-measurement-method.go
84. website-impact-measurement-method_test.go
85. website-monitoring-config-api.go
86. website-monitoring-config-rest-resource.go
87. website-monitoring-config-rest-resource_test.go
88. website-time-threshold.go

## Summary

- **Keep**: 21 files (11 core + 9 tests + 1 utility)
- **Remove**: 88 files (old API implementations)
- **Total instana/ files**: 109 files

After migration, the instana/ package will be minimal and focused only on:
1. Client configuration and initialization
2. REST client implementation  
3. Error handling
4. Logging
5. Rate limiting and retry logic