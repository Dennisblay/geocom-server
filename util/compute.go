package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ComputeHPC(input, output string) error {
	// Open the input CSV file
	file, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("error opening input CSV file: %w", err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV data: %w", err)
	}

	// Prepare the output CSV file
	outputFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("error creating output CSV file: %w", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the headers to the output CSV file
	headers := []string{"BS", "IS", "FS", "HPC", "Initial RL", "Adj", "Adj RL", "Remarks"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error writing headers to output CSV file: %w", err)
	}

	// Variables for calculations
	var prevRL, prevHPC, hpc, baseRL, lastTBM, misclosure, errorPerReading float64
	var sumBS, sumIS, sumFS float64
	cpCount := 0
	bsCount := 0 // Initialize bsCount to 0 and count all BS readings

	// First pass: Calculate the Initial RL and identify misclosure
	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		bs := ParseFloat(record[0])
		is := ParseFloat(record[1])
		fs := ParseFloat(record[2])
		remarks := record[3]

		// Accumulate sums for the checks
		sumBS += bs
		sumIS += is
		sumFS += fs

		// Initial RL set by the benchmark (TBM)
		if strings.Contains(remarks, "TBM") && i == 1 {
			baseRL = ExtractRL(remarks) // Set the base RL from TBM
			hpc = baseRL + bs           // HPC for the first row
			prevHPC = hpc               // Store HPC for subsequent calculations
			prevRL = baseRL             // Store RL for adjustment calculations
			bsCount++                   // Include the first BS reading
			continue
		} else if strings.Contains(remarks, "TBM") && i > 1 {
			// Record the final TBM value for misclosure calculation
			lastTBM = ExtractRL(remarks)
		}

		// RL Calculation
		var initialRL float64
		if fs != 0 {
			initialRL = prevHPC - fs
		} else if is != 0 {
			initialRL = prevHPC - is
		} else {
			initialRL = prevRL
		}

		// Update HPC at a Change Point (CP) with a BS
		if strings.Contains(remarks, "CP") && bs != 0 {
			hpc = initialRL + bs
			prevHPC = hpc
			cpCount++
		}

		prevRL = initialRL
		if bs != 0 {
			bsCount++
		}
	}

	// Calculate misclosure and error per reading
	misclosure = prevRL - lastTBM
	errorPerReading = -1 * (misclosure / float64(bsCount)) // Negate the error per reading to apply correctly

	// Second pass: Apply adjustments and calculate Adj RL
	prevRL = baseRL
	hpc = baseRL
	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}

		bs := ParseFloat(record[0])
		is := ParseFloat(record[1])
		fs := ParseFloat(record[2])
		remarks := record[3]

		// Initial RL set by the benchmark (TBM)
		if strings.Contains(remarks, "TBM") && i == 1 {
			hpc = baseRL + bs // HPC for the first row
			prevHPC = hpc     // Store HPC for subsequent calculations
			prevRL = baseRL   // Store RL for adjustment calculations
			adjRL := prevRL   // Initial RL is same as adjusted RL
			adj := "----"     // No adjustment for the first row

			// Write the first row
			writer.Write([]string{
				FormatFloat(bs),
				FormatFloat(is),
				FormatFloat(fs),
				FormatFloat(hpc),
				FormatFloat(prevRL),
				adj,
				FormatFloat(adjRL),
				remarks,
			})
			continue
		}

		// RL Calculation
		var initialRL float64
		if fs != 0 {
			initialRL = prevHPC - fs
		} else if is != 0 {
			initialRL = prevHPC - is
		} else {
			initialRL = prevRL
		}

		// Apply adjustment starting from the second reading
		adjustment := errorPerReading * float64(cpCount-1)
		adjustedRL := initialRL + adjustment
		adj := FormatFloat(adjustment)

		// Update HPC at a Change Point (CP) with a BS
		if strings.Contains(remarks, "CP") && bs != 0 {
			hpc = initialRL + bs
			prevHPC = hpc
			cpCount++
		}

		// Write the row to CSV
		writer.Write([]string{
			FormatFloat(bs),
			FormatFloat(is),
			FormatFloat(fs),
			FormatFloat(hpc),
			FormatFloat(initialRL),
			adj,
			FormatFloat(adjustedRL),
			remarks,
		})

		// Update previous values for the next iteration
		prevRL = initialRL
	}

	// Write the checks at the bottom of the CSV
	sumRL := sumBS - sumFS
	finalMisclosure := FormatFloat(sumBS - sumFS)
	IRL := FormatFloat(prevRL)
	FRL := FormatFloat(lastTBM)

	writer.Write([]string{"", "", "", "", ""})
	writer.Write([]string{"Checks", "ΣBS - ΣFS should be equal to Initial IRL - Final IRL", "", ""})
	writer.Write([]string{"", "", "", "", ""})
	writer.Write([]string{FormatFloat(sumBS), FormatFloat(sumIS), FormatFloat(sumFS), "", FormatFloat(sumRL), "", "", ""})
	writer.Write([]string{FormatFloat(sumBS), "ΣBS", "", "", "", IRL, "Initial IRL", "", ""})
	writer.Write([]string{FormatFloat(sumFS), "ΣFS", "", "", "", FRL, "Final IRL", "", ""})
	writer.Write([]string{finalMisclosure, "", "", "", finalMisclosure, "", "", ""})
	return nil
}
