package validator

import (
	"reflect"
	"strings"

	"github.com/Cheng1622/news_go_server/pkg/clog"
	"github.com/Cheng1622/news_go_server/pkg/code"
	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

// InitValidate validator信息翻译
func InitValidate() {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取jsonTag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Trans, ok = uni.GetTranslator(config.Conf.System.I18nLanguage)
		if !ok {
			clog.Log.Fatalln("初始化validator.v10数据校验器失败")
		}
		switch config.Conf.System.I18nLanguage {
		case "en":
			_ = enTrans.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zhTrans.RegisterDefaultTranslations(v, Trans)
		default:
			_ = enTrans.RegisterDefaultTranslations(v, Trans)
		}

	}
	clog.Log.Infoln("初始化validator.v10数据校验器完成")
}

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(c *gin.Context, err error) {
	//如何返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.Error(c, code.ValidateError, err.Error())
		return
	}
	response.Error(c, code.ValidateError, removeTopStruct(errs.Translate(Trans)))
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
