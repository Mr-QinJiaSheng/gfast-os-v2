// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/dao/internal"
)

// sysPostDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysPostDao struct {
	*internal.SysPostDao
}

var (
	// SysPost is globally public accessible object for table sys_post operations.
	SysPost sysPostDao
)

func init() {
	SysPost = sysPostDao{
		internal.NewSysPostDao(),
	}
}

// Fill with you ideas below.

// SysPostSearchParams 搜索参数
type SysPostSearchParams struct {
	PostCode string `p:"postCode"` //岗位编码
	PostName string `p:"postName"` //岗位名称
	Status   string `p:"status"`   //状态
	comModel.PageReq
}

// SysPostAddParams 添加岗位参数
type SysPostAddParams struct {
	PostCode  string `p:"postCode" v:"required#岗位编码不能为空"`
	PostName  string `p:"postName" v:"required#岗位名称不能为空"`
	PostSort  int    `p:"postSort" v:"required#岗位排序不能为空"`
	Status    string `p:"status" v:"required#状态不能为空"`
	Remark    string `p:"remark"`
	CreatedBy uint64
}

// SysPostEditParams 修改岗位参数
type SysPostEditParams struct {
	PostId int64 `p:"postId" v:"required#id必须"`
	SysPostAddParams
	UpdatedBy uint64
}
