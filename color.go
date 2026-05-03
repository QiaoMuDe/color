package color

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

var (
	// NoColor 定义输出是否着色。它根据 stdout 的文件描述符是否指向终端
	// 动态设置为 false 或 true。如果设置了 NO_COLOR 环境变量（无论其值是什么），
	// 它也会被设置为 true。这是一个全局选项，影响所有颜色。
	// 如需对每个颜色块进行更多控制，请单独使用 DisableColor() 方法。
	NoColor = noColorIsSet() || os.Getenv("TERM") == "dumb" || !stdoutIsTerminal()

	// Output 定义打印函数的标准输出。默认使用 stdOut()。
	Output = stdOut()

	// Error 定义打印函数的标准错误输出。默认使用 stdErr()。
	Error = stdErr()

	// colorsCache 用于减少创建的 Color 对象数量，并允许重用已创建的具有所需 Attribute 的对象。
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // 保护 colorsCache
)

// noColorIsSet 如果环境变量 NO_COLOR 设置为非空字符串，则返回 true。
func noColorIsSet() bool {
	return os.Getenv("NO_COLOR") != ""
}

// stdoutIsTerminal 如果 os.Stdout 是终端，则返回 true。
// 如果 os.Stdout 为 nil（例如作为 Windows 服务运行时），则返回 false。
func stdoutIsTerminal() bool {
	if os.Stdout == nil {
		return false
	}
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// stdOut 返回用于颜色输出的写入器。
// 如果 os.Stdout 为 nil（例如作为 Windows 服务运行时），则返回 io.Discard。
func stdOut() io.Writer {
	if os.Stdout == nil {
		return io.Discard
	}
	return colorable.NewColorableStdout()
}

// stdErr 返回用于颜色错误输出的写入器。
// 如果 os.Stderr 为 nil（例如作为 Windows 服务运行时），则返回 io.Discard。
func stdErr() io.Writer {
	if os.Stderr == nil {
		return io.Discard
	}
	return colorable.NewColorableStderr()
}

// Color 定义一个由 SGR 参数定义的自定义颜色对象。
type Color struct {
	params  []Attribute
	noColor *bool
}

// Attribute 定义单个 SGR 代码
type Attribute int

const escape = "\x1b"

// 基础属性
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

const (
	ResetBold Attribute = iota + 22
	ResetItalic
	ResetUnderline
	ResetBlinking
	_
	ResetReversed
	ResetConcealed
	ResetCrossedOut
)

var mapResetAttributes map[Attribute]Attribute = map[Attribute]Attribute{
	Bold:         ResetBold,
	Faint:        ResetBold,
	Italic:       ResetItalic,
	Underline:    ResetUnderline,
	BlinkSlow:    ResetBlinking,
	BlinkRapid:   ResetBlinking,
	ReverseVideo: ResetReversed,
	Concealed:    ResetConcealed,
	CrossedOut:   ResetCrossedOut,
}

// 前景文本颜色
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	// 内部用于 256 色和 24 位着色
	foreground
)

// 前景高亮文本颜色
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// 背景文本颜色
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite

	// 内部用于 256 色和 24 位着色
	background
)

// 背景高亮文本颜色
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New 返回一个新创建的颜色对象。
func New(value ...Attribute) *Color {
	c := &Color{
		params: make([]Attribute, 0),
	}

	if noColorIsSet() {
		c.noColor = boolPtr(true)
	}

	c.Add(value...)
	return c
}

// RGB 返回一个新的 24 位 RGB 前景色。
func RGB(r, g, b int) *Color {
	return New(foreground, 2, Attribute(r), Attribute(g), Attribute(b))
}

// BgRGB 返回一个新的 24 位 RGB 背景色。
func BgRGB(r, g, b int) *Color {
	return New(background, 2, Attribute(r), Attribute(g), Attribute(b))
}

// AddRGB 用于链式添加前景 RGB SGR 参数。可以使用任意数量的参数进行组合
// 并创建自定义颜色对象。示例：.Add(34, 0, 12).Add(255, 128, 0)。
func (c *Color) AddRGB(r, g, b int) *Color {
	c.params = append(c.params, foreground, 2, Attribute(r), Attribute(g), Attribute(b))
	return c
}

// AddBgRGB 用于链式添加背景 RGB SGR 参数。可以使用任意数量的参数进行组合
// 并创建自定义颜色对象。示例：.Add(34, 0, 12).Add(255, 128, 0)。
func (c *Color) AddBgRGB(r, g, b int) *Color {
	c.params = append(c.params, background, 2, Attribute(r), Attribute(g), Attribute(b))
	return c
}

