// rgb 示例 - 展示真彩 RGB 颜色方法的使用
//
// 本示例演示 color 包的 RGB 真彩色功能：
// - RGB 前景色：RGB(r, g, b) 和 SRgb(r, g, b)
// - RGB 背景色：BgRGB(r, g, b) 和 SBgbRgb(r, g, b)
// - 链式调用：AddRGB() 和 AddBgRGB()
// - 渐变色效果和彩虹色输出
//
// 注意：RGB 真彩色需要终端支持才能正常显示
package main

import (
	"fmt"
	"strings"

	"gitee.com/MM-Q/color"
)

func main() {
	println("═══════════════════════════════════════════════════════════")
	println("  RGB 真彩色示例")
	println("═══════════════════════════════════════════════════════════")
	println()
	println("注意：RGB 真彩色需要终端支持才能正常显示")
	println("      如果看到白色文字，说明您的终端不支持 RGB 颜色")
	println()

	// ============================================================
	// 第一部分：RGB 前景色基础
	// ============================================================
	printSection("1. RGB 前景色基础")

	// 纯红色
	_, _ = color.RGB(255, 0, 0).Println("纯红色 (255, 0, 0)")

	// 纯绿色
	_, _ = color.RGB(0, 255, 0).Println("纯绿色 (0, 255, 0)")

	// 纯蓝色
	_, _ = color.RGB(0, 0, 255).Println("纯蓝色 (0, 0, 255)")

	// 白色
	_, _ = color.RGB(255, 255, 255).Println("白色 (255, 255, 255)")

	// 黑色（在白色背景上显示）
	_, _ = color.New().AddBgRGB(255, 255, 255).AddRGB(0, 0, 0).Println("黑色 (0, 0, 0)")

	println()

	// ============================================================
	// 第二部分：RGB 前景色 - 常用颜色
	// ============================================================
	printSection("2. RGB 前景色 - 常用颜色")

	// 橙色
	_, _ = color.RGB(255, 165, 0).Println("橙色 (255, 165, 0)")

	// 粉色
	_, _ = color.RGB(255, 192, 203).Println("粉色 (255, 192, 203)")

	// 紫色
	_, _ = color.RGB(128, 0, 128).Println("紫色 (128, 0, 128)")

	// 青色
	_, _ = color.RGB(0, 255, 255).Println("青色 (0, 255, 255)")

	// 黄色
	_, _ = color.RGB(255, 255, 0).Println("黄色 (255, 255, 0)")

	// 棕色
	_, _ = color.RGB(165, 42, 42).Println("棕色 (165, 42, 42)")

	// 灰色
	_, _ = color.RGB(128, 128, 128).Println("灰色 (128, 128, 128)")

	println()

	// ============================================================
	// 第三部分：RGB 背景色
	// ============================================================
	printSection("3. RGB 背景色")

	// 红色背景
	_, _ = color.BgRGB(255, 0, 0).Println("红色背景")

	// 绿色背景
	_, _ = color.BgRGB(0, 255, 0).Println("绿色背景")

	// 蓝色背景
	_, _ = color.BgRGB(0, 0, 255).Println("蓝色背景")

	// 黄色背景
	_, _ = color.BgRGB(255, 255, 0).Println("黄色背景")

	// 紫色背景
	_, _ = color.BgRGB(128, 0, 128).Println("紫色背景")

	println()

	// ============================================================
	// 第四部分：前景色 + 背景色组合
	// ============================================================
	printSection("4. 前景色 + 背景色组合")

	// 白字红底
	_, _ = color.New().
		AddRGB(255, 255, 255).
		AddBgRGB(255, 0, 0).
		Println(" 白字红底 ")

	// 黑字黄底
	_, _ = color.New().
		AddRGB(0, 0, 0).
		AddBgRGB(255, 255, 0).
		Println(" 黑字黄底 ")

	// 黄字蓝底
	_, _ = color.New().
		AddRGB(255, 255, 0).
		AddBgRGB(0, 0, 139).
		Println(" 黄字蓝底 ")

	// 青字紫底
	_, _ = color.New().
		AddRGB(0, 255, 255).
		AddBgRGB(75, 0, 130).
		Println(" 青字紫底 ")

	// 粉字黑底
	_, _ = color.New().
		AddRGB(255, 105, 180).
		AddBgRGB(0, 0, 0).
		Println(" 粉字黑底 ")

	println()

	// ============================================================
	// 第五部分：RGB + 样式属性组合
	// ============================================================
	printSection("5. RGB + 样式属性组合")

	// 加粗红色
	_, _ = color.New().
		AddRGB(255, 0, 0).
		Add(color.Bold).
		Println("加粗红色")

	// 下划线绿色
	_, _ = color.New().
		AddRGB(0, 255, 0).
		Add(color.Underline).
		Println("下划线绿色")

	// 斜体蓝色
	_, _ = color.New().
		AddRGB(0, 100, 255).
		Add(color.Italic).
		Println("斜体蓝色")

	// 加粗 + 下划线 + 橙色背景
	_, _ = color.New().
		AddRGB(255, 255, 255).
		AddBgRGB(255, 140, 0).
		Add(color.Bold).
		Add(color.Underline).
		Println(" 加粗下划线白字橙底 ")

	println()

	// ============================================================
	// 第六部分：渐变色效果
	// ============================================================
	printSection("6. 渐变色效果")

	// 红色到黄色的渐变
	println("红色到黄色渐变:")
	for i := 0; i <= 10; i++ {
		r := 255
		g := int(255 * i / 10)
		b := 0
		_, _ = color.RGB(r, g, b).Printf("█")
	}
	println()
	println()

	// 蓝色到青色的渐变
	println("蓝色到青色渐变:")
	for i := 0; i <= 10; i++ {
		r := 0
		g := int(255 * i / 10)
		b := 255
		_, _ = color.RGB(r, g, b).Printf("█")
	}
	println()
	println()

	// 紫色到粉色的渐变
	println("紫色到粉色渐变:")
	for i := 0; i <= 10; i++ {
		r := 128 + int(127*i/10)
		g := int(192 * i / 10)
		b := 128 + int(75*i/10)
		_, _ = color.RGB(r, g, b).Printf("█")
	}
	println()
	println()

	// 灰度渐变
	println("灰度渐变:")
	for i := 0; i <= 10; i++ {
		v := 255 * i / 10
		_, _ = color.RGB(v, v, v).Printf("█")
	}
	println()

	println()

	// ============================================================
	// 第七部分：彩虹色
	// ============================================================
	printSection("7. 彩虹色")

	// 彩虹条
	println("彩虹条:")
	colors := []struct {
		r, g, b int
	}{
		{255, 0, 0},   // 红
		{255, 127, 0}, // 橙
		{255, 255, 0}, // 黄
		{0, 255, 0},   // 绿
		{0, 0, 255},   // 蓝
		{75, 0, 130},  // 靛
		{148, 0, 211}, // 紫
	}

	for _, c := range colors {
		_, _ = color.RGB(c.r, c.g, c.b).Printf("██████")
	}
	println()
	println()

	// 彩虹文字
	println("彩虹文字:")
	rainbowText := "Hello, RGB World!"
	for i, char := range rainbowText {
		c := colors[i%len(colors)]
		_, _ = color.RGB(c.r, c.g, c.b).Printf("%c", char)
	}
	println()

	println()

	// ============================================================
	// 第八部分：返回字符串的 RGB 方法
	// ============================================================
	printSection("8. 返回字符串的 RGB 方法")

	// Sprint - 返回带 RGB 颜色的字符串
	redStr := color.RGB(255, 0, 0).Sprint("红色字符串")
	greenStr := color.RGB(0, 255, 0).Sprint("绿色字符串")
	blueStr := color.RGB(0, 0, 255).Sprint("蓝色字符串")
	fmt.Printf("组合输出: %s %s %s\n", redStr, greenStr, blueStr)

	// 构建带颜色的标签
	tag1 := color.RGB(255, 165, 0).Sprint("[警告]")
	tag2 := color.RGB(50, 205, 50).Sprint("[成功]")
	tag3 := color.RGB(255, 69, 0).Sprint("[错误]")
	fmt.Printf("标签组合: %s 系统负载高 | %s 操作完成 | %s 连接失败\n", tag1, tag2, tag3)

	println()

	// ============================================================
	// 第九部分：实际应用场景
	// ============================================================
	printSection("9. 实际应用场景")

	// 场景1：温度显示
	printTemperature(25.5)
	printTemperature(35.8)
	printTemperature(15.2)
	printTemperature(-5.0)

	println()

	// 场景2：CPU 使用率
	printCPUUsage(15)
	printCPUUsage(45)
	printCPUUsage(78)
	printCPUUsage(95)

	println()

	// 场景3：彩色进度条
	printRGBProgressBar(65, 100, 30)

	println()

	// 场景4：彩色表格
	printRGBTable()

	println()

	// ============================================================
	// 第十部分：RGB 颜色选择器展示
	// ============================================================
	printSection("10. RGB 颜色选择器展示")

	// 展示不同 RGB 值的红色
	println("红色系:")
	for i := 0; i <= 255; i += 51 {
		_, _ = color.RGB(i, 0, 0).Printf(" %3d ", i)
	}
	println()

	// 展示不同 RGB 值的绿色
	println("绿色系:")
	for i := 0; i <= 255; i += 51 {
		_, _ = color.RGB(0, i, 0).Printf(" %3d ", i)
	}
	println()

	// 展示不同 RGB 值的蓝色
	println("蓝色系:")
	for i := 0; i <= 255; i += 51 {
		_, _ = color.RGB(0, 0, i).Printf(" %3d ", i)
	}
	println()

	println()

	// ============================================================
	// 结束
	// ============================================================
	printSection("示例结束")
	_, _ = color.RGB(50, 205, 50).Println("感谢使用 RGB 真彩色功能！")
	_, _ = color.RGB(100, 149, 237).Println("文档地址: https://gitee.com/MM-Q/color")
}

