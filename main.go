package main

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("simsun", "", "./ttf/simsun.ttf")
	pdf.SetFont("simsun", "", 10)

	// Column widths
	w := []float64{10.0, 25.0, 25.0, 25.0,25.0,50.0}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	left := (210 - wSum) / 2
	// 	Header
	pdf.SetX(left)

	header := []string{"序号", "字段名称","字段类型","是否可空", "默认值","描述"}

	for j, str := range header {
		pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetX(left)
	texts := pdf.SplitText("描述描述描述描述描述描述描述描述描述描述描述描描述描述描述描述描描述描述描述描述描", w[5])

	pdf.CellFormat(w[0], (float64)(6 * len(texts)), "1", "1", 0, "C", false, 0, "")
	pdf.CellFormat(w[1], (float64)(6 * len(texts)), "name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(w[2], (float64)(6 * len(texts)), "varchar(100)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(w[3], (float64)(6 * len(texts)), "false", "1", 0, "C", false, 0, "")
	pdf.CellFormat(w[4], (float64)(6 * len(texts)), "", "1", 0, "C", false, 0, "")


	//pdf.CellFormat(w[5], (float64)(6 * len(texts)), strings.Join(texts,"\n"), "1", 0, "C", false, 0, "")
	pdf.MultiCell(w[5], 6, strings.Join(texts,"\n"), "1","C", false )
	pdf.Ln(-1)

	err := pdf.OutputFileAndClose("h.pdf")
	if err != nil {
		panic(err)
	}


}
