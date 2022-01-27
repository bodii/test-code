package mysqlConf

import (
	"fmt"
)

var conf MysqlConfType = MysqlConfType{
	user:         "root",
	password:     "123456",
	host:         "127.0.0.1",
	port:         "3306",
	databaseName: "yt_go_auth",
	charset:      "utf8mb4",
	parseTime:    "True",
	loc:          "Local",
}

func New() string {
	// 配置内容
	return formatDSN(conf)
}

type MysqlConfType struct {
	user         string
	password     string
	host         string
	port         string
	databaseName string
	charset      string
	parseTime    string
	loc          string
}

func formatDSN(conf MysqlConfType) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.databaseName,
		conf.charset,
		conf.parseTime,
		conf.loc,
	)
}
