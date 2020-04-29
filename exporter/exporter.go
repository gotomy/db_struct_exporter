package exporter

import "db_struct_exporter/model"

type Exporter interface {
	Exporter(exportTable []*model.Table) error
}
