package tracy

import (
	"context"
	"fmt"
	"time"

	"easyagent/internal/server/log"
	"easyagent/internal/server/publisher"
)

const (
	LogTsLayout = "2006-01-02 15:04:05.000000"
)

const (
	AGENT_INSTALL_PROGRESS_LOG = "agent-install"
	AGENT_CONTROL_PROGRESS_LOG = "agent-control"
)

func TracyOutput2Path(path string, toOutput bool, format string, args ...interface{}) error {

	if toOutput {
		format = time.Now().Format(LogTsLayout) + " " + format
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		index := "dtlog-1-easyagent-nodelete-" + time.Now().Format("2006.01.02") + "_000001.alias"
		body := struct {
			Msg        string `json:"msg"`
			LastUpdate string `json:"last_update_date"`
		}{fmt.Sprintf(format, args...), time.Now().Format(LogTsLayout)}
		if err := publisher.Publish.OutputJson(ctx, "", index, "dt_agent_install_progress",
			body, []byte{}); err != nil {
		}
		cancel()
	}
	return log.Output2Path(path, format, args...)
}

func InstallProgressLog(format string, args ...interface{}) {
	TracyOutput2Path(AGENT_INSTALL_PROGRESS_LOG, true, format, args...)
}

func ControlProgressLog(format string, args ...interface{}) {
	TracyOutput2Path(AGENT_CONTROL_PROGRESS_LOG, true, format, args...)
}
