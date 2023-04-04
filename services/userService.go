package services

import (
	"github.com/kataras/iris/v12"
	"iris-init/global"
	"iris-init/model"
	"iris-init/repositories"
	"iris-init/repositories/repoComm"
	"iris-init/repositories/repoInterface"
	"iris-init/sErr"
	"reflect"
)

func NewUserService() UserService {
	return UserService{repo: repositories.NewUserRepo()}
}

type UserService struct {
	repo repoInterface.UserRepo
}

func (userServ *UserService) ListPage(ctx iris.Context) ([]model.User, *global.Pager) {
	where := repoInterface.UserSearchWhere{}
	_ = ctx.ReadQuery(&where)
	pager := global.NewPager(ctx)
	pager.SetTotal(userServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.User{}, pager
	}
	where.SelectParams = repoComm.SelectFrom{
		Offset:  pager.Offset,
		Limit:   pager.Size,
		RetSize: pager.Size,
		OrderBy: []repoComm.OrderByParams{
			{
				Column: "ID",
				Desc:   true,
			},
		},
	}
	return userServ.repo.GetList(where), pager
}

func (userServ *UserService) ListAvailable(_select ...string) []model.User {
	if len(_select) == 0 {
		_select = nil
	}
	return userServ.repo.GetList(repoInterface.UserSearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
		},
	})
}

func (userServ *UserService) ListByWhere(where repoInterface.UserSearchWhere) []model.User {
	return userServ.repo.GetList(where)
}

func (userServ *UserService) TotalCount(where repoInterface.UserSearchWhere) int64 {
	return userServ.repo.GetTotalCount(where)
}

// 获取一条数据根据ctx
// 这里条件为ID 传入ctx是方便后续修改参数条件
func (userServ *UserService) GetItem(ctx iris.Context, _select ...string) model.User {
	return userServ.repo.GetByID(ctx.URLParamUint64("ID"), _select...)
}

func (userServ *UserService) GetByID(id uint64, _select ...string) model.User {
	if id == 0 {
		return model.User{}
	}
	return userServ.repo.GetByID(id, _select...)
}
func (userServ *UserService) GetByPhone(phone sql.NullString, _select ...string) model.User {
	var v reflect.Value
	v = reflect.ValueOf(phone)
	if !v.IsValid() { // 值不存在
		return model.User{}
	}
	return userServ.repo.GetByPhone(phone, _select...)
}

func (userServ *UserService) CheckPhoneValid(user model.User) error {
	var v reflect.Value
	v = reflect.ValueOf(user.Phone)
	if !v.IsValid() { // 值不存在
		return sErr.New("无效的Phone")
	}
	_user := userServ.GetByPhone(user.Phone, "id")
	if _user.ID > 0 && user.ID != _user.ID {
		return sErr.NewFmt("Phone.已存在:%s.", user.Phone)
	}
	return nil
}
func (userServ *UserService) GetByToken(token sql.NullString, _select ...string) model.User {
	var v reflect.Value
	v = reflect.ValueOf(token)
	if !v.IsValid() { // 值不存在
		return model.User{}
	}
	return userServ.repo.GetByToken(token, _select...)
}

func (userServ *UserService) CheckTokenValid(user model.User) error {
	var v reflect.Value
	v = reflect.ValueOf(user.Token)
	if !v.IsValid() { // 值不存在
		return sErr.New("无效的Token")
	}
	_user := userServ.GetByToken(user.Token, "id")
	if _user.ID > 0 && user.ID != _user.ID {
		return sErr.NewFmt("Token.已存在:%s.", user.Token)
	}
	return nil
}
func (userServ *UserService) GetByUsername(username string, _select ...string) model.User {
	var v reflect.Value
	v = reflect.ValueOf(username)
	if !v.IsValid() { // 值不存在
		return model.User{}
	}
	return userServ.repo.GetByUsername(username, _select...)
}

