package model

import (
	apibase "easyagent/go-common/api-base"
	dbhelper "easyagent/go-common/db-helper"
	"easyagent/go-common/utils"
	"easyagent/internal/server/log"

	uuid "github.com/satori/go.uuid"
)

type agentList struct {
	dbhelper.DbTable
}

var AgentList = &agentList{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_AGENT_LIST},
}

type AgentInfo struct {
	ID            string            `db:"id"`
	SidecarId     string            `db:"sidecar_id"`
	Type          int               `db:"type"`
	Name          string            `db:"name"`
	Version       string            `db:"version"`
	IsUninstalled int               `db:"is_uninstalled"`
	DeployDate    dbhelper.NullTime `db:"deploy_date"`
	AutoDeploy    bool              `db:"auto_deployment"`
	UpdateDate    dbhelper.NullTime `db:"last_update_date"`
	AutoUpdate    bool              `db:"auto_updated"`
}

func (l *agentList) InsertAgentRecord(sidecarId, agentId uuid.UUID, agentType int, name, version string) uuid.UUID {
	if _, err := l.InsertWhere(dbhelper.UpdateFields{
		"id":         agentId.String(),
		"sidecar_id": sidecarId.String(),
		"name":       name,
		"type":       agentType,
		"version":    version,
	}); err != nil {
		apibase.ThrowDBModelError(err)
	}
	return agentId
}

func (l *agentList) NewAgentRecord(sidecarId uuid.UUID, agentType int, name, version string) uuid.UUID {
	id := uuid.NewV4()
	if _, err := l.InsertWhere(dbhelper.UpdateFields{
		"id":         id.String(),
		"sidecar_id": sidecarId.String(),
		"name":       name,
		"type":       agentType,
		"version":    version,
	}); err != nil {
		apibase.ThrowDBModelError(err)
	}
	return id
}

func (l *agentList) DeleteByagentId(agentId string) {
	query := "DELETE from " + TBL_AGENT_LIST + " "
	query += "WHERE id='" + agentId + "'"
	_, err := l.GetDB().Exec(query)
	if err != nil {
		log.Errorf("DeleteByagentId:%v", agentId)
	}
	return
}

func (l *agentList) GetAgentInfo(id uuid.UUID) *AgentInfo {
	whereCause := dbhelper.WhereCause{}
	info := AgentInfo{}
	err := l.GetWhere(nil, whereCause.Equal("id", id.String()), &info)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}
	return &info
}

var _getAgentListFields = utils.GetTagValues(AgentInfo{}, "db")

func (l *agentList) GetAgentList(pagination *apibase.Pagination) ([]AgentInfo, int) {
	rows, totalRecords, err := l.SelectWhere(_getAgentListFields, nil, pagination)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}

	list := []AgentInfo{}
	for rows.Next() {
		info := AgentInfo{}
		err = rows.StructScan(&info)
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		list = append(list, info)
	}
	return list, totalRecords
}

func (l *agentList) GetAgentsBySidecarId(pagination *apibase.Pagination, id uuid.UUID) ([]AgentInfo, int) {
	whereCause := dbhelper.WhereCause{}
	rows, totalRecords, err := l.SelectWhere(_getAgentListFields, whereCause.Equal("sidecar_id", id.String()), pagination)

	if err != nil {
		apibase.ThrowDBModelError(err)
	}

	list := []AgentInfo{}
	for rows.Next() {
		info := AgentInfo{}
		err = rows.StructScan(&info)
		if err != nil {
			apibase.ThrowDBModelError(err)
		}
		list = append(list, info)
	}
	return list, totalRecords
}

func (l *agentList) GetAgentSidecarId(id uuid.UUID) uuid.UUID {
	var sidecarId string
	err := l.GetWhere([]string{"sidecar_id"}, dbhelper.MakeWhereCause().Equal("id", id.String()), &sidecarId)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}
	id, err = uuid.FromString(sidecarId)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}
	return id
}

func (l *agentList) CheckAgentId(id uuid.UUID) (uuid.UUID, error) {
	var sidecarId string
	err := l.GetWhere([]string{"sidecar_id"}, dbhelper.MakeWhereCause().Equal("id", id.String()), &sidecarId)
	if err != nil {
		return uuid.Nil, err
	}
	id, err = uuid.FromString(sidecarId)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
