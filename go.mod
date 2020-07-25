module app

go 1.13

require (
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
	xorm.io/xorm v0.8.2
)

replace github.com/go-xorm/core => xorm.io/core v0.7.3
