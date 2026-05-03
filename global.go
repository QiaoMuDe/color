package color

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// GlobalColor 是全局颜色实例的类型
// 包含独立的配置和输出设置，与普通的 Color 实例完全分离
type GlobalColor struct {
	config *StyleConfig // 样式配置
	color  *Color       // 复用的 Color 对象
	mu     sync.RWMutex // 读写锁，保证线程安全
}

// StyleConfig 定义颜色样式配置
type StyleConfig struct {
	NoColor    bool      // 是否禁用颜色
	Bold       bool      // 加粗
	Underline  bool      // 下划线
	Italic     bool      // 斜体
	Blink      bool      // 闪烁
	Faint      bool      // 暗淡
	CrossedOut bool      // 删除线
	Output     io.Writer // 输出目标
}

// Clone 创建样式配置的深拷贝
// 返回一个新的 StyleConfig 实例，复制当前配置的所有字段
// 注意: Output 字段是接口类型，只会复制引用，不会深拷贝底层的写入器
//
// 返回值:
//   - *StyleConfig: 新的样式配置实例
func (s *StyleConfig) Clone() *StyleConfig {
	if s == nil {
		return defaultStyleConfig()
	}
	return &StyleConfig{
		NoColor:    s.NoColor,
		Bold:       s.Bold,
		Underline:  s.Underline,
		Italic:     s.Italic,
		Blink:      s.Blink,
		Faint:      s.Faint,
		CrossedOut: s.CrossedOut,
		Output:     s.Output,
	}
}

// defaultStyleConfig 返回默认样式配置
// 默认启用颜色输出和加粗样式
// NoColor 会根据全局 NoColor 变量自动判断（考虑环境变量和终端检测）
func defaultStyleConfig() *StyleConfig {
	return &StyleConfig{
		NoColor:    NoColor, // 使用全局 NoColor 判断（包含 NO_COLOR 环境变量和终端检测）
		Bold:       true,    // 默认启用加粗
		Underline:  false,
		Italic:     false,
		Blink:      false,
		Faint:      false,
		CrossedOut: false,
		Output:     os.Stdout,
	}
}

// 全局实例
var (
	globalOnce sync.Once
	globalInst *GlobalColor
)

// initGlobal 初始化全局实例
func initGlobal() {
	globalOnce.Do(func() {
		globalInst = &GlobalColor{
			config: defaultStyleConfig(),
			color:  New(),
		}
	})
}

// GetGlobal 返回全局颜色实例
func GetGlobal() *GlobalColor {
	initGlobal()
	return globalInst
}

// G 是 GetGlobal 的快捷方式，返回全局颜色实例
// 使用更短的函数名，方便频繁调用
//
// 示例:
//
//	c := color.G()
//	c.Red("红色文字")
//	c.Info("信息日志")
func G() *GlobalColor {
	return GetGlobal()
}

// ResetGlobal 重置全局实例到默认状态
func ResetGlobal() {
	globalOnce = sync.Once{}
	globalInst = nil
	initGlobal()
}

// ===========================================================
// 配置方法
// ===========================================================

// SetConfig 设置样式配置
// 会自动克隆传入的配置，防止外部修改影响全局实例
// 如果传入 nil，则使用默认配置
//
// 参数:
//   - config: 要设置的样式配置
//
// 返回:
//   - *GlobalColor: 当前 GlobalColor 对象，支持链式调用
func (g *GlobalColor) SetConfig(config *StyleConfig) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	if config == nil {
		g.config = defaultStyleConfig()
	} else {
		g.config = config.Clone()
	}
	return g
}

// GetConfig 获取当前样式配置
// 返回的是配置对象的指针，直接修改会影响全局实例
// 如果需要修改配置，建议使用 GetConfigClone() 或 SetConfig() 方法
//
// 返回值:
//   - *StyleConfig: 当前样式配置
func (g *GlobalColor) GetConfig() *StyleConfig {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.config
}

// GetConfigClone 获取当前样式配置的克隆
// 返回一个新的配置对象，修改它不会影响全局实例
// 适合用于基于当前配置创建新配置的场景
//
// 返回值:
//   - *StyleConfig: 当前样式配置的克隆
func (g *GlobalColor) GetConfigClone() *StyleConfig {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.config.Clone()
}

// SetOutput 设置输出目标
func (g *GlobalColor) SetOutput(w io.Writer) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	if w != nil {
		g.config.Output = w
	}
	return g
}

// SetNoColor 设置是否禁用颜色
func (g *GlobalColor) SetNoColor(noColor bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.NoColor = noColor
	return g
}

// SetBold 设置是否启用加粗
func (g *GlobalColor) SetBold(bold bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.Bold = bold
	return g
}

// SetUnderline 设置是否启用下划线
func (g *GlobalColor) SetUnderline(underline bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.Underline = underline
	return g
}

// SetItalic 设置是否启用斜体
func (g *GlobalColor) SetItalic(italic bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.Italic = italic
	return g
}

// SetBlink 设置是否启用闪烁
func (g *GlobalColor) SetBlink(blink bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.Blink = blink
	return g
}