// printSection 打印章节标题
func printSection(title string) {
	println()
	c := color.New(color.FgHiWhite, color.BgBlue, color.Bold)
	_, _ = c.Printf(" %s ", title)
	println()
	println()
}

// printTemperature 打印带颜色的温度
func printTemperature(temp float64) {
	var r, g int
	if temp > 30 {
		// 高温：红色到橙色
		r = 255
		g = int(165 * (40 - temp) / 10)
		if g < 0 {
			g = 0
		}
	} else if temp > 15 {
		// 中温：橙色到黄色
		r = 255
		g = 165 + int(90*(temp-15)/15)
	} else {
		// 低温：青色到蓝色
		r = 0
		g = int(255 * (temp + 10) / 25)
		if g > 255 {
			g = 255
		}
	}

	tempStr := color.RGB(r, g, 0).Sprint(fmt.Sprintf("%.1f°C", temp))
	fmt.Printf("当前温度: %s\n", tempStr)
}

// printCPUUsage 打印带颜色的 CPU 使用率
func printCPUUsage(usage int) {
	var r, g int
	if usage < 50 {
		// 低使用率：绿色
		r = int(255 * usage / 50)
		g = 255
	} else {
		// 高使用率：黄色到红色
		r = 255
		g = int(255 * (100 - usage) / 50)
	}

	// 构建进度条
	filled := usage / 5
	bar := strings.Repeat("█", filled) + strings.Repeat("░", 20-filled)

	usageStr := color.RGB(r, g, 0).Sprint(fmt.Sprintf("%3d%%", usage))
	barStr := color.RGB(r, g, 0).Sprint(bar)
	fmt.Printf("CPU 使用率: %s %s\n", usageStr, barStr)
}

