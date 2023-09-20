// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	classify "github.com/yongxin/zen/service/workflow/api/internal/handler/classify"
	process "github.com/yongxin/zen/service/workflow/api/internal/handler/process"
	task "github.com/yongxin/zen/service/workflow/api/internal/handler/task"
	template "github.com/yongxin/zen/service/workflow/api/internal/handler/template"
	ticket "github.com/yongxin/zen/service/workflow/api/internal/handler/ticket"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/classify",
				Handler: classify.ClassifyListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/classify/:classifyId",
				Handler: classify.ClassifyRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/classify",
				Handler: classify.ClassifyAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/classify/:classifyId",
				Handler: classify.ClassifyUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/classify/:classifyId",
				Handler: classify.ClassifyDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/workflow"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/template",
				Handler: template.TemplateListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/template/:templateId",
				Handler: template.TemplateRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/template",
				Handler: template.TemplateAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/template/:templateId",
				Handler: template.TemplateUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/template/:templateId",
				Handler: template.TemplateDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/workflow"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/task",
				Handler: task.TaskListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/task/:taskId",
				Handler: task.TaskRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/task",
				Handler: task.TaskAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/task/:taskId",
				Handler: task.TaskUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/task/:taskId",
				Handler: task.TaskDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/workflow"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/process",
				Handler: process.ProcessListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/process/:processId",
				Handler: process.ProcessRetrieveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/process",
				Handler: process.ProcessAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/process/:processId",
				Handler: process.ProcessUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/process/:processId",
				Handler: process.ProcessDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/process/classify",
				Handler: process.ProcessClassifyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/workflow"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ticket",
				Handler: ticket.TicketAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ticket",
				Handler: ticket.TicketListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ticket/process-structure",
				Handler: ticket.TicketInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/ticket/handle",
				Handler: ticket.TicketHandleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1/workflow"),
	)
}
