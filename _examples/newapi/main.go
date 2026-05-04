// newapi 示例 - 展示新版便捷函数和全局实例方法
//
// 本示例演示 color 包 v4.0 版本的新 API：
// - 便捷函数：Red(), Redf(), SRed(), SRedf() 等
// - 全局实例：GetGlobal() 获取实例，支持方法链式调用
//
// API 变化说明：
// - 不带 f 后缀的方法：只接受字符串，自动追加换行符
// - 带 f 后缀的方法：支持格式化字符串，不自动追加换行符
// - S 前缀的方法：返回带颜色的字符串（不打印）
// - S + f 后缀的方法：返回格式化后的带颜色字符串
package main

import (
	"fmt"
	"os"

	"gitee.com/MM-Q/color"
)

func main() {
	println("═══════════════════════════════════════════════════════════")
	println("  color 包 v4.0 新版 API 示例")
	println("═══════════════════════════════════════════════════════════")
	println()

	// ============================================================
	// 第一部分：便捷函数（直接调用）
	// ============================================================
	printSection("1. 便捷函数 - 基础颜色")

	// Red - 只接受字符串，自动追加换行
	color.Red("这是红色文本（自动换行）")

	// Redf - 支持格式化，需手动添加换行
	color.Redf("这是红色格式化文本，数字: %d，字符串: %s\n", 42, "hello")

	// SRed - 返回带颜色的字符串
	redStr := color.SRed("红色字符串")
	fmt.Printf("组合输出: %s %s\n", redStr, color.SGreen("绿色字符串"))

	// SRedf - 返回格式化后的带颜色字符串
	redfStr := color.SRedf("格式化红色: %d", 100)
	fmt.Printf("格式化组合: [%s] [%s]\n", redfStr, color.SGreenf("格式化绿色: %s", "OK"))

	println()

	// ============================================================
	// 第二部分：便捷函数 - 更多颜色
	// ============================================================
	printSection("2. 便捷函数 - 其他颜色")

	// 绿色系
	color.Green("绿色文本")
	color.Greenf("绿色格式化: %.2f\n", 3.14159)
	fmt.Println("字符串:", color.SGreen("绿"), color.SGreenf("值: %d", 10))

	// 蓝色系
	color.Blue("蓝色文本")
	color.Bluef("蓝色格式化: %s\n", "test")
	fmt.Println("字符串:", color.SBlue("蓝"), color.SBluef("值: %v", true))

	// 黄色系
	color.Yellow("黄色文本")
	color.Yellowf("黄色格式化: %d/%d\n", 1, 100)
	fmt.Println("字符串:", color.SYellow("黄"), color.SYellowf("进度: %d%%", 75))

	// 青色系
	color.Cyan("青色文本")
	color.Cyanf("青色格式化: %x\n", 255)
	fmt.Println("字符串:", color.SCyan("青"), color.SCyanf("十六进制: 0x%X", 255))

	// 洋红色系
	color.Magenta("洋红色文本")
	color.Magentaf("洋红色格式化: %t\n", true)
	fmt.Println("字符串:", color.SMagenta("洋红"), color.SMagentaf("布尔: %v", false))

	// 白色系
	color.White("白色文本")
	color.Whitef("白色格式化: %q\n", "quoted")
	fmt.Println("字符串:", color.SWhite("白"), color.SWhitef("引用: %s", "text"))

	// 灰色系
	color.Gray("灰色文本")
	color.Grayf("灰色格式化: %d\n", 50)
	fmt.Println("字符串:", color.SGray("灰"), color.SGrayf("数值: %d", 50))

	println()

	// ============================================================
	// 第三部分：高亮颜色
	// ============================================================
	printSection("3. 便捷函数 - 高亮颜色")

	// 高亮红色
	color.HiRed("高亮红色")
	color.HiRedf("高亮红格式化: %s\n", "HIGH")
	fmt.Println("字符串:", color.SHiRed("亮红"), color.SHiRedf("模式: %s", "ALERT"))

	// 高亮绿色
	color.HiGreen("高亮绿色")
	color.HiGreenf("高亮绿格式化: %s\n", "SUCCESS")
	fmt.Println("字符串:", color.SHiGreen("亮绿"), color.SHiGreenf("状态: %s", "OK"))

	// 高亮蓝色
	color.HiBlue("高亮蓝色")
	color.HiBluef("高亮蓝格式化: %s\n", "INFO")
	fmt.Println("字符串:", color.SHiBlue("亮蓝"), color.SHiBluef("级别: %s", "DEBUG"))

	println()

	// ============================================================
	// 第四部分：全局实例基础用法
	// ============================================================
	printSection("4. 全局实例 - 基础用法")

	// 获取全局实例
	c := color.GetGlobal()

	// 使用全局实例的颜色方法（不带 f，自动换行）
	c.Red("全局实例 - 红色")
	c.Green("全局实例 - 绿色")
	c.Blue("全局实例 - 蓝色")

	// 使用格式化方法（带 f，需手动换行）
	c.Redf("全局格式化红色: %d\n", 1)
	c.Greenf("全局格式化绿色: %s\n", "pass")
	c.Bluef("全局格式化蓝色: %.1f\n", 98.5)

	println()

	// ============================================================
	// 第五部分：全局实例返回字符串
	// ============================================================
	printSection("5. 全局实例 - 返回字符串")

	// 返回字符串方法
	gRed := c.SRed("全局红")
	gGreen := c.SGreen("全局绿")
	gBlue := c.SBlue("全局蓝")
	fmt.Printf("组合: %s | %s | %s\n", gRed, gGreen, gBlue)

	// 返回格式化字符串方法
	gRedf := c.SRedf("错误码: %d", 500)
	gGreenf := c.SGreenf("成功率: %.1f%%", 99.9)
	gBluef := c.SBluef("耗时: %dms", 42)
	fmt.Printf("格式化组合: %s, %s, %s\n", gRedf, gGreenf, gBluef)

	println()

	// ============================================================
	// 第六部分：全局实例样式配置
	// ============================================================
	printSection("6. 全局实例 - 样式配置")

	// 启用加粗
	c.SetBold(true)
	c.Yellow("加粗黄色")
	c.Yellowf("加粗黄色格式化: %s\n", "WARNING")
	fmt.Println("加粗字符串:", c.SYellow("粗黄"), c.SYellowf("粗黄格式化: %s", "BOLD"))

	// 启用下划线
	c.SetUnderline(true)
	c.Cyan("加粗+下划线青色")
	c.Cyanf("加粗+下划线青色格式化: %s\n", "IMPORTANT")

	// 重置配置
	c.SetConfig(&color.StyleConfig{})
	c.Green("恢复默认样式")

	println()

	// ============================================================
	// 第七部分：全局实例输出方向
	// ============================================================
	printSection("7. 全局实例 - 输出方向控制")

	// 输出到 stdout（默认）
	c.SetOutput(os.Stdout)
	c.Blue("输出到标准输出")

	// 输出到 stderr
	c.SetOutput(os.Stderr)
	c.Red("输出到标准错误")

	// 恢复 stdout
	c.SetOutput(os.Stdout)

	println()

	// ============================================================
	// 第八部分：实际应用场景
	// ============================================================
	printSection("8. 实际应用场景")

	// 场景1：日志级别输出
	printLog("INFO", "系统启动成功")
	printLog("WARN", "内存使用率超过 80%")
	printLog("ERROR", "数据库连接失败")
	printLog("DEBUG", "正在处理请求 #12345")

	println()

	// 场景2：状态标签
	printStatusTag("运行中", "running")
	printStatusTag("已停止", "stopped")
	printStatusTag("警告", "warning")
	printStatusTag("错误", "error")

	println()

	// 场景3：表格行输出
	printTableRow("服务A", "运行中", "2026-05-04 10:00")
	printTableRow("服务B", "已停止", "2026-05-04 09:30")
	printTableRow("服务C", "运行中", "2026-05-04 10:15")

	println()

	// ============================================================
	// 第九部分：混合使用便捷函数和全局实例
	// ============================================================
	printSection("9. 混合使用便捷函数和全局实例")

	// 便捷函数快速输出
	color.Cyan("=== 系统报告 ===")

	// 全局实例精细控制
	c.SetBold(true)
	c.Greenf("✓ 所有检查通过\n")

	// 组合字符串
	status := c.SGreen("正常")
	color.Whitef("系统状态: %s\n", status)

	// 重置并输出
	c.SetConfig(&color.StyleConfig{})
	color.Yellow("=== 报告结束 ===")

	println()

	// ============================================================
	// 结束
	// ============================================================
	printSection("示例结束")
	color.HiGreen("感谢使用 color 包 v4.0！")
	color.HiGreenf("文档地址: %s\n", "https://gitee.com/MM-Q/color")
}

