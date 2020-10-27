package express

import (
	"fmt"
	"regexp"
)

//比较表达式
type Comparable string
func(this Comparable) filter() string {
	reg,err:=regexp.Compile(ComparePattern)
	if err!=nil{
		return ""
	}
	ret:=reg.FindStringSubmatch(string(this))
	if ret!=nil && len(ret)==4{
		token:=getCompareToken(ret[2])
		if token==""{
			return ""
		}
		return fmt.Sprintf("%s %s %s",token,parseToken(ret[1]),parseToken(ret[3]))
	}
	return ""
}

func IsComparable(expr string ) bool {
	reg,err:=regexp.Compile(ComparePattern)
	if err!=nil{
		return false
	}
	return reg.MatchString(expr)
}
