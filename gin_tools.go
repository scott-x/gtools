package gtools

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var trans ut.Translator //国际化翻译器

// 初始化翻译器
func init() {
	if err := initTrans("zh"); err != nil {
		log.Panic("初始化翻译器失败")
	}
	log.Println("成功初始化翻译器")
}

// 优化Validator的error eg: "SignUpForm.email"` 改成 `"email"`
func ValidateBindError(c *gin.Context, err error) {
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
		"code":  2002,
		"error": fixStructKey(errs.Translate(trans)),
	})
}
