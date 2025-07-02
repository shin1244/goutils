package goutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestOpenTxt(t *testing.T) {
	// Create a temporary text file for testing
	content := "line1\nline2\nline3"
	tmpFile, err := os.CreateTemp("", "test_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	assert.NoError(t, tmpFile.Close())

	// Call OpenTxt and assert the result
	lines, err := OpenTxt(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, []string{"line1", "line2", "line3"}, lines)

	// Test with a non-existent file
	_, err = OpenTxt("non_existent_file.txt")
	assert.Error(t, err)
}

func TestOpenXlsx(t *testing.T) {
	// Create a temporary XLSX file for testing
	f := excelize.NewFile()
	// Create a sheet and write some data
	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", "World")
	f.SetCellValue("Sheet1", "A2", 123)
	f.SetCellValue("Sheet1", "B2", true)

	// Create another sheet
	f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A1", "Test")

	tmpFile, err := os.CreateTemp("", "test_*.xlsx")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file

	assert.NoError(t, f.SaveAs(tmpFile.Name()))
	assert.NoError(t, tmpFile.Close())

	// Call OpenXlsx and assert the result
	allSheetsData, err := OpenXlsx(tmpFile.Name())
	assert.NoError(t, err)

	expectedSheet1 := [][]string{
		{"Hello", "World"},
		{"123", "TRUE"},
	}
	expectedSheet2 := [][]string{
		{"Test"},
	}

	assert.Len(t, allSheetsData, 2) // Expecting 2 sheets
	assert.Equal(t, expectedSheet1, allSheetsData[0])
	assert.Equal(t, expectedSheet2, allSheetsData[1])

	// Test with a non-existent file
	_, err = OpenXlsx("non_existent_file.xlsx")
	assert.Error(t, err)
}
