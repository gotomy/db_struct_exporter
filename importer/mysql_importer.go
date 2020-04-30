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

type ShowIndecies struct {
	NonUnique    sql.NullInt32
	IndexName      sql.NullString
	SeqInIndex   sql.NullInt32
	ColumnName   sql.NullString
	IndexType    sql.NullString
	IndexComment sql.NullString
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
	err := export.tables()
	if err != nil {
		panic("get table info error")
	}

	// 获取表的创建语句
	export.createSql()

	// 获取索引信息
	export.indecies()

	// 获取表的字段
	export.columns()
}

/**
 * 获取数据库的所有表名
 */
func (export *MysqlImporter) tables() error {
	selectTableSql := fmt.Sprintf("SELECT table_name,engine,table_collation,table_comment from information_schema.tables where table_schema = '%s'", export.dbName)
	rows, err := export.db.Query(selectTableSql)
	if err != nil {
		fmt.Printf("get tables error: %s", err)
		return err
	}

	for rows.Next() {
		table := &model.Table{}
		rows.Scan(&table.Name, &table.Engine, &table.Charset, &table.Comment)
		export.ExportTables = append(export.ExportTables, table)
	}

	return nil
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
 * 获取表的索引信息
 */
func (export *MysqlImporter) indecies() {
	var showIndexSql string
	for _, item := range export.ExportTables {
		showIndexSql = fmt.Sprintf("SELECT NON_UNIQUE, INDEX_NAME, SEQ_IN_INDEX,COLUMN_NAME,INDEX_TYPE,INDEX_COMMENT FROM INFORMATION_SCHEMA.STATISTICS WHERE TABLE_NAME = '%s'", item.Name)
		rows, err := export.db.Query(showIndexSql)
		if err != nil {
			panic("show index error")
		}
		var index int32 = 1
		var showIndex []*ShowIndecies
		for rows.Next() {
			var indexInfo = &ShowIndecies{}
			error := rows.Scan(&indexInfo.NonUnique, &indexInfo.IndexName, &indexInfo.SeqInIndex, &indexInfo.ColumnName, &indexInfo.IndexType,
				&indexInfo.IndexComment)
			if error != nil {
				fmt.Printf("index info error: %s", error)
			}

			index ++
			showIndex = append(showIndex, indexInfo)
		}

		// 合并索引名称相同的行
		indexMap := make(map[string]*model.Index)
		index = 1
		for _, item := range showIndex {
			val, ok := indexMap[item.IndexName.String]
			if ok {
				val.ContainKey = val.ContainKey + "," + item.ColumnName.String
				indexMap[item.IndexName.String] = val
			} else {
				unique := true
				if item.NonUnique.Int32 == 1 {
					unique = false
				}
				indexMap[item.IndexName.String] = &model.Index{
					Order:      index,
					Name:       item.IndexName.String,
					Unique:     unique,
					IndexType:  item.IndexType.String,
					ContainKey: item.ColumnName.String,
					Comment:    item.IndexComment.String,
				}
			}
			index ++
		}

		for _, v := range indexMap {
			item.Indexes = append(item.Indexes, v)
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
