// utils.go - 终端 UI 辅助函数库
// 本文件展示了如何使用 color 库构建常用的终端 UI 组件
// 这些函数可以作为参考，根据项目需求进行调整

package main

import (
	"fmt"
	"strings"

	"gitee.com/MM-Q/color"
)

// ============================================================
// 标题和分隔线
// ============================================================

// PrintSection 打印章节标题
// 使用蓝色背景和白色文字，带加粗样式
//
// 参数:
//   - title: 章节标题文本
//
// 示例:
//
//	PrintSection("第一部分: 基础用法")
func PrintSection(title string) {
	fmt.Println()
	_, _ = color.New(color.FgHiWhite, color.BgBlue, color.Bold).Printf(" %s ", title)
	fmt.Println()
	fmt.Println()
}

// PrintSubSection 打印子章节标题
// 使用青色前景和加粗样式
//
// 参数:
//   - title: 子章节标题文本
func PrintSubSection(title string) {
	fmt.Println()
	_, _ = color.New(color.FgCyan, color.Bold).Printf("▸ %s", title)
	fmt.Println()
	fmt.Println()
}

// PrintSeparator 打印分隔线
// 使用指定颜色的虚线分隔符
//
// 参数:
//   - c: 颜色属性
//   - length: 分隔线长度
func PrintSeparator(c color.Attribute, length int) {
	separator := strings.Repeat("-", length)
	_, _ = color.New(c).Println(separator)
}

// PrintDoubleSeparator 打印双分隔线
// 使用指定颜色的等号分隔符
//
// 参数:
//   - c: 颜色属性
//   - length: 分隔线长度
func PrintDoubleSeparator(c color.Attribute, length int) {
	separator := strings.Repeat("=", length)
	_, _ = color.New(c).Println(separator)
}

// ============================================================
// 日志级别打印
// ============================================================

// LogInfo 打印信息级别日志
// 格式: [INFO] 消息内容
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func LogInfo(format string, a ...interface{}) {
	color.Cyanf("[INFO]  "+format+"\n", a...)
}

// LogSuccess 打印成功级别日志
// 格式: [SUCCESS] 消息内容
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func LogSuccess(format string, a ...interface{}) {
	color.Greenf("[SUCCESS] "+format+"\n", a...)
}

// LogWarn 打印警告级别日志
// 格式: [WARN] 消息内容
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func LogWarn(format string, a ...interface{}) {
	color.Yellowf("[WARN]  "+format+"\n", a...)
}

// LogError 打印错误级别日志
// 格式: [ERROR] 消息内容
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func LogError(format string, a ...interface{}) {
	color.Redf("[ERROR] "+format+"\n", a...)
}

// LogDebug 打印调试级别日志
// 格式: [DEBUG] 消息内容
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func LogDebug(format string, a ...interface{}) {
	color.Magentaf("[DEBUG] "+format+"\n", a...)
}

// ============================================================
// 表格输出
// ============================================================

// Table 表格结构体
type Table struct {
	Headers []string
	Rows    [][]string
	Widths  []int
}

// NewTable 创建新表格
//
// 参数:
//   - headers: 表头数组
//   - widths: 每列宽度数组
//
// 返回值:
//   - *Table: 表格实例
//
// 示例:
//
//	table := NewTable([]string{"名称", "状态", "时间"}, []int{12, 10, 20})
func NewTable(headers []string, widths []int) *Table {
	return &Table{
		Headers: headers,
		Rows:    make([][]string, 0),
		Widths:  widths,
	}
}

// AddRow 添加一行数据
//
// 参数:
//   - row: 行数据数组
func (t *Table) AddRow(row []string) {
	t.Rows = append(t.Rows, row)
}

