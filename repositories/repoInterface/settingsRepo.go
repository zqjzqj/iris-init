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
	DeleteByID(id ...uint64) (rowsAffected int64, err error)
	Save(_model *model.Settings, _select ...string) error
	SaveOmit(_model *model.Settings, _omit ...string) error
	Create(_model *[]model.Settings) error
	GetByID(id uint64, _select ...string) model.Settings
	GetByIDLock(id uint64, _select ...string) (model.Settings, repoComm.ReleaseLock)
	GetByWhere(where SettingsSearchWhere) model.Settings
	GetIDByWhere(where SettingsSearchWhere) []uint64
	GetByKey(key string, _select ...string) model.Settings
	DeleteByKey(key string) (rowsAffected int64, err error)
	UpdateByWhere(where SettingsSearchWhere, data interface{}) (rowsAffected int64, err error)
}

type SettingsSearchWhere struct {
	ID            string
	IDLt          string // ID < IDLt
	IDGt          string // ID > IDGt
	IDElt         string // ID <= IDElt
	IDEgt         string // ID >= IDEgt
	IDSort        string // 排序
	Key           string
	KeyLike       string
	Name          string
	NameLike      string
	Desc          string
	DescLike      string
	Value         string
	ValueLike     string
	CreatedAt     string
	CreatedAtLt   string // CreatedAt < CreatedAtLt
	CreatedAtGt   string // CreatedAt > CreatedAtGt
	CreatedAtElt  string // CreatedAt <= CreatedAtElt
	CreatedAtEgt  string // CreatedAt >= CreatedAtEgt
	CreatedAtSort string // 排序
	UpdatedAt     string
	UpdatedAtLt   string // UpdatedAt < UpdatedAtLt
	UpdatedAtGt   string // UpdatedAt > UpdatedAtGt
	UpdatedAtElt  string // UpdatedAt <= UpdatedAtElt
	UpdatedAtEgt  string // UpdatedAt >= UpdatedAtEgt
	UpdatedAtSort string // 排序
	SelectParams  repoComm.SelectFrom
}