// printRGBProgressBar 打印 RGB 彩色进度条
func printRGBProgressBar(current, total, width int) {
	if total <= 0 {
		return
	}

	percent := float64(current) / float64(total)
	filled := int(float64(width) * percent)

	fmt.Print("进度: ")

	// 打印彩色进度条（从绿色渐变到红色）
	for i := 0; i < width; i++ {
		if i < filled {
			// 根据位置计算颜色（绿 -> 黄 -> 红）
			progress := float64(i) / float64(width)
			var r, g int
			if progress < 0.5 {
				// 绿到黄
				r = int(255 * progress * 2)
				g = 255
			} else {
				// 黄到红
				r = 255
				g = int(255 * (1 - progress) * 2)
			}
			_, _ = color.RGB(r, g, 0).Printf("█")
		} else {
			_, _ = color.RGB(100, 100, 100).Printf("░")
		}
	}

	// 打印百分比
	_, _ = color.RGB(0, 255, 255).Printf(" %.1f%%", percent*100)
	println()
}

// printRGBTable 打印 RGB 彩色表格
func printRGBTable() {
	// 表头
	headerBg := color.New().AddBgRGB(50, 50, 50)
	_, _ = headerBg.AddRGB(255, 255, 255).Printf(" %-12s ", "服务名称")
	_, _ = headerBg.AddRGB(255, 255, 255).Printf(" %-10s ", "状态")
	_, _ = headerBg.AddRGB(255, 255, 255).Printf(" %-12s ", "响应时间")
	_, _ = headerBg.AddRGB(255, 255, 255).Printf(" %-8s ", "负载")
	println()

	// 数据行
	printRGBTableRow("Web 服务器", "运行中", 45, 35)
	printRGBTableRow("数据库", "运行中", 12, 65)
	printRGBTableRow("缓存服务", "运行中", 8, 25)
	printRGBTableRow("消息队列", "警告", 120, 85)
}

