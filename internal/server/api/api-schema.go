package api

import (
	apibase "easyagent/go-common/api-base"
)

var ApiV1Schema = apibase.Route{
	Path: "/api/v1",
	SubRoutes: []apibase.Route{
		SystemApiRoutes,
		SidecarManagementRoutes,
		AgentListManagementRoutes,
		DeploymentManageApis,
		ServerManageApis,
		SshhApiRoutes,
	},
}
