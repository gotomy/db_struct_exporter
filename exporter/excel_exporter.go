package exporter

import (
	"db_struct_exporter/model"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

type ExcelExport struct {
	fileName string
	sheet    string
}

func NewExcelExport(fileName string) *ExcelExport {
	return &ExcelExport{
		fileName: fileName,
	}
}

func (export *ExcelExport) Exporter(exporterTables []*model.Table) {
	const sheetName = "表结构"
	excelFile := excelize.NewFile()
	sheetIndex := excelFile.NewSheet(sheetName)
	excelFile.SetActiveSheet(sheetIndex)

	// 设置列宽
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(1), ColumnIndexToLetter(1), 15)
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(2), ColumnIndexToLetter(2), 35)
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(3), ColumnIndexToLetter(3), 20)
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(4), ColumnIndexToLetter(4), 20)
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(5), ColumnIndexToLetter(5), 20)
	excelFile.SetColWidth(sheetName, ColumnIndexToLetter(6), ColumnIndexToLetter(6), 60)

	rowStyle, _ := excelFile.NewStyle(`{"font":{"size":12},"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)

	rowIndex := 1
	for _, table := range exporterTables {
		export.buildTableTitle(excelFile, sheetName, rowIndex, table.Name, table.Comment)
		rowIndex++
		export.buildTableHeader(excelFile, sheetName, rowIndex)
		rowIndex++
		for _, column := range table.Columns {
			// 设置列值
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(1)+strconv.Itoa(rowIndex), column.Order)
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(2)+strconv.Itoa(rowIndex), column.Name)
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(3)+strconv.Itoa(rowIndex), column.Type)
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(4)+strconv.Itoa(rowIndex), column.CanNull)
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(5)+strconv.Itoa(rowIndex), column.DefaultValue)
			excelFile.SetCellValue(sheetName, ColumnIndexToLetter(6)+strconv.Itoa(rowIndex), column.Comment)
			excelFile.SetCellStyle(sheetName, ColumnIndexToLetter(1)+strconv.Itoa(rowIndex), ColumnIndexToLetter(6)+strconv.Itoa(rowIndex), rowStyle)
			rowIndex++
		}
		rowIndex++;
	}

	if err := excelFile.SaveAs(export.fileName); err != nil {
		panic(err)
	}
}

/**
  设置标题
 */
func (export *ExcelExport) buildTableTitle(file *excelize.File, sheet string, rowIndex int, tableName string, tableComment string) {
	file.MergeCell(sheet, ColumnIndexToLetter(1)+strconv.Itoa(rowIndex), ColumnIndexToLetter(6)+strconv.Itoa(rowIndex))
	titlePosition := ColumnIndexToLetter(1) + strconv.Itoa(rowIndex)
	file.SetCellValue(sheet, titlePosition, fmt.Sprintf("表名：%s, 注释：%s", tableName, tableComment))

	titleStyle, _ := file.NewStyle(`{"font":{"size":16},"fill":{"type":"pattern","color":["#9ACD32"],"pattern":1},"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)
	file.SetCellStyle(sheet, ColumnIndexToLetter(1)+strconv.Itoa(rowIndex), ColumnIndexToLetter(6)+strconv.Itoa(rowIndex), titleStyle)
}

/**
  设置表头
 */
func (export *ExcelExport) buildTableHeader(file *excelize.File, sheet string, rowIndex int) {
	titles := []string{
		"序号", "字段名称", "字段类型", "是否可空", "默认值", "描述",
	}

	for columnNum, v := range titles {
		sheetPosition := ColumnIndexToLetter(columnNum+1) + strconv.Itoa(rowIndex)
		file.SetCellValue(sheet, sheetPosition, v)
	}

	headerStyle, _ := file.NewStyle(`{"font":{"size":14},"fill":{"type":"pattern","color":["#9ACD32"],"pattern":1},"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)
	file.SetCellStyle(sheet, ColumnIndexToLetter(1)+strconv.Itoa(rowIndex), ColumnIndexToLetter(6)+strconv.Itoa(rowIndex), headerStyle)
}

/**
  * 将列序号转换为excel的字母编号.
 */
func ColumnIndexToLetter(columnIndex int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if columnIndex > 26 { //数据大于26需要进行拆分
		for {
			k = columnIndex % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			columnIndex = (columnIndex - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if columnIndex <= 26 { //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, columnIndex)
				break
			}
		}
	} else {
		return Slice[columnIndex]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