// Set 立即设置给定的参数。它将使用给定的 SGR 参数更改输出颜色，
// 直到调用 color.Unset() 为止。
func Set(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset 重置所有转义属性并清除输出。通常在 Set() 之后调用。
func Unset() {
	if NoColor {
		return
	}

	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set 设置 SGR 序列。
func (c *Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprint(Output, c.format())
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}

// SetWriter 用于使用给定的 io.Writer 设置 SGR 序列。这是一个底层函数，
// 用户应该使用更高级的函数，如 color.Fprint、color.Print 等。
func (c *Color) SetWriter(w io.Writer) *Color {
	_, _ = c.setWriter(w)
	return c
}

func (c *Color) setWriter(w io.Writer) (int, error) {
	if c.isNoColorSet() {
		return 0, nil
	}

	return fmt.Fprint(w, c.format())
}

// UnsetWriter 使用给定的 io.Writer 重置所有转义属性并清除输出。
// 通常在 SetWriter() 之后调用。
func (c *Color) UnsetWriter(w io.Writer) {
	_, _ = c.unsetWriter(w)
}

func (c *Color) unsetWriter(w io.Writer) (int, error) {
	if c.isNoColorSet() {
		return 0, nil
	}

	return fmt.Fprintf(w, "%s[%dm", escape, Reset)
}

// Add 用于链式添加 SGR 参数。可以使用任意数量的参数进行组合
// 并创建自定义颜色对象。示例：Add(color.FgRed, color.Underline)。
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

// Fprint 使用其操作数的默认格式进行格式化并写入 w。
// 当操作数都不是字符串时，在它们之间添加空格。
// 它返回写入的字节数和遇到的任何写入错误。
// 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。
func (c *Color) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	n, err = c.setWriter(w)
	if err != nil {
		return n, err
	}

	nn, err := fmt.Fprint(w, a...)
	n += nn
	if err != nil {
		return
	}

	nn, err = c.unsetWriter(w)
	n += nn
	return n, err
}

// Print 使用其操作数的默认格式进行格式化并写入标准输出。
// 当操作数都不是字符串时，在它们之间添加空格。
// 它返回写入的字节数和遇到的任何写入错误。
// 这是用给定颜色包装的标准 fmt.Print() 方法。
func (c *Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

// Fprintf 根据格式说明符进行格式化并写入 w。
// 它返回写入的字节数和遇到的任何写入错误。
// 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	n, err = c.setWriter(w)
	if err != nil {
		return n, err
	}

	nn, err := fmt.Fprintf(w, format, a...)
	n += nn
	if err != nil {
		return
	}

	nn, err = c.unsetWriter(w)
	n += nn
	return n, err
}

// Printf 根据格式说明符进行格式化并写入标准输出。
// 它返回写入的字节数和遇到的任何写入错误。
// 这是用给定颜色包装的标准 fmt.Printf() 方法。
func (c *Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

// Fprintln 使用其操作数的默认格式进行格式化并写入 w。
// 操作数之间始终添加空格，并追加换行符。
// 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, c.wrap(sprintln(a...)))
}

// Println 使用其操作数的默认格式进行格式化并写入标准输出。
// 操作数之间始终添加空格，并追加换行符。
// 它返回写入的字节数和遇到的任何写入错误。
// 这是用给定颜色包装的标准 fmt.Print() 方法。
func (c *Color) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Output, c.wrap(sprintln(a...)))
}

// Sprint 类似于 Print，但返回字符串而不是打印它。
func (c *Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln 类似于 Println，但返回字符串而不是打印它。
func (c *Color) Sprintln(a ...interface{}) string {
	return c.wrap(sprintln(a...)) + "\n"
}

// Sprintf 类似于 Printf，但返回字符串而不是打印它。
func (c *Color) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// FprintFunc 返回一个新函数，该函数使用 color.Fprint() 将传入的参数打印为彩色。
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprint(w, a...)
	}
}

// PrintFunc 返回一个新函数，该函数使用 color.Print() 将传入的参数打印为彩色。
func (c *Color) PrintFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Print(a...)
	}
}

// FprintfFunc 返回一个新函数，该函数使用 color.Fprintf() 将传入的参数打印为彩色。
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{}) {
	return func(w io.Writer, format string, a ...interface{}) {
		c.Fprintf(w, format, a...)
	}
}

// PrintfFunc 返回一个新函数，该函数使用 color.Printf() 将传入的参数打印为彩色。
func (c *Color) PrintfFunc() func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		c.Printf(format, a...)
	}
}

// FprintlnFunc 返回一个新函数，该函数使用 color.Fprintln() 将传入的参数打印为彩色。
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		c.Fprintln(w, a...)
	}
}

// PrintlnFunc 返回一个新函数，该函数使用 color.Println() 将传入的参数打印为彩色。
func (c *Color) PrintlnFunc() func(a ...interface{}) {
	return func(a ...interface{}) {
		c.Println(a...)
	}
}

