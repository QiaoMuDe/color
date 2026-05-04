// 全局实例使用示例
// 演示如何使用 color 包的 GlobalColor 和 StyleConfig
// 运行方式: go run main.go

package main

import (
	"fmt"
	"os"

	"gitee.com/MM-Q/color"
)

func main() {
	// 获取全局实例（默认启用颜色 + 加粗）
	c := color.GetGlobal()

	// ============================================================
	// 第一部分：默认配置（启用颜色 + 加粗）
	// ============================================================
	printSection("1. 默认配置（启用颜色 + 加粗）")

	c.Red("红色加粗文本")
	c.Green("绿色加粗文本")
	c.Blue("蓝色加粗文本")

	println()

	// ============================================================
	// 第二部分：添加下划线（加粗保持）
	// ============================================================
	printSection("2. 添加下划线（加粗保持）")

	c.SetUnderline(true)

	c.Red("红色加粗下划线")
	c.Green("绿色加粗下划线")
	c.Cyan("青色加粗下划线")

	println()

	// ============================================================
	// 第三部分：禁用加粗
	// ============================================================
	printSection("3. 禁用加粗，只保留颜色")

	c.SetBold(false)
	c.SetUnderline(false)

	c.Red("红色文本（无加粗）")
	c.Green("绿色文本（无加粗）")
	c.Blue("蓝色文本（无加粗）")

	println()

	// ============================================================
	// 第四部分：禁用颜色（所有样式都被禁用）
	// ============================================================
	printSection("4. 禁用颜色（所有样式都被禁用）")

	// 先设置一些样式
	c.SetBold(true)
	c.SetUnderline(true)
	c.SetItalic(true)

	// 然后禁用颜色 - 这会同时禁用所有样式效果
	c.SetNoColor(true)

	c.Red("这段文字纯文本（颜色、加粗、下划线、斜体都被禁用）")
	c.Green("这段文字纯文本（所有样式效果都被禁用）")
	c.Red("错误信息也是纯文本")

	println()

	// 重新启用颜色
	c.SetNoColor(false)

	// ============================================================
	// 第五部分：重新启用颜色后，样式恢复
	// ============================================================
	printSection("5. 重新启用颜色（样式恢复）")

	c.Red("红色加粗下划线斜体（样式恢复）")
	c.Green("绿色加粗下划线斜体（样式恢复）")

	println()

	// 重置配置
	c.SetConfig(&color.StyleConfig{})

	// ============================================================
	// 第六部分：完整样式效果
	// ============================================================
	printSection("6. 完整样式效果（加粗 + 下划线 + 斜体 + 闪烁 + 暗淡 + 删除线）")

	c.SetBold(true).
		SetUnderline(true).
		SetItalic(true).
		SetBlink(true).
		SetFaint(true).
		SetCrossedOut(true)

	c.Red("红色 - 全样式")
	c.Yellow("黄色 - 全样式")

	println()

	// ============================================================
	// 第七部分：禁用颜色后，所有样式都被清除
	// ============================================================
	printSection("7. 禁用颜色后，所有样式都被清除")

	c.SetNoColor(true)

	c.Red("纯文本（所有样式被清除）")
	c.Yellow("纯文本（所有样式被清除）")

	// 重新启用
	c.SetNoColor(false)

	println()

	// ============================================================
	// 第八部分：输出到 stderr
	// ============================================================
	printSection("8. 输出到 stderr（默认加粗）")

	// 重置配置
	c.SetConfig(&color.StyleConfig{
		Output: os.Stderr,
	})

	c.Red("这条输出到 stderr（默认加粗红色）")
	c.Redf("错误信息也输出到 stderr\n")

	// 恢复输出到 stdout
	c.SetOutput(os.Stdout)

	println()

	// ============================================================
	// 第九部分：日志级别方法
	// ============================================================
	printSection("9. 日志级别方法（默认加粗）")

	// 重置为默认配置
	c.SetConfig(&color.StyleConfig{})

	c.Cyan("[INFO] 系统启动成功")
	c.Green("[OK] 数据库连接成功")
	c.Yellow("[WARN] 内存使用率超过 80%")
	c.Red("[ERROR] 无法连接到服务器")
	c.Magenta("[DEBUG] 正在处理请求 #12345")

	println()

	// ============================================================
	// 第十部分：高亮颜色
	// ============================================================
	printSection("10. 高亮颜色（默认加粗）")

	c.HiRed("高亮红色加粗")
	c.HiGreen("高亮绿色加粗")
	c.HiYellow("高亮黄色加粗")
	c.HiBlue("高亮蓝色加粗")

	println()

	// ============================================================
	// 第十一部分：实际应用场景
	// ============================================================
	printSection("11. 实际应用场景 - CLI 工具")

	// 模拟一个命令行工具的输出
	c.SetConfig(&color.StyleConfig{}) // 使用默认配置（加粗）

	c.Cyan("========================================")
	c.Cyan("  欢迎使用 MyApp CLI 工具 v1.0.0")
	c.Cyan("========================================")
	println()

	c.Cyan("[INFO] 正在检查环境...")
	c.Green("[OK] ✓ Go 版本: 1.25.0")
	c.Green("[OK] ✓ 配置文件: 已加载")

	// 添加下划线用于警告
	c.SetUnderline(true)
	c.Yellow("[WARN] ⚠ 配置文件使用了默认设置")

	// 恢复默认
	c.SetConfig(&color.StyleConfig{})
	c.Cyan("[INFO] 正在执行任务...")
	c.Magenta("[DEBUG] 处理文件: config.yaml")
	c.Green("[OK] 任务完成！处理了 2 个文件")

	println()

	c.Red("[ERROR] 发现 1 个错误:")
	c.Red("  - 无法读取 secret.key: 权限 denied")

	println()

	c.Yellow("提示: 使用 --help 查看所有可用命令")

	println()

	// ============================================================
	// 第十二部分：返回字符串的方法
	// ============================================================
	printSection("12. 返回字符串的方法（用于组合输出）")

	// 重置配置
	c.SetConfig(&color.StyleConfig{})

	// 组合多个带颜色的字符串
	redStr := c.SRed("红色")
	greenStr := c.SGreen("绿色")
	blueStr := c.SBlue("蓝色")

	// 使用普通的 fmt 打印组合后的字符串
	fmt.Printf("组合输出: %s + %s + %s\n", redStr, greenStr, blueStr)

	// 构建表格样式的输出
	fmt.Println("\n构建表格:")
	fmt.Printf("| %s | %s | %s |\n", c.SRed("错误"), c.SYellow("警告"), c.SGreen("成功"))
	fmt.Printf("| %s | %s | %s |\n", c.SHiRed("高亮错误"), c.SHiYellow("高亮警告"), c.SHiGreen("高亮成功"))

	// 用于日志前缀
	fmt.Println("\n日志前缀:")
	infoPrefix := c.SCyan("[INFO]")
	warnPrefix := c.SYellow("[WARN]")
	errPrefix := c.SRed("[ERROR]")
	fmt.Printf("%s 系统启动成功\n", infoPrefix)
	fmt.Printf("%s 内存使用率超过 80%%\n", warnPrefix)
	fmt.Printf("%s 连接失败\n", errPrefix)

	// 与标准库结合使用
	fmt.Println("\n与标准库结合:")
	result := fmt.Sprintf("状态: %s, 进度: %s", c.SGreen("正常"), c.SBlue("75%%"))
	fmt.Println(result)

	println()

	// ============================================================
	// 第十三部分：终端提示信息前缀方法
	// ============================================================
	printSection("13. 终端提示信息前缀方法")

	// 重置配置
	c.SetConfig(&color.StyleConfig{})

	// 使用颜色模拟 ok/warn/err 前缀效果
	c.Green("[OK] 服务启动成功")
	c.Green("[OK] 配置文件加载完成")
	c.Yellow("[WARN] 内存使用率超过 80%")
	c.Red("[ERR] 数据库连接失败")

	println()

	// 实际应用场景：命令执行流程
	printSection("14. 实际应用场景 - 命令执行流程")

	c.Green("[OK] 初始化环境")
	c.Green("[OK] 加载配置文件")
	c.Green("[OK] 连接数据库")
	c.Yellow("[WARN] 发现 3 个可更新包")
	c.Red("[ERR] 无法连接到远程服务器")
	c.Green("[OK] 使用本地缓存")

	println()

	// 测试场景
	printSection("15. 测试场景")

	c.Green("[OK] TestAdd: 通过 (5ms)")
	c.Green("[OK] TestSub: 通过 (3ms)")
	c.Yellow("[WARN] TestComplex: 跳过 (需要外部服务)")
	c.Red("[ERR] TestDivide: 失败 - 除零错误")
	c.Yellow("[WARN] 覆盖率: 87.5%")
	c.Green("[OK] 测试执行完成")

	println()

	// ============================================================
	// 结束
	// ============================================================
	printSection("示例结束")
	c.HiGreen("感谢使用 color 包的 GlobalColor 功能！")
}

// printSection 打印章节标题
func printSection(title string) {
	println()
	c := color.GetGlobal()
	c.SetConfig(&color.StyleConfig{}) // 使用默认配置
	c.Bluef("【%s】\n", title)
}
