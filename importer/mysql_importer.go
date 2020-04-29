package importer

import (
	"database/sql"
	"db_struct_exporter/model"
	"fmt"
)

type ShowColumns struct {
	Field      sql.NullString
	Type       sql.NullString
	Collation  sql.NullString
	Null       sql.NullString
	Key        sql.NullString
	Default    sql.NullString
	Extra      sql.NullString
	Privileges sql.NullString
	Comment    sql.NullString
}

type MysqlImporter struct {
	db           *sql.DB
	dbName       string
	ExportTables []*model.Table
}

func NewMysqlImporter(db *sql.DB, dbName string) *MysqlImporter {
	return &MysqlImporter{
		db:     db,
		dbName: dbName,
	}
}

func (export *MysqlImporter) Importer() {
	tables, err := export.tables()
	if err != nil {

	}
	for _, table := range tables {
		export.ExportTables = append(export.ExportTables, &model.Table{Name: table})
	}

	// 获取表的创建语句
	export.createSql()

	// 获取表的字段
	export.columns()
}

/**
 * 获取数据库的所有表名
 */
func (export *MysqlImporter) tables() ([]string, error) {
	showTableSql := fmt.Sprintf("SHOW TABLES FROM %s", export.dbName)
	rows, err := export.db.Query(showTableSql)
	if err != nil {
		fmt.Printf("get tables error: %s", err)
		return nil, err
	}
	var tableName string
	var tables []string
	for rows.Next() {
		rows.Scan(&tableName)
		tables = append(tables, tableName)
	}

	return tables, nil
}

/**
 * 获取表的字段
 */
func (export *MysqlImporter) columns() {
	var showColumnSql string
	for _, item := range export.ExportTables {
		showColumnSql = fmt.Sprintf("SHOW FULL COLUMNS FROM %s", item.Name)
		rows, err := export.db.Query(showColumnSql)
		if err != nil {
			panic("show columns error")
		}
		var index int32 = 1
		for rows.Next() {
			var cols = &ShowColumns{}
			error := rows.Scan(&cols.Field, &cols.Type, &cols.Collation, &cols.Null, &cols.Key, &cols.Default, &cols.Extra, &cols.Privileges, &cols.Comment)
			if error != nil {
				fmt.Printf("scan error: %s", error)
			}

			column := &model.Column{
				Order:        index,
				Name:         cols.Field.String,
				Type:         cols.Type.String,
				CanNull:      cols.Null.String,
				DefaultValue: cols.Default.String,
				Comment:      cols.Comment.String,
			}
			index++

			item.Columns = append(item.Columns, column)
		}
	}
}

/**
 * 获取创建表的sql语句
 */
func (export *MysqlImporter) createSql() {
	var showCreateSql string
	for _, item := range export.ExportTables {
		showCreateSql = fmt.Sprintf("SHOW CREATE TABLE %s", item.Name)
		row := export.db.QueryRow(showCreateSql)
		var name string
		row.Scan(&name, &item.Sql)
	}
}
