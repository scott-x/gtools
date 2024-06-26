package gtools

import "strings"

// 优化Validator的error eg: "SignUpForm.email"` 改成 `"email"`
func fixStructKey(fileds map[string]string) map[string]string {
	rsp := make(map[string]string)
	for field, err := range fileds {
		rsp[field[strings.LastIndex(field, ".")+1:]] = err
	}
	return rsp
}
