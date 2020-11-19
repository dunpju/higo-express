package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dengpju/higo-express/express/fnucExpr"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Me struct {

}

func (this *Me)Age() int {
	return 21
}

var FuncMap map[string]interface{}
func main()  {
	/**
	tpl:=template.New("test").Funcs(map[string]interface{}{
		"echo": func(params ...interface{}) interface{}{
			return fmt.Sprintf("echo:%v", params[0])
		},
	})
	t,err:=tpl.Parse("{{ gt .age 20 | echo }}")
	if err!=nil {
		log.Fatal(err)
	}
	var buf = &bytes.Buffer{}
	err=t.Execute(buf,map[string]interface{}{
		"age":21,
	})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(buf.String())
	*/
	//fmt.Println(express.Exec(".me.Age", map[string]interface{}{"myage":19.0,"me":&Me{}}))

	FuncMap= map[string]interface{}{
		"test": func(name string,age int64) {
			log.Println("this is test", name, age)
		},
		"User":&User{Adm:&Admin{}},
	}

	is:=antlr.NewInputStream("User.Adm.Abc()")
	lexer:=fnucExpr.NewFnucExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := fnucExpr.NewFnucExprParser(ts)
	lis:=&FuncListener{}
	antlr.ParseTreeWalkerDefault.Walk(lis, p.Start())
	lis.Run()
}

type FuncListener struct {
	*fnucExpr.BaseFnucExprListener
	funcName   string
	args       []reflect.Value
	methodName string // 方法名 eg: user.abc.bcd.age()
	execType   uint8 // 执行类型 0代表函数(默认值) 1代表struct执行
}

func (this *FuncListener) ExitFuncArgs(ctx *fnucExpr.FuncArgsContext) {
	for i:=0;i<ctx.GetChildCount();i++{
		token:=ctx.GetChild(i).GetPayload().(*antlr.CommonToken)
		var value reflect.Value
		switch  token.GetTokenType() {
		case fnucExpr.FnucExprLexerStringArg:
			stringArg:=strings.Trim(token.GetText(),"'")
			value=reflect.ValueOf(stringArg)
			break
		case fnucExpr.FnucExprLexerIntArg:
			v,err:=strconv.ParseInt(token.GetText(),10,64)
			if err!=nil{
				panic("parse int64 error")
			}
			value=reflect.ValueOf(v)
			break
		case fnucExpr.FnucExprLexerFloatArg:
			v,err:=strconv.ParseFloat(token.GetText(),64)
			if err!=nil{
				panic("parse float64 error")
			}
			value=reflect.ValueOf(v)
			break
		default:
			continue
		}
		this.args=append(this.args,value)
	}

}

func (this *FuncListener) ExitFuncCall(ctx *fnucExpr.FuncCallContext) {
	this.funcName=ctx.GetStart().GetText()
}

func (this *FuncListener) ExitMethodCall(ctx *fnucExpr.MethodCallContext) {
	this.execType=1
	this.methodName=ctx.GetStart().GetText()
}

func (this *FuncListener) Run() {
	switch this.execType {
	case 0:
		if f,ok:=FuncMap[this.funcName];ok{
			v:=reflect.ValueOf(f)
			if v.Kind()==reflect.Func{
				v.Call(this.args)
			}
		}
	case 1:
		ms:=strings.Split(this.methodName,".")
		if obj,ok:=FuncMap[ms[0]];ok {
			objv:=reflect.ValueOf(obj)
			current:=objv
			for i:=1;i<len(ms);i++{
				if i==len(ms)-1{//最后一个是方法名
					if method:=current.MethodByName(ms[i]);!method.IsValid(){
						panic("method error:"+ms[i])
					}else{
						method.Call(this.args)
					}
					break
				}
				field:=this.findField(ms[i],current)
				if field.IsValid(){
					current=field
				}else{
					panic("field error:"+ms[i])
				}
			}
		}
	default:
		log.Println("nothing to do")
	}

}

func(this *FuncListener) findField (method string,v reflect.Value) reflect.Value  {
	if v.Kind()==reflect.Ptr{
		v=v.Elem()
	}
	if field:=v.FieldByName(method);field.IsValid(){
		return field
	}
	return reflect.Value{}
}

type Admin struct {

}

func (this *Admin)Abc()  {
	fmt.Println("abc")
}

type User struct {
	Adm *Admin
}

func (this *User)Name(name string,age int64)  {
	fmt.Println("user name", name, age)
}