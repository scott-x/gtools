# gtools
gin tools

### api
- `func HandleValidatorError(c *gin.Context, err error, trans ut.Translator)`: 表单验证

### translation使用方法

- 1，全局初始化翻译器
- 2，调用`HandleValidatorError(c *gin.Context, err error, trans ut.Translator)`进行表单验证

