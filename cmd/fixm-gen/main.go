package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Root schema directory
	schemaDir := "./schemas/fixm/core"
	// Use a completely new output directory
	outputDir := "./internal/fixm_final"

	// Path to the local xgen binary
	xgenPath := "./xgen/bin/xgen"

	// Clean up any existing directory
	os.RemoveAll(outputDir)

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Process the root schema first
	rootSchema := filepath.Join(schemaDir, "Fixm.xsd")
	processSchema(rootSchema, outputDir, xgenPath)

	// Process all other schemas
	err := filepath.Walk(schemaDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".xsd") {
			// Skip the root schema as we've already processed it
			if path == rootSchema {
				return nil
			}

			// Process each XSD file
			processSchema(path, outputDir, xgenPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through schemas: %v\n", err)
	}

	fmt.Println("Processing complete!")
}

func processSchema(schemaPath, outputDir, xgenPath string) {
	// Get relative path to maintain package structure
	relPath, err := filepath.Rel("./schemas/fixm/core", schemaPath)
	if err != nil {
		fmt.Printf("Error getting relative path for %s: %v\n", schemaPath, err)
		return
	}

	// Create output directory structure
	outPath := filepath.Join(outputDir, filepath.Dir(relPath))
	if err := os.MkdirAll(outPath, 0755); err != nil {
		fmt.Printf("Error creating output directory %s: %v\n", outPath, err)
		return
	}

	// Generate unique output file name - add unique suffix
	baseName := strings.TrimSuffix(filepath.Base(schemaPath), ".xsd")
	safeName := strings.ReplaceAll(relPath, "/", "_")
	safeName = strings.ReplaceAll(safeName, ".", "_")
	outFile := filepath.Join(outPath, strings.ToLower(baseName)+"_"+safeName+".go")

	// Generate package name from directory structure
	pkgPath := filepath.Dir(relPath)
	pkgName := "fixm"
	if pkgPath != "." {
		pkgName = strings.Replace(filepath.Base(pkgPath), "-", "", -1)
	}

	fmt.Printf("Processing %s -> %s (package: %s)\n", schemaPath, outFile, pkgName)

	// Run xgen command using the local binary with correct capitalization
	cmd := exec.Command(xgenPath,
		"-l", "Go", // Use capitalized "Go"
		"-i", schemaPath,
		"-o", outFile,
		"-p", pkgName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error generating code for %s: %v\n", schemaPath, err)
		fmt.Println(string(output))
	} else {
		fmt.Printf("Successfully generated %s\n", outFile)
	}
}
