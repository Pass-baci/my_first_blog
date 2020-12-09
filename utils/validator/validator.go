package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	untTranslations "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	reflect "reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := untTranslations.New(zh_Hans_CN.New())
	translator, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTranslations.RegisterDefaultTranslations(validate, translator)
	if err != nil {
		fmt.Println("err:",err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		get := field.Tag.Get("label")
		return get
	})
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(translator), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
