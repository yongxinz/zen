// Code generated by goctl. DO NOT EDIT.
package types

type PageReq struct {
	PageIndex int64 `form:"pageIndex,default=1"`
	PageSize  int64 `form:"pageSize,default=20"`
}

type Pagination struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Count     int64 `json:"count"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
	Code     string `json:"code"`
}

type LoginResp struct {
	CurrentAuthority string `json:"currentAuthority"`
	Expire           int64  `json:"expire"`
	Token            string `json:"token"`
}

type CaptchaResp struct {
	Data string `json:"data"`
	Id   string `json:"id"`
}

type UserInfoResp struct {
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	UserName     string   `json:"userName"`
	UserId       int64    `json:"userId"`
	DeptId       int64    `json:"deptId"`
	Name         string   `json:"name"`
	Buttons      []string `json:"buttons"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
}

type UserListReq struct {
	PageReq
}

type UserListData struct {
	UserId    int64  `json:"userId"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Status    int64  `json:"status"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Sex       int64  `json:"sex"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
	RoleId    int64  `json:"roleId"`
	DeptId    int64  `json:"deptId"`
	PostId    int64  `json:"postId"`
	DeptName  string `json:"deptName"`
}

type UserListResp struct {
	List []UserListData `json:"list"`
	Pagination
}

type UserAddReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Status   int64  `json:"status"`
	Email    string `json:"email"`
	Sex      int64  `json:"sex,optional"`
	Remark   string `json:"remark,optional"`
	RoleId   int64  `json:"roleId,optional"`
	DeptId   int64  `json:"deptId"`
	PostId   int64  `json:"postId,optional"`
	CreateBy int64  `json:"createBy,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type UserRetrieveReq struct {
	UserId int64 `path:"userId"`
}

type UserRetrieveResp struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Status   int64  `json:"status"`
	Email    string `json:"email"`
	Sex      int64  `json:"sex"`
	Remark   string `json:"remark"`
	RoleId   int64  `json:"roleId"`
	DeptId   int64  `json:"deptId"`
	PostId   int64  `json:"postId"`
}

type UserUpdateReq struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Status   int64  `json:"status"`
	Email    string `json:"email"`
	Sex      int64  `json:"sex,optional"`
	Remark   string `json:"remark,optional"`
	RoleId   int64  `json:"roleId,optional"`
	DeptId   int64  `json:"deptId"`
	PostId   int64  `json:"postId,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type UserUpdateStatusReq struct {
	UserId   int64 `json:"userId"`
	Status   int64 `json:"status"`
	UpdateBy int64 `json:"updateBy,optional"`
}

type UserUpdatePwdReq struct {
	UserId   int64  `json:"userId"`
	Password string `json:"password"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type UserDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type RoleListReq struct {
	PageReq
}

type RoleListData struct {
	RoleId    int64  `json:"roleId"`
	RoleName  string `json:"roleName"`
	RoleKey   string `json:"roleKey"`
	Sort      int64  `json:"sort"`
	Status    int64  `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type RoleListResp struct {
	List []RoleListData `json:"list"`
	Pagination
}

type RoleRetrieveReq struct {
	RoleId int64 `path:"roleId"`
}

type RoleRetrieveResp struct {
	RoleId   int64   `json:"roleId"`
	RoleName string  `json:"roleName"`
	RoleKey  string  `json:"roleKey"`
	Sort     int64   `json:"sort"`
	Status   int64   `json:"status"`
	Remark   string  `json:"remark"`
	MenuIds  []int64 `json:"menuIds"`
}

type RoleAddReq struct {
	RoleName string  `json:"roleName"`
	RoleKey  string  `json:"roleKey"`
	Sort     int64   `json:"sort"`
	Status   int64   `json:"status,optional"`
	Remark   string  `json:"remark,optional"`
	CreateBy int64   `json:"createBy,optional"`
	UpdateBy int64   `json:"updateBy,optional"`
	MenuIds  []int64 `json:"menuIds,optional"`
}

type RoleUpdateReq struct {
	RoleId   int64   `json:"roleId"`
	RoleName string  `json:"roleName,optional"`
	RoleKey  string  `json:"roleKey,optional"`
	Sort     int64   `json:"sort,optional"`
	Status   int64   `json:"status,optional"`
	Remark   string  `json:"remark,optional"`
	UpdateBy int64   `json:"updateBy,optional"`
	MenuIds  []int64 `json:"menuIds,optional"`
}

type RoleDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type RoleMenuTreeReq struct {
	RoleId int64 `path:"roleId"`
}

type RoleMenuTreeData struct {
	Id       int64              `json:"id"`
	Label    string             `json:"label"`
	Children []RoleMenuTreeData `json:"children"`
}

type RoleMenuTreeResp struct {
	Menus       []RoleMenuTreeData `json:"menus"`
	CheckedKeys []int64            `json:"checkedKeys"`
}

type DeptTreeResp struct {
	Id       int64          `json:"deptId"`
	Label    string         `json:"label"`
	Children []DeptTreeResp `json:"children"`
}

type DeptListData struct {
	DeptId    int64  `json:"deptId"`
	DeptPath  string `json:"deptPath"`
	DeptName  string `json:"deptName"`
	Phone     string `json:"phone"`
	Status    int64  `json:"status"`
	Email     string `json:"email"`
	Leader    string `json:"leader"`
	Sort      int64  `json:"sort"`
	ParentId  int64  `json:"parentId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type DeptListResp struct {
	DeptId    int64          `json:"deptId"`
	DeptPath  string         `json:"deptPath"`
	DeptName  string         `json:"deptName"`
	Phone     string         `json:"phone"`
	Status    int64          `json:"status"`
	Email     string         `json:"email"`
	Leader    string         `json:"leader"`
	Sort      int64          `json:"sort"`
	ParentId  int64          `json:"parentId"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	CreateBy  int64          `json:"createBy"`
	UpdateBy  int64          `json:"updateBy"`
	Children  []DeptListResp `json:"children"`
}

type DeptRetrieveReq struct {
	DeptId int64 `path:"deptId"`
}

type DeptRetrieveResp struct {
	DeptId   int64  `json:"deptId"`
	DeptName string `json:"deptName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Leader   string `json:"leader"`
	Sort     int64  `json:"sort"`
	Status   int64  `json:"status"`
	ParentId int64  `json:"parentId"`
}

