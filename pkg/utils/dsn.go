package utils

import (
	"strings"

	"github.com/jiuxia211/ice-pomelo/config"
)

func GetMysqlDsn() string {
	// if config.ConfigInfo.Mysql == nil {
	// 	klog.Fatal("config not found")
	// }

	dsn := strings.Join([]string{config.ConfigInfo.Mysql.Username, ":", config.ConfigInfo.Mysql.Password, "@tcp(", config.ConfigInfo.Mysql.Addr, ")/", config.ConfigInfo.Mysql.Database, "?charset=" + config.ConfigInfo.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}
