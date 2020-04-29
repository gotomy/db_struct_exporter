# 数据库表结构导出工具

这是一个命令行工具，提供快速导出数据库结构到各类文件，或直接提供web服务，在线显示数据库表结构。

- 使用

可通过`go run cmd/main.go -help` 获取使用帮助。

```
database table struct
Usage: dbexport []
Options:
  -db database
    	the need export database of mysql
  -dsfile file
    	the destination of export file (default "export.xlsx")
  -extype export type
    	export type,only can: xlsx,md,pdf (default "xlsx")
  -h host
    	the mysql host contain ip and port, ex: localhost:3306 (default "localhost:3306")
  -help
    	help
  -p password
    	the password of mysql
  -u username
    	the username of mysql
```


- 计划支持导出的文件类型：
- [x] excel文件
- [ ] markdown文件
- [ ] pdf文件
- [ ] online直接上线浏览表结构方式



## excel导出

```
go run cmd/main.go -extype xlsx -h localhost:3306 -u root -p root -db boss -dsfile test.xlsx
```
## markdown导出

TODO 

## pdf导出

TODO
 
## online在线浏览

TODO