// Print 打印表格
func (t *Table) Print() {
	// 打印表头
	headerColor := color.New(color.FgWhite, color.BgBlue, color.Bold)
	for i, header := range t.Headers {
		if i < len(t.Widths) {
			_, _ = headerColor.Printf(" %-*s ", t.Widths[i], header)
		}
	}
	fmt.Println()

	// 打印分隔线
	separatorColor := color.New(color.FgBlue)
	totalWidth := 0
	for _, w := range t.Widths {
		totalWidth += w + 2 // +2 for spaces
	}
	_, _ = separatorColor.Println(strings.Repeat("─", totalWidth))

	// 打印数据行
	for _, row := range t.Rows {
		for i, cell := range row {
			if i < len(t.Widths) {
				// 使用 Sprint 获取字符串，避免自动换行
				coloredCell := color.New(color.FgWhite).Sprintf(" %-*s ", t.Widths[i], cell)
				fmt.Print(coloredCell)
			}
		}
		fmt.Println()
	}
}

// ============================================================
// 列表输出
// ============================================================

// PrintBulletList 打印无序列表
//
// 参数:
//   - items: 列表项数组
//   - bulletColor: 项目符号颜色
func PrintBulletList(items []string, bulletColor color.Attribute) {
	c := color.New(bulletColor)
	for _, item := range items {
		_, _ = c.Print("• ")
		fmt.Println(item)
	}
}

// PrintNumberedList 打印有序列表
//
// 参数:
//   - items: 列表项数组
//   - numberColor: 序号颜色
func PrintNumberedList(items []string, numberColor color.Attribute) {
	c := color.New(numberColor)
	for i, item := range items {
		_, _ = c.Printf("%d. ", i+1)
		fmt.Println(item)
	}
}

// ============================================================
// 状态标签
// ============================================================

// StatusTag 状态标签结构体
type StatusTag struct {
	Text    string
	Color   color.Attribute
	BgColor color.Attribute
}

// PrintStatus 打印状态标签
// 在文本周围添加颜色背景和前景
//
// 参数:
//   - text: 标签文本
//   - fg: 前景色
//   - bg: 背景色
func PrintStatus(text string, fg, bg color.Attribute) {
	_, _ = color.New(fg, bg, color.Bold).Printf(" %s ", text)
}

// PrintStatusRunning 打印"运行中"状态标签
func PrintStatusRunning() {
	PrintStatus("运行中", color.FgHiWhite, color.BgGreen)
}

// PrintStatusStopped 打印"已停止"状态标签
func PrintStatusStopped() {
	PrintStatus("已停止", color.FgHiWhite, color.BgRed)
}

// PrintStatusWarning 打印"警告"状态标签
func PrintStatusWarning() {
	PrintStatus("警告", color.FgHiWhite, color.BgYellow)
}

// PrintStatusPending 打印"待处理"状态标签
func PrintStatusPending() {
	PrintStatus("待处理", color.FgHiWhite, color.BgBlue)
}

// ============================================================
// 进度指示
// ============================================================

// PrintProgressBar 打印进度条
//
// 参数:
//   - current: 当前进度
//   - total: 总进度
//   - width: 进度条宽度
func PrintProgressBar(current, total, width int) {
	if total <= 0 {
		return
	}

	percent := float64(current) / float64(total)
	filled := int(float64(width) * percent)
	empty := width - filled

	// 打印进度条
	_, _ = color.New(color.FgGreen).Print("[")
	_, _ = color.New(color.FgHiGreen).Print(strings.Repeat("█", filled))
	_, _ = color.New(color.FgHiBlack).Print(strings.Repeat("░", empty))
	_, _ = color.New(color.FgGreen).Print("]")

	// 打印百分比
	_, _ = color.New(color.FgCyan).Printf(" %.1f%%", percent*100)
}

// PrintSpinner 打印旋转指示器（静态展示）
// 实际使用时需要配合 goroutine 实现动画效果
//
// 参数:
//   - frame: 当前帧索引
func PrintSpinner(frame int) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	c := color.New(color.FgCyan)
	_, _ = c.Printf("%s ", frames[frame%len(frames)])
}

// ============================================================
// 代码块和引用
// ============================================================

