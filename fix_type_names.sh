#!/bin/bash
# Fix type name mismatches

# maintenancewindow: MaintenanceWindowConfig -> MaintenanceWindow
sed -i '' 's/type MaintenanceWindowConfig struct/type MaintenanceWindow struct/g' api/maintenancewindow/maintenancewindow.go
sed -i '' 's/func (m \*MaintenanceWindowConfig)/func (m *MaintenanceWindow)/g' api/maintenancewindow/maintenancewindow.go

# sliconfig: SLIConfig -> SliConfig  
sed -i '' 's/type SLIConfig struct/type SliConfig struct/g' api/sliconfig/sliconfig.go
sed -i '' 's/func (s \*SLIConfig)/func (s *SliConfig)/g' api/sliconfig/sliconfig.go

# sloconfig: SLOConfig -> SloConfig
sed -i '' 's/type SLOConfig struct/type SloConfig struct/g' api/sloconfig/sloconfig.go
sed -i '' 's/func (s \*SLOConfig)/func (s *SloConfig)/g' api/sloconfig/sloconfig.go

# slocorrection: SLOCorrection -> SloCorrection
sed -i '' 's/type SLOCorrection struct/type SloCorrection struct/g' api/slocorrection/slocorrection.go
sed -i '' 's/func (s \*SLOCorrection)/func (s *SloCorrection)/g' api/slocorrection/slocorrection.go

echo "✅ Type names fixed"
