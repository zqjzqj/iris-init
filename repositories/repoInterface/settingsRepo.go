package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type SettingsRepo interface {
	repoComm.RepoInterface
	GetTotalCount(where SettingsSearchWhere) int64
	GetList(where SettingsSearchWhere) []model.Settings
	Delete(_model model.Settings) (rowsAffected int64, err error)
	DeleteByID(ID ...uint64) (rowsAffected int64, err error)
	Save(_model *model.Settings, _select ...string) error
	SaveOmit(_model *model.Settings, _omit ...string) error
	Create(_model *[]model.Settings) error
	GetByID(ID uint64, _select ...string) model.Settings
	GetByIDLock(ID uint64, _select ...string) (model.Settings, repoComm.ReleaseLock)
	GetByWhere(where SettingsSearchWhere) model.Settings
	GetIDByWhere(where SettingsSearchWhere) []uint64
	GetByKey(key string, _select ...string) model.Settings
	DeleteByKey(key string) (rowsAffected int64, err error)
	UpdateByWhere(where SettingsSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where SettingsSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where SettingsSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...SettingsSearchWhere) error
}

type SettingsSearchWhere struct {
	ID               string
	IDNeq            string //不等于条件
	IDNull           bool
	IDNotNull        bool
	IDLt             string   // ID < IDLt
	IDGt             string   // ID > IDGt
	IDElt            string   // ID <= IDElt
	IDEgt            string   // ID >= IDEgt
	IDSort           string   // 排序
	IDNotIn          []uint64 // not in查询
	IDIn             []uint64 // in查询
	Key              string
	KeyNeq           string //不等于条件
	KeyNull          bool
	KeyNotNull       bool
	KeyLike          string
	KeyNotIn         []string // not in查询
	KeyIn            []string // in查询
	Name             string
	NameNeq          string //不等于条件
	NameNull         bool
	NameNotNull      bool
	NameLike         string
	NameNotIn        []string // not in查询
	NameIn           []string // in查询
	Desc             string
	DescNeq          string //不等于条件
	DescNull         bool
	DescNotNull      bool
	DescLike         string
	DescNotIn        []string // not in查询
	DescIn           []string // in查询
	Value            string
	ValueNeq         string //不等于条件
	ValueNull        bool
	ValueNotNull     bool
	ValueLike        string
	ValueNotIn       []string // not in查询
	ValueIn          []string // in查询
	InputType        string
	InputTypeNeq     string //不等于条件
	InputTypeNull    bool
	InputTypeNotNull bool
	InputTypeLike    string
	InputTypeNotIn   []string // not in查询
	InputTypeIn      []string // in查询
	CreatedAt        string
	CreatedAtNeq     string //不等于条件
	CreatedAtNull    bool
	CreatedAtNotNull bool
	CreatedAtLt      string  // CreatedAt < CreatedAtLt
	CreatedAtGt      string  // CreatedAt > CreatedAtGt
	CreatedAtElt     string  // CreatedAt <= CreatedAtElt
	CreatedAtEgt     string  // CreatedAt >= CreatedAtEgt
	CreatedAtSort    string  // 排序
	CreatedAtNotIn   []int64 // not in查询
	CreatedAtIn      []int64 // in查询
	UpdatedAt        string
	UpdatedAtNeq     string //不等于条件
	UpdatedAtNull    bool
	UpdatedAtNotNull bool
	UpdatedAtLt      string  // UpdatedAt < UpdatedAtLt
	UpdatedAtGt      string  // UpdatedAt > UpdatedAtGt
	UpdatedAtElt     string  // UpdatedAt <= UpdatedAtElt
	UpdatedAtEgt     string  // UpdatedAt >= UpdatedAtEgt
	UpdatedAtSort    string  // 排序
	UpdatedAtNotIn   []int64 // not in查询
	UpdatedAtIn      []int64 // in查询
	SelectParams     repoComm.SelectFrom
}