// PrintCodeBlock 打印代码块
//
// 参数:
//   - code: 代码内容
//   - lang: 语言标识（显示在右上角）
func PrintCodeBlock(code, lang string) {
	// 计算最大行长度
	maxLen := 40 // 最小宽度
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// 打印顶部边框
	fmt.Println()
	_, _ = color.New(color.FgHiBlack).Print("┌─ ")
	_, _ = color.New(color.FgYellow).Print(lang)
	_, _ = color.New(color.FgHiBlack).Print(" ")
	_, _ = color.New(color.FgHiBlack).Print(strings.Repeat("─", maxLen-len(lang)-1))
	_, _ = color.New(color.FgHiBlack).Println("┐")

	// 打印代码内容
	for _, line := range lines {
		_, _ = color.New(color.FgHiBlack).Print("│ ")
		_, _ = color.New(color.FgWhite).Print(line)
		// 填充空格对齐右边框
		if len(line) < maxLen {
			fmt.Print(strings.Repeat(" ", maxLen-len(line)))
		}
		_, _ = color.New(color.FgHiBlack).Println(" │")
	}

	// 打印底部边框
	_, _ = color.New(color.FgHiBlack).Print("└")
	_, _ = color.New(color.FgHiBlack).Print(strings.Repeat("─", maxLen+2))
	_, _ = color.New(color.FgHiBlack).Println("┘")
	fmt.Println()
}

// PrintQuote 打印引用文本
//
// 参数:
//   - text: 引用内容
//   - author: 作者（可选）
func PrintQuote(text, author string) {
	// 计算最大行长度
	maxLen := 40 // 最小宽度
	if len(text) > maxLen {
		maxLen = len(text)
	}
	if author != "" && len(author)+3 > maxLen { // "— " + author
		maxLen = len(author) + 3
	}

	fmt.Println()
	// 顶部边框
	_, _ = color.New(color.FgHiBlack).Print("┌")
	_, _ = color.New(color.FgHiBlack).Print(strings.Repeat("─", maxLen+2))
	_, _ = color.New(color.FgHiBlack).Println("┐")

	// 引用内容
	_, _ = color.New(color.FgHiBlack).Print("│ ")
	_, _ = color.New(color.FgCyan).Print(text)
	// 填充空格对齐右边框
	if len(text) < maxLen {
		fmt.Print(strings.Repeat(" ", maxLen-len(text)))
	}
	_, _ = color.New(color.FgHiBlack).Println(" │")

	// 作者
	if author != "" {
		authorLine := "— " + author
		_, _ = color.New(color.FgHiBlack).Print("│ ")
		_, _ = color.New(color.FgHiBlack).Print(authorLine)
		// 填充空格对齐右边框
		if len(authorLine) < maxLen {
			fmt.Print(strings.Repeat(" ", maxLen-len(authorLine)))
		}
		_, _ = color.New(color.FgHiBlack).Println(" │")
	}

	// 底部边框
	_, _ = color.New(color.FgHiBlack).Print("└")
	_, _ = color.New(color.FgHiBlack).Print(strings.Repeat("─", maxLen+2))
	_, _ = color.New(color.FgHiBlack).Println("┘")
	fmt.Println()
}

// ============================================================
// 提示框
// ============================================================

// PrintTip 打印提示框
//
// 参数:
//   - title: 提示标题
//   - message: 提示内容
func PrintTip(title, message string) {
	fmt.Println()
	_, _ = color.New(color.FgHiBlue, color.Bold).Printf("💡 %s: ", title)
	_, _ = color.New(color.FgWhite).Println(message)
	fmt.Println()
}

// PrintWarning 打印警告框
//
// 参数:
//   - message: 警告内容
func PrintWarning(message string) {
	fmt.Println()
	_, _ = color.New(color.FgHiYellow, color.Bold).Print("⚠️  警告: ")
	_, _ = color.New(color.FgYellow).Println(message)
	fmt.Println()
}

// PrintError 打印错误框
//
// 参数:
//   - message: 错误内容
func PrintError(message string) {
	fmt.Println()
	_, _ = color.New(color.FgHiRed, color.Bold).Print("❌ 错误: ")
	_, _ = color.New(color.FgRed).Println(message)
	fmt.Println()
}

// PrintSuccess 打印成功框
//
// 参数:
//   - message: 成功内容
func PrintSuccess(message string) {
	fmt.Println()
	_, _ = color.New(color.FgHiGreen, color.Bold).Print("✅ 成功: ")
	_, _ = color.New(color.FgGreen).Println(message)
	fmt.Println()
}
