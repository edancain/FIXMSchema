package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// XSD parsing structures
type XSDSchema struct {
	XMLName      xml.Name         `xml:"schema"`
	Namespace    string           `xml:"targetNamespace,attr"`
	SimpleTypes  []XSDSimpleType  `xml:"simpleType"`
	ComplexTypes []XSDComplexType `xml:"complexType"`
	Elements     []XSDElement     `xml:"element"`
	Imports      []XSDImport      `xml:"import"`
	Includes     []XSDInclude     `xml:"include"`
}

type XSDImport struct {
	Namespace      string `xml:"namespace,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type XSDInclude struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type XSDSimpleType struct {
	Name        string         `xml:"name,attr"`
	Restriction XSDRestriction `xml:"restriction"`
	List        XSDList        `xml:"list"`
	Annotation  XSDAnnotation  `xml:"annotation"`
}

type XSDList struct {
	ItemType string `xml:"itemType,attr"`
}

type XSDRestriction struct {
	Base         string           `xml:"base,attr"`
	Enumerations []XSDEnumeration `xml:"enumeration"`
	Pattern      XSDPattern       `xml:"pattern"`
}

type XSDPattern struct {
	Value string `xml:"value,attr"`
}

type XSDEnumeration struct {
	Value      string        `xml:"value,attr"`
	Annotation XSDAnnotation `xml:"annotation"`
}

type XSDComplexType struct {
	Name       string        `xml:"name,attr"`
	Sequence   XSDSequence   `xml:"sequence"`
	Choice     XSDChoice     `xml:"choice"`
	Annotation XSDAnnotation `xml:"annotation"`
}

type XSDChoice struct {
	Elements []XSDElement `xml:"element"`
}

type XSDSequence struct {
	Elements []XSDElement `xml:"element"`
	Choice   XSDChoice    `xml:"choice"`
}

type XSDElement struct {
	Name       string        `xml:"name,attr"`
	Type       string        `xml:"type,attr"`
	MinOccurs  string        `xml:"minOccurs,attr"`
	MaxOccurs  string        `xml:"maxOccurs,attr"`
	Nillable   string        `xml:"nillable,attr"`
	Annotation XSDAnnotation `xml:"annotation"`
}

type XSDAnnotation struct {
	Documentation string `xml:"documentation"`
}

// Namespaces map
var namespaces = map[string]string{
	"http://www.fixm.aero/base/4.3":   "base",
	"http://www.fixm.aero/flight/4.3": "flight",
}

func main() {
	// Configuration
	schemaRoot := "schemas/fixm/core"
	outputDir := "generated/fixm"

	fmt.Println("Starting FIXM Go code generator...")

	// Clean up existing directory
	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, 0755)

	// Create base directory
	os.MkdirAll(filepath.Join(outputDir, "base"), 0755)

	// Process base XSD files
	baseDir := filepath.Join(schemaRoot, "base")
	baseFiles, _ := filepath.Glob(filepath.Join(baseDir, "*.xsd"))
	for _, xsdFile := range baseFiles {
		processXSDFile(xsdFile, filepath.Join(outputDir, "base"), "base")
	}

	// Create flight directory and subdirectories
	flightDir := filepath.Join(outputDir, "flight")
	os.MkdirAll(flightDir, 0755)

	// Process Flight.xsd
	flightXsd := filepath.Join(schemaRoot, "flight", "Flight.xsd")
	processXSDFile(flightXsd, flightDir, "flight")

	// Process flight subdirectories
	flightSubdirs, _ := filepath.Glob(filepath.Join(schemaRoot, "flight", "*"))
	for _, subdirPath := range flightSubdirs {
		info, err := os.Stat(subdirPath)
		if err == nil && info.IsDir() {
			subdir := filepath.Base(subdirPath)
			subOutputDir := filepath.Join(flightDir, subdir)
			os.MkdirAll(subOutputDir, 0755)

			// Process XSD files in subdirectory
			xsdFiles, _ := filepath.Glob(filepath.Join(subdirPath, "*.xsd"))
			for _, xsdFile := range xsdFiles {
				processXSDFile(xsdFile, subOutputDir, subdir)
			}
		}
	}

	fmt.Println("FIXM Go code generation completed successfully!")
}

func processXSDFile(xsdFile, outputDir, packageName string) {
	fmt.Printf("Processing %s -> %s (package: %s)\n", filepath.Base(xsdFile), outputDir, packageName)

	// Read XSD content
	content, err := ioutil.ReadFile(xsdFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse XSD
	var schema XSDSchema
	err = xml.Unmarshal(content, &schema)
	if err != nil {
		fmt.Printf("Error parsing XSD: %v\n", err)
		return
	}

	// Generate Go code
	outputFile := filepath.Join(outputDir, strings.ToLower(filepath.Base(xsdFile))+".go")
	outputFile = strings.Replace(outputFile, ".xsd.go", ".go", 1)

	var goCode strings.Builder
	goCode.WriteString(fmt.Sprintf("// Code generated from %s; DO NOT EDIT.\n\n", filepath.Base(xsdFile)))
	goCode.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	// Add imports based on the XSD imports and includes
	needsImport := (len(schema.Imports) > 0 || packageName != "base")

	if needsImport {
		goCode.WriteString("import (\n")
		if packageName != "base" {
			goCode.WriteString("\t\"github.com/edancain/FIXMSchema.git/generated/fixm/base\"\n")
		}
		// Add other imports if needed
		goCode.WriteString(")\n\n")
	}

	// Process simple types
	for _, simpleType := range schema.SimpleTypes {
		if simpleType.Name == "" {
			continue
		}

		goTypeName := simpleType.Name

		// Add documentation comment if available
		if simpleType.Annotation.Documentation != "" {
			doc := formatDocComment(simpleType.Annotation.Documentation)
			goCode.WriteString(fmt.Sprintf("// %s\n", doc))
		}

		// Process list types
		if simpleType.List.ItemType != "" {
			itemType := resolveType(simpleType.List.ItemType, packageName)
			goCode.WriteString(fmt.Sprintf("type %s []%s\n\n", goTypeName, itemType))
			continue
		}

		// Process restriction types
		if simpleType.Restriction.Base != "" {
			baseType := resolveType(simpleType.Restriction.Base, packageName)

			// Handle enumerations differently - they become string constants
			if len(simpleType.Restriction.Enumerations) > 0 {
				goCode.WriteString(fmt.Sprintf("type %s %s\n\n", goTypeName, baseType))

				// Add constants for each enumeration value
				goCode.WriteString(fmt.Sprintf("const (\n"))
				for _, enum := range simpleType.Restriction.Enumerations {
					constName := fmt.Sprintf("%s%s", goTypeName, toTitleCase(enum.Value))
					goCode.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", constName, goTypeName, enum.Value))
				}
				goCode.WriteString(")\n\n")
			} else {
				// For pattern-based restrictions, just use the base type
				goCode.WriteString(fmt.Sprintf("type %s %s\n\n", goTypeName, baseType))
			}
		}
	}

	// Process complex types
	for _, complexType := range schema.ComplexTypes {
		if complexType.Name == "" {
			continue
		}

		goTypeName := complexType.Name

		// Add documentation comment if available
		if complexType.Annotation.Documentation != "" {
			doc := formatDocComment(complexType.Annotation.Documentation)
			goCode.WriteString(fmt.Sprintf("// %s\n", doc))
		}

		// Begin struct definition
		goCode.WriteString(fmt.Sprintf("type %s struct {\n", goTypeName))

		// Process sequence elements
		for _, elem := range complexType.Sequence.Elements {
			processElement(&goCode, elem, packageName)
		}

		// Process choice elements
		for _, elem := range complexType.Sequence.Choice.Elements {
			processElement(&goCode, elem, packageName)
		}

		// Process direct choice elements
		for _, elem := range complexType.Choice.Elements {
			processElement(&goCode, elem, packageName)
		}

		// End struct definition
		goCode.WriteString("}\n\n")
	}

	// Process elements
	for _, elem := range schema.Elements {
		if elem.Type != "" && elem.Name != "" {
			// Add documentation comment if available
			if elem.Annotation.Documentation != "" {
				doc := formatDocComment(elem.Annotation.Documentation)
				goCode.WriteString(fmt.Sprintf("// %s\n", doc))
			}

			goType := resolveType(elem.Type, packageName)
			goCode.WriteString(fmt.Sprintf("var %s %s\n\n", toTitleCase(elem.Name), goType))
		}
	}

	// Write output file
	err = ioutil.WriteFile(outputFile, []byte(goCode.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated %s\n", outputFile)
}

func processElement(sb *strings.Builder, elem XSDElement, packageName string) {
	if elem.Name == "" || elem.Type == "" {
		return
	}

	// Convert element name to Go field name
	fieldName := toTitleCase(elem.Name)

	// Resolve element type to Go type
	goType := resolveType(elem.Type, packageName)

	// Determine if field is a slice
	isSlice := (elem.MaxOccurs == "unbounded" || elem.MaxOccurs > "1")
	if isSlice {
		goType = "[]" + goType
	}

	// Determine if field is a pointer
	isPointer := (elem.MinOccurs == "0" || elem.Nillable == "true")
	if isPointer && !isSlice {
		goType = "*" + goType
	}

	// Add documentation comment if available
	if elem.Annotation.Documentation != "" {
		doc := formatDocComment(elem.Annotation.Documentation)
		sb.WriteString(fmt.Sprintf("\t// %s\n", doc))
	}

	// Add field with XML tag
	sb.WriteString(fmt.Sprintf("\t%s %s `xml:\"%s\"`\n", fieldName, goType, elem.Name))
}

