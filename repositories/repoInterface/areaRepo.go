package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type AreaRepo interface {
	repoComm.RepoInterface
	GetTotalCount(where AreaSearchWhere) int64
	GetList(where AreaSearchWhere) []model.Area
	Delete(_model model.Area) (rowsAffected int64, err error)
	DeleteByID(ID ...uint64) (rowsAffected int64, err error)
	Save(_model *model.Area, _select ...string) error
	SaveOmit(_model *model.Area, _omit ...string) error
	Create(_model *[]model.Area) error
	GetByID(ID uint64, _select ...string) model.Area
	GetByIDLock(ID uint64, _select ...string) (model.Area, repoComm.ReleaseLock)
	GetByWhere(where AreaSearchWhere) model.Area
	GetIDByWhere(where AreaSearchWhere) []uint64
	GetByFirst(first string, _select ...string) []model.Area
	GetByLevel(level uint8, _select ...string) []model.Area
	GetByPid(pid uint, _select ...string) []model.Area
	DeleteByFirst(first string) (rowsAffected int64, err error)
	DeleteByLevel(level uint8) (rowsAffected int64, err error)
	DeleteByPid(pid uint) (rowsAffected int64, err error)
	UpdateByWhere(where AreaSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where AreaSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where AreaSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...AreaSearchWhere) error
	GetListByPID(pid uint, _select ...string) []model.Area
}

type AreaSearchWhere struct {
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
	Pid               string
	PidNeq            string //不等于条件
	PidNull           bool
	PidNotNull        bool
	PidLt             string // Pid < PidLt
	PidGt             string // Pid > PidGt
	PidElt            string // Pid <= PidElt
	PidEgt            string // Pid >= PidEgt
	PidSort           string // 排序
	PidNotIn          []uint // not in查询
	PidIn             []uint // in查询
	ShortName         string
	ShortNameNeq      string //不等于条件
	ShortNameNull     bool
	ShortNameNotNull  bool
	ShortNameLike     string
	ShortNameNotIn    []string // not in查询
	ShortNameIn       []string // in查询
	Name              string
	NameNeq           string //不等于条件
	NameNull          bool
	NameNotNull       bool
	NameLike          string
	NameNotIn         []string // not in查询
	NameIn            []string // in查询
	MergerName        string
	MergerNameNeq     string //不等于条件
	MergerNameNull    bool
	MergerNameNotNull bool
	MergerNameLike    string
	MergerNameNotIn   []string // not in查询
	MergerNameIn      []string // in查询
	Level             string
	LevelNeq          string //不等于条件
	LevelNull         bool
	LevelNotNull      bool
	LevelLt           string // Level < LevelLt
	LevelGt           string // Level > LevelGt
	LevelElt          string // Level <= LevelElt
	LevelEgt          string // Level >= LevelEgt
	LevelSort         string // 排序
	LevelNotIn        []int  // not in查询
	LevelIn           []int  // in查询
	PinYin            string
	PinYinNeq         string //不等于条件
	PinYinNull        bool
	PinYinNotNull     bool
	PinYinLike        string
	PinYinNotIn       []string // not in查询
	PinYinIn          []string // in查询
	Code              string
	CodeNeq           string //不等于条件
	CodeNull          bool
	CodeNotNull       bool
	CodeLike          string
	CodeNotIn         []string // not in查询
	CodeIn            []string // in查询
	ZipCode           string
	ZipCodeNeq        string //不等于条件
	ZipCodeNull       bool
	ZipCodeNotNull    bool
	ZipCodeLike       string
	ZipCodeNotIn      []string // not in查询
	ZipCodeIn         []string // in查询
	First             string
	FirstNeq          string //不等于条件
	FirstNull         bool
	FirstNotNull      bool
	FirstLike         string
	FirstNotIn        []string // not in查询
	FirstIn           []string // in查询
	Lng               string
	LngNeq            string //不等于条件
	LngNull           bool
	LngNotNull        bool
	LngLike           string
	LngNotIn          []string // not in查询
	LngIn             []string // in查询
	Lat               string
	LatNeq            string //不等于条件
	LatNull           bool
	LatNotNull        bool
	LatLike           string
	LatNotIn          []string // not in查询
	LatIn             []string // in查询
	Children          string
	ChildrenNeq       string //不等于条件
	ChildrenNull      bool
	ChildrenNotNull   bool
	ChildrenNotIn     [][]model.Area // not in查询
	ChildrenIn        [][]model.Area // in查询
	SelectParams      repoComm.SelectFrom
}
