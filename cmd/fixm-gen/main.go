package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"aqwari.net/xml/xsdgen"
)

func main() {
	var (
		schemaDir  = flag.String("schema", "./schemas/fixm/core", "Directory containing FIXM XSD files")
		outputPath = flag.String("output", "./internal/fixm", "Output directory for generated code")
		pkgName    = flag.String("package", "fixm", "Package name for generated code")
	)
	flag.Parse()

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(*outputPath, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Configure the code generator
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.PackageName(*pkgName),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.LogOutput(log.New(os.Stderr, "", 0)),
		xsdgen.LogLevel(3),
		xsdgen.UseFieldNames(),
	)

	// Find all .xsd files in the schema directory
	var xsdFiles []string
	err := filepath.Walk(*schemaDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".xsd" {
			xsdFiles = append(xsdFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking schema directory: %v", err)
	}

	// Generate code
	fmt.Printf("Generating Go code from %d FIXM XSD files...\n", len(xsdFiles))
	source, err := cfg.GenSource(xsdFiles...)
	if err != nil {
		log.Fatalf("Error generating code: %v", err)
	}

	// Write the generated code to a file
	outputFile := filepath.Join(*outputPath, "fixm_gen.go")
	if err := os.WriteFile(outputFile, source, 0644); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Printf("Successfully generated Go code at %s\n", outputFile)
}