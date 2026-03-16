#!/bin/bash
# Copy missing types from Terraform provider to shared/types

TERRAFORM_RESTAPI="../terraform-provider-instana/internal/restapi"
SHARED_TYPES="shared/types"

echo "Copying missing types to shared/types..."

# Copy and convert each file
for file in custom-payload-field access-rule boundary-scope included-application website-time-threshold website-impact-measurement-method access-type relation-type; do
    if [ -f "$TERRAFORM_RESTAPI/${file}.go" ]; then
        echo "  Processing ${file}.go..."
        # Copy file and change package name
        sed 's/^package restapi$/package types/' "$TERRAFORM_RESTAPI/${file}.go" > "$SHARED_TYPES/${file}.go"
        echo "    ✅ Created $SHARED_TYPES/${file}.go"
    fi
done

# Also need to copy Scheduling type from slo-correction-config-api.go
echo "  Extracting Scheduling type..."
cat > "$SHARED_TYPES/scheduling.go" << 'SCHEDEOF'
package types

// Scheduling represents scheduling configuration for automation policies and SLO corrections
type Scheduling struct {
	Start      int64        `json:"start"`
	End        int64        `json:"end"`
	Recurrence Recurrence   `json:"recurrence"`
	Duration   Duration     `json:"duration"`
}

// Recurrence represents recurrence configuration
type Recurrence struct {
	Type     string `json:"type"`
	Interval int    `json:"interval"`
}

// Duration represents duration configuration
type Duration struct {
	Amount int          `json:"amount"`
	Unit   DurationUnit `json:"unit"`
}

// DurationUnit represents the unit of duration
type DurationUnit string

const (
	// DurationUnitMinute represents minute duration unit
	DurationUnitMinute DurationUnit = "MINUTE"
	// DurationUnitHour represents hour duration unit
	DurationUnitHour DurationUnit = "HOUR"
	// DurationUnitDay represents day duration unit
	DurationUnitDay DurationUnit = "DAY"
)
SCHEDEOF
echo "    ✅ Created $SHARED_TYPES/scheduling.go"

echo ""
echo "✅ All missing types copied to shared/types"
