package exporter

import (
	"db_struct_exporter/model"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strconv"
	"strings"
)

type PdfExporter struct {
	fileName string
	dbName   string
}

func NewPdfExporter(fileName, dbName string) *PdfExporter {
	return &PdfExporter{
		fileName: fileName,
		dbName:   dbName,
	}
}

func (export *PdfExporter) Exporter(exportTable []*model.Table) {
	// create a pdf
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	// 设置宋体以显示中文
	pdf.AddUTF8Font("simsun", "", "./ttf/simsun.ttf")
	pdf.SetFont("simsun", "", 10)

	export.buildPageTitle(pdf, export.dbName)

	order := 1
	for _, table := range exportTable {
		export.buildTableTitle(pdf, table, order)
		export.buildTable(pdf, table)
		pdf.Ln(2)
		export.buildIndex(pdf, table.Indexes)
		order ++
	}

	err := pdf.OutputFileAndClose(export.fileName)
	if err != nil {
		panic(err)
	}
}

/**
 * 设置标题.
 */
func (export *PdfExporter) buildPageTitle(pdf *gofpdf.Fpdf, dbname string) {
	title := fmt.Sprintf("%s数据库表结构", dbname)
	left := (210 - pdf.GetStringWidth(title)) / 2
	pdf.SetX(left)
	pdf.Write(24, title)
	pdf.Ln(20)
}

/**
 * 设置表结构标题.
 */
func (export *PdfExporter) buildTableTitle(pdf *gofpdf.Fpdf, table *model.Table, order int) {
	tableTitle := fmt.Sprintf("%d.表名：%s，注释：%s，字符集：%s，引擎：%s", order, table.Name, table.Comment, table.Charset, table.Engine)
	pdf.SetX((210 - 160) / 2)
	pdf.Write(18, tableTitle)
	pdf.Ln(15)
}

/**
 * 设置表格字段.
 */
func (export *PdfExporter) buildTable(pdf *gofpdf.Fpdf, table *model.Table) {
	// Column widths
	w := []float64{10.0, 35.0, 25.0, 20.0, 20.0, 50.0}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	left := (210 - wSum) / 2

	// 	Header
	pdf.SetX(left)
	header := []string{"序号", "字段名称", "字段类型", "是否可空", "默认值", "描述"}

	for j, str := range header {
		pdf.CellFormat(w[j], 7, str, "1", 0, "L", false, 0, "")
	}
	pdf.Ln(-1)

	for _, column := range table.Columns {
		pdf.SetX(left)
		texts := pdf.SplitText(column.Comment, w[5])
		height := float64(6 * len(texts))
		if len(texts) == 0 {
			height = 6
		}

		pdf.CellFormat(w[0], height, strconv.Itoa(int(column.Order)), "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[1], height, column.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[2], height, column.Type, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[3], height, column.CanNull, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[4], height, column.DefaultValue, "1", 0, "L", false, 0, "")
		//pdf.CellFormat(w[5], (float64)(6 * len(texts)), strings.Join(texts, "\n"), "1", 0, "L", false, 0, "")
		pdf.MultiCell(w[5], 6, strings.Join(texts, "\n"), "1", "L", false)

	}

}

/**
 * 设置表格索引.
 */
func (export *PdfExporter) buildIndex(pdf *gofpdf.Fpdf, index []*model.Index) {
	w := []float64{10.0, 35.0, 50.0, 20.0, 20.0, 25.0}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	left := (210 - wSum) / 2

	// 	Header
	pdf.SetX(left)
	header := []string{"序号", "索引名称", "包含字段", "索引类型", "是否惟一", "描述"}

	for j, str := range header {
		pdf.CellFormat(w[j], 7, str, "1", 0, "L", false, 0, "")
	}
	pdf.Ln(-1)

	height := 6.0

	for _, item := range index {
		pdf.SetX(left)
		pdf.CellFormat(w[0], height, strconv.Itoa(int(item.Order)), "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[1], height, item.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[2], height, item.ContainKey, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[3], height, item.IndexType, "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[4], height, strconv.FormatBool(item.Unique), "1", 0, "L", false, 0, "")
		pdf.CellFormat(w[5], height, item.Comment, "1", 0, "L", false, 0, "")
		pdf.Ln(-1)
	}
}
