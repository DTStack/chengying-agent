package api

import (
	apibase "easyagent/go-common/api-base"
	"easyagent/internal/server/api/impl"
)

var LabelApis = apibase.Route{
	Path: "label",
	SubRoutes: []apibase.Route{{
		Path: "{label:string}",
		SubRoutes: []apibase.Route{{
			Path: "startScript",
			POST: impl.StartScriptByLabel,
			Docs: apibase.Docs{
				POST: &apibase.ApiDoc{
					Name: "根据label执行脚本任务",
				},
			},
		}},
	}},
}
