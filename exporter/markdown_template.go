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

const TplTableIndexTitle = `
|序号		|索引名	    |包含字段    |索引类型	|是否唯一	|描述	   |
|-----------|-----------|-----------|-----------|-----------|----------|
{params}`

const TplTableIndexParam = `|{order}   |{name}	    |{keys}    |{type}		|{unique}	  |{comment}	   |`

const TplTableSql = `### 建表语句`
