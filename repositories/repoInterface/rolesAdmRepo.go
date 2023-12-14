package repoInterface

import (
	"iris-init/model"
	"iris-init/repositories/repoComm"
)

type RolesAdmRepo interface {
	repoComm.RepoInterface
	SaveByAdm(adm model.Admin) error //当adm.RolesID == ''时 应当清空对应的数据
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	DeleteByRoleID(roleID ...uint64) (rowsAffected int64, err error)
	DeleteByAdmID(admID ...uint64) (rowsAffected int64, err error)
}
