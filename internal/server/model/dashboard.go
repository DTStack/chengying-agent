package model

import (
	apibase "easyagent/go-common/api-base"
	dbhelper "easyagent/go-common/db-helper"
)

type dashboardListTable struct {
	dbhelper.DbTable
}

var DashboardList = dashboardListTable{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_DASHBOARD_LIST},
}

type DashboardListInfo struct {
	ID          int    `db:"id"`
	ClusterID   string `db:"clusterId"`
	ServicesID  string `db:"servicesId"`
	ContainerID string `db:"containerId"`
	HostID      string `db:"hostId"`
	Url         string `db:"url"`
}

func (dl *dashboardListTable) RetUrlByClusterID(clusterId string) string {
	where := dbhelper.MakeWhereCause()
	if len(clusterId) > 0 {
		where = where.Equal("clusterId", clusterId)
	}
	row := dl.SelectOneWhere(nil, where)
	if row != nil {
		info := DashboardListInfo{}
		err := row.StructScan(&info)
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		return info.Url
	} else {
		apibase.ThrowDBModelError("Null result where clusterId = %s", clusterId)
	}
	return ""
}

func (dl *dashboardListTable) RetUrlByServicesID(servicesId string) string {
	where := dbhelper.MakeWhereCause()
	if len(servicesId) > 0 {
		where = where.Equal("servicesId", servicesId)
	}
	row := dl.SelectOneWhere(nil, where)
	if row != nil {
		info := DashboardListInfo{}
		err := row.StructScan(&info)
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		return info.Url
	} else {
		apibase.ThrowDBModelError("Null result where servicesId = %s", servicesId)
	}
	return ""
}
