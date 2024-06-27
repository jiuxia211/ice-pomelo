package db

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jiuxia211/ice-pomelo/pkg/constants"
	"github.com/jiuxia211/ice-pomelo/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(utils.GetMysqlDsn()), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		klog.Fatalf("mysql open error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		klog.Fatalf("mysql setting error: %v", err)
	}

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间

	DB = db.Table(constants.VideoTableName)

}
