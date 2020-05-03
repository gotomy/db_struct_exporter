package exporter

import (
	"db_struct_exporter/model"
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
)

type OnlineExporter struct {
	dbName string
	port   int
}

func NewOnlineExporter(dbName string, port int) *OnlineExporter {
	return &OnlineExporter{
		dbName: dbName,
		port:   port,
	}
}

func (exporter *OnlineExporter) Exporter(exportTable []*model.Table) {
	// start a html server
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		mdExporter := NewMarkdownExporter("", exporter.dbName)
		body := mdExporter.ExporterJson(exportTable)

		output := blackfriday.Run(body, blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(
			blackfriday.HTMLRendererParameters{
				Flags: blackfriday.TOC,
			},
		)))

		tpl := template.New("")
		tpl = tpl.Funcs(template.FuncMap{"unescaped": unescaped})
		tpl, err := template.ParseFiles("./tpl/index.html")
		err = tpl.Execute(w, string(output))
		if err != nil {
			panic(err)
		}
	})

	fmt.Printf("server is listening at %d", exporter.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", exporter.port), nil)

	if err != nil {
		panic(err)
	}
}

func unescaped(x string) interface{} { return template.HTML(x) }
