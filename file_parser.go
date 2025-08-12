package goutils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// This file contains functions for parsing data from various file formats (e.g., txt, xlsx, xml).

// OpenXlsx는 지정된 경로의 XLSX 파일을 열고 모든 시트의 데이터를 읽어옵니다.
// 반환 값은 [시트 인덱스][행 인덱스][열 인덱스] 형태의 문자열 슬라이스입니다.
func OpenXlsx(path string) ([][][]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("XLSX 파일 열기 실패: %w", err)
	}
	defer f.Close()

	var allSheetsData [][][]string

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, fmt.Errorf("XLSX에 시트가 존재하지 않음")
	}

	for _, sheetName := range sheetList {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("시트에서 row를 읽을 수 없음 '%s': %w", sheetName, err)
		}
		allSheetsData = append(allSheetsData, rows)
	}

	return allSheetsData, nil
}

// OpenTxt는 지정된 경로의 텍스트 파일을 한 줄씩 읽어 문자열 슬라이스로 반환합니다.
func OpenTxt(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("텍스트 파일 열기 실패: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("텍스트 읽기 실패: %w", err)
	}

	return lines, nil
}

// MakeJsonFile는 주어진 데이터를 JSON 형식으로 파일에 저장합니다.
func MakeJsonFile(data any, filename string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON 마샬 실패:", err)
		return
	}

	jsonFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("파일 생성 실패:", err)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
	}
}

// FindXlsx는 현재 작업 디렉토리에서 첫 번째 XLSX 파일을 찾습니다.
// 찾은 파일의 이름을 반환합니다.
func FindXlsx() string {
	cwd, _ := os.Getwd()
	entries, _ := os.ReadDir(cwd)

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".xlsx") {
			return entry.Name()
		}
	}
	fmt.Println("XLSX 파일을 찾을 수 없습니다. 현재 디렉토리에 XLSX 파일이 있는지 확인하세요.")
	fmt.Scanln()
	panic("XLSX 파일을 찾을 수 없습니다.")
}
