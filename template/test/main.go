package test

import (
	"os"
	"text/template"
)

type Student struct {
	Name string
}

func basicExample() {
	//数据
	name := "Tom"
	//定义模板
	muban := "hello, {{.}}\n"
	//解析模板
	tmpl, err := template.New("test").Parse(muban)
	if err != nil {
		panic(err)
	}
	//数据驱动模板
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}

func structExample() {
	tom := Student{"Tom"}
	muban := "hello, {{.Name}}\n"
	tmpl, err := template.New("test").Parse(muban)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, tom)
	if err != nil {
		panic(err)
	}
}

func removeSpaces() {
	name := "jack"
	//定义模板
	muban := `
my name is {{- $name -}} HAHA
my name is {{- $name}} HAHA
my name is {{$name -}} HAHA
`
	//解析模板
	tmpl, err := template.New("test").Parse(muban)
	if err != nil {
		panic(err)
	}
	//数据驱动模板
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}

func main() {
	basicExample()
	structExample()
	removeSpaces()
}
