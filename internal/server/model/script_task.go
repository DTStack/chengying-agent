package model

import (
	"database/sql"
	apibase "easyagent/go-common/api-base"
	dbhelper "easyagent/go-common/db-helper"
	"strconv"
	"time"
)

type scriptTask struct {
	dbhelper.DbTable
}

var ScriptTask = &scriptTask{
	dbhelper.DbTable{USE_MYSQL_DB, TBL_SCRIPT_TASK},
}

type ScriptTaskInfo struct {
	ID         int               `db:"id"`
	SeqNo      int               `db:"seq_no"`
	SidecarId  string            `db:"sidecar_id"`
	Status     int               `db:"status"`
	Message    sql.NullString    `db:"message"`
	CreateTime dbhelper.NullTime `db:"create_time"`
	FinishTime dbhelper.NullTime `db:"finish_time"`
}

type ErrMsg struct {
	String string `json:"String"`
	Valid  bool   `json:"Valid"`
}

func (t *scriptTask) InsertScriptTask(seqNo int64, status int, sidecarId string) (int64, error) {
	result, err := t.InsertWhere(dbhelper.UpdateFields{
		"sidecar_id":  sidecarId,
		"seq_no":      seqNo,
		"status":      status,
		"create_time": time.Now(),
	})
	if err != nil {
		apibase.ThrowDBModelError(err)
	}
	return result.LastInsertId()
}

func (t *scriptTask) GetScriptTask(taskId string) (*ScriptTaskInfo, error) {
	whereClause := dbhelper.WhereCause{}
	taskIdStr, err := strconv.Atoi(taskId)
	if err != nil {
		apibase.ThrowDBModelError(err)
	}
	info := ScriptTaskInfo{}
	err = t.GetWhere(nil, whereClause.Equal("id", taskIdStr), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (t *scriptTask) UpdateScriptTask(seq uint32, updateFields dbhelper.UpdateFields) error {
	if !t.Exists(int(seq)) {
		return nil
	}
	return t.UpdateWhere(dbhelper.MakeWhereCause().Equal("seq_no", seq), updateFields, false)
}

func (t *scriptTask) Exists(seq int) bool {
	whereClause := dbhelper.WhereCause{}
	task := ScriptTaskInfo{}
	err := t.GetWhere(nil, whereClause.Equal("seq_no", seq), &task)
	if err != nil {
		return false
	}
	return true
}

func (t *scriptTask) UpdateTaskStatusBySid(whereClause dbhelper.WhereCause, updateFields dbhelper.UpdateFields) error {
	return t.UpdateWhere(whereClause, updateFields, false)
}
