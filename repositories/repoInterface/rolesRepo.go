package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type RolesRepo interface {
	repoComm.RepoInterface
	repoComm.RepoInterface
	GetTotalCount(where RolesSearchWhere) int64
	GetList(where RolesSearchWhere) []model.Roles
	Delete(_model model.Roles) (rowsAffected int64, err error)
	DeleteByID(ID ...uint64) (rowsAffected int64, err error)
	Save(_model *model.Roles, _select ...string) error
	SaveOmit(_model *model.Roles, _omit ...string) error
	Create(_model *[]model.Roles) error
	GetByID(ID uint64, _select ...string) model.Roles
	GetByIDLock(ID uint64, _select ...string) (model.Roles, repoComm.ReleaseLock)
	GetByWhere(where RolesSearchWhere) model.Roles
	GetIDByWhere(where RolesSearchWhere) []uint64
	GetByName(name string, _select ...string) []model.Roles
	DeleteByName(name string) (rowsAffected int64, err error)
	UpdateByWhere(where RolesSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where RolesSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where RolesSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...RolesSearchWhere) error
	GetRolesByID(id ...uint64) []model.Roles
}

type RolesSearchWhere struct {
	ID                string
	IDNeq             string //不等于条件
	IDNull            bool
	IDNotNull         bool
	IDLt              string   // ID < IDLt
	IDGt              string   // ID > IDGt
	IDElt             string   // ID <= IDElt
	IDEgt             string   // ID >= IDEgt
	IDSort            string   // 排序
	IDNotIn           []uint64 // not in查询
	IDIn              []uint64 // in查询
	Name              string
	NameNeq           string //不等于条件
	NameNull          bool
	NameNotNull       bool
	NameLike          string
	NameNotIn         []string // not in查询
	NameIn            []string // in查询
	Remark            string
	RemarkNeq         string //不等于条件
	RemarkNull        bool
	RemarkNotNull     bool
	RemarkLike        string
	RemarkNotIn       []string // not in查询
	RemarkIn          []string // in查询
	PermIdents        string
	PermIdentsNeq     string //不等于条件
	PermIdentsNull    bool
	PermIdentsNotNull bool
	PermIdentsNotIn   [][]string // not in查询
	PermIdentsIn      [][]string // in查询
	CreatedAt         string
	CreatedAtNeq      string //不等于条件
	CreatedAtNull     bool
	CreatedAtNotNull  bool
	CreatedAtLt       string  // CreatedAt < CreatedAtLt
	CreatedAtGt       string  // CreatedAt > CreatedAtGt
	CreatedAtElt      string  // CreatedAt <= CreatedAtElt
	CreatedAtEgt      string  // CreatedAt >= CreatedAtEgt
	CreatedAtSort     string  // 排序
	CreatedAtNotIn    []int64 // not in查询
	CreatedAtIn       []int64 // in查询
	UpdatedAt         string
	UpdatedAtNeq      string //不等于条件
	UpdatedAtNull     bool
	UpdatedAtNotNull  bool
	UpdatedAtLt       string  // UpdatedAt < UpdatedAtLt
	UpdatedAtGt       string  // UpdatedAt > UpdatedAtGt
	UpdatedAtElt      string  // UpdatedAt <= UpdatedAtElt
	UpdatedAtEgt      string  // UpdatedAt >= UpdatedAtEgt
	UpdatedAtSort     string  // 排序
	UpdatedAtNotIn    []int64 // not in查询
	UpdatedAtIn       []int64 // in查询
	SelectParams      repoComm.SelectFrom
}
