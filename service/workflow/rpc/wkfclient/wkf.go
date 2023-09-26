// Code generated by goctl. DO NOT EDIT.
// Source: wkf.proto

package wkfclient

import (
	"context"

	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ClassifyAddReq       = wkf.ClassifyAddReq
	ClassifyAddResp      = wkf.ClassifyAddResp
	ClassifyDeleteReq    = wkf.ClassifyDeleteReq
	ClassifyDeleteResp   = wkf.ClassifyDeleteResp
	ClassifyListData     = wkf.ClassifyListData
	ClassifyListReq      = wkf.ClassifyListReq
	ClassifyListResp     = wkf.ClassifyListResp
	ClassifyRetrieveReq  = wkf.ClassifyRetrieveReq
	ClassifyRetrieveResp = wkf.ClassifyRetrieveResp
	ClassifyUpdateReq    = wkf.ClassifyUpdateReq
	ClassifyUpdateResp   = wkf.ClassifyUpdateResp
	ProcessAddReq        = wkf.ProcessAddReq
	ProcessAddResp       = wkf.ProcessAddResp
	ProcessClassifyData  = wkf.ProcessClassifyData
	ProcessClassifyReq   = wkf.ProcessClassifyReq
	ProcessClassifyResp  = wkf.ProcessClassifyResp
	ProcessDeleteReq     = wkf.ProcessDeleteReq
	ProcessDeleteResp    = wkf.ProcessDeleteResp
	ProcessListData      = wkf.ProcessListData
	ProcessListReq       = wkf.ProcessListReq
	ProcessListResp      = wkf.ProcessListResp
	ProcessRetrieveReq   = wkf.ProcessRetrieveReq
	ProcessRetrieveResp  = wkf.ProcessRetrieveResp
	ProcessUpdateReq     = wkf.ProcessUpdateReq
	ProcessUpdateResp    = wkf.ProcessUpdateResp
	TaskAddReq           = wkf.TaskAddReq
	TaskAddResp          = wkf.TaskAddResp
	TaskDeleteReq        = wkf.TaskDeleteReq
	TaskDeleteResp       = wkf.TaskDeleteResp
	TaskListData         = wkf.TaskListData
	TaskListReq          = wkf.TaskListReq
	TaskListResp         = wkf.TaskListResp
	TaskRetrieveReq      = wkf.TaskRetrieveReq
	TaskRetrieveResp     = wkf.TaskRetrieveResp
	TaskUpdateReq        = wkf.TaskUpdateReq
	TaskUpdateResp       = wkf.TaskUpdateResp
	TemplateAddReq       = wkf.TemplateAddReq
	TemplateAddResp      = wkf.TemplateAddResp
	TemplateDeleteReq    = wkf.TemplateDeleteReq
	TemplateDeleteResp   = wkf.TemplateDeleteResp
	TemplateListData     = wkf.TemplateListData
	TemplateListReq      = wkf.TemplateListReq
	TemplateListResp     = wkf.TemplateListResp
	TemplateRetrieveReq  = wkf.TemplateRetrieveReq
	TemplateRetrieveResp = wkf.TemplateRetrieveResp
	TemplateUpdateReq    = wkf.TemplateUpdateReq
	TemplateUpdateResp   = wkf.TemplateUpdateResp
	TicketAddReq         = wkf.TicketAddReq
	TicketAddResp        = wkf.TicketAddResp
	TicketDeleteReq      = wkf.TicketDeleteReq
	TicketDeleteResp     = wkf.TicketDeleteResp
	TicketFinishReq      = wkf.TicketFinishReq
	TicketFinishResp     = wkf.TicketFinishResp
	TicketHandleReq      = wkf.TicketHandleReq
	TicketHandleResp     = wkf.TicketHandleResp
	TicketListData       = wkf.TicketListData
	TicketListReq        = wkf.TicketListReq
	TicketListResp       = wkf.TicketListResp
	TicketProcessReq     = wkf.TicketProcessReq
	TicketProcessResp    = wkf.TicketProcessResp
	TicketTransferReq    = wkf.TicketTransferReq
	TicketTransferResp   = wkf.TicketTransferResp

	Wkf interface {
		ClassifyList(ctx context.Context, in *ClassifyListReq, opts ...grpc.CallOption) (*ClassifyListResp, error)
		ClassifyRetrieve(ctx context.Context, in *ClassifyRetrieveReq, opts ...grpc.CallOption) (*ClassifyRetrieveResp, error)
		ClassifyAdd(ctx context.Context, in *ClassifyAddReq, opts ...grpc.CallOption) (*ClassifyAddResp, error)
		ClassifyUpdate(ctx context.Context, in *ClassifyUpdateReq, opts ...grpc.CallOption) (*ClassifyUpdateResp, error)
		ClassifyDelete(ctx context.Context, in *ClassifyDeleteReq, opts ...grpc.CallOption) (*ClassifyDeleteResp, error)
		TemplateList(ctx context.Context, in *TemplateListReq, opts ...grpc.CallOption) (*TemplateListResp, error)
		TemplateRetrieve(ctx context.Context, in *TemplateRetrieveReq, opts ...grpc.CallOption) (*TemplateRetrieveResp, error)
		TemplateAdd(ctx context.Context, in *TemplateAddReq, opts ...grpc.CallOption) (*TemplateAddResp, error)
		TemplateUpdate(ctx context.Context, in *TemplateUpdateReq, opts ...grpc.CallOption) (*TemplateUpdateResp, error)
		TemplateDelete(ctx context.Context, in *TemplateDeleteReq, opts ...grpc.CallOption) (*TemplateDeleteResp, error)
		TaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error)
		TaskRetrieve(ctx context.Context, in *TaskRetrieveReq, opts ...grpc.CallOption) (*TaskRetrieveResp, error)
		TaskAdd(ctx context.Context, in *TaskAddReq, opts ...grpc.CallOption) (*TaskAddResp, error)
		TaskUpdate(ctx context.Context, in *TaskUpdateReq, opts ...grpc.CallOption) (*TaskUpdateResp, error)
		TaskDelete(ctx context.Context, in *TaskDeleteReq, opts ...grpc.CallOption) (*TaskDeleteResp, error)
		ProcessList(ctx context.Context, in *ProcessListReq, opts ...grpc.CallOption) (*ProcessListResp, error)
		ProcessRetrieve(ctx context.Context, in *ProcessRetrieveReq, opts ...grpc.CallOption) (*ProcessRetrieveResp, error)
		ProcessAdd(ctx context.Context, in *ProcessAddReq, opts ...grpc.CallOption) (*ProcessAddResp, error)
		ProcessUpdate(ctx context.Context, in *ProcessUpdateReq, opts ...grpc.CallOption) (*ProcessUpdateResp, error)
		ProcessDelete(ctx context.Context, in *ProcessDeleteReq, opts ...grpc.CallOption) (*ProcessDeleteResp, error)
		ProcessClassify(ctx context.Context, in *ProcessClassifyReq, opts ...grpc.CallOption) (*ProcessClassifyResp, error)
		TicketProcess(ctx context.Context, in *TicketProcessReq, opts ...grpc.CallOption) (*TicketProcessResp, error)
		TicketList(ctx context.Context, in *TicketListReq, opts ...grpc.CallOption) (*TicketListResp, error)
		TicketAdd(ctx context.Context, in *TicketAddReq, opts ...grpc.CallOption) (*TicketAddResp, error)
		TicketHandle(ctx context.Context, in *TicketHandleReq, opts ...grpc.CallOption) (*TicketHandleResp, error)
		TicketDelete(ctx context.Context, in *TicketDeleteReq, opts ...grpc.CallOption) (*TicketDeleteResp, error)
		TicketFinish(ctx context.Context, in *TicketFinishReq, opts ...grpc.CallOption) (*TicketFinishResp, error)
		TicketTransfer(ctx context.Context, in *TicketTransferReq, opts ...grpc.CallOption) (*TicketTransferResp, error)
	}

	defaultWkf struct {
		cli zrpc.Client
	}
)

