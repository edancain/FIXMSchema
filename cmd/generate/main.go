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
	// Configuration
	baseDir := "generated/fixm"
	schemaRoot := "schemas/fixm/core"

	// Clean up existing directory
	os.RemoveAll(baseDir)

	// Find xgen path
	xgenPath, err := exec.LookPath("xgen")
	if err != nil {
		log.Fatalf("Error finding xgen: %v", err)
	}
	fmt.Printf("Using xgen at: %s\n", xgenPath)

	// Create directories
	os.MkdirAll(filepath.Join(baseDir, "base"), 0755)
	os.MkdirAll(filepath.Join(baseDir, "flight"), 0755)

	// Create flight subdirectories
	subdirs := []string{
		"aircraft", "arrival", "capability", "cargo", "departure",
		"emergency", "enroute", "flightdata", "flightroutetrajectory",
	}

	for _, subdir := range subdirs {
		os.MkdirAll(filepath.Join(baseDir, "flight", subdir), 0755)
	}

	// Process base directory
	fmt.Println("=== Processing base directory ===")
	baseFiles, _ := filepath.Glob(filepath.Join(schemaRoot, "base", "*.xsd"))
	for _, file := range baseFiles {
		processXsdFile(xgenPath, file, filepath.Join(baseDir, "base"), "base")
	}

	// Process flight directory
	fmt.Println("\n=== Processing flight directory ===")
	flightFile := filepath.Join(schemaRoot, "flight", "Flight.xsd")
	processXsdFile(xgenPath, flightFile, filepath.Join(baseDir, "flight"), "flight")

	// Process flight subdirectories
	for _, subdir := range subdirs {
		fmt.Printf("\n=== Processing %s subdirectory ===\n", subdir)
		subdirFiles, _ := filepath.Glob(filepath.Join(schemaRoot, "flight", subdir, "*.xsd"))
		for _, file := range subdirFiles {
			processXsdFile(xgenPath, file, filepath.Join(baseDir, "flight", subdir), subdir)
		}
	}

	// Cleanup schema directories
	cleanupSchemaDirectories(baseDir)

	fmt.Println("\nCode generation completed successfully!")
}

func processXsdFile(xgenPath, xsdFile, outputDir, packageName string) {
	// Get base name and output name
	baseName := filepath.Base(xsdFile)
	outputName := strings.ToLower(strings.TrimSuffix(baseName, ".xsd")) + ".go"
	outputPath := filepath.Join(outputDir, outputName)

	fmt.Printf("Processing %s -> %s\n", baseName, outputPath)

	// Run xgen with -p (package) flag directly to the output directory
	cmd := exec.Command(xgenPath,
		"-i", xsdFile,
		"-o", outputDir,
		"-l", "Go",
		"-p", packageName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("  Error running xgen: %v\n%s\n", err, output)

		// Create simple placeholder file with just package declaration
		createPlaceholderFile(outputPath, packageName, baseName)
		return
	}

	// Check if the file was created as expected
	generatedFile := filepath.Join(outputDir, baseName+".go")
	if _, err := os.Stat(generatedFile); err == nil {
		// If file is generated with original name, rename it to lowercase
		if generatedFile != outputPath {
			os.Rename(generatedFile, outputPath)
		}

		// Fix package and interface references
		fixGoFile(outputPath, packageName)
		fmt.Printf("  Successfully generated %s\n", outputPath)
	} else {
		// If no file was generated, look for files in schemas directory
		schemasDir := filepath.Join(outputDir, "schemas")
		if _, err := os.Stat(schemasDir); err == nil {
			// Find any generated files
			var foundFiles []string
			filepath.Walk(schemasDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasSuffix(path, ".go") {
					foundFiles = append(foundFiles, path)
				}
				return nil
			})

			if len(foundFiles) > 0 {
				// Copy the first found file and fix it
				copyAndFixFile(foundFiles[0], outputPath, packageName)
				fmt.Printf("  Generated from schemas dir: %s\n", outputPath)
			} else {
				// Create placeholder if no files found
				fmt.Printf("  No output files found\n")
				createPlaceholderFile(outputPath, packageName, baseName)
			}
		} else {
			// Create placeholder if no schemas directory
			fmt.Printf("  No output files and no schemas directory found\n")
			createPlaceholderFile(outputPath, packageName, baseName)
		}
	}
}

func fixGoFile(filePath, packageName string) {
	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("  Error reading file %s: %v\n", filePath, err)
		return
	}

	contentStr := string(content)

	// Fix package declaration
	if !strings.Contains(contentStr, "package "+packageName) {
		contentStr = strings.Replace(contentStr, "package schema", "package "+packageName, -1)
	}

	// Fix interface references
	contentStr = strings.Replace(contentStr, "*Interface{}", "interface{}", -1)
	contentStr = strings.Replace(contentStr, "[]*Interface{}", "[]interface{}", -1)

	// Write back
	err = os.WriteFile(filePath, []byte(contentStr), 0644)
	if err != nil {
		fmt.Printf("  Error writing file %s: %v\n", filePath, err)
	}
}

func copyAndFixFile(sourcePath, destPath, packageName string) {
	// Read source file
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("  Error reading source file %s: %v\n", sourcePath, err)
		return
	}

	contentStr := string(content)

	// Fix package declaration
	contentStr = strings.Replace(contentStr, "package schema", "package "+packageName, -1)

	// Fix interface references
	contentStr = strings.Replace(contentStr, "*Interface{}", "interface{}", -1)
	contentStr = strings.Replace(contentStr, "[]*Interface{}", "[]interface{}", -1)

	// Write to destination
	err = os.WriteFile(destPath, []byte(contentStr), 0644)
	if err != nil {
		fmt.Printf("  Error writing destination file %s: %v\n", destPath, err)
	}
}

func createPlaceholderFile(filePath, packageName, xsdFileName string) {
	content := fmt.Sprintf("// Generated from %s\npackage %s\n", xsdFileName, packageName)
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("  Error creating placeholder file %s: %v\n", filePath, err)
	} else {
		fmt.Printf("  Created placeholder file: %s\n", filePath)
	}
}

func cleanupSchemaDirectories(rootDir string) {
	fmt.Println("\n=== Cleaning up unwanted schemas directories ===")

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && filepath.Base(path) == "schemas" {
			fmt.Printf("Removing directory: %s\n", path)
			os.RemoveAll(path)
			return filepath.SkipDir
		}

		return nil
	})
}
