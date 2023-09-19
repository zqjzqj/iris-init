package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type AdminRepo interface {
	repoComm.RepoInterface
	GetTotalCount(where AdminSearchWhere) int64
	GetList(where AdminSearchWhere) []model.Admin
	Delete(_model model.Admin) (rowsAffected int64, err error)
	DeleteByID(ID ...uint64) (rowsAffected int64, err error)
	Save(_model *model.Admin, _select ...string) error
	SaveOmit(_model *model.Admin, _omit ...string) error
	Create(_model *[]model.Admin) error
	GetByID(ID uint64, _select ...string) model.Admin
	GetByIDLock(ID uint64, _select ...string) (model.Admin, repoComm.ReleaseLock)
	GetByWhere(where AdminSearchWhere) model.Admin
	GetIDByWhere(where AdminSearchWhere) []uint64
	GetByPhone(phone string, _select ...string) model.Admin
	GetByToken(token string, _select ...string) model.Admin
	GetByUsername(username string, _select ...string) model.Admin
	GetByRealName(realName string, _select ...string) []model.Admin
	GetByStatus(status uint8, _select ...string) []model.Admin
	DeleteByPhone(phone string) (rowsAffected int64, err error)
	DeleteByToken(token string) (rowsAffected int64, err error)
	DeleteByUsername(username string) (rowsAffected int64, err error)
	DeleteByRealName(realName string) (rowsAffected int64, err error)
	DeleteByStatus(status uint8) (rowsAffected int64, err error)
	UpdateByWhere(where AdminSearchWhere, data interface{}) (rowsAffected int64, err error)
	DeleteByWhere(where AdminSearchWhere) (rowsAffected int64, err error)
	ScanByWhere(where AdminSearchWhere, dest any) error
	ScanByOrWhere(dest any, where ...AdminSearchWhere) error
}

