package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Get the path to the directory where the executable is located
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting the path of the executable:", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// Set the paths of the input and output directories
	inputDir := filepath.Join(exeDir, "input")
	outputDir := filepath.Join(exeDir, "output")

	// Check if the input and output directories exist, and create them if they don't
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		if err := os.Mkdir(inputDir, 0755); err != nil {
			fmt.Println("Error creating the input directory:", err)
			return
		}
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0755); err != nil {
			fmt.Println("Error creating the output directory:", err)
			return
		}
	}

	// Find all the PNG files in the input directory and its subdirectories
	pngFiles := []string{}
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".png" {
			pngFiles = append(pngFiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error finding PNG files in the input directory:", err)
		return
	}

	// Display information about the files found
	fmt.Println("PNG files found in the input directory:")
	for _, f := range pngFiles {
		fmt.Println("-", f)
	}

	// Process the PNG files found and create the same directory structure in the output directory
	startTime := time.Now()

	for _, f := range pngFiles {
		inputFile, err := os.Open(f)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", f, err)
			continue
		}
		defer inputFile.Close()

		inputImage, err := png.Decode(inputFile)
		if err != nil {
			fmt.Printf("Error decoding file %s: %s\n", f, err)
			continue
		}

		inputFileInfo, err := os.Stat(f)
		if err != nil {
			fmt.Printf("Error getting file information for %s: %s\n", f, err)
			continue
		}

		outputPath := filepath.Join(outputDir, f[len(inputDir):])
		outputDir := filepath.Dir(outputPath)
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				fmt.Printf("Error creating output directory %s: %s\n", outputDir, err)
				continue
			}
		}

		outputImage := resize.Resize(0, 0, inputImage, resize.MitchellNetravali)

		outputFile, err := os.Create(outputPath)
		if err != nil {
			fmt.Printf("Error creating output file %s: %s\n", outputPath, err)
			continue
		}

		if err := png.Encode(outputFile, outputImage); err != nil {
			fmt.Printf("Error encoding file %s: %s\n", outputPath, err)
			continue
		}

		outputFileInfo, err := os.Stat(outputPath)
		if err != nil {
			fmt.Printf("Error getting file information for %s: %s\n", f, err)
			continue
		}

		outputSize := float64(outputFileInfo.Size()) / 1024.0
		inputSize := float64(inputFileInfo.Size()) / 1024.0
		compressionRatio := outputSize / inputSize

		fmt.Printf("Processed file %s: %.2f KB -> %.2f KB (compression ratio: %.2f)\n", f, inputSize, outputSize, compressionRatio)
	}

	// Display information about the processing time and the number of files processed
	elapsedTime := time.Since(startTime)
	fmt.Printf("Processed %d PNG files in %s\n", len(pngFiles), elapsedTime)
}
