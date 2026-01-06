package gorm_db

import (
	"dxxproject/config_prepare/app_config"
	"fmt"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDb(injector do.Injector) (db *gorm.DB, err error) {
	cfg := do.MustInvoke[*app_config.AppConfig](injector).MysqlConfig

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	) //data source name
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		err = errors.Wrapf(err, "gorm初始化失败")
		return
	}

	return
}