func NewWkf(cli zrpc.Client) Wkf {
	return &defaultWkf{
		cli: cli,
	}
}

func (m *defaultWkf) ClassifyList(ctx context.Context, in *ClassifyListReq, opts ...grpc.CallOption) (*ClassifyListResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ClassifyList(ctx, in, opts...)
}

func (m *defaultWkf) ClassifyRetrieve(ctx context.Context, in *ClassifyRetrieveReq, opts ...grpc.CallOption) (*ClassifyRetrieveResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ClassifyRetrieve(ctx, in, opts...)
}

func (m *defaultWkf) ClassifyAdd(ctx context.Context, in *ClassifyAddReq, opts ...grpc.CallOption) (*ClassifyAddResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ClassifyAdd(ctx, in, opts...)
}

func (m *defaultWkf) ClassifyUpdate(ctx context.Context, in *ClassifyUpdateReq, opts ...grpc.CallOption) (*ClassifyUpdateResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ClassifyUpdate(ctx, in, opts...)
}

func (m *defaultWkf) ClassifyDelete(ctx context.Context, in *ClassifyDeleteReq, opts ...grpc.CallOption) (*ClassifyDeleteResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ClassifyDelete(ctx, in, opts...)
}

func (m *defaultWkf) TemplateList(ctx context.Context, in *TemplateListReq, opts ...grpc.CallOption) (*TemplateListResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TemplateList(ctx, in, opts...)
}

