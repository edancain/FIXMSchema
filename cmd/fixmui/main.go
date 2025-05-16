package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	clearScreen()
	fmt.Println("==== FIXM XML Parser ====")
	fmt.Println("Using Go-generated FIXM schema files")
	fmt.Println()

	// List sample files
	files := listSampleFiles()
	if len(files) == 0 {
		fmt.Println("No FIXM sample files found in ./sample_messages directory")
		os.Exit(1)
	}

	// Display file menu
	fmt.Println("Available FIXM files:")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}

	fmt.Println("\nAdditional options:")
	fmt.Printf("%d. Process all files\n", len(files)+1)
	fmt.Printf("%d. Exit\n", len(files)+2)

	// Get user selection
	var choice int
	for {
		fmt.Print("\nEnter choice (1-" + strconv.Itoa(len(files)+2) + "): ")
		var input string
		fmt.Scanln(&input)

		var err error
		choice, err = strconv.Atoi(input)
		if err == nil && choice >= 1 && choice <= len(files)+2 {
			break
		}
		fmt.Println("Invalid choice, please try again")
	}

	// Process selection
	if choice == len(files)+2 {
		fmt.Println("Exiting...")
		return
	} else if choice == len(files)+1 {
		// Process all files
		fmt.Println("\nProcessing all files...")
		processFiles(files)
	} else {
		// Process selected file
		selectedFile := files[choice-1]
		fmt.Printf("\nProcessing file: %s\n", selectedFile)
		processFile(selectedFile)
	}

	// Wait for user to press enter before exiting
	fmt.Print("\nPress Enter to exit...")
	fmt.Scanln()
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func listSampleFiles() []string {
	var files []string

	// Look for sample_messages directory
	samplesDir := "sample_messages/4.3"
	if _, err := os.Stat(samplesDir); err != nil {
		return files
	}

	// Walk directory and collect XML files
	filepath.Walk(samplesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".xml") {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func processFiles(files []string) {
	for i, file := range files {
		if i > 0 {
			fmt.Println("\n---------------------------------------------------\n")
		}
		fmt.Printf("File %d/%d: %s\n", i+1, len(files), file)
		processFile(file)
	}
}

func processFile(filePath string) {
	// Create output file path
	baseName := filepath.Base(filePath)
	outputFileName := strings.TrimSuffix(baseName, filepath.Ext(baseName)) + "_output.txt"
	outputDir := "output"

	// Create output directory if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	outputPath := filepath.Join(outputDir, outputFileName)

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Build and execute the parser command
	cmd := exec.Command("go", "run", "cmd/fixmparser/main.go", filePath)

	// Capture command output to a buffer
	output, err := cmd.CombinedOutput()

	// Display output on the console
	fmt.Println(string(output))

	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		// Error is already in the output variable
	}

	// Write output to the file (includes both stdout and stderr)
	_, err = outputFile.Write(output)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Printf("Output saved to %s\n", outputPath)
}
