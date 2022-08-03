package report

import (
	"bytes"
	apibase "easyagent/go-common/api-base"
	"easyagent/internal/server/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
 @Author: zhijian
 @Date: 2021/5/31 10:55
 @Description:
*/
var ReportHost, ReportSeqUri, IsShowLogUri, ShellStatusUri string

type ReqRequst struct {
	ExecId string `json:"execId"`
	SeqNo  uint32 `json:"seq"`
}

func ReportSeq(execId string, seq uint32) error {

	reqStruct := ReqRequst{
		ExecId: execId,
		SeqNo:  seq,
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s%s", ReportHost, ReportSeqUri), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	result := apibase.ApiResult{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf("report seq error:%s", result.Msg)
	}
	return nil
}

func IsShowLog(seq uint32) (bool, error) {
	if seq == 0 {
		return false, nil
	}
	log.Debugf("IsShowLog: seq %d", seq)
	url := fmt.Sprintf(fmt.Sprintf("http://%s%s?seq=%d", ReportHost, IsShowLogUri, seq))
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	result := apibase.ApiResult{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return false, err
	}
	if result.Code != 0 {
		return false, fmt.Errorf("report seq error:%s", result.Msg)
	}
	if res, ok := result.Data.(bool); ok {
		return res, nil
	} else {
		return false, fmt.Errorf("IsShowLog bool 类型转换 err")
	}

}

func ShellStatusReport(seq uint32, status int) error {
	reqStruct := struct {
		Seq    uint32 `json:"seq"`
		Status int    `json:"status"`
	}{
		Seq:    seq,
		Status: status,
	}
	if reqStruct.Seq == 0 {
		return nil
	}
	log.Debugf("ShellStatusReport: seq %d status %s", seq, status)
	body, err := json.Marshal(reqStruct)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("http://%s%s", ReportHost, ShellStatusUri), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	result := apibase.ApiResult{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	if result.Code != 0 {
		return fmt.Errorf("shell status report error:%s", result.Msg)
	}
	return nil
}
