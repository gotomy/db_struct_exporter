package exporter

import (
	"bufio"
	"db_struct_exporter/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MarkDownExporter struct {
	fileName string
	dbName   string
}

func NewMarkdownExporter(fileName, dbName string) *MarkDownExporter {
	return &MarkDownExporter{
		fileName: fileName,
		dbName:   dbName,
	}
}

func (export *MarkDownExporter) Exporter(exporterTables []*model.Table) {
	var content []string
	content = append(content, export.renderPageTitle())
	for _, table := range exporterTables {
		content = append(content, export.renderTableSection(table))
	}
	finalString := strings.Join(content, "\n")
	err := export.createFile([]byte(finalString))

	if err != nil {
		fmt.Printf("%v", err)
		panic("export markdown err")
	}

}

func (export *MarkDownExporter) renderPageTitle() string {
	ts := TplTitle
	ts = strings.Replace(ts, "{dbname}", export.dbName, 1)
	return ts
}

func (export *MarkDownExporter) renderTableSection(table *model.Table) string {
	var tableSection []string
	ts := TplTableSection
	ts = strings.Replace(ts, "{tablename}", table.Name, 1)
	ts = strings.Replace(ts, "{tablecomment}", table.Comment, 1)
	ts = strings.Replace(ts, "{charset}", table.Charset, 1)
	ts = strings.Replace(ts, "{engine}", table.Engine, 1)

	tableSection = append(tableSection, ts)

	var tableColumns []string
	for _, column := range table.Columns {
		ts := TplTableColumnParam
		ts = strings.Replace(ts, "{order}", strconv.Itoa(int(column.Order)), 1)
		ts = strings.Replace(ts, "{name}", column.Name, 1)
		ts = strings.Replace(ts, "{type}", column.Type, 1)
		ts = strings.Replace(ts, "{cannull}", column.CanNull, 1)
		ts = strings.Replace(ts, "{default}", column.DefaultValue, 1)
		// 处理comment的内部，将换行符替换成空格，有换行的情况下markdown table显示有异常
		handleComment, _ := mergeLines(column.Comment)
		ts = strings.Replace(ts, "{comment}", handleComment, 1)
		tableColumns = append(tableColumns, ts)
	}
	tableContent := strings.Join(tableColumns, "\n")

	tableSection = append(tableSection, strings.Replace(TplTableColumnTitle, "{params}", tableContent, 1))

	tableSection = append(tableSection, TplTableIndex)

	tableSection = append(tableSection, TplTableSql)
	tableSection = append(tableSection, "```", table.Sql, "```")

	return strings.Join(tableSection, "\n")
}

func (export *MarkDownExporter) createFile(body []byte) error {
	f, err := os.Create(export.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(body)
	if err != nil {
		return err
	}
	return err
}

func mergeLines(comment string) (string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(comment))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return strings.Join(lines, " "), nil
}
