package api

import (
	apibase "easyagent/go-common/api-base"
	"easyagent/internal/server/api/impl"
)

var SshhApiRoutes = apibase.Route{
	Path: "ssh",
	SubRoutes: []apibase.Route{{
		Path: "checkByUserPwd",
		POST: impl.CheckByUserPwd,
		Docs: apibase.Docs{
			POST: &apibase.ApiDoc{
				Name: "ssh连通性(用户名密码)",
				Body: apibase.ApiParams{
					"$.host":     apibase.ApiParam{"string", "主机域名orIP", "", true},
					"$.port":     apibase.ApiParam{"int", "端口", "", true},
					"$.user":     apibase.ApiParam{"string", "用户名", "", true},
					"$.password": apibase.ApiParam{"string", "登录密码", "", true},
				},
			},
		},
	}, {
		Path: "runWithUserPwd",
		POST: impl.RunWithUserPwd,
		Docs: apibase.Docs{
			POST: &apibase.ApiDoc{
				Name: "ssh安装(用户名密码)",
				Body: apibase.ApiParams{
					"$.host":     apibase.ApiParam{"string", "主机域名orIP", "", true},
					"$.port":     apibase.ApiParam{"int", "端口", "", true},
					"$.user":     apibase.ApiParam{"string", "用户名", "", true},
					"$.password": apibase.ApiParam{"string", "登录密码", "", true},
					"$.cmd":      apibase.ApiParam{"string", "一键安装脚本", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$.result": apibase.ApiReturn{"string", "执行结果"},
					},
				}},
			},
		},
	}, {
		Path: "checkByUserPk",
		POST: impl.CheckByUserPk,
		Docs: apibase.Docs{
			POST: &apibase.ApiDoc{
				Name: "ssh安装(用户名密钥)",
				Body: apibase.ApiParams{
					"$.host": apibase.ApiParam{"string", "主机域名orIP", "", true},
					"$.port": apibase.ApiParam{"int", "端口", "", true},
					"$.user": apibase.ApiParam{"string", "用户名", "", true},
					"$.pk":   apibase.ApiParam{"string", "秘钥内容", "", true},
				},
			},
		},
	}, {
		Path: "runWithUserPk",
		POST: impl.RunWithUserPk,
		Docs: apibase.Docs{
			POST: &apibase.ApiDoc{
				Name: "ssh安装(用户名密钥)",
				Body: apibase.ApiParams{
					"$.host": apibase.ApiParam{"string", "主机域名orIP", "", true},
					"$.port": apibase.ApiParam{"int", "端口", "", true},
					"$.user": apibase.ApiParam{"string", "用户名", "", true},
					"$.pk":   apibase.ApiParam{"string", "秘钥内容", "", true},
					"$.cmd":  apibase.ApiParam{"string", "一键安装脚本", "", true},
				},
				Returns: []apibase.ApiReturnGroup{{
					Fields: apibase.ResultFields{
						"$.result": apibase.ApiReturn{"string", "执行结果"},
					},
				}},
			},
		},
	}},
}
