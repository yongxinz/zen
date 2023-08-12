// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/yongxin/zen/service/sys/rpc/internal/logic"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

func (s *SysServer) Login(ctx context.Context, in *sys.LoginRequest) (*sys.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *SysServer) UserInfo(ctx context.Context, in *sys.UserInfoReq) (*sys.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *SysServer) UserList(ctx context.Context, in *sys.UserListReq) (*sys.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *SysServer) UserRetrieve(ctx context.Context, in *sys.UserInfoReq) (*sys.UserRetrieveResp, error) {
	l := logic.NewUserRetrieveLogic(ctx, s.svcCtx)
	return l.UserRetrieve(in)
}

func (s *SysServer) UserAdd(ctx context.Context, in *sys.UserAddReq) (*sys.UserAddResp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *SysServer) UserUpdate(ctx context.Context, in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *SysServer) UserUpdateStatus(ctx context.Context, in *sys.UserUpdateStatusReq) (*sys.UserUpdateStatusResp, error) {
	l := logic.NewUserUpdateStatusLogic(ctx, s.svcCtx)
	return l.UserUpdateStatus(in)
}

func (s *SysServer) UserUpdatePwd(ctx context.Context, in *sys.UserUpdatePwdReq) (*sys.UserUpdatePwdResp, error) {
	l := logic.NewUserUpdatePwdLogic(ctx, s.svcCtx)
	return l.UserUpdatePwd(in)
}

func (s *SysServer) UserDelete(ctx context.Context, in *sys.UserDeleteReq) (*sys.UserDeleteResp, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

func (s *SysServer) RoleMenuTree(ctx context.Context, in *sys.RoleMenuTreeReq) (*sys.RoleMenuTreeResp, error) {
	l := logic.NewRoleMenuTreeLogic(ctx, s.svcCtx)
	return l.RoleMenuTree(in)
}

func (s *SysServer) RoleList(ctx context.Context, in *sys.RoleListReq) (*sys.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

func (s *SysServer) RoleAdd(ctx context.Context, in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	l := logic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

func (s *SysServer) RoleRetrieve(ctx context.Context, in *sys.RoleRetrieveReq) (*sys.RoleRetrieveResp, error) {
	l := logic.NewRoleRetrieveLogic(ctx, s.svcCtx)
	return l.RoleRetrieve(in)
}

func (s *SysServer) RoleUpdate(ctx context.Context, in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	l := logic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

func (s *SysServer) RoleDelete(ctx context.Context, in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	l := logic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

func (s *SysServer) MenuAdd(ctx context.Context, in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	l := logic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

func (s *SysServer) MenuList(ctx context.Context, in *sys.MenuListReq) (*sys.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *SysServer) MenuRetrieve(ctx context.Context, in *sys.MenuRetrieveReq) (*sys.MenuRetrieveResp, error) {
	l := logic.NewMenuRetrieveLogic(ctx, s.svcCtx)
	return l.MenuRetrieve(in)
}

func (s *SysServer) MenuUpdate(ctx context.Context, in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *SysServer) MenuDelete(ctx context.Context, in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

func (s *SysServer) MenuRole(ctx context.Context, in *sys.MenuRoleReq) (*sys.MenuRoleResp, error) {
	l := logic.NewMenuRoleLogic(ctx, s.svcCtx)
	return l.MenuRole(in)
}

func (s *SysServer) DeptTree(ctx context.Context, in *sys.DeptTreeReq) (*sys.DeptTreeResp, error) {
	l := logic.NewDeptTreeLogic(ctx, s.svcCtx)
	return l.DeptTree(in)
}

func (s *SysServer) DeptList(ctx context.Context, in *sys.DeptListReq) (*sys.DeptListResp, error) {
	l := logic.NewDeptListLogic(ctx, s.svcCtx)
	return l.DeptList(in)
}

func (s *SysServer) DeptRetrieve(ctx context.Context, in *sys.DeptRetrieveReq) (*sys.DeptRetrieveResp, error) {
	l := logic.NewDeptRetrieveLogic(ctx, s.svcCtx)
	return l.DeptRetrieve(in)
}

func (s *SysServer) DeptAdd(ctx context.Context, in *sys.DeptAddReq) (*sys.DeptAddResp, error) {
	l := logic.NewDeptAddLogic(ctx, s.svcCtx)
	return l.DeptAdd(in)
}

func (s *SysServer) DeptUpdate(ctx context.Context, in *sys.DeptUpdateReq) (*sys.DeptUpdateResp, error) {
	l := logic.NewDeptUpdateLogic(ctx, s.svcCtx)
	return l.DeptUpdate(in)
}

func (s *SysServer) DeptDelete(ctx context.Context, in *sys.DeptDeleteReq) (*sys.DeptDeleteResp, error) {
	l := logic.NewDeptDeleteLogic(ctx, s.svcCtx)
	return l.DeptDelete(in)
}

func (s *SysServer) PostList(ctx context.Context, in *sys.PostListReq) (*sys.PostListResp, error) {
	l := logic.NewPostListLogic(ctx, s.svcCtx)
	return l.PostList(in)
}

func (s *SysServer) PostRetrieve(ctx context.Context, in *sys.PostRetrieveReq) (*sys.PostRetrieveResp, error) {
	l := logic.NewPostRetrieveLogic(ctx, s.svcCtx)
	return l.PostRetrieve(in)
}

func (s *SysServer) PostAdd(ctx context.Context, in *sys.PostAddReq) (*sys.PostAddResp, error) {
	l := logic.NewPostAddLogic(ctx, s.svcCtx)
	return l.PostAdd(in)
}

func (s *SysServer) PostUpdate(ctx context.Context, in *sys.PostUpdateReq) (*sys.PostUpdateResp, error) {
	l := logic.NewPostUpdateLogic(ctx, s.svcCtx)
	return l.PostUpdate(in)
}

func (s *SysServer) PostDelete(ctx context.Context, in *sys.PostDeleteReq) (*sys.PostDeleteResp, error) {
	l := logic.NewPostDeleteLogic(ctx, s.svcCtx)
	return l.PostDelete(in)
}

func (s *SysServer) LoginLogList(ctx context.Context, in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}