// printRGBTableRow 打印 RGB 表格行
func printRGBTableRow(name, status string, responseTime, load int) {
	// 名称（白色）
	nameStr := color.RGB(200, 200, 200).Sprint(fmt.Sprintf(" %-12s ", name))
	fmt.Print(nameStr)
	fmt.Print(" ")

	// 状态（根据状态着色）
	var statusStr string
	switch status {
	case "运行中":
		statusStr = color.RGB(50, 205, 50).Sprint(fmt.Sprintf(" %-10s ", status))
	case "警告":
		statusStr = color.RGB(255, 165, 0).Sprint(fmt.Sprintf(" %-10s ", status))
	case "错误":
		statusStr = color.RGB(255, 69, 0).Sprint(fmt.Sprintf(" %-10s ", status))
	default:
		statusStr = fmt.Sprintf(" %-10s ", status)
	}
	fmt.Print(statusStr)
	fmt.Print(" ")

	// 响应时间（根据时间着色）
	var rtR, rtG int
	if responseTime < 50 {
		rtR, rtG = 50, 205
	} else if responseTime < 100 {
		rtR, rtG = 255, 165
	} else {
		rtR, rtG = 255, 69
	}
	rtStr := color.RGB(rtR, rtG, 0).Sprint(fmt.Sprintf(" %-12s ", fmt.Sprintf("%dms", responseTime)))
	fmt.Print(rtStr)
	fmt.Print(" ")

	// 负载（渐变色）
	var loadR, loadG int
	if load < 50 {
		loadR = int(255 * load / 50)
		loadG = 255
	} else {
		loadR = 255
		loadG = int(255 * (100 - load) / 50)
	}
	loadStr := color.RGB(loadR, loadG, 0).Sprint(fmt.Sprintf(" %-8s ", fmt.Sprintf("%d%%", load)))
	fmt.Println(loadStr)
}
