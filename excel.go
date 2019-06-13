package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func ReadExcel() {
	excelFileName := "./20190117.xlsx"
	// excelFileName = "./test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {

	}
	sheet := xlFile.Sheets[0]
	// 主表内容
	var contents [][]string
	// 采样内容（最后一列）
	samples := make([]string, 0, 20)
	// 表头（指标）行
	keyMap := make(map[int]string)

	colNum := 0

	for i := 0; i < len(sheet.Rows); i++ {
		row := sheet.Rows[i]
		rows := make([]string, 0, len(row.Cells))
		colNum = len(row.Cells)

		for j := 0; j < len(row.Cells); j++ {
			cell := row.Cells[j]
			text := cell.String()
			rows = append(rows, text)
			if i == 2 && j > 0 && j < colNum-1 {
				keyMap[j] = text
			}
			if j == colNum-1 && i > 2 {
				samples = append(samples, text)
			}
		}
		contents = append(contents, rows)
	}

	for i := 0; i < len(samples); i++ {
		if samples[i] == "0" {
			continue
		}
		for j := 1; j < colNum-1; j++ {
			if contents[i+3][j] != "+" {
				delete(keyMap, j)
			}
		}
	}

	fmt.Println("result:", keyMap)
	// err = xlFile.Save("./test.xlsx")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	// xlFile, err = xlsx.OpenFile("./test.xlsx")

}

func SetCellStyle(cell *xlsx.Cell, color string) {
	style := cell.GetStyle()
	style.Font.Color = color //"FFFF00FF"
	style.Fill.BgColor = "00FFFF00"
	style.Fill.FgColor = "FF00FF00"
	// type 类型参考文档 https://documentation.devexpress.com/OfficeFileAPI/DevExpress.Spreadsheet.PatternType.enum
	style.Fill.PatternType = "Gray125"
	cell.SetStyle(style)
}
