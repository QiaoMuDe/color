//go:build ignore

// 此文件用于生成 helper.go 中的便捷函数
// 运行方式: go generate 或 go run gen/helper_gen.go

package main

import (
	"os"
	"text/template"
)

// 颜色定义列表
type ColorDef struct {
	Name      string // 函数名（如 Red）
	Attr      string // 属性常量（如 FgRed）
	Desc      string // 中文描述
	IsHiColor bool   // 是否高亮色
}

var colors = []ColorDef{
	{"Black", "FgBlack", "黑色", false},
	{"Red", "FgRed", "红色", false},
	{"Green", "FgGreen", "绿色", false},
	{"Yellow", "FgYellow", "黄色", false},
	{"Blue", "FgBlue", "蓝色", false},
	{"Magenta", "FgMagenta", "洋红色", false},
	{"Cyan", "FgCyan", "青色", false},
	{"White", "FgWhite", "白色", false},
	{"Gray", "FgHiBlack", "灰色", false},
	{"HiBlack", "FgHiBlack", "高亮黑色", true},
	{"HiRed", "FgHiRed", "高亮红色", true},
	{"HiGreen", "FgHiGreen", "高亮绿色", true},
	{"HiYellow", "FgHiYellow", "高亮黄色", true},
	{"HiBlue", "FgHiBlue", "高亮蓝色", true},
	{"HiMagenta", "FgHiMagenta", "高亮洋红色", true},
	{"HiCyan", "FgHiCyan", "高亮青色", true},
	{"HiWhite", "FgHiWhite", "高亮白色", true},
}

const tmpl = `package color

// 代码由 gen/helper_gen.go 自动生成，请勿手动修改
// 生成命令: go generate . || go run gen/helper_gen.go
{{range .}}
// {{.Name}} 以{{.Desc}}前景打印文本
func {{.Name}}(message string) { colorPrint({{.Attr}}, "%s\n", message) }

// {{.Name}}f 以{{.Desc}}前景打印文本
func {{.Name}}f(format string, a ...interface{}) { colorPrint({{.Attr}}, format, a...) }

// S{{.Name}} 返回{{.Desc}}前景字符串
func S{{.Name}}(message string) string { return colorString({{.Attr}}, "%s", message) }

// S{{.Name}}f 返回{{.Desc}}前景字符串
func S{{.Name}}f(format string, a ...interface{}) string { return colorString({{.Attr}}, format, a...) }
{{end}}
`

func main() {
	t := template.Must(template.New("helper").Parse(tmpl))

	f, err := os.Create("helper.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, colors)
	if err != nil {
		panic(err)
	}
}
