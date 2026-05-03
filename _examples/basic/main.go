// color 包基础使用示例
// 本文件演示 color 包的各种常用 API 和函数
// 运行方式: go run main.go

package main

import (
	"os"

	"gitee.com/MM-Q/color"
)

func main() {
	// ============================================================
	// 第一部分: 便捷函数（最简单的方式）
	// ============================================================
	printSection("1. 便捷函数")

	// 基础颜色打印
	color.Red("红色文本 - 用于错误信息")
	color.Green("绿色文本 - 用于成功信息")
	color.Yellow("黄色文本 - 用于警告信息")
	color.Blue("蓝色文本 - 用于提示信息")
	color.Cyan("青色文本 - 用于信息展示")
	color.Magenta("洋红色文本 - 用于特殊信息")

	println() // 空行

	// 高亮颜色打印（更明亮）
	color.HiRed("高亮红色 - 更醒目的错误")
	color.HiGreen("高亮绿色 - 更醒目的成功")
	color.HiYellow("高亮黄色 - 更醒目的警告")
	color.HiBlue("高亮蓝色 - 更醒目的提示")
	color.HiCyan("高亮青色 - 更醒目的信息")
	color.HiMagenta("高亮洋红色 - 更醒目的特殊信息")

	println()

	// ============================================================
	// 第二部分: 格式化输出
	// ============================================================
	printSection("2. 格式化输出")

	// 使用 Printf 风格的格式化
	color.Red("错误代码: %d, 消息: %s", 500, "服务器内部错误")
	color.Green("成功处理了 %d 条记录", 42)
	color.Yellow("警告: 磁盘使用率已达 %.1f%%", 85.5)

	println()

	// 获取带颜色的字符串（不直接打印）
	redStr := color.RedString("红色字符串")
	greenStr := color.GreenString("绿色字符串")
	println("组合输出:", redStr, "+", greenStr)

	println()

	// ============================================================
	// 第三部分: 链式调用（组合样式）
	// ============================================================
	printSection("3. 链式调用组合样式")

	// 加粗红色
	boldRed := color.New(color.FgRed).Add(color.Bold)
	boldRed.Println("加粗红色文本")

	// 下划线绿色
	underlineGreen := color.New(color.FgGreen, color.Underline)
	underlineGreen.Println("下划线绿色文本")

	// 高亮蓝色 + 加粗 + 下划线
	fancy := color.New(color.FgHiBlue, color.Bold, color.Underline)
	fancy.Println("高亮蓝色加粗下划线")

	// 红色前景 + 白色背景
	redOnWhite := color.New(color.FgRed, color.BgWhite)
	redOnWhite.Println("红字白底")

	// 高亮黄色 + 黑色背景
	yellowOnBlack := color.New(color.FgHiYellow, color.BgBlack)
	yellowOnBlack.Println("高亮黄字黑底")

	println()

	// ============================================================
	// 第四部分: RGB 真彩色
	// ============================================================
	printSection("4. RGB 真彩色")

	// RGB 前景色
	color.RGB(255, 128, 0).Println("橙色前景 (255, 128, 0)")
	color.RGB(255, 0, 128).Println("粉色前景 (255, 0, 128)")
	color.RGB(0, 255, 128).Println("薄荷绿前景 (0, 255, 128)")

	println()

	// RGB 背景色
	color.BgRGB(0, 128, 255).Println("蓝色背景 (0, 128, 255)")

	// 组合 RGB 前景和背景
	color.New().
		AddRGB(255, 255, 255). // 白色前景
		AddBgRGB(255, 0, 0).   // 红色背景
		Add(color.Bold).       // 加粗
		Println("白字红底加粗")

	// 渐变色效果演示
	println("渐变色效果:")
	for i := 0; i < 5; i++ {
		g := 50 + i*40
		color.RGB(0, g, 255).Printf("█")
	}
	println()

	println()

	// ============================================================
	// 第五部分: 输出到不同目标
	// ============================================================
	printSection("5. 输出到不同目标")

	// 输出到标准错误
	color.New(color.FgRed).Fprintln(os.Stderr, "这是一条错误信息（输出到 stderr）")

	// 输出到文件
	file, err := os.CreateTemp("", "color-example-*.txt")
	if err == nil {
		defer file.Close()
		color.New(color.FgGreen).Fprintln(file, "这是写入文件的内容")
		color.Cyan("内容已写入文件: %s", file.Name())
	}

	println()

	// ============================================================
	// 第六部分: 函数生成器
	// ============================================================
	printSection("6. 函数生成器")

	// 创建可复用的打印函数
	danger := color.New(color.FgRed, color.Bold).PrintfFunc()
	success := color.New(color.FgGreen).PrintlnFunc()
	info := color.New(color.FgCyan).PrintlnFunc()

	// 多次使用这些函数
	danger("[%s] ", "错误")
	println("数据库连接失败")

	danger("[%s] ", "错误")
	println("网络请求超时")

	success("✓ 数据备份完成")
	success("✓ 缓存清理完成")

	info("ℹ 系统运行正常")
	info("ℹ 当前版本: v1.0.0")

	println()

	// ============================================================
	// 第七部分: 颜色控制
	// ============================================================
	printSection("7. 颜色控制")

	// 创建颜色对象
	c := color.New(color.FgYellow)

	// 正常输出
	c.Println("正常颜色输出")

	// 禁用颜色
	c.DisableColor()
	c.Println("颜色已禁用（纯文本）")

	// 重新启用颜色
	c.EnableColor()
	c.Println("颜色已重新启用")

	println()

	// ============================================================
	// 第八部分: 实际应用场景
	// ============================================================
	printSection("8. 实际应用场景")

	// 日志级别样式
	printLog("INFO", "系统启动成功")
	printLog("WARN", "内存使用率超过 80%")
	printLog("ERROR", "无法连接到数据库")
	printLog("DEBUG", "正在处理请求 #12345")
	printLog("SUCCESS", "操作完成")

	println()

	// 表格样式输出
	printTable()

	println()

	// ============================================================
	// 结束
	// ============================================================
	printSection("示例结束")
	color.HiGreen("感谢使用 color 包！")
}

// printSection 打印章节标题
func printSection(title string) {
	println()
	color.New(color.FgHiWhite, color.BgBlue, color.Bold).Printf(" %s ", title)
	println()
	println()
}

// printLog 打印带级别的日志
func printLog(level, message string) {
	switch level {
	case "INFO":
		color.Cyan("[INFO]  %s", message)
	case "WARN":
		color.Yellow("[WARN]  %s", message)
	case "ERROR":
		color.Red("[ERROR] %s", message)
	case "DEBUG":
		color.Magenta("[DEBUG] %s", message)
	case "SUCCESS":
		color.Green("[OK]    %s", message)
	default:
		println("["+level+"]", message)
	}
}

// printTable 打印示例表格
func printTable() {
	// 表头
	header := color.New(color.FgWhite, color.BgBlue, color.Bold)
	header.Printf(" %-10s ", "名称")
	header.Printf(" %-8s ", "状态")
	header.Printf(" %-15s ", "时间")
	println()

	// 数据行
	printRow("服务A", "运行中", "2026-05-03 10:00", true)
	printRow("服务B", "已停止", "2026-05-03 09:30", false)
	printRow("服务C", "运行中", "2026-05-03 10:15", true)
}

// printRow 打印表格行
func printRow(name, status, timeStr string, running bool) {
	color.White(" %-10s ", name)
	if running {
		color.Green(" %-8s ", status)
	} else {
		color.Red(" %-8s ", status)
	}
	color.Cyan(" %-15s ", timeStr)
	println()
}
