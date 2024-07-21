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
	//http.StatusBadRequest 改成了 http.StatusOK 否则react前端请求 network会报400错误，导致无法获取后端接口数据 而无法正确展示error msg
	//详情请看 https://statics.scott-xiong.com/docusaurus/0aaa2bb2642443619e1f51863aa7884e.png
	//c.JSON(http.StatusBadRequest, gin.H{
	c.JSON(http.StatusOK, gin.H{
		//errs.Translate(trans)的本质就是map[string]string
		"code":  2002,
		"error": fixStructKey(errs.Translate(trans)),
	})
}
