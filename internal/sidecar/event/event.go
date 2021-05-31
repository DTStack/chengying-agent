package event

import (
	"easyagent/internal/proto"
	"easyagent/internal/sidecar/base"
	"easyagent/internal/sidecar/client"
)

var eventDefaultClient client.EaClienter

func SetEventDefaultClient(c client.EaClienter) {
	eventDefaultClient = c
}

func ReportEvent(event interface{}) {
	if eventDefaultClient == nil {
		return
	}

	e := &proto.Event{}
	switch inst := event.(type) {
	case *proto.Event_OperationProgress:
		e.EventType = proto.EVT_OP_PROGRESS
		e.Details = &proto.Event_OpProgress{inst}
	case *proto.Event_AgentError:
		e.EventType = proto.EVT_AGENT_ERR
		e.Details = &proto.Event_AgentError_{inst}
	case *proto.Event_OsResourceUsages:
		e.EventType = proto.EVT_OS_RESOURCE_USAGES
		e.Details = &proto.Event_OsResourceUsages_{inst}
	case *proto.Event_ProcessResourceUsages:
		e.EventType = proto.EVT_PROC_RESOURCE_USAGES
		e.Details = &proto.Event_ProcResourceUsages{inst}
	case *proto.Event_ExecScriptResponse:
		e.EventType = proto.EVT_EXEC_SCRIPT
		e.Details = &proto.Event_ExecScriptResponse_{inst}
	case *proto.Event_ExecRestResponse:
		e.EventType = proto.EVT_REST_SCRIPT
		e.Details = &proto.Event_ExecRestResponse_{inst}
	case *proto.Event_AgentHealthCheck:
		e.EventType = proto.EVT_AGENT_HEALTH_CHECK
		e.Details = &proto.Event_AgentHealthCheck_{inst}
	default:
		return
	}

	if err := eventDefaultClient.ReportEvent(e); err != nil {
		base.Errorf("ReportEvent error: %v", err)
	}
}
