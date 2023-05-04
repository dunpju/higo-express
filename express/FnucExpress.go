package express

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dunpju/higo-express/express/fnucExpr"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type FuncMap map[string]interface{}

var funcMap FuncMap

type FuncListener struct {
	*fnucExpr.BaseFnucExprListener
	funcName   string
	args       []reflect.Value
	methodName string // 方法名 eg: user.abc.bcd.age()
	execType   uint8  // 执行类型 0代表函数(默认值) 1代表struct执行
}

func init() {
	funcMap = make(map[string]interface{})
}

// 执行表达式
func Run(express string) Result {
	is := antlr.NewInputStream(express)
	lexer := fnucExpr.NewFnucExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := fnucExpr.NewFnucExprParser(ts)
	lis := &FuncListener{}
	antlr.ParseTreeWalkerDefault.Walk(lis, p.Start())
	return lis.Run()
}

// 设置函数容器
func SetFuncMap(key string, value interface{}) {
	funcMap[key] = value
}

func (this *FuncListener) ExitFuncArgs(ctx *fnucExpr.FuncArgsContext) {
	for i := 0; i < ctx.GetChildCount(); i++ {
		token := ctx.GetChild(i).GetPayload().(*antlr.CommonToken)
		var value reflect.Value
		switch token.GetTokenType() {
		case fnucExpr.FnucExprLexerStringArg:
			stringArg := strings.Trim(token.GetText(), "'")
			value = reflect.ValueOf(stringArg)
			break
		case fnucExpr.FnucExprLexerIntArg:
			v, err := strconv.ParseInt(token.GetText(), 10, 64)
			if err != nil {
				panic("parse int64 error")
			}
			value = reflect.ValueOf(v)
			break
		case fnucExpr.FnucExprLexerFloatArg:
			v, err := strconv.ParseFloat(token.GetText(), 64)
			if err != nil {
				panic("parse float64 error")
			}
			value = reflect.ValueOf(v)
			break
		default:
			continue
		}
		this.args = append(this.args, value)
	}
}

func (this *FuncListener) ExitFuncCall(ctx *fnucExpr.FuncCallContext) {
	this.funcName = ctx.GetStart().GetText()
}

func (this *FuncListener) ExitMethodCall(ctx *fnucExpr.MethodCallContext) {
	this.execType = 1
	this.methodName = ctx.GetStart().GetText()
}

func (this *FuncListener) Run() Result {
	if funcMap == nil {
		panic("exprMap required")
	}
	switch this.execType {
	case 0:
		if f, ok := funcMap[this.funcName]; ok {
			v := reflect.ValueOf(f)
			if v.Kind() == reflect.Func {
				return result(v.Call(this.args))
			}
		}
		break
	case 1:
		ms := strings.Split(this.methodName, ".")
		if obj, ok := funcMap[ms[0]]; ok {
			objv := reflect.ValueOf(obj)
			current := objv
			for i := 1; i < len(ms); i++ {
				if i == len(ms)-1 { //最后一个是方法名
					if method := current.MethodByName(ms[i]); !method.IsValid() {
						panic("method error:" + ms[i])
					} else {
						return result(method.Call(this.args))
					}
				}
				field := this.findField(ms[i], current)
				if field.IsValid() {
					current = field
				} else {
					panic("field error:" + ms[i])
				}
			}
		}
		break
	default:
		log.Println("nothing to do")
	}
	return newResult()
}

func (this *FuncListener) findField(method string, v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if field := v.FieldByName(method); field.IsValid() {
		return field
	}
	return reflect.Value{}
}
