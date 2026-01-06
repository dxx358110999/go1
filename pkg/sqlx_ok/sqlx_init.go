package sqlx_ok

import (
	"dxxproject/config_prepare/app_config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlxDb struct {
	db *sqlx.DB
}

func (r *SqlxDb) Init(dbConfig *app_config.MysqlConfig) (err error) {
	/*
		初始化MySQL连接
	*/

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB,
	) //构造连接
	r.db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	r.db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	r.db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	return
}

func (r *SqlxDb) Close() {
	r.db.Close()
}