type AdminSearchWhere struct {
	ID                   string
	IDNeq                string //不等于条件
	IDNull               bool
	IDNotNull            bool
	IDLt                 string   // ID < IDLt
	IDGt                 string   // ID > IDGt
	IDElt                string   // ID <= IDElt
	IDEgt                string   // ID >= IDEgt
	IDSort               string   // 排序
	IDNotIn              []uint64 // not in查询
	IDIn                 []uint64 // in查询
	Username             string
	UsernameNeq          string //不等于条件
	UsernameNull         bool
	UsernameNotNull      bool
	UsernameLike         string
	UsernameNotIn        []string // not in查询
	UsernameIn           []string // in查询
	Phone                string
	PhoneNeq             string //不等于条件
	PhoneNull            bool
	PhoneNotNull         bool
	PhoneLike            string
	PhoneNotIn           []string // not in查询
	PhoneIn              []string // in查询
	QQ                   string
	QQNeq                string //不等于条件
	QQNull               bool
	QQNotNull            bool
	QQLike               string
	QQNotIn              []string // not in查询
	QQIn                 []string // in查询
	Status               string
	StatusNeq            string //不等于条件
	StatusNull           bool
	StatusNotNull        bool
	StatusLt             string // Status < StatusLt
	StatusGt             string // Status > StatusGt
	StatusElt            string // Status <= StatusElt
	StatusEgt            string // Status >= StatusEgt
	StatusSort           string // 排序
	StatusNotIn          []int  // not in查询
	StatusIn             []int  // in查询
	RealName             string
	RealNameNeq          string //不等于条件
	RealNameNull         bool
	RealNameNotNull      bool
	RealNameLike         string
	RealNameNotIn        []string // not in查询
	RealNameIn           []string // in查询
	Avatar               string
	AvatarNeq            string //不等于条件
	AvatarNull           bool
	AvatarNotNull        bool
	AvatarLike           string
	AvatarNotIn          []string // not in查询
	AvatarIn             []string // in查询
	Sex                  string
	SexNeq               string //不等于条件
	SexNull              bool
	SexNotNull           bool
	SexLt                string // Sex < SexLt
	SexGt                string // Sex > SexGt
	SexElt               string // Sex <= SexElt
	SexEgt               string // Sex >= SexEgt
	SexSort              string // 排序
	SexNotIn             []int  // not in查询
	SexIn                []int  // in查询
	Password             string
	PasswordNeq          string //不等于条件
	PasswordNull         bool
	PasswordNotNull      bool
	PasswordLike         string
	PasswordNotIn        []string // not in查询
	PasswordIn           []string // in查询
	Salt                 string
	SaltNeq              string //不等于条件
	SaltNull             bool
	SaltNotNull          bool
	SaltLike             string
	SaltNotIn            []string // not in查询
	SaltIn               []string // in查询
	Token                string
	TokenNeq             string //不等于条件
	TokenNull            bool
	TokenNotNull         bool
	TokenLike            string
	TokenNotIn           []string // not in查询
	TokenIn              []string // in查询
	TokenStatus          string
	TokenStatusNeq       string //不等于条件
	TokenStatusNull      bool
	TokenStatusNotNull   bool
	TokenStatusLt        string // TokenStatus < TokenStatusLt
	TokenStatusGt        string // TokenStatus > TokenStatusGt
	TokenStatusElt       string // TokenStatus <= TokenStatusElt
	TokenStatusEgt       string // TokenStatus >= TokenStatusEgt
	TokenStatusSort      string // 排序
	TokenStatusNotIn     []int  // not in查询
	TokenStatusIn        []int  // in查询
	LastLoginTime        string
	LastLoginTimeNeq     string //不等于条件
	LastLoginTimeNull    bool
	LastLoginTimeNotNull bool
	LastLoginTimeLt      string  // LastLoginTime < LastLoginTimeLt
	LastLoginTimeGt      string  // LastLoginTime > LastLoginTimeGt
	LastLoginTimeElt     string  // LastLoginTime <= LastLoginTimeElt
	LastLoginTimeEgt     string  // LastLoginTime >= LastLoginTimeEgt
	LastLoginTimeSort    string  // 排序
	LastLoginTimeNotIn   []int64 // not in查询
	LastLoginTimeIn      []int64 // in查询
	RolesID              string
	RolesIDNeq           string //不等于条件
	RolesIDNull          bool
	RolesIDNotNull       bool
	RolesIDLike          string
	RolesIDNotIn         []string // not in查询
	RolesIDIn            []string // in查询
	Desc                 string
	DescNeq              string //不等于条件
	DescNull             bool
	DescNotNull          bool
	DescLike             string
	DescNotIn            []string // not in查询
	DescIn               []string // in查询
	CreatedAt            string
	CreatedAtNeq         string //不等于条件
	CreatedAtNull        bool
	CreatedAtNotNull     bool
	CreatedAtLt          string  // CreatedAt < CreatedAtLt
	CreatedAtGt          string  // CreatedAt > CreatedAtGt
	CreatedAtElt         string  // CreatedAt <= CreatedAtElt
	CreatedAtEgt         string  // CreatedAt >= CreatedAtEgt
	CreatedAtSort        string  // 排序
	CreatedAtNotIn       []int64 // not in查询
	CreatedAtIn          []int64 // in查询
	UpdatedAt            string
	UpdatedAtNeq         string //不等于条件
	UpdatedAtNull        bool
	UpdatedAtNotNull     bool
	UpdatedAtLt          string  // UpdatedAt < UpdatedAtLt
	UpdatedAtGt          string  // UpdatedAt > UpdatedAtGt
	UpdatedAtElt         string  // UpdatedAt <= UpdatedAtElt
	UpdatedAtEgt         string  // UpdatedAt >= UpdatedAtEgt
	UpdatedAtSort        string  // 排序
	UpdatedAtNotIn       []int64 // not in查询
	UpdatedAtIn          []int64 // in查询
	RolesIDSlices        string
	RolesIDSlicesNeq     string //不等于条件
	RolesIDSlicesNull    bool
	RolesIDSlicesNotNull bool
	RolesIDSlicesNotIn   [][]string // not in查询
	RolesIDSlicesIn      [][]string // in查询
	Permissions          string
	PermissionsNeq       string //不等于条件
	PermissionsNull      bool
	PermissionsNotNull   bool
	PermissionsNotIn     [][]string // not in查询
	PermissionsIn        [][]string // in查询
	RolesName            string
	RolesNameNeq         string //不等于条件
	RolesNameNull        bool
	RolesNameNotNull     bool
	RolesNameNotIn       [][]string // not in查询
	RolesNameIn          [][]string // in查询
	SelectParams         repoComm.SelectFrom
}
