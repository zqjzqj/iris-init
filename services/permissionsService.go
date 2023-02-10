package services

import (
	"github.com/kataras/iris/v12"
	"iris-init/global"
	"iris-init/logs"
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoInterface"
	"strconv"
	"strings"
)

type PermissionService struct {
	repo repoInterface.PermissionsRepo
}

func NewPermissionService() PermissionService {
	return PermissionService{repo: repositories.NewPermissionsRepo()}
}

func (permServ PermissionService) GetPremAsMenu(idents []string) []model.Permissions {
	if len(idents) == 1 && idents[0] == model.RoleAdmin {
		idents = nil
	}
	return permServ.repo.GetListAsMenu(idents)
}

func (permServ PermissionService) GetPermTree(ident ...string) []map[string]interface{} {
	if len(ident) == 1 && ident[0] == model.RoleAdmin {
		ident = nil
	}
	perms := permServ.repo.GetListPreloadChildren_2()
	return permServ.ShowPermTreeMap(perms, ident...)
}

func (permServ PermissionService) RefreshChildren(perm *model.Permissions, force bool) {
	if !force && perm.Children != nil {
		return
	}
	perm.Children = permServ.repo.GetListPreloadChildren(repoInterface.PermissionsSearchWhere{
		Pid: int64(perm.ID),
	})
}

func (permServ PermissionService) ShowPermTreeMap(perms []model.Permissions, ident ...string) []map[string]interface{} {
	l := len(perms)
	if l == 0 {
		return []map[string]interface{}{}
	}
	r := make([]map[string]interface{}, 0, l)
	for _, perm := range perms {
		rr := map[string]interface{}{
			"id":     perm.Ident,
			"title":  perm.Name,
			"field":  "Idents",
			"spread": true,
		}

		//这里因为前端layui的tree主键会根据children的勾选情况来决定自身是否勾选
		//所以如果有children的节点不勾选
		if len(perm.Children) == 0 && global.InSlice(perm.Ident, ident) {
			rr["checked"] = true
		}
		//permServ.RefreshChildren(&perm, false)
		rr["children"] = permServ.ShowPermTreeMap(perm.Children, ident...)
		r = append(r, rr)
	}
	return r
}

func (permServ PermissionService) GenerateAdminPermissionsByRoutes(app *iris.Application) {
	//先截断表
	permServ.repo.TruncateTable()
	var pid uint64
	for _, r := range app.GetRoutes() {
		logs.PrintlnInfo("create permission ", r.Name, r.Method, r.Path)

		if !strings.Contains(r.Name, "@") {
			logs.PrintlnWarning("continue Contains Name.... ", r.Name, r.Method, r.Path)
			continue
		}
		//拆分失败 则跳过
		pName := strings.Split(r.Name, "@")
		if len(pName) != 2 {
			logs.PrintlnWarning("continue Split  .... ", r.Name, r.Method, r.Path, pName)
			continue
		}
		//分为目录和菜单
		dir, mBtn := pName[0], pName[1]
		if dir != "" {
			_dir, sort := permServ.getNameAndSortByName(dir)
			perm, err := permServ.repo.GetOrCreatePermissionByName(_dir, 0, model.PermissionsLevelDir, sort)
			pid = perm.ID
			if err != nil {
				logs.PrintErr("get dir pid fail ", _dir, err)
				continue
			}
		}
		//这里要查找一下菜单下是否有按钮
		mBts := strings.Split(mBtn, ":")
		lenMBts := len(mBts)
		var path model.Permissions
		path = permServ.repo.GetByIdent(permServ.GeneratePermissionAuthIdentify(r.Method, r.Path))
		if lenMBts >= 2 { //有按钮存在
			_btn, sort := permServ.getNameAndSortByName(mBts[0])
			perm, err := permServ.repo.GetOrCreatePermissionByName(_btn, pid, model.PermissionsLevelMenu, sort)
			pid = perm.ID
			if err != nil {
				logs.PrintErr(err)
				continue
			}
			//拿出最后一个作为path 并且排除第一个已创建的菜单
			mBtn = mBts[lenMBts-1]
			mBts = mBts[:lenMBts-1]
			for _, v := range mBts[1:] {
				logs.PrintlnSuccess("get or create ...", v)
				_v, _sort := permServ.getNameAndSortByName(v)
				perm, _ = permServ.repo.GetOrCreatePermissionByName(_v, pid, model.PermissionsLevelBtn, _sort)
				pid = perm.ID
			}
			path = model.Permissions{Level: model.PermissionsLevelBtn}
		} else {
			path = model.Permissions{Level: model.PermissionsLevelMenu}
		}

		//没有修改
		if path.ID > 0 && path.Pid == pid && path.Path == r.Path && path.Method == r.Method && path.Name == mBtn {
			logs.PrintlnSuccess("exist path ", mBtn, path.Name, path.Method, path.Path)
			continue
		}
		//排序处理
		path.Name, path.Sort = permServ.getNameAndSortByName(mBtn)
		path.Pid = pid
		path.Path = r.Path
		path.Method = r.Method
		path.GenerateIdent()
		if err := permServ.repo.Save(&path); err != nil {
			logs.PrintErr("save path fail ", path.Name, err)
			continue
		}
		logs.PrintlnSuccess("save path success ", mBtn, path.Name, path.Method, path.Path)
	}
	logs.PrintlnSuccess("GenerateAdminPermissionsByRoutes OK.")
}

func (permServ PermissionService) getNameAndSortByName(name string) (newName string, sort uint) {
	if strings.Contains(name, ".Sort{") {
		_mBtn := strings.Split(name, ".Sort{")
		newName = _mBtn[0]
		_mBtn[1] = strings.TrimRight(_mBtn[1], "}")
		_sort, err := strconv.Atoi(_mBtn[1])
		if err != nil {
			_sort = 100
		}
		sort = uint(_sort)
		return newName, sort
	}
	return name, 100
}

func (permServ PermissionService) GeneratePermissionAuthIdentify(method, path string) string {
	p := model.Permissions{Method: method, Path: path}
	p.GenerateIdent()
	return p.Ident
}

func (permServ PermissionService) IdentifyExists(ident string) bool {
	perm := permServ.repo.GetByIdent(ident, "id")
	if perm.ID > 0 {
		return true
	}
	return false
}

func (permServ PermissionService) GetByIdent(ident string, _select ...string) model.Permissions {
	return permServ.repo.GetByIdent(ident, _select...)
}

func (permServ PermissionService) GetByID(id uint64, _select ...string) model.Permissions {
	return permServ.repo.GetByID(id, _select...)
}

func (permServ PermissionService) GetPermParentsByIdent(ident string) []model.Permissions {
	permIdent := permServ.GetByIdent(ident)
	if permIdent.Pid == 0 {
		return nil
	}
	r := permServ.GetByID(permIdent.Pid)
	if r.Pid == 0 {
		return []model.Permissions{r}
	}
	perms := make([]model.Permissions, 0, 2)
	perms = append(perms, r)
	perms = append(perms, permServ.GetPermParentsByIdent(r.Ident)...)
	return perms
}
