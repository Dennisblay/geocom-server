package util

import (
	"fmt"
	"strconv"
	"strings"
)

// ExtractRL Helper function to extract RL value from the TBM remark
func ExtractRL(remark string) float64 {
	// Example remark: "TBM 49.873"
	parts := strings.Fields(remark)
	if len(parts) > 1 {
		return ParseFloat(parts[1])
	}
	return 0
}

// ParseFloat Helper function to parse float values from strings
func ParseFloat(value string) float64 {
	if value == "" {
		return 0
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return result
}

// FormatFloat Helper function to format float values as strings
func FormatFloat(value float64) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf("%.3f", value)
}
