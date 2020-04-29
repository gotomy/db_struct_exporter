package main

import (
	"database/sql"
	"db_struct_exporter/exporter"
	"db_struct_exporter/importer"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	xlsx = "xlsx"
	md   = "md"
	pdf  = "pdf"
)

var (
	help   bool
	h      string
	u      string
	p      string
	db     string
	extype string
	dsfile string
)

func init() {
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&h, "h", "localhost:3306", "the mysql `host` contain ip and port, ex: localhost:3306")
	flag.StringVar(&u, "u", "", "the `username` of mysql")
	flag.StringVar(&p, "p", "", "the `password` of mysql")
	flag.StringVar(&db, "db", "", "the need export `database` of mysql")
	flag.StringVar(&extype, "extype", "xlsx", "`export type`,only can: xlsx,md,pdf")
	flag.StringVar(&dsfile, "dsfile", "export.xlsx", "the destination of export `file`")

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `database table struct 
Usage: dbexport []
Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	switch extype {
	case xlsx:
		xlsxExport()
	case md:

	case pdf:

	default:
		fmt.Printf("the export type of %s is not supported\n", extype)
	}

}

func xlsxExport() {
	if xlsxValid() {
		fmt.Printf("xlsx export param invalid, please check\n")
		return
	}
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", u, p, h, db)
	datadb, err := sql.Open("mysql", datasource)
	defer datadb.Close()
	if err != nil {
		fmt.Printf("open database error: %s", err)
		panic("open database error")
	}

	importer := importer.NewMysqlImporter(datadb, db)
	importer.Importer()

	excelExport := exporter.NewExcelExport(dsfile)
	excelExport.Exporter(importer.ExportTables)
}

func xlsxValid() bool {
	return len(h) == 0 || len(u) == 0 || len(dsfile) == 0 || len(db) == 0
}
