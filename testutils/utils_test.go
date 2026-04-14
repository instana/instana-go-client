package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGetRootFolder tests the GetRootFolder function
func TestGetRootFolder(t *testing.T) {
	rootFolder, err := GetRootFolder()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if rootFolder == "" {
		t.Fatal("Expected non-empty root folder")
	}

	// Verify go.mod exists in the root folder
	goModPath := filepath.Join(rootFolder, "go.mod")
	if !fileExists(goModPath) {
		t.Errorf("Expected go.mod to exist at %s", goModPath)
	}
}

// TestLookupRootFolder_WithGoMod tests finding root folder with go.mod
func TestLookupRootFolder_WithGoMod(t *testing.T) {
	// Create a temporary directory structure
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "subdir", "nested")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directories: %v", err)
	}

	// Create go.mod in root
	goModPath := filepath.Join(tempDir, "go.mod")
	err = os.WriteFile(goModPath, []byte("module test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Test from nested directory
	result, err := lookupRootFolder(subDir, 0)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != tempDir {
		t.Errorf("Expected root folder %s, got %s", tempDir, result)
	}
}

// TestLookupRootFolder_WithMainGo tests finding root folder with main.go
func TestLookupRootFolder_WithMainGo(t *testing.T) {
	// Create a temporary directory structure
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "subdir")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directories: %v", err)
	}

	// Create main.go in root (no go.mod)
	mainGoPath := filepath.Join(tempDir, "main.go")
	err = os.WriteFile(mainGoPath, []byte("package main"), 0644)
	if err != nil {
		t.Fatalf("Failed to create main.go: %v", err)
	}

	// Test from subdirectory
	result, err := lookupRootFolder(subDir, 0)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != tempDir {
		t.Errorf("Expected root folder %s, got %s", tempDir, result)
	}
}

// TestLookupRootFolder_MaxDepthExceeded tests error when max depth is exceeded
func TestLookupRootFolder_MaxDepthExceeded(t *testing.T) {
	// Create a deep directory structure without go.mod or main.go
	tempDir := t.TempDir()
	deepDir := filepath.Join(tempDir, "a", "b", "c", "d", "e", "f")
	err := os.MkdirAll(deepDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directories: %v", err)
	}

	// Test from deep directory (should fail after 5 levels)
	_, err = lookupRootFolder(deepDir, 0)

	if err == nil {
		t.Fatal("Expected error for max depth exceeded, got nil")
	}

	expectedError := "failed to find root folder"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

// TestLookupRootFolder_AlreadyAtRoot tests when already at root
func TestLookupRootFolder_AlreadyAtRoot(t *testing.T) {
	// Create a temporary directory with go.mod
	tempDir := t.TempDir()
	goModPath := filepath.Join(tempDir, "go.mod")
	err := os.WriteFile(goModPath, []byte("module test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Test from root directory itself
	result, err := lookupRootFolder(tempDir, 0)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != tempDir {
		t.Errorf("Expected root folder %s, got %s", tempDir, result)
	}
}

// TestLookupRootFolder_PreferGoModOverMainGo tests that go.mod is preferred over main.go
func TestLookupRootFolder_PreferGoModOverMainGo(t *testing.T) {
	// Create a temporary directory with both go.mod and main.go
	tempDir := t.TempDir()

	goModPath := filepath.Join(tempDir, "go.mod")
	err := os.WriteFile(goModPath, []byte("module test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	mainGoPath := filepath.Join(tempDir, "main.go")
	err = os.WriteFile(mainGoPath, []byte("package main"), 0644)
	if err != nil {
		t.Fatalf("Failed to create main.go: %v", err)
	}

	// Test - should find go.mod first
	result, err := lookupRootFolder(tempDir, 0)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != tempDir {
		t.Errorf("Expected root folder %s, got %s", tempDir, result)
	}
}

// TestFileExists_FileExists tests fileExists with an existing file
func TestFileExists_FileExists(t *testing.T) {
	// Create a temporary file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")
	err := os.WriteFile(testFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := fileExists(testFile)

	if !result {
		t.Error("Expected fileExists to return true for existing file")
	}
}

// TestFileExists_FileDoesNotExist tests fileExists with a non-existent file
func TestFileExists_FileDoesNotExist(t *testing.T) {
	result := fileExists("/nonexistent/path/to/file.txt")

	if result {
		t.Error("Expected fileExists to return false for non-existent file")
	}
}

// TestFileExists_Directory tests fileExists with a directory (should return false)
func TestFileExists_Directory(t *testing.T) {
	tempDir := t.TempDir()

	result := fileExists(tempDir)

	if result {
		t.Error("Expected fileExists to return false for directory")
	}
}

// TestFileExists_EmptyPath tests fileExists with empty path
func TestFileExists_EmptyPath(t *testing.T) {
	result := fileExists("")

	if result {
		t.Error("Expected fileExists to return false for empty path")
	}
}

// TestLookupRootFolder_MultipleNestedLevels tests finding root from various depths
func TestLookupRootFolder_MultipleNestedLevels(t *testing.T) {
	// Create a directory structure with go.mod at root
	tempDir := t.TempDir()
	goModPath := filepath.Join(tempDir, "go.mod")
	err := os.WriteFile(goModPath, []byte("module test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Test from different nesting levels
	testCases := []struct {
		name  string
		depth string
	}{
		{"level1", "a"},
		{"level2", filepath.Join("a", "b")},
		{"level3", filepath.Join("a", "b", "c")},
		{"level4", filepath.Join("a", "b", "c", "d")},
		{"level5", filepath.Join("a", "b", "c", "d", "e")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nestedDir := filepath.Join(tempDir, tc.depth)
			err := os.MkdirAll(nestedDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create nested directory: %v", err)
			}

			result, err := lookupRootFolder(nestedDir, 0)

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if result != tempDir {
				t.Errorf("Expected root folder %s, got %s", tempDir, result)
			}
		})
	}
}

// TestLookupRootFolder_WithInitialLevel tests lookupRootFolder with non-zero initial level
func TestLookupRootFolder_WithInitialLevel(t *testing.T) {
	// Create a shallow directory structure
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "sub")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directories: %v", err)
	}

	// No go.mod or main.go, starting at level 4 (should fail quickly)
	_, err = lookupRootFolder(subDir, 4)

	if err == nil {
		t.Fatal("Expected error when starting at high level")
	}
}

// TestGetRootFolder_Integration tests GetRootFolder in actual project context
func TestGetRootFolder_Integration(t *testing.T) {
	// This test runs in the actual project context
	rootFolder, err := GetRootFolder()

	if err != nil {
		t.Fatalf("Expected no error in actual project, got %v", err)
	}

	// Verify the root folder contains expected files
	goModPath := filepath.Join(rootFolder, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		t.Errorf("Expected go.mod to exist in project root at %s", goModPath)
	}

	// Verify testutils directory exists
	testutilsPath := filepath.Join(rootFolder, "testutils")
	if _, err := os.Stat(testutilsPath); os.IsNotExist(err) {
		t.Errorf("Expected testutils directory to exist at %s", testutilsPath)
	}
}