func (userServ *UserService) CheckUsernameValid(user model.User) error {
	var v reflect.Value
	v = reflect.ValueOf(user.Username)
	if !v.IsValid() { // 值不存在
		return sErr.New("无效的账户名")
	}
	_user := userServ.GetByUsername(user.Username, "id")
	if _user.ID > 0 && user.ID != _user.ID {
		return sErr.NewFmt("账户名.已存在:%s.", user.Username)
	}
	return nil
}

func (userServ *UserService) GetByIDLock(id uint64, _select ...string) (model.User, repoComm.ReleaseLock) {
	return userServ.repo.GetByIDLock(id, _select...)
}

// 通过请求ctx编辑/新增一条数据
func (userServ *UserService) SaveByCtx(ctx iris.Context) (model.User, error) {
	userValidator := UserValidator{}
	err := global.ScanValidatorByRequestPost(ctx, &userValidator)
	if err != nil {
		return model.User{}, err
	}
	return userServ.SaveByValidator(userValidator)
}

func (userServ *UserService) SaveByValidator(userValidator UserValidator) (model.User, error) {
	user, err := userServ.GetUserByValidate(userValidator)
	if err != nil {
		return user, err
	}
	err = userServ.repo.Save(&user)
	return user, err
}

func (userServ *UserService) Save(user *model.User) error {
	return userServ.repo.Save(user)
}

func (userServ *UserService) Create(user *[]model.User) error {
	return userServ.repo.Create(user)
}

func (userServ *UserService) DeleteByCtx(ctx iris.Context) error {
	return userServ.DeleteByID(uint64(ctx.PostValueInt64Default("ID", 0)))
}

func (userServ *UserService) Delete(user model.User) error {
	if user.ID == 0 {
		return nil
	}
	_, err := userServ.repo.DeleteByID(user.ID)
	return err
}

func (userServ *UserService) DeleteByID(id ...uint64) error {
	if len(id) == 0 {
		return nil
	}
	_, err := userServ.repo.DeleteByID(id...)
	return err
}

func (userServ *UserService) ShowMapList(user []model.User) []map[string]interface{} {
	_user := []map[string]interface{}{}
	for _, v := range user {
		_user = append(_user, v.ShowMap())
	}
	return _user
}

// 验证参数 并返回到一个新的User model
func (userServ *UserService) GetUserByValidate(userValidator UserValidator) (model.User, error) {
	err := userValidator.Validate()
	if err != nil {
		return model.User{}, err
	}
	var user model.User
	if userValidator.ID > 0 {
		user = userServ.repo.GetByID(userValidator.ID)
		if user.ID == 0 {
			return user, sErr.New("无效的ID")
		}
	} else {
		user = model.User{}
	}
	//完成其他的赋值逻辑处理...
	user.ID = userValidator.ID
	user.Username = userValidator.Username
	user.Phone = userValidator.Phone
	user.Avatar = userValidator.Avatar
	user.RealName = userValidator.RealName
	user.Sex = userValidator.Sex
	user.Token = userValidator.Token
	user.TokenStatus = userValidator.TokenStatus
	user.TokenExpired = userValidator.TokenExpired
	user.LastLoginTime = userValidator.LastLoginTime
	if err = userServ.CheckPhoneValid(user); err != nil {
		return user, err
	}
	if err = userServ.CheckTokenValid(user); err != nil {
		return user, err
	}
	if err = userServ.CheckUsernameValid(user); err != nil {
		return user, err
	}
	return user, nil
}

type UserValidator struct {
	ID            uint64         `label:"ID"`
	Username      string         `validate:"max=30,required" label:"账户名"`
	Phone         sql.NullString `validate:"max=15" label:"Phone"`
	Avatar        string         `validate:"required" label:"Avatar"`
	RealName      string         `validate:"max=20" label:"RealName"`
	Sex           uint8          `label:"Sex"`
	Token         sql.NullString `validate:"max=32,required" label:"Token"`
	TokenStatus   uint8          `label:"TokenStatus"`
	TokenExpired  int64          `label:"TokenExpired"`
	LastLoginTime int64          `label:"LastLoginTime"`
}

func (userValidator *UserValidator) Validate() error {
	err := global.ValidateV9Struct(userValidator)
	if err != nil {
		return err
	}
	return nil
}
