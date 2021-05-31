package model

import (
	dbhelper "easyagent/go-common/db-helper"
	"time"
)

type resourceUsageRecordModel struct {
	dbhelper.DbTable
}

var ResourceUsageRecords = resourceUsageRecordModel{
	DbTable: dbhelper.DbTable{
		USE_MYSQL_DB,
		TBL_RESOURCE_USAGES,
	},
}

func (u *resourceUsageRecordModel) SaveRecord(ts time.Time) {
}
