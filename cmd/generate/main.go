package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Create output directory structure
	baseDir := "generated/fixm"
	os.MkdirAll(filepath.Join(baseDir, "base"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "flight"), 0755)

	// Create flight subdirectories
	flightSubdirs := []string{
		"aircraft", "arrival", "capability", "cargo", "departure",
		"emergency", "enroute", "flightdata", "flightroutetrajectory",
	}

	for _, dir := range flightSubdirs {
		os.MkdirAll(filepath.Join(baseDir, "flight", dir), 0755)
	}

	// Find the path to the xgen executable
	xgenPath, err := exec.LookPath("xgen")
	if err != nil {
		log.Fatalf("xgen not found in PATH: %v", err)
	}

	fmt.Printf("Using xgen at: %s\n", xgenPath)

	// Process base schemas
	fmt.Println("Processing base schemas...")
	baseFiles, err := filepath.Glob("schemas/fixm/core/base/*.xsd")
	if err != nil {
		log.Fatalf("Error finding base schemas: %v", err)
	}

	for _, file := range baseFiles {
		baseName := filepath.Base(file)
		fmt.Printf("  Processing %s...\n", baseName)

		cmd := exec.Command(xgenPath,
			"-i", file,
			"-o", filepath.Join(baseDir, "base"),
			"-l", "Go",
			"-p", "base")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Printf("Warning: Issue processing %s: %v", file, err)
		}
	}

	// Process Flight.xsd
	fmt.Println("Processing Flight.xsd...")
	flightCmd := exec.Command(xgenPath,
		"-i", "schemas/fixm/core/flight/Flight.xsd",
		"-o", filepath.Join(baseDir, "flight"),
		"-l", "Go",
		"-p", "flight")

	flightCmd.Stdout = os.Stdout
	flightCmd.Stderr = os.Stderr

	if err := flightCmd.Run(); err != nil {
		log.Printf("Warning: Issue processing Flight.xsd: %v", err)
	}

	// Process flight subdirectory schemas
	fmt.Println("Processing flight subdirectory schemas...")

	for _, subdir := range flightSubdirs {
		subdirPath := filepath.Join("schemas/fixm/core/flight", subdir)
		files, err := filepath.Glob(filepath.Join(subdirPath, "*.xsd"))
		if err != nil {
			log.Printf("Warning: Error finding schemas in %s: %v", subdirPath, err)
			continue
		}

		for _, file := range files {
			baseName := filepath.Base(file)
			fmt.Printf("  Processing %s/%s...\n", subdir, baseName)

			cmd := exec.Command(xgenPath,
				"-i", file,
				"-o", filepath.Join(baseDir, "flight", subdir),
				"-l", "Go",
				"-p", fmt.Sprintf("flight/%s", subdir))

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Printf("Warning: Issue processing %s: %v", file, err)
			}
		}
	}

	// Process root Fixm.xsd
	fmt.Println("Processing root Fixm.xsd...")
	rootCmd := exec.Command(xgenPath,
		"-i", "schemas/fixm/core/Fixm.xsd",
		"-o", baseDir,
		"-l", "Go",
		"-p", "fixm")

	rootCmd.Stdout = os.Stdout
	rootCmd.Stderr = os.Stderr

	if err := rootCmd.Run(); err != nil {
		log.Printf("Warning: Issue processing Fixm.xsd: %v", err)
	}

	// Fix import references and packages
	fmt.Println("Fixing import references...")

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			fixGoFile(path)
		}

		return nil
	})

	if err != nil {
		log.Printf("Warning: Error fixing import references: %v", err)
	}

	// Remove unwanted Base.xsd.go files in subdirectories
	fmt.Println("Cleaning up unwanted files...")

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Remove Base.xsd.go files in schemas subdirectories
		if !info.IsDir() &&
			strings.HasSuffix(path, "Base.xsd.go") &&
			strings.Contains(path, "schemas") {
			os.Remove(path)
			fmt.Printf("  Removed %s\n", path)
		}

		return nil
	})

	if err != nil {
		log.Printf("Warning: Error cleaning up files: %v", err)
	}

	// Create fixm.go in root directory to import all packages
	fmt.Println("Creating fixm.go in root directory...")

	fixmContent := `package fixm

import (
	_ "github.com/yourusername/fixm/generated/fixm/base"
	_ "github.com/yourusername/fixm/generated/fixm/flight"
)
`

	err = os.WriteFile(filepath.Join(baseDir, "fixm.go"), []byte(fixmContent), 0644)
	if err != nil {
		log.Printf("Warning: Error creating fixm.go: %v", err)
	}

	fmt.Println("Code generation completed")
}

func fixGoFile(filePath string) {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Warning: Could not read file %s: %v", filePath, err)
		return
	}

	contentStr := string(content)

	// Fix package declaration
	relPath, err := filepath.Rel("generated/fixm", filePath)
	if err != nil {
		log.Printf("Warning: Could not get relative path for %s: %v", filePath, err)
		return
	}

	packagePath := filepath.Dir(relPath)
	packageName := strings.Replace(packagePath, "/", "_", -1)
	if packageName == "." {
		packageName = "fixm"
	}

	// Fix package declaration
	contentStr = strings.Replace(contentStr, "package schema", "package "+packageName, -1)

	// Fix interface references
	contentStr = strings.Replace(contentStr, "*Interface{}", "interface{}", -1)
	contentStr = strings.Replace(contentStr, "[]*Interface{}", "[]interface{}", -1)

	// Add imports for base types if needed
	if strings.Contains(contentStr, "CharacterStringType") ||
		strings.Contains(contentStr, "CountPositiveType") ||
		strings.Contains(contentStr, "AircraftExtensionType") {
		importStr := `
import (
	"github.com/yourusername/fixm/generated/fixm/base"
)
`
		// Add import after package declaration
		lines := strings.Split(contentStr, "\n")
		contentStr = lines[0] + "\n" + importStr + strings.Join(lines[1:], "\n")

		// Fix references to base types
		contentStr = strings.Replace(contentStr, "CharacterStringType", "base.CharacterStringType", -1)
		contentStr = strings.Replace(contentStr, "CountPositiveType", "base.CountPositiveType", -1)
		contentStr = strings.Replace(contentStr, "AircraftExtensionType", "base.AircraftExtensionType", -1)
		contentStr = strings.Replace(contentStr, "AircraftTypeExtensionType", "base.AircraftTypeExtensionType", -1)
		contentStr = strings.Replace(contentStr, "AircraftTypeDesignatorType", "base.AircraftTypeDesignatorType", -1)
	}

	// Write the fixed content back
	err = os.WriteFile(filePath, []byte(contentStr), 0644)
	if err != nil {
		log.Printf("Warning: Could not write fixed file %s: %v", filePath, err)
		return
	}

	fmt.Printf("  Fixed %s\n", filePath)
}