// SprintFunc 返回一个新函数，该函数使用 fmt.Sprint() 为给定参数返回彩色字符串。
// 可用于放入或混合到其他字符串中。Windows 用户应将其与 color.Output 一起使用，示例：
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// SprintfFunc 返回一个新函数，该函数使用 fmt.Sprintf() 为给定参数返回彩色字符串。
// 可用于放入或混合到其他字符串中。Windows 用户应将其与 color.Output 一起使用。
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// SprintlnFunc 返回一个新函数，该函数使用 fmt.Sprintln() 为给定参数返回彩色字符串。
// 可用于放入或混合到其他字符串中。Windows 用户应将其与 color.Output 一起使用。
func (c *Color) SprintlnFunc() func(a ...interface{}) string {
	return func(a ...interface{}) string {
		return c.wrap(sprintln(a...)) + "\n"
	}
}

// sequence 返回格式化的 SGR 序列，用于插入 "\x1b[...m"
// 示例输出可能是："1;36" -> 粗体青色
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap 用颜色属性包装字符串 s。该字符串可以直接打印。
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	//return fmt.Sprintf("%s[%dm", escape, Reset)
	//对于序列中的每个元素，让我们使用特定的重置转义，如果未找到则使用通用重置
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(Reset))
		ra, ok := mapResetAttributes[v]
		if ok {
			format[i] = strconv.Itoa(int(ra))
		}
	}

	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

// DisableColor 禁用颜色输出。可用于在不更改任何现有代码的情况下仍然能够输出。
// 可用于 "--no-color" 等标志。要重新启用，请使用 EnableColor() 方法。
func (c *Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor 启用颜色输出。与 DisableColor() 一起使用。否则，此方法没有副作用。
func (c *Color) EnableColor() {
	c.noColor = boolPtr(false)
}

func (c *Color) isNoColorSet() bool {
	// 首先检查是否有用户设置的选项
	if c.noColor != nil {
		return *c.noColor
	}

	// 如果没有，则返回全局选项，默认情况下是禁用的
	return NoColor
}

// Equals 返回一个布尔值，指示两种颜色是否相等。
func (c *Color) Equals(c2 *Color) bool {
	if c == nil && c2 == nil {
		return true
	}
	if c == nil || c2 == nil {
		return false
	}

	if len(c.params) != len(c2.params) {
		return false
	}

	counts := make(map[Attribute]int, len(c.params))
	for _, attr := range c.params {
		counts[attr]++
	}

	for _, attr := range c2.params {
		if counts[attr] == 0 {
			return false
		}
		counts[attr]--
	}

	return true
}

func boolPtr(v bool) *bool {
	return &v
}

func getCachedColor(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		c = New(p)
		colorsCache[p] = c
	}

	return c
}

func colorPrint(format string, p Attribute, a ...interface{}) {
	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	if len(a) == 0 {
		c.Print(format)
	} else {
		c.Printf(format, a...)
	}
}

