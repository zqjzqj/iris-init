package global

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
)

// GenerateXLSX 生成xlsx文件，参数titles和data分别对应xlsx文件的标题和内容
func GenerateXLSX(titles map[string]string, data []map[string]interface{}) (*excelize.File, error) {
	// 创建xlsx文件
	f := excelize.NewFile()

	// 设置表头样式
	titleStyle, err := f.NewStyle(`{"font":{"bold":true,"size":14},"fill":{"type":"pattern","color":["#9FB6CD"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"},"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}],"alignment":{"wrap_text":true}}`)
	if err != nil {
		return nil, err
	}

	// 设置单元格样式
	cellStyle, err := f.NewStyle(`{"font":{"size":12},"alignment":{"horizontal":"center","vertical":"center"},"border":[{"type":"left","color":"#000000","style":1},{"type":"top","color":"#000000","style":1},{"type":"right","color":"#000000","style":1},{"type":"bottom","color":"#000000","style":1}],"alignment":{"wrap_text":true}}`)
	if err != nil {
		return nil, err
	}

	// 添加表头
	for k, v := range titles {
		cell := fmt.Sprintf("%s1", k)
		f.SetCellValue("Sheet1", cell, v)
		f.SetColWidth("Sheet1", k, k, 25)
		f.SetCellStyle("Sheet1", cell, cell, titleStyle)
	}

	// 添加表格内容
	for i, row := range data {
		for k, v := range row {
			cell := fmt.Sprintf("%s%d", k, i+2)
			f.SetCellValue("Sheet1", cell, v)
			f.SetCellStyle("Sheet1", cell, cell, cellStyle)
		}
	}

	// 设置列宽和行高
	for i := 1; i <= len(titles); i++ {
		col := excelize.ToAlphaString(i)
		f.SetColWidth("Sheet1", col, col, 30)
	}
	for i := 1; i <= len(data)+1; i++ {
		f.SetRowHeight("Sheet1", i, 20)
	}

	return f, nil
}

func ParseXLSXByReader(reader io.Reader) ([]map[string]string, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	// 获取第一个Sheet
	sheetName := f.GetSheetName(1)
	rows := f.GetRows(sheetName)

	// 获取表头
	var headers []string
	if len(rows) > 0 {
		for _, cell := range rows[0] {
			headers = append(headers, cell)
		}
	}

	// 获取数据
	var data []map[string]string
	for _, row := range rows[1:] {
		rowData := make(map[string]string)
		for i, cell := range row {
			rowData[headers[i]] = cell
		}
		data = append(data, rowData)
	}

	return data, nil
}

func ParseXLSX(path string) ([]map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return ParseXLSXByReader(f)
}