type DeptAddReq struct {
	DeptName string `json:"deptName"`
	Phone    string `json:"phone,optional"`
	Email    string `json:"email,optional"`
	Leader   string `json:"leader"`
	Sort     int64  `json:"sort,optional"`
	Status   int64  `json:"status,optional"`
	ParentId int64  `json:"parentId"`
	CreateBy int64  `json:"createBy,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type DeptUpdateReq struct {
	DeptId   int64  `json:"deptId"`
	DeptName string `json:"deptName"`
	Phone    string `json:"phone,optional"`
	Email    string `json:"email,optional"`
	Leader   string `json:"leader"`
	Sort     int64  `json:"sort,optional"`
	Status   int64  `json:"status,optional"`
	ParentId int64  `json:"parentId"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type DeptDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type MenuRoleResp struct {
	MenuId     int64          `json:"menuId"`
	MenuName   string         `json:"menuName"`
	MenuType   string         `json:"menuType"`
	Title      string         `json:"title"`
	Permission string         `json:"permission"`
	Params     string         `json:"params"`
	Path       string         `json:"path"`
	Paths      string         `json:"paths"`
	Action     string         `json:"action"`
	Apis       string         `json:"apis"`
	SysApi     string         `json:"sysApi"`
	Breadcrumb string         `json:"breadcrumb"`
	Component  string         `json:"component"`
	ParentId   int64          `json:"parentId"`
	Sort       int64          `json:"sort"`
	DataScope  string         `json:"dataScope"`
	Icon       string         `json:"icon"`
	IsFrame    string         `json:"isFrame"`
	Visible    string         `json:"visible"`
	Is_select  bool           `json:"is_select"`
	NoCache    int64          `json:"noCache"`
	CreateBy   int64          `json:"createBy"`
	CreatedAt  string         `json:"createdAt"`
	UpdateBy   int64          `json:"updateBy"`
	UpdatedAt  string         `json:"updatedAt"`
	Children   []MenuRoleResp `json:"children"`
}

type MenuListData struct {
	MenuId     int64  `json:"menuId"`
	MenuName   string `json:"menuName"`
	MenuType   string `json:"menuType"`
	Title      string `json:"title"`
	Permission string `json:"permission"`
	Params     string `json:"params"`
	Path       string `json:"path"`
	Paths      string `json:"paths"`
	Action     string `json:"action"`
	Apis       string `json:"apis"`
	SysApi     string `json:"sysApi"`
	Breadcrumb string `json:"breadcrumb"`
	Component  string `json:"component"`
	ParentId   int64  `json:"parentId"`
	Sort       int64  `json:"sort"`
	DataScope  string `json:"dataScope"`
	Icon       string `json:"icon"`
	IsFrame    string `json:"isFrame"`
	Visible    string `json:"visible"`
	Is_select  bool   `json:"is_select"`
	NoCache    int64  `json:"noCache"`
	CreateBy   int64  `json:"createBy"`
	CreatedAt  string `json:"createdAt"`
	UpdateBy   int64  `json:"updateBy"`
	UpdatedAt  string `json:"updatedAt"`
}

type MenuListResp struct {
	MenuId     int64          `json:"menuId"`
	MenuName   string         `json:"menuName"`
	MenuType   string         `json:"menuType"`
	Title      string         `json:"title"`
	Permission string         `json:"permission"`
	Params     string         `json:"params"`
	Path       string         `json:"path"`
	Paths      string         `json:"paths"`
	Action     string         `json:"action"`
	Apis       string         `json:"apis"`
	SysApi     string         `json:"sysApi"`
	Breadcrumb string         `json:"breadcrumb"`
	Component  string         `json:"component"`
	ParentId   int64          `json:"parentId"`
	Sort       int64          `json:"sort"`
	DataScope  string         `json:"dataScope"`
	Icon       string         `json:"icon"`
	IsFrame    string         `json:"isFrame"`
	Visible    string         `json:"visible"`
	Is_select  bool           `json:"is_select"`
	NoCache    int64          `json:"noCache"`
	CreateBy   int64          `json:"createBy"`
	CreatedAt  string         `json:"createdAt"`
	UpdateBy   int64          `json:"updateBy"`
	UpdatedAt  string         `json:"updatedAt"`
	Children   []MenuListResp `json:"children"`
}

type MenuRetrieveReq struct {
	MenuId int64 `path:"menuId"`
}

type MenuRetrieveResp struct {
	MenuId     int64  `json:"menuId"`
	MenuName   string `json:"menuName"`
	MenuType   string `json:"menuType"`
	Title      string `json:"title"`
	Permission string `json:"permission"`
	Params     string `json:"params"`
	Path       string `json:"path"`
	Paths      string `json:"paths"`
	Action     string `json:"action"`
	Apis       string `json:"apis"`
	SysApi     string `json:"sysApi"`
	Breadcrumb string `json:"breadcrumb"`
	Component  string `json:"component"`
	ParentId   int64  `json:"parentId"`
	Sort       int64  `json:"sort"`
	DataScope  string `json:"dataScope"`
	Icon       string `json:"icon"`
	IsFrame    string `json:"isFrame"`
	Visible    string `json:"visible"`
	Is_select  bool   `json:"is_select"`
	NoCache    int64  `json:"noCache"`
}

type MenuAddReq struct {
	MenuName   string `json:"menuName"`
	MenuType   string `json:"menuType"`
	Title      string `json:"title"`
	Permission string `json:"permission"`
	Params     string `json:"params,optional"`
	Path       string `json:"path,optional"`
	Paths      string `json:"paths,optional"`
	Action     string `json:"action,optional"`
	Breadcrumb string `json:"breadcrumb,optional"`
	Component  string `json:"component,optional"`
	ParentId   int64  `json:"parentId,optional"`
	Sort       int64  `json:"sort"`
	Icon       string `json:"icon"`
	IsFrame    string `json:"isFrame,optional"`
	Visible    string `json:"visible"`
	NoCache    int64  `json:"noCache,optional"`
	CreateBy   int64  `json:"createBy,optional"`
	UpdateBy   int64  `json:"updateBy,optional"`
}

type MenuUpdateReq struct {
	MenuId     int64  `json:"menuId"`
	MenuName   string `json:"menuName"`
	MenuType   string `json:"menuType"`
	Title      string `json:"title"`
	Permission string `json:"permission"`
	Params     string `json:"params"`
	Path       string `json:"path"`
	Paths      string `json:"paths"`
	Action     string `json:"action"`
	Apis       string `json:"apis"`
	SysApi     string `json:"sysApi"`
	Breadcrumb string `json:"breadcrumb"`
	Component  string `json:"component"`
	ParentId   int64  `json:"parentId"`
	Sort       int64  `json:"sort"`
	DataScope  string `json:"dataScope"`
	Icon       string `json:"icon"`
	IsFrame    string `json:"isFrame"`
	Visible    string `json:"visible"`
	Is_select  bool   `json:"is_select"`
	NoCache    int64  `json:"noCache"`
	UpdateBy   int64  `json:"updateBy,optional"`
}

type MenuDeleteReq struct {
	MenuId int64 `path:"menuId"`
}

type PostListReq struct {
	PageReq
}

type PostListData struct {
	PostId    int64  `json:"postId"`
	PostName  string `json:"postName"`
	PostCode  string `json:"postCode"`
	Sort      int64  `json:"sort"`
	Status    int64  `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreateBy  int64  `json:"createBy"`
	UpdateBy  int64  `json:"updateBy"`
}

type PostListResp struct {
	List []PostListData `json:"list"`
	Pagination
}

type PostRetrieveReq struct {
	PostId int64 `path:"postId"`
}

type PostRetrieveResp struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int64  `json:"sort"`
	Status   int64  `json:"status"`
	Remark   string `json:"remark"`
}

type PostAddReq struct {
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int64  `json:"sort"`
	Status   int64  `json:"status,optional"`
	Remark   string `json:"remark,optional"`
	CreateBy int64  `json:"createBy,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type PostUpdateReq struct {
	PostId   int64  `json:"postId"`
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int64  `json:"sort"`
	Status   int64  `json:"status,optional"`
	Remark   string `json:"remark,optional"`
	UpdateBy int64  `json:"updateBy,optional"`
}

type PostDeleteReq struct {
	Ids []int64 `json:"ids"`
}
