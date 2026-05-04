// markdown 示例 - 展示 Markdown 语法的终端 UI 组件
//
// 本示例演示如何在终端中渲染 Markdown 语法的 UI 组件：
// - 标题渲染（H1-H6）
// - 列表渲染（有序/无序）
// - 代码块渲染
// - 引用块渲染
// - 表格渲染
// - 强调文本（粗体、斜体、删除线）
// - 链接和分隔线
//
// 使用 color 包为不同的 Markdown 元素添加颜色样式
package main

import (
	"fmt"
	"strings"

	"gitee.com/MM-Q/color"
)

func main() {
	println("═══════════════════════════════════════════════════════════")
	println("  Markdown 语法终端渲染示例")
	println("═══════════════════════════════════════════════════════════")
	println()

	// ============================================================
	// 第一部分：标题渲染
	// ============================================================
	printSection("1. 标题渲染")

	renderH1("这是一级标题")
	renderH2("这是二级标题")
	renderH3("这是三级标题")
	renderH4("这是四级标题")
	renderH5("这是五级标题")
	renderH6("这是六级标题")

	println()

	// ============================================================
	// 第二部分：强调文本
	// ============================================================
	printSection("2. 强调文本")

	// 粗体
	renderBold("这是粗体文本")

	// 斜体
	renderItalic("这是斜体文本")

	// 粗斜体
	renderBoldItalic("这是粗斜体文本")

	// 删除线
	renderStrikethrough("这是删除线文本")

	// 行内代码
	renderInlineCode("fmt.Println()")

	// 组合使用
	println("组合效果:")
	renderMixedEmphasis()

	println()

	// ============================================================
	// 第三部分：列表渲染
	// ============================================================
	printSection("3. 列表渲染")

	// 无序列表
	items := []string{"第一项", "第二项", "第三项", "嵌套项示例"}
	renderUnorderedList(items)

	println()

	// 有序列表
	renderOrderedList([]string{"第一步", "第二步", "第三步"})

	println()

	// 任务列表
	tasks := []struct {
		text string
		done bool
	}{
		{"已完成任务", true},
		{"未完成任务", false},
		{"另一个已完成", true},
	}
	renderTaskList(tasks)

	println()

	// ============================================================
	// 第四部分：代码块渲染
	// ============================================================
	printSection("4. 代码块渲染")

	goCode := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`
	renderCodeBlock("go", goCode)

	jsonCode := `{
  "name": "color",
  "version": "4.0.0",
  "description": "终端颜色库"
}`
	renderCodeBlock("json", jsonCode)

	shellCode := `# 安装 color 包
go get gitee.com/MM-Q/color

# 运行示例
cd _examples/basic
go run main.go`
	renderCodeBlock("bash", shellCode)

	println()

	// ============================================================
	// 第五部分：引用块渲染
	// ============================================================
	printSection("5. 引用块渲染")

	renderBlockquote("这是一段引用文本，用于强调重要信息或引用他人的话。", "")

	renderBlockquote(
		"编程的艺术就是处理复杂性的艺术。",
		"— Edsger W. Dijkstra",
	)

	renderBlockquote(
		"简单是可靠的先决条件。",
		"— Tony Hoare",
	)

	println()

	// ============================================================
	// 第六部分：表格渲染
	// ============================================================
	printSection("6. 表格渲染")

	// Markdown 风格表格
	headers := []string{"功能", "支持", "说明"}
	rows := [][]string{
		{"基础颜色", "✓", "16 种标准颜色"},
		{"高亮颜色", "✓", "16 种高亮颜色"},
		{"RGB 真彩色", "✓", "24 位真彩色"},
		{"样式属性", "✓", "加粗、下划线等"},
	}
	renderMarkdownTable(headers, rows)

	println()

	// 带状态的表格
	statusHeaders := []string{"服务", "状态", "延迟"}
	statusRows := [][]string{
		{"API 服务", "运行中", "12ms"},
		{"数据库", "运行中", "8ms"},
		{"缓存", "警告", "45ms"},
		{"消息队列", "错误", "超时"},
	}
	renderStatusTable(statusHeaders, statusRows)

	println()

	// ============================================================
	// 第七部分：链接和分隔线
	// ============================================================
	printSection("7. 链接和分隔线")

	renderLink("color 包文档", "https://gitee.com/MM-Q/color")
	renderLink("Go 语言官网", "https://golang.org")
	renderLink("GitHub 镜像", "https://github.com/fatih/color")

	println()

	renderHorizontalRule()
	println("上方是分隔线")
	renderHorizontalRule()

	println()

	// ============================================================
	// 第八部分：综合示例 - Markdown 文档渲染
	// ============================================================
	printSection("8. 综合示例 - 完整 Markdown 文档")

	renderMarkdownDocument()

	println()

	// ============================================================
	// 结束
	// ============================================================
	printSection("示例结束")
	_, _ = color.New(color.FgHiGreen).Println("感谢使用 Markdown 终端渲染组件！")
}

// renderH1 渲染一级标题
func renderH1(text string) {
	// 蓝色粗体，下方有双下划线效果
	c := color.New(color.FgHiBlue, color.Bold)
	_, _ = c.Println(text)
	_, _ = color.New(color.FgBlue).Println(strings.Repeat("═", len(text)*2))
}

// renderH2 渲染二级标题
func renderH2(text string) {
	// 青色粗体，下方有单下划线效果
	c := color.New(color.FgHiCyan, color.Bold)
	_, _ = c.Println(text)
	_, _ = color.New(color.FgCyan).Println(strings.Repeat("─", len(text)*2))
}

// renderH3 渲染三级标题
func renderH3(text string) {
	// 绿色粗体
	_, _ = color.New(color.FgHiGreen, color.Bold).Printf("### %s\n", text)
}

// renderH4 渲染四级标题
func renderH4(text string) {
	// 黄色
	_, _ = color.New(color.FgHiYellow).Printf("#### %s\n", text)
}

// renderH5 渲染五级标题
func renderH5(text string) {
	// 洋红色
	_, _ = color.New(color.FgHiMagenta).Printf("##### %s\n", text)
}

// renderH6 渲染六级标题
func renderH6(text string) {
	// 灰色斜体
	_, _ = color.New(color.FgGray, color.Italic).Printf("###### %s\n", text)
}

// renderBold 渲染粗体
func renderBold(text string) {
	_, _ = color.New(color.Bold).Printf("**%s**\n", text)
}

// renderItalic 渲染斜体
func renderItalic(text string) {
	_, _ = color.New(color.Italic).Printf("*%s*\n", text)
}

// renderBoldItalic 渲染粗斜体
func renderBoldItalic(text string) {
	_, _ = color.New(color.Bold, color.Italic).Printf("***%s***\n", text)
}

// renderStrikethrough 渲染删除线
func renderStrikethrough(text string) {
	// 使用灰色模拟删除线效果
	_, _ = color.New(color.CrossedOut, color.FgGray).Printf("~~%s~~\n", text)
}

// renderInlineCode 渲染行内代码
func renderInlineCode(code string) {
	// 红色背景，白色文字
	_, _ = color.New(color.FgHiWhite, color.BgRed).Printf(" `%s` \n", code)
}

// renderMixedEmphasis 渲染组合强调
func renderMixedEmphasis() {
	// 这是 **粗体** 和 *斜体* 的 ***组合***
	_, _ = color.New().Print("这是 ")
	_, _ = color.New(color.Bold).Print("粗体")
	_, _ = color.New().Print(" 和 ")
	_, _ = color.New(color.Italic).Print("斜体")
	_, _ = color.New().Print(" 的 ")
	_, _ = color.New(color.Bold, color.Italic).Print("组合")
	println()
}

// renderUnorderedList 渲染无序列表
func renderUnorderedList(items []string) {
	bulletColor := color.New(color.FgHiCyan)
	textColor := color.New(color.FgWhite)

	for _, item := range items {
		_, _ = bulletColor.Print("  ● ")
		_, _ = textColor.Println(item)
	}
}

// renderOrderedList 渲染有序列表
func renderOrderedList(items []string) {
	numberColor := color.New(color.FgHiGreen)
	textColor := color.New(color.FgWhite)

	for i, item := range items {
		_, _ = numberColor.Printf("  %d. ", i+1)
		_, _ = textColor.Println(item)
	}
}

// renderTaskList 渲染任务列表
func renderTaskList(tasks []struct {
	text string
	done bool
}) {
	for _, task := range tasks {
		if task.done {
			_, _ = color.New(color.FgHiGreen).Print("  [✓] ")
			_, _ = color.New(color.FgGray).Print("~~")
			_, _ = color.New(color.FgGray, color.CrossedOut).Print(task.text)
			_, _ = color.New(color.FgGray).Println("~~")
		} else {
			_, _ = color.New(color.FgHiYellow).Print("  [ ] ")
			_, _ = color.New(color.FgWhite).Println(task.text)
		}
	}
}

// renderCodeBlock 渲染代码块
func renderCodeBlock(lang, code string) {
	// 顶部边框，显示语言
	langTag := fmt.Sprintf(" %s ", lang)
	_, _ = color.New(color.FgHiBlack).Print("┌─")
	_, _ = color.New(color.FgYellow, color.BgHiBlack).Print(langTag)
	_, _ = color.New(color.FgHiBlack).Println("─┐")

	// 代码内容
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		// 语法高亮（简单版本）
		highlighted := syntaxHighlight(line, lang)
		_, _ = color.New(color.FgHiBlack).Print("│ ")
		fmt.Print(highlighted)
		// 填充空格对齐右边框
		padding := 50 - len(line)
		if padding > 0 {
			fmt.Print(strings.Repeat(" ", padding))
		}
		_, _ = color.New(color.FgHiBlack).Println(" │")
	}

	// 底部边框
	_, _ = color.New(color.FgHiBlack).Print("└")
	_, _ = color.New(color.FgHiBlack).Println(strings.Repeat("─", 54) + "┘")
	println()
}

// syntaxHighlight 简单的语法高亮
func syntaxHighlight(line, lang string) string {
	switch lang {
	case "go":
		// 关键字高亮
		keywords := []string{"package", "import", "func", "var", "const", "type", "struct", "interface", "return", "if", "else", "for", "range", "switch", "case", "default"}
		result := line
		for _, kw := range keywords {
			if strings.Contains(result, kw) {
				result = strings.ReplaceAll(result, kw, color.SHiMagenta(kw))
			}
		}
		// 字符串高亮
		if strings.Contains(result, "\"") {
			parts := strings.Split(result, "\"")
			for i := 1; i < len(parts)-1; i += 2 {
				parts[i] = color.SGreen("\"" + parts[i] + "\"")
			}
			result = strings.Join(parts, "")
		}
		return result
	case "json":
		// JSON 键高亮
		if strings.Contains(line, "\"") && strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return color.SHiCyan(parts[0]) + ":" + parts[1]
			}
		}
		return line
	case "bash":
		// 注释高亮
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			return color.SGray(line)
		}
		// 命令高亮
		if strings.HasPrefix(line, "go ") {
			return color.SHiYellow("go") + line[2:]
		}
		return line
	default:
		return line
	}
}

// renderBlockquote 渲染引用块
func renderBlockquote(text, author string) {
	// 引用符号和左边框
	quoteColor := color.New(color.FgHiBlue)
	textColor := color.New(color.FgWhite)
	authorColor := color.New(color.FgGray, color.Italic)

	// 计算文本宽度以便对齐
	maxWidth := 50

	// 顶部装饰
	_, _ = quoteColor.Println("  ┌" + strings.Repeat("─", maxWidth) + "┐")

	// 引用内容
	_, _ = quoteColor.Print("  │ ")
	_, _ = textColor.Printf("%-*s", maxWidth-2, text)
	_, _ = quoteColor.Println(" │")

	// 作者（如果有）
	if author != "" {
		_, _ = quoteColor.Print("  │ ")
		_, _ = authorColor.Printf("%-*s", maxWidth-2, author)
		_, _ = quoteColor.Println(" │")
	}

	// 底部装饰
	_, _ = quoteColor.Println("  └" + strings.Repeat("─", maxWidth) + "┘")
	println()
}

// renderMarkdownTable 渲染 Markdown 风格表格
func renderMarkdownTable(headers []string, rows [][]string) {
	// 计算列宽
	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h) + 2
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell)+2 > colWidths[i] {
				colWidths[i] = len(cell) + 2
			}
		}
	}

	// 表头（蓝色背景）
	headerColor := color.New(color.FgHiWhite, color.BgBlue, color.Bold)
	_, _ = headerColor.Print("│")
	for i, h := range headers {
		_, _ = headerColor.Printf(" %-*s │", colWidths[i]-2, h)
	}
	println()

	// 分隔线
	sepColor := color.New(color.FgBlue)
	_, _ = sepColor.Print("├")
	for i, w := range colWidths {
		_, _ = sepColor.Print(strings.Repeat("─", w))
		if i < len(colWidths)-1 {
			_, _ = sepColor.Print("┼")
		} else {
			_, _ = sepColor.Print("┤")
		}
	}
	println()

	// 数据行
	for _, row := range rows {
		_, _ = color.New(color.FgWhite).Print("│")
		for i, cell := range row {
			if i < len(colWidths) {
				_, _ = color.New(color.FgWhite).Printf(" %-*s │", colWidths[i]-2, cell)
			}
		}
		println()
	}

	// 底部边框
	_, _ = sepColor.Print("└")
	for i, w := range colWidths {
		_, _ = sepColor.Print(strings.Repeat("─", w))
		if i < len(colWidths)-1 {
			_, _ = sepColor.Print("┴")
		} else {
			_, _ = sepColor.Print("┘")
		}
	}
	println()
}

// renderStatusTable 渲染带状态颜色的表格
func renderStatusTable(headers []string, rows [][]string) {
	// 表头
	headerColor := color.New(color.FgHiWhite, color.BgHiBlack, color.Bold)
	_, _ = headerColor.Print("  ")
	for _, h := range headers {
		_, _ = headerColor.Printf(" %-12s ", h)
	}
	println()

	// 数据行
	for _, row := range rows {
		_, _ = color.New(color.FgWhite).Print("  ")

		// 名称
		_, _ = color.New(color.FgWhite).Printf(" %-12s ", row[0])

		// 状态（带颜色）
		status := row[1]
		switch status {
		case "运行中":
			_, _ = color.New(color.FgHiGreen).Printf(" %-12s ", status)
		case "警告":
			_, _ = color.New(color.FgHiYellow).Printf(" %-12s ", status)
		case "错误":
			_, _ = color.New(color.FgHiRed).Printf(" %-12s ", status)
		default:
			_, _ = color.New(color.FgWhite).Printf(" %-12s ", status)
		}

		// 延迟
		_, _ = color.New(color.FgCyan).Printf(" %-12s ", row[2])
		println()
	}
}

// renderLink 渲染链接
func renderLink(text, url string) {
	_, _ = color.New(color.FgHiBlue, color.Underline).Print(text)
	_, _ = color.New(color.FgGray).Print(" → ")
	_, _ = color.New(color.FgCyan).Printf("%s\n", url)
}

// renderHorizontalRule 渲染分隔线
func renderHorizontalRule() {
	_, _ = color.New(color.FgHiBlack).Println(strings.Repeat("─", 60))
}

// renderMarkdownDocument 渲染完整的 Markdown 文档
func renderMarkdownDocument() {
	// 模拟一个 README.md 的渲染
	renderH1("Color 包")

	_, _ = color.New(color.FgWhite).Println("一个强大的 Go 语言终端颜色库，支持 16 色、256 色和 RGB 真彩色。")
	println()

	renderH2("特性")

	features := []string{
		"支持 16 种标准颜色",
		"支持 16 种高亮颜色",
		"支持 RGB 真彩色（24 位）",
		"支持样式属性（加粗、下划线等）",
		"支持链式调用",
		"跨平台支持（Windows、Linux、macOS）",
	}
	renderUnorderedList(features)

	println()
	renderH2("安装")

	installCode := `go get gitee.com/MM-Q/color`
	renderCodeBlock("bash", installCode)

	renderH2("快速开始")

	quickCode := `package main

import "gitee.com/MM-Q/color"

func main() {
    color.Red("红色文本")
    color.Green("绿色文本")
}`
	renderCodeBlock("go", quickCode)

	renderH2("链接")

	renderLink("项目主页", "https://gitee.com/MM-Q/color")
	renderLink("文档", "https://gitee.com/MM-Q/color/blob/main/README.md")
}

// printSection 打印章节标题
func printSection(title string) {
	println()
	c := color.New(color.FgHiWhite, color.BgBlue, color.Bold)
	_, _ = c.Printf(" %s ", title)
	println()
	println()
}
