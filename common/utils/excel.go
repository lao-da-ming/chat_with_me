package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// ExcelExportConfig 导出配置
type ExcelExportConfig struct {
	FileName    string      // 文件名
	SheetName   string      // 工作表名
	Headers     []string    // 表头
	Data        [][]string  // 数据内容
	ColWith     float64     // 列宽（0表示自动列宽）
	MergeCells  []MergeCell //合并单元格设置
	HeaderStyle *excelize.Style
	DataStyle   *excelize.Style
}
type MergeCell struct {
	//如想要合并A2~A7的单元格 BottomRightCell = A2  TopLeftCell = A7
	TopLeftCell     string //合并截止单元格
	BottomRightCell string //合并到目标的单元格(只保留这个单元格数据)
}

// ExportExcel 导出Excel主方法
func ExportExcel(cfg *ExcelExportConfig) error {
	f := excelize.NewFile()
	defer f.Close()

	// 创建工作表
	index, err := f.NewSheet(cfg.SheetName)
	if err != nil {
		return fmt.Errorf("创建Sheet失败: %v", err)
	}

	// 设置默认样式
	if cfg.HeaderStyle == nil || cfg.DataStyle == nil {
		headerStyle, dataStyle := defaultStyles()
		if cfg.HeaderStyle == nil {
			cfg.HeaderStyle = headerStyle
		}
		if cfg.DataStyle == nil {
			cfg.DataStyle = dataStyle
		}
	}

	// 注册样式
	headerStyleID, _ := f.NewStyle(cfg.HeaderStyle)
	dataStyleID, _ := f.NewStyle(cfg.DataStyle)

	// ==== 性能优化点1：批量写入表头 ====
	headerCells := make([]interface{}, len(cfg.Headers))
	for i, header := range cfg.Headers {
		headerCells[i] = header
	}
	// 批量写入表头
	startHeaderCell, _ := excelize.CoordinatesToCellName(1, 1)
	if err := f.SetSheetRow(cfg.SheetName, startHeaderCell, &headerCells); err != nil {
		return err
	}
	// 设置整行表头样式
	endHeaderCell, _ := excelize.CoordinatesToCellName(len(cfg.Headers), 1)
	if err := f.SetCellStyle(cfg.SheetName, startHeaderCell, endHeaderCell, headerStyleID); err != nil {
		return err
	}

	// ==== 性能优化点2：批量写入数据 ====
	for rowIdx, rowData := range cfg.Data {
		startDataCell, _ := excelize.CoordinatesToCellName(1, rowIdx+2)
		row := make([]interface{}, len(rowData))
		for i, v := range rowData {
			row[i] = v
		}
		if err := f.SetSheetRow(cfg.SheetName, startDataCell, &row); err != nil {
			return err
		}
	}
	// 处理合并单元格
	for _, item := range cfg.MergeCells {
		if err = f.MergeCell(cfg.SheetName, item.TopLeftCell, item.BottomRightCell); err != nil {
			return err
		}
	}
	// 设置整个数据区域样式（性能关键）
	if len(cfg.Data) > 0 {
		startDataCell, _ := excelize.CoordinatesToCellName(1, 2)
		endDataCell, _ := excelize.CoordinatesToCellName(len(cfg.Headers), len(cfg.Data)+1)
		if err := f.SetCellStyle(cfg.SheetName, startDataCell, endDataCell, dataStyleID); err != nil {
			return err
		}
	}

	// ==== 列宽设置优化 ====
	if cfg.ColWith > 0 {
		if err := setUniformColumnWidth(f, cfg.SheetName, len(cfg.Headers), cfg.ColWith); err != nil {
			return fmt.Errorf("调整列宽失败: %v", err)
		}
	} else {
		if err := setAutoColumnWidth(f, cfg.SheetName, cfg.Headers, cfg.Data); err != nil {
			return fmt.Errorf("自动列宽设置失败: %v", err)
		}
	}

	// 设置默认工作表
	f.SetActiveSheet(index)
	// 保存文件
	if err := f.SaveAs(cfg.FileName); err != nil {
		return fmt.Errorf("文件保存失败: %v", err)
	}
	return nil
}

// 设置统一列宽
func setUniformColumnWidth(f *excelize.File, sheetName string, colCount int, width float64) error {
	if width < 8 || width > 50 {
		width = 8
	}
	for col := 1; col <= colCount; col++ {
		colName, _ := excelize.ColumnNumberToName(col)
		if err := f.SetColWidth(sheetName, colName, colName, width); err != nil {
			return err
		}
	}
	return nil
}

// 设置自适应列宽
func setAutoColumnWidth(f *excelize.File, sheetName string, headers []string, data [][]string) error {
	colMaxWidth := make([]float64, len(headers))

	// 计算表头宽度
	for colIdx, header := range headers {
		width := float64(len(header))*1.2 + 2 // 中文字符补偿
		if width > colMaxWidth[colIdx] {
			colMaxWidth[colIdx] = width
		}
	}

	// 计算数据内容最大宽度
	for _, row := range data {
		for colIdx, val := range row {
			width := float64(len(val))*1.1 + 1
			if width > colMaxWidth[colIdx] {
				colMaxWidth[colIdx] = width
			}
		}
	}

	// 应用列宽（限制在8-50之间）
	for colIdx, width := range colMaxWidth {
		if width < 8 {
			width = 8
		} else if width > 50 {
			width = 50
		}
		colName, _ := excelize.ColumnNumberToName(colIdx + 1)
		if err := f.SetColWidth(sheetName, colName, colName, width); err != nil {
			return err
		}
	}
	return nil
}

// DefaultStyles 默认样式配置
func defaultStyles() (*excelize.Style, *excelize.Style) {
	// 标题样式 (居中 + 加粗)
	headerStyle := &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold: true,
		},
	}
	// 数据样式 (居中)
	dataStyle := &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	return headerStyle, dataStyle
}

// 生成A-Z数组
func GenerateAtoZ() []string {
	letters := make([]string, 26)
	for i := range letters {
		letters[i] = string('A' + i)
	}
	return letters
}
