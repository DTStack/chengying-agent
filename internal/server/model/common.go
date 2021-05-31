package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var (
	UICDB   *sqlx.DB
	MYSQLDB *sqlx.DB
)

const (
	TBL_PRODUCT_LIST     = "product_list"
	TBL_SIDECAR_LIST     = "sidecar_list"
	TBL_AGENT_LIST       = "agent_list"
	TBL_OP_HISTORY       = "operation_history"
	TBL_UPDATE_HISTORY   = "update_history"
	TBL_PROGRESS_HISTORY = "progress_history"
	TBL_RESOURCE_USAGES  = "resource_usages"
	TBL_DEPLOY_CALLBACK  = "deploy_callback"
	TBL_DASHBOARD_LIST   = "dashboard_list"

	TBL_TRIGGER_LIST  = "trigger_list"
	TBL_STRATEGY_LIST = "strategy_list"
)

func USE_MYSQL_DB() *sqlx.DB {
	return MYSQLDB
}

func USE_UIC_DB() *sqlx.DB {
	return UICDB
}

func connectDatabase(host, user, password, dbname string, port int) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true", user, password, host, port, dbname))
}

func ConfigureUicDatabase(host string, port int, user, password, dbname string) error {
	var err error
	UICDB, err = connectDatabase(host, user, password, dbname, port)
	return err
}

func ConfigureMysqlDatabase(host string, port int, user, password, dbname string) error {
	var err error
	MYSQLDB, err = connectDatabase(host, user, password, dbname, port)
	return err
}
