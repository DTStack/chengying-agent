package api

import (
	apibase "easyagent/go-common/api-base"
	"easyagent/internal/server/api/impl"
)

var ServerManageApis = apibase.Route{
	Path: "server",
	SubRoutes: []apibase.Route{{
		Path: "dashboard",
		SubRoutes: []apibase.Route{{
			//http://xxxx/api/v1/server/dashboard/url
			Path: "url",
			GET:  impl.RetDashboardUrl,
			Docs: apibase.Docs{
				GET: &apibase.ApiDoc{
					Name: "提供给出 dashboard的 URL列表API",
					Query: apibase.ApiParams{
						"type": apibase.ApiParam{"string", "'cluster' or 'services'", "", true},
						"id":   apibase.ApiParam{"string", "服务器组ID", "", true},
					},
					Returns: []apibase.ApiReturnGroup{{
						Fields: apibase.ResultFields{
							"$[*].url": apibase.ApiReturn{"string", "dashboard仪表盘链接URL"},
						},
					}},
				},
			},
		}},
	}},
}