func colorString(format string, p Attribute, a ...interface{}) string {
	c := getCachedColor(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

// Black 是一个便捷的辅助函数，用于以黑色前景打印。默认会在 format 末尾追加换行符。
func Black(format string, a ...interface{}) { colorPrint(format, FgBlack, a...) }

// Red 是一个便捷的辅助函数，用于以红色前景打印。默认会在 format 末尾追加换行符。
func Red(format string, a ...interface{}) { colorPrint(format, FgRed, a...) }

// Green 是一个便捷的辅助函数，用于以绿色前景打印。默认会在 format 末尾追加换行符。
func Green(format string, a ...interface{}) { colorPrint(format, FgGreen, a...) }

// Yellow 是一个便捷的辅助函数，用于以黄色前景打印。默认会在 format 末尾追加换行符。
func Yellow(format string, a ...interface{}) { colorPrint(format, FgYellow, a...) }

// Blue 是一个便捷的辅助函数，用于以蓝色前景打印。默认会在 format 末尾追加换行符。
func Blue(format string, a ...interface{}) { colorPrint(format, FgBlue, a...) }

// Magenta 是一个便捷的辅助函数，用于以洋红色前景打印。默认会在 format 末尾追加换行符。
func Magenta(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }

// Cyan 是一个便捷的辅助函数，用于以青色前景打印。默认会在 format 末尾追加换行符。
func Cyan(format string, a ...interface{}) { colorPrint(format, FgCyan, a...) }

// White 是一个便捷的辅助函数，用于以白色前景打印。默认会在 format 末尾追加换行符。
func White(format string, a ...interface{}) { colorPrint(format, FgWhite, a...) }

// BlackString 是一个便捷的辅助函数，用于返回带有黑色前景的字符串。
func BlackString(format string, a ...interface{}) string { return colorString(format, FgBlack, a...) }

// RedString 是一个便捷的辅助函数，用于返回带有红色前景的字符串。
func RedString(format string, a ...interface{}) string { return colorString(format, FgRed, a...) }

// GreenString 是一个便捷的辅助函数，用于返回带有绿色前景的字符串。
func GreenString(format string, a ...interface{}) string { return colorString(format, FgGreen, a...) }

// YellowString 是一个便捷的辅助函数，用于返回带有黄色前景的字符串。
func YellowString(format string, a ...interface{}) string { return colorString(format, FgYellow, a...) }

// BlueString 是一个便捷的辅助函数，用于返回带有蓝色前景的字符串。
func BlueString(format string, a ...interface{}) string { return colorString(format, FgBlue, a...) }

// MagentaString 是一个便捷的辅助函数，用于返回带有洋红色前景的字符串。
func MagentaString(format string, a ...interface{}) string {
	return colorString(format, FgMagenta, a...)
}

// CyanString 是一个便捷的辅助函数，用于返回带有青色前景的字符串。
func CyanString(format string, a ...interface{}) string { return colorString(format, FgCyan, a...) }

// WhiteString 是一个便捷的辅助函数，用于返回带有白色前景的字符串。
func WhiteString(format string, a ...interface{}) string { return colorString(format, FgWhite, a...) }

// HiBlack 是一个便捷的辅助函数，用于以高亮黑色前景打印。默认会在 format 末尾追加换行符。
func HiBlack(format string, a ...interface{}) { colorPrint(format, FgHiBlack, a...) }

// HiRed 是一个便捷的辅助函数，用于以高亮红色前景打印。默认会在 format 末尾追加换行符。
func HiRed(format string, a ...interface{}) { colorPrint(format, FgHiRed, a...) }

// HiGreen 是一个便捷的辅助函数，用于以高亮绿色前景打印。默认会在 format 末尾追加换行符。
func HiGreen(format string, a ...interface{}) { colorPrint(format, FgHiGreen, a...) }

// HiYellow 是一个便捷的辅助函数，用于以高亮黄色前景打印。默认会在 format 末尾追加换行符。
func HiYellow(format string, a ...interface{}) { colorPrint(format, FgHiYellow, a...) }

// HiBlue 是一个便捷的辅助函数，用于以高亮蓝色前景打印。默认会在 format 末尾追加换行符。
func HiBlue(format string, a ...interface{}) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta 是一个便捷的辅助函数，用于以高亮洋红色前景打印。默认会在 format 末尾追加换行符。
func HiMagenta(format string, a ...interface{}) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan 是一个便捷的辅助函数，用于以高亮青色前景打印。默认会在 format 末尾追加换行符。
func HiCyan(format string, a ...interface{}) { colorPrint(format, FgHiCyan, a...) }

// HiWhite 是一个便捷的辅助函数，用于以高亮白色前景打印。默认会在 format 末尾追加换行符。
func HiWhite(format string, a ...interface{}) { colorPrint(format, FgHiWhite, a...) }

// HiBlackString 是一个便捷的辅助函数，用于返回带有高亮黑色前景的字符串。
func HiBlackString(format string, a ...interface{}) string {
	return colorString(format, FgHiBlack, a...)
}

// HiRedString 是一个便捷的辅助函数，用于返回带有高亮红色前景的字符串。
func HiRedString(format string, a ...interface{}) string { return colorString(format, FgHiRed, a...) }

// HiGreenString 是一个便捷的辅助函数，用于返回带有高亮绿色前景的字符串。
func HiGreenString(format string, a ...interface{}) string {
	return colorString(format, FgHiGreen, a...)
}

// HiYellowString 是一个便捷的辅助函数，用于返回带有高亮黄色前景的字符串。
func HiYellowString(format string, a ...interface{}) string {
	return colorString(format, FgHiYellow, a...)
}

// HiBlueString 是一个便捷的辅助函数，用于返回带有高亮蓝色前景的字符串。
func HiBlueString(format string, a ...interface{}) string { return colorString(format, FgHiBlue, a...) }

// HiMagentaString 是一个便捷的辅助函数，用于返回带有高亮洋红色前景的字符串。
func HiMagentaString(format string, a ...interface{}) string {
	return colorString(format, FgHiMagenta, a...)
}

// HiCyanString 是一个便捷的辅助函数，用于返回带有高亮青色前景的字符串。
func HiCyanString(format string, a ...interface{}) string { return colorString(format, FgHiCyan, a...) }

// HiWhiteString 是一个便捷的辅助函数，用于返回带有高亮白色前景的字符串。
func HiWhiteString(format string, a ...interface{}) string {
	return colorString(format, FgHiWhite, a...)
}

// sprintln 是一个辅助函数，用于使用 fmt.Sprintln 格式化字符串并去除末尾的换行符。
func sprintln(a ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}
