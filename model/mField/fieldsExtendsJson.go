package mField

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"iris-init/global"
)

type FieldsExtendsJsonType struct {
	ExtraData       string                 `gorm:"type:text;comment:扩展字段" mapstructure:"extra_data"`
	extraDataMap    map[string]interface{} `gorm:"-"`
	extraDataResult gjson.Result           `gorm:"-"`
}

func (e *FieldsExtendsJsonType) GetResult() gjson.Result {
	if e.ExtraData == "" {
		return gjson.Result{}
	}
	if !e.extraDataResult.Exists() {
		e.extraDataResult = gjson.Parse(e.ExtraData)
	}
	return e.extraDataResult
}

func (e *FieldsExtendsJsonType) GetExtendsMap() map[string]interface{} {
	if e.ExtraData == "" {
		return nil
	}
	if e.extraDataMap != nil {
		return e.extraDataMap
	}
	e.extraDataMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(e.ExtraData), &e.extraDataMap)
	if err != nil {
		return nil
	}
	return e.extraDataMap
}

func (e *FieldsExtendsJsonType) GetExtendsJson(key string) gjson.Result {
	if !e.extraDataResult.Exists() {
		e.ReloadExtraDataResult()
	}
	return e.extraDataResult.Get(key)
}

func (e *FieldsExtendsJsonType) ReloadExtraDataMap() {
	if e.ExtraData == "" {
		e.extraDataMap = nil
	}
	e.extraDataMap = global.Json2Map(e.ExtraData)
}

func (e *FieldsExtendsJsonType) ReloadExtraDataResult() {
	if e.ExtraData == "" {
		e.extraDataResult = gjson.Result{}
	}
	e.extraDataResult = gjson.Parse(e.ExtraData)
}

func (e *FieldsExtendsJsonType) SetExtendsJson(key string, value interface{}) {
	if e.extraDataMap == nil {
		e.extraDataMap = global.Json2Map(e.ExtraData)
	}
	e.extraDataMap[key] = value
	eJson, _ := json.Marshal(e.extraDataMap)
	e.ExtraData = string(eJson)
	e.ReloadExtraDataResult()
}