func (m *defaultWkf) TemplateRetrieve(ctx context.Context, in *TemplateRetrieveReq, opts ...grpc.CallOption) (*TemplateRetrieveResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TemplateRetrieve(ctx, in, opts...)
}

func (m *defaultWkf) TemplateAdd(ctx context.Context, in *TemplateAddReq, opts ...grpc.CallOption) (*TemplateAddResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TemplateAdd(ctx, in, opts...)
}

func (m *defaultWkf) TemplateUpdate(ctx context.Context, in *TemplateUpdateReq, opts ...grpc.CallOption) (*TemplateUpdateResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TemplateUpdate(ctx, in, opts...)
}

func (m *defaultWkf) TemplateDelete(ctx context.Context, in *TemplateDeleteReq, opts ...grpc.CallOption) (*TemplateDeleteResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TemplateDelete(ctx, in, opts...)
}

func (m *defaultWkf) TaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TaskList(ctx, in, opts...)
}

func (m *defaultWkf) TaskRetrieve(ctx context.Context, in *TaskRetrieveReq, opts ...grpc.CallOption) (*TaskRetrieveResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TaskRetrieve(ctx, in, opts...)
}

func (m *defaultWkf) TaskAdd(ctx context.Context, in *TaskAddReq, opts ...grpc.CallOption) (*TaskAddResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TaskAdd(ctx, in, opts...)
}

func (m *defaultWkf) TaskUpdate(ctx context.Context, in *TaskUpdateReq, opts ...grpc.CallOption) (*TaskUpdateResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TaskUpdate(ctx, in, opts...)
}

func (m *defaultWkf) TaskDelete(ctx context.Context, in *TaskDeleteReq, opts ...grpc.CallOption) (*TaskDeleteResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TaskDelete(ctx, in, opts...)
}

func (m *defaultWkf) ProcessList(ctx context.Context, in *ProcessListReq, opts ...grpc.CallOption) (*ProcessListResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessList(ctx, in, opts...)
}

func (m *defaultWkf) ProcessRetrieve(ctx context.Context, in *ProcessRetrieveReq, opts ...grpc.CallOption) (*ProcessRetrieveResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessRetrieve(ctx, in, opts...)
}

func (m *defaultWkf) ProcessAdd(ctx context.Context, in *ProcessAddReq, opts ...grpc.CallOption) (*ProcessAddResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessAdd(ctx, in, opts...)
}

func (m *defaultWkf) ProcessUpdate(ctx context.Context, in *ProcessUpdateReq, opts ...grpc.CallOption) (*ProcessUpdateResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessUpdate(ctx, in, opts...)
}

func (m *defaultWkf) ProcessDelete(ctx context.Context, in *ProcessDeleteReq, opts ...grpc.CallOption) (*ProcessDeleteResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessDelete(ctx, in, opts...)
}

func (m *defaultWkf) ProcessClassify(ctx context.Context, in *ProcessClassifyReq, opts ...grpc.CallOption) (*ProcessClassifyResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.ProcessClassify(ctx, in, opts...)
}

func (m *defaultWkf) TicketProcess(ctx context.Context, in *TicketProcessReq, opts ...grpc.CallOption) (*TicketProcessResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketProcess(ctx, in, opts...)
}

func (m *defaultWkf) TicketList(ctx context.Context, in *TicketListReq, opts ...grpc.CallOption) (*TicketListResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketList(ctx, in, opts...)
}

func (m *defaultWkf) TicketAdd(ctx context.Context, in *TicketAddReq, opts ...grpc.CallOption) (*TicketAddResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketAdd(ctx, in, opts...)
}

func (m *defaultWkf) TicketHandle(ctx context.Context, in *TicketHandleReq, opts ...grpc.CallOption) (*TicketHandleResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketHandle(ctx, in, opts...)
}

func (m *defaultWkf) TicketDelete(ctx context.Context, in *TicketDeleteReq, opts ...grpc.CallOption) (*TicketDeleteResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketDelete(ctx, in, opts...)
}

func (m *defaultWkf) TicketFinish(ctx context.Context, in *TicketFinishReq, opts ...grpc.CallOption) (*TicketFinishResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketFinish(ctx, in, opts...)
}

func (m *defaultWkf) TicketTransfer(ctx context.Context, in *TicketTransferReq, opts ...grpc.CallOption) (*TicketTransferResp, error) {
	client := wkf.NewWkfClient(m.cli.Conn())
	return client.TicketTransfer(ctx, in, opts...)
}
