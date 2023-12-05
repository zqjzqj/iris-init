package repoInterface

import (
	"big_data_new/model"
	"big_data_new/repositories/repoComm"
)

type RolesAdmRepo interface {
	repoComm.RepoInterface
	SaveByAdm(adm model.Admin) error //当adm.RolesId == ''时 应当清空对应的数据
	Delete(query string, args ...interface{}) (rowsAffected int64, err error)
	DeleteByRoleID(roleID ...uint64) (rowsAffected int64, err error)
	DeleteByAdmID(admID ...uint64) (rowsAffected int64, err error)
}
