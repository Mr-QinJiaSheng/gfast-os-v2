// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// SysModelInfo is the golang structure for table sys_model_info.
type SysModelInfo struct {
	ModelId         uint   `orm:"model_id,primary"  json:"modelId"`         // 模型ID
	ModelCategoryId uint   `orm:"model_category_id" json:"modelCategoryId"` // 模板分类id
	ModelName       string `orm:"model_name,unique" json:"modelName"`       // 模型标识
	ModelTitle      string `orm:"model_title"       json:"modelTitle"`      // 模型名称
	ModelPk         string `orm:"model_pk"          json:"modelPk"`         // 主键字段
	ModelOrder      string `orm:"model_order"       json:"modelOrder"`      // 默认排序字段
	ModelSort       string `orm:"model_sort"        json:"modelSort"`       // 表单字段排序
	ModelList       string `orm:"model_list"        json:"modelList"`       // 列表显示字段，为空显示全部
	ModelEdit       string `orm:"model_edit"        json:"modelEdit"`       // 可编辑字段，为空则除主键外均可以编辑
	ModelIndexes    string `orm:"model_indexes"     json:"modelIndexes"`    // 索引字段
	SearchList      string `orm:"search_list"       json:"searchList"`      // 高级搜索的字段
	CreateTime      uint64 `orm:"create_time"       json:"createTime"`      // 创建时间
	UpdateTime      uint64 `orm:"update_time"       json:"updateTime"`      // 更新时间
	ModelStatus     uint   `orm:"model_status"      json:"modelStatus"`     // 状态
	ModelEngine     string `orm:"model_engine"      json:"modelEngine"`     // 数据库引擎
	CreateBy        uint64 `orm:"create_by"         json:"createBy"`        // 创建人
	UpdateBy        uint64 `orm:"update_by"         json:"updateBy"`        // 修改人
}
