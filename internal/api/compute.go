package api

import (
	"fmt"
	"github.com/Dennisblay/ordering-app-server/util"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *Server) computeController(ctx *gin.Context) {
	// Parse the incoming multipart form to get the file
    fmt.Println("Starting Computation")

    file, header, err := ctx.Request.FormFile("file")
    if err != nil {
        fmt.Println("Error retrieving the file:", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
        return
    }
    defer file.Close()

    // Log the file name for debugging purposes
    fmt.Println("File Name:", header.Filename)

    // Continue with your computation logic...

	// Create a temporary file to store the uploaded CSV
	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "uploaded-*.csv")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temp file: " + err.Error()})
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up the file after processing
	defer tempFile.Close()

	// Copy the uploaded file data to the temp file
	if _, err := io.Copy(tempFile, file); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy uploaded file to temp file: " + err.Error()})
		return
	}

	// Define the output file path
	outputFilePath := filepath.Join(tempDir, "output.csv")

	// Call the ComputeHPC function with the paths to the temporary files
	if err := util.ComputeHPC(tempFile.Name(), outputFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compute HPC: " + err.Error()})
		return
	}

	// Open the generated output file
	outputFile, err := os.Open(outputFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open output file: " + err.Error()})
		return
	}
	defer outputFile.Close()

	// Read the output file content
	outputData, err := io.ReadAll(outputFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read output file: " + err.Error()})
		return
	}

	// Send the processed CSV back to the frontend
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=output.csv")
    fmt.Println("computation completed")

	ctx.Data(http.StatusOK, "text/csv", outputData)
}
