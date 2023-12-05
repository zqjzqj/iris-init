package global

import (
	"big_data_new/sErr"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
	zh2 "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
)

var ValidateV9 *validator.Validate
var ZhTrans ut.Translator

func init() {
	zhCh := zh.New()
	ValidateV9 = validator.New()
	ValidateV9.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	uni := ut.New(zhCh)
	ZhTrans, _ = uni.GetTranslator("zh")
	_ = zh2.RegisterDefaultTranslations(ValidateV9, ZhTrans)
}

func ValidateV9Var(_var interface{}, tag string) error {
	err := ValidateV9.Var(_var, tag)
	if err != nil {
		errStr := ""
		for _, errVal := range err.(validator.ValidationErrors) {
			errStr = fmt.Sprintf("%s %s", errStr, errVal.Translate(ZhTrans))
		}
		return sErr.New(errStr)
	}
	return nil
}

func ValidateV9Struct(s interface{}) error {
	err := ValidateV9.Struct(s)
	if err != nil {
		errStr := ""
		for _, errVal := range err.(validator.ValidationErrors) {
			errStr = fmt.Sprintf("%s %s", errStr, errVal.Translate(ZhTrans))
		}
		return sErr.New(errStr)
	}
	return nil
}

func ValidateV9Struct_FieldErrors(s interface{}) []sErr.FieldErr {
	err := ValidateV9.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		_errors := make([]sErr.FieldErr, 0, len(validationErrors))
		for _, errVal := range validationErrors {
			_errors = append(_errors, sErr.FieldErr{
				Field: errVal.Field(),
				Err:   errVal.Translate(ZhTrans),
			})
		}
		return _errors
	}
	return nil
}

type ValidatorV9Interface interface {
	Validate() error //这里只做 不操作orm 的基础验证
}

func ScanValidatorByRequestPost(ctx iris.Context, v ValidatorV9Interface) error {
	_ = ctx.ReadBody(v)
	return nil
}
