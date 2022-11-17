package repoInterface

import "jd-fxl/model"

type RolesAdmRepo interface {
	RepoInterface
	SaveByAdm(adm model.Admin) error //当adm.RolesId == ''时 应当清空对应的数据
}
