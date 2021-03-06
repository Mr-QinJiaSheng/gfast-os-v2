// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysRoleDeptDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysRoleDeptDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns sysRoleDeptColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysRoleDeptColumns defines and stores column names for table sys_role_dept.
type sysRoleDeptColumns struct {
	RoleId string // 角色ID
	DeptId string // 部门ID
}

var (
	// SysRoleDept is globally public accessible object for table sys_role_dept operations.
	SysRoleDept = SysRoleDeptDao{
		M:     g.DB("default").Model("sys_role_dept").Safe(),
		DB:    g.DB("default"),
		Table: "sys_role_dept",
		Columns: sysRoleDeptColumns{
			RoleId: "role_id",
			DeptId: "dept_id",
		},
	}
)
