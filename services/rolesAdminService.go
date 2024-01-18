package services

import (
	"github.com/kataras/iris/v12"
	"iris-init/global"
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
)

func NewRolesAdminService() RolesAdminService {
	return RolesAdminService{repo: repositories.NewRolesAdminRepo()}
}

func NewRolesAdminServiceByOrm(orm any) RolesAdminService {
	r := RolesAdminService{repo: repositories.NewRolesAdminRepo()}
	r.repo.SetOrm(orm)
	return r
}

func NewRolesAdminServiceByRepo(repo repoInterface.RolesAdminRepo) RolesAdminService {
	return RolesAdminService{repo: repo}
}

type RolesAdminService struct {
	repo repoInterface.RolesAdminRepo
}

func (rolesAdminServ RolesAdminService) ListPage(ctx iris.Context) ([]model.RolesAdmin, *global.Pager) {
	where := repoInterface.RolesAdminSearchWhere{}
	_ = ctx.ReadQuery(&where)
	where.SelectParams = repoComm.SelectFrom{
		OrderBy: []repoComm.OrderByParams{
			{
				Column: "ID",
				Desc:   true,
			},
		},
	}
	pager := global.NewPager(ctx)
	if pager.Size < 0 {
		return rolesAdminServ.repo.GetList(where), nil
	}
	pager.SetTotal(rolesAdminServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.RolesAdmin{}, pager
	}
	where.SelectParams.Offset = pager.Offset
	where.SelectParams.Limit = pager.Size
	where.SelectParams.RetSize = pager.Size
	return rolesAdminServ.repo.GetList(where), pager
}

func (rolesAdminServ RolesAdminService) ListAvailable(_select ...string) []model.RolesAdmin {
	if len(_select) == 0 {
		_select = nil
	}
	return rolesAdminServ.repo.GetList(repoInterface.RolesAdminSearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
		},
	})
}

func (rolesAdminServ RolesAdminService) ListByWhere(where repoInterface.RolesAdminSearchWhere) []model.RolesAdmin {
	return rolesAdminServ.repo.GetList(where)
}

func (rolesAdminServ RolesAdminService) TotalCount(where repoInterface.RolesAdminSearchWhere) int64 {
	return rolesAdminServ.repo.GetTotalCount(where)
}

func (rolesAdminServ RolesAdminService) GetByID(ID []uint8, _select ...string) model.RolesAdmin {
	return rolesAdminServ.repo.GetByID(ID, _select...)
}

func (rolesAdminServ RolesAdminService) GetByWhere(where repoInterface.RolesAdminSearchWhere) model.RolesAdmin {
	return rolesAdminServ.repo.GetByWhere(where)
}

func (rolesAdminServ RolesAdminService) ScanByWhere(where repoInterface.RolesAdminSearchWhere, dest any) error {
	return rolesAdminServ.repo.ScanByWhere(where, dest)
}

func (rolesAdminServ RolesAdminService) ScanByOrWhere(dest any, where ...repoInterface.RolesAdminSearchWhere) error {
	return rolesAdminServ.repo.ScanByOrWhere(dest, where...)
}

func (rolesAdminServ RolesAdminService) UpdateByWhere(where repoInterface.RolesAdminSearchWhere, data interface{}) (rowsAffected int64, err error) {
	return rolesAdminServ.repo.UpdateByWhere(where, data)
}
func (rolesAdminServ RolesAdminService) GetByAdminID(adminID uint64, _select ...string) []model.RolesAdmin {
	return rolesAdminServ.repo.GetByAdminID(adminID, _select...)
}

func (rolesAdminServ RolesAdminService) DeleteByAdminID(adminID uint64) error {
	_, err := rolesAdminServ.repo.DeleteByAdminID(adminID)
	return err
}
func (rolesAdminServ RolesAdminService) GetByRoleID(roleID uint64, _select ...string) []model.RolesAdmin {
	return rolesAdminServ.repo.GetByRoleID(roleID, _select...)
}

func (rolesAdminServ RolesAdminService) DeleteByRoleID(roleID uint64) error {
	_, err := rolesAdminServ.repo.DeleteByRoleID(roleID)
	return err
}

func (rolesAdminServ RolesAdminService) GetByIDLock(ID []uint8, _select ...string) (model.RolesAdmin, repoComm.ReleaseLock) {
	return rolesAdminServ.repo.GetByIDLock(ID, _select...)
}

func (rolesAdminServ RolesAdminService) ReloadAdmin(rolesAdmin *model.RolesAdmin) {
	if rolesAdmin.AdminID > 0 {
		rolesAdmin.Admin = NewAdminServiceByOrm(rolesAdminServ.repo.GetOrm()).
			GetByID(rolesAdmin.AdminID)
	}
}

func (rolesAdminServ RolesAdminService) SaveByAdm(adm model.Admin) error {
	return rolesAdminServ.repo.SaveByAdm(adm)
}

func (rolesAdminServ RolesAdminService) Save(rolesAdmin *model.RolesAdmin) error {
	return rolesAdminServ.repo.Save(rolesAdmin)
}

func (rolesAdminServ RolesAdminService) Create(rolesAdmin *[]model.RolesAdmin) error {
	return rolesAdminServ.repo.Create(rolesAdmin)
}

// 这个方法目前是与DeleteByID功能一致 主要是用来扩展的 根据model的多条件作删除 需要开发者自己完成业务逻辑
func (rolesAdminServ RolesAdminService) Delete(rolesAdmin model.RolesAdmin) error {
	_, err := rolesAdminServ.repo.DeleteByID(rolesAdmin.ID)
	return err
}

func (rolesAdminServ RolesAdminService) ShowMapList(rolesAdmin []model.RolesAdmin) []map[string]interface{} {
	_rolesAdmin := []map[string]interface{}{}
	for _, v := range rolesAdmin {
		_rolesAdmin = append(_rolesAdmin, v.ShowMap())
	}
	return _rolesAdmin
}
