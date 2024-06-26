package gtools

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

var trans ut.Translator //国际化翻译器

func init() {
	initTranslation() //初始化翻译器
}

func HandleValidatorError(c *gin.Context, err error) {
	log.Println("trans:", trans)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":  2001,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		//errs.Translate(trans)的本质就是map[string]string
		"error": fixStructKey(errs.Translate(trans)),
	})
}
