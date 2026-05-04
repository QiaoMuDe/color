// main.go - 终端 UI 辅助函数演示
// 本文件演示如何使用 utils.go 中定义的辅助函数
// 运行方式: go run .

package main

import (
	"fmt"
	"time"

	"gitee.com/MM-Q/color"
)

func main() {
	// ============================================================
	// 演示 1: 章节标题
	// ============================================================
	PrintSection("1. 章节标题示例")

	fmt.Println("这是普通文本内容...")
	fmt.Println("章节标题可以帮助组织输出内容，使其更易读")

	PrintSubSection("1.1 子章节标题")
	fmt.Println("子章节标题使用青色和加粗样式")

	PrintSubSection("1.2 另一个子章节")
	fmt.Println("通过层级标题，可以清晰地展示内容结构")

	// ============================================================
	// 演示 2: 分隔线
	// ============================================================
	PrintSection("2. 分隔线示例")

	fmt.Println("上面的内容...")
	PrintSeparator(color.FgCyan, 50)
	fmt.Println("分隔线可以区分不同部分")
	PrintDoubleSeparator(color.FgBlue, 50)
	fmt.Println("双分隔线用于更重要的分隔")

	// ============================================================
	// 演示 3: 日志级别打印
	// ============================================================
	PrintSection("3. 日志级别打印示例")

	LogInfo("系统启动成功，版本: v1.0.0")
	LogSuccess("数据库连接已建立")
	LogWarn("内存使用率超过 80%，当前: 85%")
	LogError("无法连接到远程服务器: connection timeout")
	LogDebug("正在处理请求 #12345，参数: {...}")

	fmt.Println()
	fmt.Println("这些日志级别函数可以统一项目中的日志输出样式")

	// ============================================================
	// 演示 4: 表格输出
	// ============================================================
	PrintSection("4. 表格输出示例")

	// 创建服务状态表格
	serviceTable := NewTable(
		[]string{"服务名称", "状态", "运行时间", "版本"},
		[]int{12, 10, 15, 10},
	)

	serviceTable.AddRow([]string{"API Gateway", "运行中", "3天 2小时", "v2.1.0"})
	serviceTable.AddRow([]string{"Auth Service", "运行中", "3天 2小时", "v1.5.2"})
	serviceTable.AddRow([]string{"Database", "运行中", "15天 8小时", "v14.2"})
	serviceTable.AddRow([]string{"Cache", "警告", "2天 12小时", "v6.0.1"})
	serviceTable.AddRow([]string{"Worker", "已停止", "-", "v1.0.0"})

	serviceTable.Print()

	fmt.Println()
	fmt.Println("表格可以清晰地展示结构化数据")

	// 创建用户信息表格
	PrintSubSection("用户信息表")

	userTable := NewTable(
		[]string{"ID", "用户名", "角色", "状态"},
		[]int{8, 15, 12, 10},
	)

	userTable.AddRow([]string{"001", "admin", "管理员", "活跃"})
	userTable.AddRow([]string{"002", "zhangsan", "开发者", "活跃"})
	userTable.AddRow([]string{"003", "lisi", "测试员", "离线"})

	userTable.Print()

	// ============================================================
	// 演示 5: 列表输出
	// ============================================================
	PrintSection("5. 列表输出示例")

	PrintSubSection("无序列表")
	items := []string{
		"安装 Go 1.25 或更高版本",
		"运行 go mod tidy 下载依赖",
		"执行 go run main.go 启动应用",
		"访问 http://localhost:8080 查看结果",
	}
	PrintBulletList(items, color.FgCyan)

	fmt.Println()
	PrintSubSection("有序列表")
	tasks := []string{
		"需求分析和设计",
		"数据库建模",
		"API 接口开发",
		"前端页面实现",
		"测试和部署",
	}
	PrintNumberedList(tasks, color.FgGreen)

	// ============================================================
	// 演示 6: 状态标签
	// ============================================================
	PrintSection("6. 状态标签示例")

	fmt.Print("服务A: ")
	PrintStatusRunning()
	fmt.Println()

	fmt.Print("服务B: ")
	PrintStatusStopped()
	fmt.Println()

	fmt.Print("服务C: ")
	PrintStatusWarning()
	fmt.Println()

	fmt.Print("任务队列: ")
	PrintStatusPending()
	fmt.Println()

	fmt.Println()
	fmt.Println("状态标签可以快速识别系统组件的状态")

	// ============================================================
	// 演示 7: 进度条
	// ============================================================
	PrintSection("7. 进度条示例")

	fmt.Println("文件下载进度:")
	for i := 0; i <= 10; i++ {
		fmt.Print("\r") // 回到行首
		PrintProgressBar(i, 10, 30)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("数据处理进度:")
	PrintProgressBar(75, 100, 40)
	fmt.Println(" 75/100 记录")

	fmt.Println()
	fmt.Println("进度条可以直观地展示任务完成情况")

	// ============================================================
	// 演示 8: 代码块
	// ============================================================
	PrintSection("8. 代码块示例")

	code := `func main() {
    // 创建颜色对象
    c := color.New(color.FgRed)
    
    // 打印红色文本
    c.Println("Hello, World!")
}`

	PrintCodeBlock(code, "go")

	sqlCode := `SELECT id, name, email
FROM users
WHERE status = 'active'
ORDER BY created_at DESC;`

	PrintCodeBlock(sqlCode, "sql")

	// ============================================================
	// 演示 9: 引用文本
	// ============================================================
	PrintSection("9. 引用文本示例")

	PrintQuote(
		"Talk is cheap. Show me the code.",
		"Linus Torvalds",
	)

	PrintQuote(
		"程序必须首先为人类阅读而写，其次才是为机器执行而写。",
		"Harold Abelson",
	)

	// ============================================================
	// 演示 10: 提示框
	// ============================================================
	PrintSection("10. 提示框示例")

	PrintTip("使用提示", "你可以通过设置 NO_COLOR=1 环境变量来禁用所有颜色输出")

	PrintWarning("配置文件未找到，将使用默认配置")

	PrintError("连接数据库失败，请检查网络连接")

	PrintSuccess("部署成功！应用已在 http://localhost:8080 运行")

	// ============================================================
	// 演示 11: 综合示例 - 系统监控面板
	// ============================================================
	PrintSection("11. 综合示例: 系统监控面板")

	printSystemDashboard()

	// ============================================================
	// 结束
	// ============================================================
	PrintSection("演示结束")
	PrintSuccess("所有演示已完成！")
	fmt.Println()
	fmt.Println("提示: 这些辅助函数都在 utils.go 中定义，")
	fmt.Println("      你可以根据项目需求进行修改和扩展。")
}

// printSystemDashboard 打印系统监控面板
func printSystemDashboard() {
	// 标题
	PrintDoubleSeparator(color.FgHiBlue, 60)
	_, _ = color.New(color.FgHiWhite, color.Bold).Println("                    系统监控面板")
	PrintDoubleSeparator(color.FgHiBlue, 60)
	fmt.Println()

	// 系统状态概览
	PrintSubSection("系统状态")
	fmt.Print("整体状态: ")
	PrintStatusRunning()
	fmt.Println()

	fmt.Print("系统负载: ")
	PrintProgressBar(45, 100, 30)
	fmt.Println()

	fmt.Println()

	// 服务列表
	PrintSubSection("服务状态")
	serviceTable := NewTable(
		[]string{"服务", "状态", "CPU", "内存"},
		[]int{15, 10, 10, 10},
	)
	serviceTable.AddRow([]string{"Web Server", "运行中", "12%", "256MB"})
	serviceTable.AddRow([]string{"Database", "运行中", "35%", "1.2GB"})
	serviceTable.AddRow([]string{"Cache", "运行中", "8%", "128MB"})
	serviceTable.AddRow([]string{"Queue Worker", "警告", "78%", "512MB"})
	serviceTable.Print()

	fmt.Println()

	// 最近日志
	PrintSubSection("最近日志")
	LogInfo("系统启动完成，耗时 2.3s")
	LogSuccess("用户 admin 登录成功")
	LogWarn("数据库连接池使用率超过 80%")
	LogDebug("处理 HTTP 请求: GET /api/users")

	fmt.Println()
	PrintSeparator(color.FgHiBlack, 60)
}
