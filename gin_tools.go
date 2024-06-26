package gtools

import (
	"log"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"github.com/scott-x/resp"
)

var trans ut.Translator //国际化翻译器

func init() {
	initTranslation() //初始化翻译器
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		resp.Error(c, 2001, err.Error())
		return
	}
	//errs.Translate(trans)的本质就是map[string]string
	log.Panicln("trans:", trans)
	resp.Error(c, 2002, fixStructKey(errs.Translate(trans)))
}