// printSection 打印章节标题
func printSection(title string) {
	println()
	c := color.GetGlobal()
	c.SetBold(true)
	c.Cyanf("【%s】\n", title)
	c.SetConfig(&color.StyleConfig{})
}

// printLog 打印带级别的日志
func printLog(level, message string) {
	switch level {
	case "INFO":
		color.Cyanf("[INFO]  %s\n", message)
	case "WARN":
		color.Yellowf("[WARN]  %s\n", message)
	case "ERROR":
		color.Redf("[ERROR] %s\n", message)
	case "DEBUG":
		color.Magentaf("[DEBUG] %s\n", message)
	default:
		fmt.Printf("[%s] %s\n", level, message)
	}
}

// printStatusTag 打印状态标签
func printStatusTag(name, status string) {
	c := color.GetGlobal()
	switch status {
	case "running":
		tag := c.SGreen("●")
		fmt.Printf("%s %s\n", tag, color.SGreenf(" %-8s ", name))
	case "stopped":
		tag := c.SRed("●")
		fmt.Printf("%s %s\n", tag, color.SRedf(" %-8s ", name))
	case "warning":
		tag := c.SYellow("●")
		fmt.Printf("%s %s\n", tag, color.SYellowf(" %-8s ", name))
	case "error":
		tag := c.SHiRed("●")
		fmt.Printf("%s %s\n", tag, color.SHiRedf(" %-8s ", name))
	}
}

// printTableRow 打印表格行
func printTableRow(name, status, timeStr string) {
	c := color.GetGlobal()
	nameStr := c.SWhitef(" %-10s ", name)
	fmt.Print(nameStr)
	fmt.Print(" ")

	var statusStr string
	if status == "运行中" {
		statusStr = c.SGreenf(" %-8s ", status)
	} else {
		statusStr = c.SRedf(" %-8s ", status)
	}
	fmt.Print(statusStr)
	fmt.Print(" ")

	timeColorStr := c.SBlue(timeStr)
	fmt.Println(timeColorStr)
}