// SetFaint 设置是否启用暗淡效果
func (g *GlobalColor) SetFaint(faint bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.Faint = faint
	return g
}

// SetCrossedOut 设置是否启用删除线
func (g *GlobalColor) SetCrossedOut(crossedOut bool) *GlobalColor {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.config.CrossedOut = crossedOut
	return g
}

// ===========================================================
// 内部方法
// ===========================================================

// buildParams 根据前景色和配置构建 SGR 参数列表
func (g *GlobalColor) buildParams(fgColor Attribute) []Attribute {
	config := g.config
	params := make([]Attribute, 0, 7) // 预分配容量：1个颜色 + 最多6个样式

	// 如果禁用颜色，返回空参数（所有样式效果都被禁用）
	if config.NoColor {
		return params
	}

	// 添加前景色
	params = append(params, fgColor)

	// 添加样式属性
	if config.Bold {
		params = append(params, Bold)
	}
	if config.Faint {
		params = append(params, Faint)
	}
	if config.Italic {
		params = append(params, Italic)
	}
	if config.Underline {
		params = append(params, Underline)
	}
	if config.Blink {
		params = append(params, BlinkSlow)
	}
	if config.CrossedOut {
		params = append(params, CrossedOut)
	}

	return params
}

// output 获取当前输出目标
func (g *GlobalColor) output() io.Writer {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if g.config.Output != nil {
		return g.config.Output
	}
	return os.Stdout
}

// printf 内部格式化输出方法
func (g *GlobalColor) printf(format string, a ...interface{}) {
	g.mu.RLock()
	c := g.color
	noColor := g.config.NoColor
	g.mu.RUnlock()

	if c == nil || noColor {
		_, _ = fmt.Fprintf(g.output(), format+"\n", a...)
		return
	}
	_, _ = fmt.Fprint(g.output(), c.Sprintf(format+"\n", a...))
}

// setColor 设置颜色并打印
func (g *GlobalColor) setColor(fgColor Attribute, format string, a ...interface{}) {
	g.mu.Lock()
	g.color.params = g.buildParams(fgColor)
	g.mu.Unlock()
	g.printf(format, a...)
}

// sprintColor 设置颜色并返回格式化字符串（不换行）
func (g *GlobalColor) sprintColor(fgColor Attribute, format string, a ...interface{}) string {
	g.mu.Lock()
	g.color.params = g.buildParams(fgColor)
	c := g.color
	noColor := g.config.NoColor
	g.mu.Unlock()

	if c == nil || noColor {
		return fmt.Sprintf(format, a...)
	}
	return c.Sprintf(format, a...)
}

// ===========================================================
// 颜色快捷方法
// ===========================================================

// Red 使用红色样式打印
func (g *GlobalColor) Red(format string, a ...interface{}) {
	g.setColor(FgRed, format, a...)
}

// Green 使用绿色样式打印
func (g *GlobalColor) Green(format string, a ...interface{}) {
	g.setColor(FgGreen, format, a...)
}

// Yellow 使用黄色样式打印
func (g *GlobalColor) Yellow(format string, a ...interface{}) {
	g.setColor(FgYellow, format, a...)
}

// Blue 使用蓝色样式打印
func (g *GlobalColor) Blue(format string, a ...interface{}) {
	g.setColor(FgBlue, format, a...)
}

// Cyan 使用青色样式打印
func (g *GlobalColor) Cyan(format string, a ...interface{}) {
	g.setColor(FgCyan, format, a...)
}

// Magenta 使用洋红色样式打印
func (g *GlobalColor) Magenta(format string, a ...interface{}) {
	g.setColor(FgMagenta, format, a...)
}

// White 使用白色样式打印
func (g *GlobalColor) White(format string, a ...interface{}) {
	g.setColor(FgWhite, format, a...)
}

// Black 使用黑色样式打印
func (g *GlobalColor) Black(format string, a ...interface{}) {
	g.setColor(FgBlack, format, a...)
}

// ===========================================================
// 返回字符串的颜色方法（不换行）
// ===========================================================

// SRed 返回红色样式的字符串
func (g *GlobalColor) SRed(format string, a ...interface{}) string {
	return g.sprintColor(FgRed, format, a...)
}

// SGreen 返回绿色样式的字符串
func (g *GlobalColor) SGreen(format string, a ...interface{}) string {
	return g.sprintColor(FgGreen, format, a...)
}

// SYellow 返回黄色样式的字符串
func (g *GlobalColor) SYellow(format string, a ...interface{}) string {
	return g.sprintColor(FgYellow, format, a...)
}

// SBlue 返回蓝色样式的字符串
func (g *GlobalColor) SBlue(format string, a ...interface{}) string {
	return g.sprintColor(FgBlue, format, a...)
}

// SCyan 返回青色样式的字符串
func (g *GlobalColor) SCyan(format string, a ...interface{}) string {
	return g.sprintColor(FgCyan, format, a...)
}

// SMagenta 返回洋红色样式的字符串
func (g *GlobalColor) SMagenta(format string, a ...interface{}) string {
	return g.sprintColor(FgMagenta, format, a...)
}

