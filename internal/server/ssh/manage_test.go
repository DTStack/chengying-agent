package sshs

import (
	"testing"

	"easyagent/internal/server/log"
	"fmt"
)

func init() {
	log.ConfigureLogger("/tmp", 0, 0, 0)
}

func TestManage_RunWithSSH(t *testing.T) {
	param := &SshParam{
		Host: "172.16.10.108",
		User: "dtstack",
		Pass: "abc123",
		Port: 22,
		Mode: 1,
		Cmd:  "sh /opt/dtstack/easymanager/easyagent/easyagent.sh restart",
	}
	result, err := SSHManager.RunWithSSH(param, true)
	if err != nil {
		t.Errorf("RunWithSSH err:%v, %v", err, result)
	}
	fmt.Println(result)
}

func TestManage_RunWithSSHS(t *testing.T) {
	param := &SshParam{
		Host: "172.16.10.108",
		User: "dtstack",
		Pass: "abc123",
		Port: 22,
		Mode: 1,
		Cmd:  "sudo systemctl status network",
	}
	params := []*SshParam{param}
	SSHManager.RunWithSSHS(params, true)
}