func resolveType(xsdType, packageName string) string {
	// Handle namespaced types
	if strings.Contains(xsdType, ":") {
		parts := strings.Split(xsdType, ":")
		ns := parts[0]
		typeName := parts[1]

		// Map namespace prefix to package
		if ns == "fb" {
			if packageName == "base" {
				return typeName
			}
			return "base." + typeName
		} else if ns == "fx" {
			if packageName == "flight" {
				return typeName
			}
			return "flight." + typeName
		}

		return typeName
	}

	// Handle XSD built-in types
	switch xsdType {
	case "xs:string":
		return "string"
	case "xs:int", "xs:integer":
		return "int"
	case "xs:boolean":
		return "bool"
	case "xs:decimal", "xs:double", "xs:float":
		return "float64"
	case "xs:date", "xs:dateTime", "xs:time":
		return "string" // Or use time.Time with custom XML unmarshal
	default:
		return xsdType
	}
}

func formatDocComment(doc string) string {
	// Trim whitespace and newlines
	doc = strings.TrimSpace(doc)

	// Replace copyright notice with a simpler reference
	if strings.Contains(doc, "Copyright (c)") {
		copyrightIdx := strings.Index(doc, "===========================================")
		if copyrightIdx > 0 {
			doc = strings.TrimSpace(doc[:copyrightIdx]) + " [Copyright details omitted]"
		}
	}

	// Replace newlines with space
	doc = regexp.MustCompile(`\s+`).ReplaceAllString(doc, " ")

	// Truncate if too long
	if len(doc) > 80 {
		return doc[:77] + "..."
	}

	return doc
}

func toTitleCase(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}