// SWhite 返回白色样式的字符串
func (g *GlobalColor) SWhite(format string, a ...interface{}) string {
	return g.sprintColor(FgWhite, format, a...)
}

// SBlack 返回黑色样式的字符串
func (g *GlobalColor) SBlack(format string, a ...interface{}) string {
	return g.sprintColor(FgBlack, format, a...)
}

// ===========================================================
// 高亮颜色方法
// ===========================================================

// HiRed 使用高亮红色样式打印
func (g *GlobalColor) HiRed(format string, a ...interface{}) {
	g.setColor(FgHiRed, format, a...)
}

// HiGreen 使用高亮绿色样式打印
func (g *GlobalColor) HiGreen(format string, a ...interface{}) {
	g.setColor(FgHiGreen, format, a...)
}

// HiYellow 使用高亮黄色样式打印
func (g *GlobalColor) HiYellow(format string, a ...interface{}) {
	g.setColor(FgHiYellow, format, a...)
}

// HiBlue 使用高亮蓝色样式打印
func (g *GlobalColor) HiBlue(format string, a ...interface{}) {
	g.setColor(FgHiBlue, format, a...)
}

// HiCyan 使用高亮青色样式打印
func (g *GlobalColor) HiCyan(format string, a ...interface{}) {
	g.setColor(FgHiCyan, format, a...)
}

// HiMagenta 使用高亮洋红色样式打印
func (g *GlobalColor) HiMagenta(format string, a ...interface{}) {
	g.setColor(FgHiMagenta, format, a...)
}

// HiWhite 使用高亮白色样式打印
func (g *GlobalColor) HiWhite(format string, a ...interface{}) {
	g.setColor(FgHiWhite, format, a...)
}

// ===========================================================
// 返回字符串的高亮颜色方法（不换行）
// ===========================================================

// SHiRed 返回高亮红色样式的字符串
func (g *GlobalColor) SHiRed(format string, a ...interface{}) string {
	return g.sprintColor(FgHiRed, format, a...)
}

// SHiGreen 返回高亮绿色样式的字符串
func (g *GlobalColor) SHiGreen(format string, a ...interface{}) string {
	return g.sprintColor(FgHiGreen, format, a...)
}

// SHiYellow 返回高亮黄色样式的字符串
func (g *GlobalColor) SHiYellow(format string, a ...interface{}) string {
	return g.sprintColor(FgHiYellow, format, a...)
}

// SHiBlue 返回高亮蓝色样式的字符串
func (g *GlobalColor) SHiBlue(format string, a ...interface{}) string {
	return g.sprintColor(FgHiBlue, format, a...)
}

// SHiCyan 返回高亮青色样式的字符串
func (g *GlobalColor) SHiCyan(format string, a ...interface{}) string {
	return g.sprintColor(FgHiCyan, format, a...)
}

// SHiMagenta 返回高亮洋红色样式的字符串
func (g *GlobalColor) SHiMagenta(format string, a ...interface{}) string {
	return g.sprintColor(FgHiMagenta, format, a...)
}

// SHiWhite 返回高亮白色样式的字符串
func (g *GlobalColor) SHiWhite(format string, a ...interface{}) string {
	return g.sprintColor(FgHiWhite, format, a...)
}

// ===========================================================
// 日志级别方法
// ===========================================================

// Info 以信息级别（青色）打印
func (g *GlobalColor) Info(format string, a ...interface{}) {
	g.Cyan("[INFO] "+format, a...)
}

// Success 以成功级别（绿色）打印
func (g *GlobalColor) Success(format string, a ...interface{}) {
	g.Green("[SUCCESS] "+format, a...)
}

// Warning 以警告级别（黄色）打印
func (g *GlobalColor) Warning(format string, a ...interface{}) {
	g.Yellow("[WARN] "+format, a...)
}

// Error 以错误级别（红色）打印
func (g *GlobalColor) Error(format string, a ...interface{}) {
	g.Red("[ERROR] "+format, a...)
}

// Debug 以调试级别（洋红色）打印
func (g *GlobalColor) Debug(format string, a ...interface{}) {
	g.Magenta("[DEBUG] "+format, a...)
}

// ===========================================================
// 终端提示信息前缀方法
// ===========================================================

// Ok 以 OK 前缀（绿色）打印成功信息
// 格式: "ok: xxx"
func (g *GlobalColor) Ok(format string, a ...interface{}) {
	g.Green("ok: "+format, a...)
}

// Warn 以 WARN 前缀（黄色）打印警告信息
// 格式: "warn: xxx"
func (g *GlobalColor) Warn(format string, a ...interface{}) {
	g.Yellow("warn: "+format, a...)
}

// Err 以 ERR 前缀（红色）打印错误信息
// 格式: "err: xxx"
func (g *GlobalColor) Err(format string, a ...interface{}) {
	g.Red("err: "+format, a...)
}

// ===========================================================
// 初始化函数
// ===========================================================

func init() {
	initGlobal()
}
