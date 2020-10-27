package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main()  {
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
}
