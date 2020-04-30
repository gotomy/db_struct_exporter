package exporter

const TplTitle = `# {dbname}数据库表结构`

const TplTableSection = `## 表名：{tablename}，描述：{tablecomment}，字符集：{charset}，引擎：{engine}`

const TplTableTitle = `### 表结构`

const TplTableColumnTitle = `
|序号		|字段名称	|字段类型	|是否可空	|默认值	   |描述	        |
|-----------|-----------|-----------|-----------|----------|------------|
{params}`

const TplTableColumnParam = `|{order}   |{name}	    |{type}		|{cannull}		|{default}     |{comment}	   |`

const TplTableIndex = `### 索引结构`


const TplTableSql = `### 建表语句`
