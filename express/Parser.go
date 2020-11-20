package express

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

const (
	VarPattern=`[0-9a-zA-Z_]+`
	CompareSign=">|>=|<=|<|==|!="
	CompareSignToken="gt|ge|le|lt|eq|ne"
	ComparePattern=`^(`+VarPattern+`)\s*(`+CompareSign+`)\s*(`+VarPattern+`)\s*$`
)

type Express string

//根据比较符 ，获取token
func getCompareToken(sign string) string{
	for index,item:=range strings.Split(CompareSign,"|"){
		if item==sign{
			return  strings.Split(CompareSignToken,"|")[index]
		}
	}
	return ""
}
//对于数字不加.(点)
func parseToken(token string) string  {
	if IsNumeric(token){
		return token
	}else {
		return "."+token
	}
}


//执行表达式
func Exec(express Express, data map[string]interface{}) (string,error){
	tpl:=template.New("expr").Funcs(map[string]interface{}{
		"echo": func(params ...interface{}) interface{} {
			return fmt.Sprintf("echo:%v",params[0])
		},
	})

	t,err:= tpl.Parse(fmt.Sprintf("{{%s}}",express))
	if err!=nil{
		return "",err
	}
	var buf=&bytes.Buffer{}
	err=t.Execute(buf,data)
	if err!=nil{
		return "",err
	}
	return buf.String(),nil
}