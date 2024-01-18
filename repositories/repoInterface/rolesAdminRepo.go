package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type RolesAdminRepo interface {
	repoComm.RepoInterface
	GetTotalCount(where RolesAdminSearchWhere) int64
	GetList(where RolesAdminSearchWhere) []model.RolesAdmin
	Delete(_model model.RolesAdmin) (rowsAffected int64, err error)
	DeleteByID(ID ...[]uint8) (rowsAffected int64, err error)
	Save(_model *model.RolesAdmin, _select ...string) error
	SaveOmit(_model *model.RolesAdmin, _omit ...string) error
	Create(_model *[]model.RolesAdmin) error
	GetByID(ID []uint8, _select ...string) model.RolesAdmin
	GetByIDLock(ID []uint8, _select ...string) (model.RolesAdmin, repoComm.ReleaseLock)
	GetByWhere(where RolesAdminSearchWhere) model.RolesAdmin
	GetIDByWhere(where RolesAdminSearchWhere) [][]uint8
	GetByAdminID(adminID uint64, _select ...string) []model.RolesAdmin
	GetByRoleID(roleID uint64, _select ...string) []model.RolesAdmin
	DeleteByAdminID(adminID uint64) (rowsAffected int64, err error)
	DeleteByRoleID(roleID uint64) (rowsAffected int64, err error)
	UpdateByWhere(where RolesAdminSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where RolesAdminSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where RolesAdminSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...RolesAdminSearchWhere) error
	SaveByAdm(adm model.Admin) error //当adm.RolesID == ''时 应当清空对应的数据
}

type RolesAdminSearchWhere struct {
	ID               string
	IDNeq            string //不等于条件
	IDNull           bool
	IDNotNull        bool
	IDLt             string    // ID < IDLt
	IDGt             string    // ID > IDGt
	IDElt            string    // ID <= IDElt
	IDEgt            string    // ID >= IDEgt
	IDNotIn          [][]uint8 // not in查询
	IDIn             [][]uint8 // in查询
	IDSort           string    // 排序
	RoleID           string
	RoleIDNeq        string //不等于条件
	RoleIDNull       bool
	RoleIDNotNull    bool
	RoleIDLt         string   // RoleID < RoleIDLt
	RoleIDGt         string   // RoleID > RoleIDGt
	RoleIDElt        string   // RoleID <= RoleIDElt
	RoleIDEgt        string   // RoleID >= RoleIDEgt
	RoleIDNotIn      []uint64 // not in查询
	RoleIDIn         []uint64 // in查询
	RoleIDSort       string   // 排序
	AdminID          string
	AdminIDNeq       string //不等于条件
	AdminIDNull      bool
	AdminIDNotNull   bool
	AdminIDLt        string   // AdminID < AdminIDLt
	AdminIDGt        string   // AdminID > AdminIDGt
	AdminIDElt       string   // AdminID <= AdminIDElt
	AdminIDEgt       string   // AdminID >= AdminIDEgt
	AdminIDNotIn     []uint64 // not in查询
	AdminIDIn        []uint64 // in查询
	AdminIDSort      string   // 排序
	CreatedAt        string
	CreatedAtNeq     string //不等于条件
	CreatedAtNull    bool
	CreatedAtNotNull bool
	CreatedAtLt      string  // CreatedAt < CreatedAtLt
	CreatedAtGt      string  // CreatedAt > CreatedAtGt
	CreatedAtElt     string  // CreatedAt <= CreatedAtElt
	CreatedAtEgt     string  // CreatedAt >= CreatedAtEgt
	CreatedAtNotIn   []int64 // not in查询
	CreatedAtIn      []int64 // in查询
	CreatedAtSort    string  // 排序
	UpdatedAt        string
	UpdatedAtNeq     string //不等于条件
	UpdatedAtNull    bool
	UpdatedAtNotNull bool
	UpdatedAtLt      string  // UpdatedAt < UpdatedAtLt
	UpdatedAtGt      string  // UpdatedAt > UpdatedAtGt
	UpdatedAtElt     string  // UpdatedAt <= UpdatedAtElt
	UpdatedAtEgt     string  // UpdatedAt >= UpdatedAtEgt
	UpdatedAtNotIn   []int64 // not in查询
	UpdatedAtIn      []int64 // in查询
	UpdatedAtSort    string  // 排序
	SelectParams     repoComm.SelectFrom
}
