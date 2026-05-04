//go:build ignore

// 此文件用于生成 global.go 中的颜色快捷方法
// 运行方式: go generate 或 go run gen/global_gen.go

package main

import (
	"os"
	"text/template"
)

// 颜色定义列表
type ColorDef struct {
	Name string // 函数名（如 Red）
	Attr string // 属性常量（如 FgRed）
	Desc string // 中文描述
}

var colors = []ColorDef{
	{"Black", "FgBlack", "黑色"},
	{"Red", "FgRed", "红色"},
	{"Green", "FgGreen", "绿色"},
	{"Yellow", "FgYellow", "黄色"},
	{"Blue", "FgBlue", "蓝色"},
	{"Magenta", "FgMagenta", "洋红色"},
	{"Cyan", "FgCyan", "青色"},
	{"White", "FgWhite", "白色"},
	{"Gray", "FgHiBlack", "灰色"},
	{"HiBlack", "FgHiBlack", "高亮黑色"},
	{"HiRed", "FgHiRed", "高亮红色"},
	{"HiGreen", "FgHiGreen", "高亮绿色"},
	{"HiYellow", "FgHiYellow", "高亮黄色"},
	{"HiBlue", "FgHiBlue", "高亮蓝色"},
	{"HiMagenta", "FgHiMagenta", "高亮洋红色"},
	{"HiCyan", "FgHiCyan", "高亮青色"},
	{"HiWhite", "FgHiWhite", "高亮白色"},
}

const tmpl = `package color

// 代码由 gen/global_gen.go 自动生成，请勿手动修改
// 生成命令: go generate . || go run gen/global_gen.go
{{range .}}
// {{.Name}} 使用{{.Desc}}样式打印
func (g *GlobalColor) {{.Name}}(message string) { g.printColor({{.Attr}}, "%s\n", message) }

// {{.Name}}f 使用{{.Desc}}样式打印
func (g *GlobalColor) {{.Name}}f(format string, a ...interface{}) { g.printColor({{.Attr}}, format, a...) }

// S{{.Name}} 返回{{.Desc}}样式的字符串
func (g *GlobalColor) S{{.Name}}(message string) string { return g.sprintColor({{.Attr}}, "%s", message) }

// S{{.Name}}f 返回{{.Desc}}样式的字符串
func (g *GlobalColor) S{{.Name}}f(format string, a ...interface{}) string { return g.sprintColor({{.Attr}}, format, a...) }
{{end}}
`

func main() {
	t := template.Must(template.New("global").Parse(tmpl))

	f, err := os.Create("global_methods.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, colors)
	if err != nil {
		panic(err)
	}
}
