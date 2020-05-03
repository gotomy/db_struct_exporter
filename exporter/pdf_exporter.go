package exporter

import "db_struct_exporter/model"

type PdfExporter struct {
	fileName string
}

func NewPdfExporter(fileName string) *PdfExporter {
	return &PdfExporter{
		fileName: fileName,
	}
}

func (export *PdfExporter) Exporter(exportTable []*model.Table) {

}
