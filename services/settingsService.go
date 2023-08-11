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

func NewSettingsService() SettingsService {
	return SettingsService{repo: repositories.NewSettingsRepo()}
}

type SettingsService struct {
	repo repoInterface.SettingsRepo
}

func (settingsServ SettingsService) GetWebsiteTitle() string {
	return settingsServ.GetByKey(model.SettingsKeyWebsiteTitle).Value
}

func (settingsServ *SettingsService) ListPage(ctx iris.Context) ([]model.Settings, *global.Pager) {
	where := repoInterface.SettingsSearchWhere{}
	_ = ctx.ReadQuery(&where)
	pager := global.NewPager(ctx)
	if pager.Size < 0 {
		return settingsServ.repo.GetList(where), nil
	}
	pager.SetTotal(settingsServ.repo.GetTotalCount(where))
	if pager.Total == 0 {
		return []model.Settings{}, pager
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
	return settingsServ.repo.GetList(where), pager
}

func (settingsServ *SettingsService) ListAvailable(_select ...string) []model.Settings {
	if len(_select) == 0 {
		_select = nil
	}
	return settingsServ.repo.GetList(repoInterface.SettingsSearchWhere{
		SelectParams: repoComm.SelectFrom{
			Select: _select,
		},
	})
}

func (settingsServ *SettingsService) ListByWhere(where repoInterface.SettingsSearchWhere) []model.Settings {
	return settingsServ.repo.GetList(where)
}

func (settingsServ *SettingsService) TotalCount(where repoInterface.SettingsSearchWhere) int64 {
	return settingsServ.repo.GetTotalCount(where)
}

// 获取一条数据根据ctx
// 这里条件为ID 传入ctx是方便后续修改参数条件
func (settingsServ *SettingsService) GetItem(ctx iris.Context, _select ...string) model.Settings {
	return settingsServ.repo.GetByID(ctx.URLParamUint64("ID"), _select...)
}

func (settingsServ *SettingsService) GetByID(id uint64, _select ...string) model.Settings {
	if id == 0 {
		return model.Settings{}
	}
	return settingsServ.repo.GetByID(id, _select...)
}

func (settingsServ *SettingsService) UpdateByWhere(where repoInterface.SettingsSearchWhere, data interface{}) (rowsAffected int64, err error) {
	return settingsServ.repo.UpdateByWhere(where, data)
}
func (settingsServ *SettingsService) GetByKey(key string, _select ...string) model.Settings {
	var v reflect.Value
	v = reflect.ValueOf(key)
	if !v.IsValid() { // 值不存在
		return model.Settings{}
	}
	return settingsServ.repo.GetByKey(key, _select...)
}

func (settingsServ *SettingsService) CheckKeyValid(settings model.Settings) error {
	var v reflect.Value
	v = reflect.ValueOf(settings.Key)
	if !v.IsValid() { // 值不存在
		return sErr.New("无效的Key")
	}
	_settings := settingsServ.GetByKey(settings.Key, "id")
	if _settings.ID > 0 && settings.ID != _settings.ID {
		return sErr.NewFmt("Key.已存在:%s.", settings.Key)
	}
	return nil
}

func (settingsServ *SettingsService) DeleteByKey(key string) error {
	_, err := settingsServ.repo.DeleteByKey(key)
	return err
}

func (settingsServ *SettingsService) GetByIDLock(id uint64, _select ...string) (model.Settings, repoComm.ReleaseLock) {
	return settingsServ.repo.GetByIDLock(id, _select...)
}

func (settingsServ *SettingsService) SaveByCtx(ctx iris.Context) error {
	formValues := ctx.FormValues()
	for k, v := range formValues {
		if len(v) == 0 {
			continue
		}
		_, err := settingsServ.SaveByValidator(SettingsValidator{Key: k, Value: v[0]})
		if err != nil {
			return err
		}
	}
	return nil
}

func (settingsServ SettingsService) GetSettings() []model.Settings {
	return []model.Settings{
		{
			Key:  model.SettingsKeyWebsiteTitle,
			Name: "网站标题",
			Desc: "网站标题，空则使用默认标题.",
		},
	}
}

func (settingsServ SettingsService) ReloadSettings() error {
	for _, def := range settingsServ.GetSettings() {
		local := settingsServ.GetByKey(def.Key)
		if local.ID > 0 {
			local.Name = def.Name
			local.Desc = def.Desc
		} else {
			local = def
		}
		if err := settingsServ.Save(&local); err != nil {
			return err
		}
	}
	return nil
}

func (settingsServ *SettingsService) SaveByValidator(settingsValidator SettingsValidator) (model.Settings, error) {
	settings, err := settingsServ.GetSettingsByValidate(settingsValidator)
	if err != nil {
		return settings, err
	}
	err = settingsServ.Save(&settings)
	return settings, err
}

func (settingsServ *SettingsService) Save(settings *model.Settings) error {
	return settingsServ.repo.Save(settings)
}

func (settingsServ *SettingsService) Create(settings *[]model.Settings) error {
	return settingsServ.repo.Create(settings)
}

func (settingsServ *SettingsService) DeleteByCtx(ctx iris.Context) error {
	return settingsServ.DeleteByID(uint64(ctx.PostValueInt64Default("ID", 0)))
}

// 这个方法目前是与DeleteByID功能一致 主要是用来扩展的 根据model的多条件作删除 需要开发者自己完成业务逻辑
func (settingsServ *SettingsService) Delete(settings model.Settings) error {
	if settings.ID == 0 {
		return nil
	}
	_, err := settingsServ.repo.DeleteByID(settings.ID)
	return err
}

func (settingsServ *SettingsService) DeleteByID(id ...uint64) error {
	if len(id) == 0 {
		return nil
	}
	_, err := settingsServ.repo.DeleteByID(id...)
	return err
}

func (settingsServ *SettingsService) ShowMapList(settings []model.Settings) []map[string]interface{} {
	_settings := []map[string]interface{}{}
	for _, v := range settings {
		_settings = append(_settings, v.ShowMap())
	}
	return _settings
}

// 验证参数 并返回到一个新的Settings model
func (settingsServ *SettingsService) GetSettingsByValidate(settingsValidator SettingsValidator) (model.Settings, error) {
	err := settingsValidator.Validate()
	if err != nil {
		return model.Settings{}, err
	}
	var settings model.Settings
	if settingsValidator.Key == "" {
		return settings, sErr.New("无效的配置项")
	}
	settings = settingsServ.GetByKey(settingsValidator.Key)
	if settings.ID == 0 {
		return settings, sErr.New("无效的配置项")
	}
	//完成其他的赋值逻辑处理...
	settings.Value = settingsValidator.Value
	return settings, nil
}

type SettingsValidator struct {
	Key   string `validate:"max=200,required" label:"Key"`
	Value string `label:"Value"`
}

func (settingsValidator *SettingsValidator) Validate() error {
	err := global.ValidateV9Struct(settingsValidator)
	if err != nil {
		return err
	}
	return nil
}
