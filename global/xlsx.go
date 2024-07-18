package global

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"iris-init/logs"
	"os"
	"strings"
)

// GenerateXLSX 生成xlsx文件，参数titles和data分别对应xlsx文件的标题和内容
func GenerateXLSX(titles map[string]string, data []map[string]interface{}, autoImg bool) (*excelize.File, error) {
	// 创建xlsx文件
	f := excelize.NewFile()

	header_style := &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#9FB6CD"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	}
	cell_style := &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{},
		Font: &excelize.Font{
			Bold: false,
			Size: 13,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	}
	// 设置表头样式
	titleStyle, err := f.NewStyle(header_style)
	if err != nil {
		return nil, err
	}

	// 设置单元格样式
	cellStyle, err := f.NewStyle(cell_style)
	if err != nil {
		return nil, err
	}

	// 添加表头
	for k, v := range titles {
		cell := fmt.Sprintf("%s1", k)
		_ = f.SetCellValue("Sheet1", cell, v)
		_ = f.SetColWidth("Sheet1", k, k, 25)
		_ = f.SetCellStyle("Sheet1", cell, cell, titleStyle)
	}

	// 添加表格内容
	for i, row := range data {
		for k, v := range row {
			cell := fmt.Sprintf("%s%d", k, i+2)
			v_str, ok := v.(string)
			if autoImg && ok && strings.HasPrefix(v_str, "http") { //这里直接当成网络图片处理
				_bytes, _err := GetHttpBodyBytes(v_str)
				if _err == nil {
					err = f.AddPictureFromBytes("Sheet1", cell, &excelize.Picture{
						Extension: ".jpg",
						File:      _bytes,
						Format: &excelize.GraphicOptions{
							AutoFit: true,
						},
					})
				} else {
					_ = f.SetCellValue("Sheet1", cell, v)
				}
			} else {
				_ = f.SetCellValue("Sheet1", cell, v)
			}
			_ = f.SetCellStyle("Sheet1", cell, cell, cellStyle)
		}
	}

	// 设置列宽和行高
	for i := 1; i <= len(titles); i++ {
		col := fmt.Sprintf("%d", i)
		_ = f.SetColWidth("Sheet1", col, col, 30)
	}
	for i := 1; i <= len(data)+1; i++ {
		_ = f.SetRowHeight("Sheet1", i, 20)
	}

	return f, nil
}

func ParseXLSXByReader(reader io.Reader) ([]map[string]string, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	// 获取第一个Sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.Rows(sheetName)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	// 获取表头
	var headers []string
	if rows.Next() {
		headers, err = rows.Columns()
		if err != nil {
			return nil, err
		}
	}

	var data []map[string]string
	rowIndex := 1
	for rows.Next() {
		rowIndex++
		rowData := make(map[string]string)

		for colIndex := range headers {
			cellName, err := excelize.CoordinatesToCellName(colIndex+1, rowIndex)
			if err != nil {
				return nil, err
			}
			// 检查单元格是否包含图片
			pictures, err := f.GetPictures(sheetName, cellName)
			if err == nil && pictures != nil {
				//这里根据需要修改
				filename := fmt.Sprintf("./static/uploads/%s", uuid.New().String()+"."+pictures[0].Extension)
				err = UploadLocalByBytes(pictures[0].File, uuid.New().String()+"."+pictures[0].Extension)
				if err != nil {
					logs.PrintErr("upload xlsx img err ", err)
				}
				rowData[headers[colIndex]] = strings.TrimLeft(filename, ".")
			} else {
				rowData[headers[colIndex]], _ = f.GetCellValue(sheetName, cellName)
			}
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
