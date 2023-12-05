package repoInterface

import (
	"big_data_new/model"
	"big_data_new/repositories/repoComm"
)

type PermissionsRepo interface {
	repoComm.RepoInterface
	Save(perm *model.Permissions, _select ...string) error
	GetByIdent(ident string, _select ...string) model.Permissions
	GetByID(id uint64, _select ...string) model.Permissions
	GetListAsMenu(idents []string) []model.Permissions
	GetListPreloadChildren_2() []model.Permissions
	TruncateTable()
	GetOrCreatePermissionByName(name string, pid uint64, level uint8, sort uint) (model.Permissions, error)
	GetList(where PermissionsSearchWhere) []model.Permissions
	GetListPreloadChildren(where PermissionsSearchWhere) []model.Permissions
	GetTotalCount(where PermissionsSearchWhere) int64
	Delete(_model model.Permissions) (rowsAffected int64, err error)
	DeleteByID(ID ...uint64) (rowsAffected int64, err error)
	SaveOmit(_model *model.Permissions, _omit ...string) error
	Create(_model *[]model.Permissions) error
	GetByIDLock(ID uint64, _select ...string) (model.Permissions, repoComm.ReleaseLock)
	GetByWhere(where PermissionsSearchWhere) model.Permissions
	GetIDByWhere(where PermissionsSearchWhere) []uint64
	GetByPid_Name(pid uint64, name string, _select ...string) model.Permissions
	GetByLevel(level uint8, _select ...string) []model.Permissions
	DeleteByIdent(ident string) (rowsAffected int64, err error)
	DeleteByPid_Name(pid uint64, name string) (rowsAffected int64, err error)
	DeleteByLevel(level uint8) (rowsAffected int64, err error)
	UpdateByWhere(where PermissionsSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where PermissionsSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where PermissionsSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...PermissionsSearchWhere) error
}

type PermissionsSearchWhere struct {
	ID               string
	IDNeq            string //不等于条件
	IDNull           bool
	IDNotNull        bool
	IDLt             string   // ID < IDLt
	IDGt             string   // ID > IDGt
	IDElt            string   // ID <= IDElt
	IDEgt            string   // ID >= IDEgt
	IDNotIn          []uint64 // not in查询
	IDIn             []uint64 // in查询
	IDSort           string   // 排序
	Pid              string
	PidNeq           string //不等于条件
	PidNull          bool
	PidNotNull       bool
	PidLt            string   // Pid < PidLt
	PidGt            string   // Pid > PidGt
	PidElt           string   // Pid <= PidElt
	PidEgt           string   // Pid >= PidEgt
	PidNotIn         []uint64 // not in查询
	PidIn            []uint64 // in查询
	PidSort          string   // 排序
	Level            string
	LevelNeq         string //不等于条件
	LevelNull        bool
	LevelNotNull     bool
	LevelLt          string // Level < LevelLt
	LevelGt          string // Level > LevelGt
	LevelElt         string // Level <= LevelElt
	LevelEgt         string // Level >= LevelEgt
	LevelNotIn       []int  // not in查询
	LevelIn          []int  // in查询
	LevelSort        string // 排序
	Name             string
	NameNeq          string //不等于条件
	NameNull         bool
	NameNotNull      bool
	NameLike         string
	NameNotIn        []string // not in查询
	NameIn           []string // in查询
	NameSort         string   // 排序
	Method           string
	MethodNeq        string //不等于条件
	MethodNull       bool
	MethodNotNull    bool
	MethodLike       string
	MethodNotIn      []string // not in查询
	MethodIn         []string // in查询
	MethodSort       string   // 排序
	Path             string
	PathNeq          string //不等于条件
	PathNull         bool
	PathNotNull      bool
	PathLike         string
	PathNotIn        []string // not in查询
	PathIn           []string // in查询
	PathSort         string   // 排序
	Sort             string
	SortNeq          string //不等于条件
	SortNull         bool
	SortNotNull      bool
	SortLt           string // Sort < SortLt
	SortGt           string // Sort > SortGt
	SortElt          string // Sort <= SortElt
	SortEgt          string // Sort >= SortEgt
	SortNotIn        []uint // not in查询
	SortIn           []uint // in查询
	SortSort         string // 排序
	Ident            string
	IdentNeq         string //不等于条件
	IdentNull        bool
	IdentNotNull     bool
	IdentLike        string
	IdentNotIn       []string // not in查询
	IdentIn          []string // in查询
	IdentSort        string   // 排序
	Children         string
	ChildrenNeq      string //不等于条件
	ChildrenNull     bool
	ChildrenNotNull  bool
	ChildrenNotIn    [][]model.Permissions // not in查询
	ChildrenIn       [][]model.Permissions // in查询
	ChildrenSort     string                // 排序
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
