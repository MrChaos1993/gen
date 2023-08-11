module gorm.io/gen/tests

go 1.18

require (
	gorm.io/driver/mysql v1.5.2
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gen v0.3.23
	gorm.io/gorm v1.25.5
	gorm.io/plugin/dbresolver v1.4.7
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/hints v1.1.2 // indirect
)

replace gorm.io/gen => ../
