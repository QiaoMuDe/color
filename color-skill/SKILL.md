---
name: go-color-assistant
description: |
  为 Go 语言的 color 库（gitee.com/MM-Q/color）生成彩色输出代码。
  
  当用户需要以下帮助时触发此 skill：
  - 给终端输出添加颜色（如"给这段文字加上颜色"、"让输出更醒目"）
  - 使用 color 库美化输出（如"用 color 库高亮显示错误信息"）
  - 生成彩色日志、表格、进度条等终端 UI 组件
  - 询问 color 库的 API 用法（如"怎么打印红色文字"、"如何组合样式"）
  - 调试颜色相关问题（如"颜色为什么不显示"、"怎么禁用颜色"）
  
  即使只是提到"颜色"、"color"、"美化输出"、"高亮"等关键词，也应该使用此 skill 提供专业的 color 库解决方案。
  
  此 skill 专门处理 gitee.com/MM-Q/color 库，支持：
  - 便捷函数（Red/Green/Blue 等）
  - 链式调用（New().Add().Print()）
  - RGB 真彩色
  - 全局实例（G()）
  - 字符串返回方法（SRed/SGreen 等）
---

# Go Color 库助手

## 核心能力

帮助用户生成使用 `gitee.com/MM-Q/color` 库的 Go 代码，实现终端彩色输出。

## 使用方式

### 1. 分析用户需求

当用户提出颜色相关需求时：
- 理解用户想要的效果（颜色、样式、输出方式）
- 判断使用哪种 API 最合适
- 提供完整的代码示例

### 2. 选择合适的 API

根据场景推荐最佳方案：

| 场景 | 推荐 API | 示例 |
|------|----------|------|
| 简单打印一行彩色文字 | 便捷函数 | `color.Red("错误信息")` |
| 组合多种样式 | 链式调用 | `color.New(color.FgRed, color.Bold).Println(...)` |
| 需要返回字符串 | S 前缀函数 | `str := color.SRed("红色文字")` |
| 频繁使用相同样式 | 全局实例 | `c := color.G(); c.Red(...)` |
| 自定义 RGB 颜色 | RGB 函数 | `color.RGB(255, 128, 0).Println(...)` |

### 3. 生成代码

提供完整、可直接运行的代码示例：
- 包含 import 语句
- 展示多种实现方式供选择
- 添加注释说明关键点

### 4. 解释 API 差异

特别强调与原库（github.com/fatih/color）的差异：
- 字符串返回方法：`RedString()` → `SRed()`
- 新增全局实例：`color.G()`

## 代码模板

### 模板 1：基础用法

```go
package main

import "gitee.com/MM-Q/color"

func main() {
    // 便捷函数 - 最简单的方式
    color.Red("这是红色文字")
    color.Green("这是绿色文字")
    color.Blue("这是蓝色文字")
    
    // 格式化输出
    color.Yellow("警告: %s", "磁盘空间不足")
}
```

### 模板 2：链式调用

```go
package main

import "gitee.com/MM-Q/color"

func main() {
    // 组合多种样式
    color.New(color.FgRed, color.Bold, color.Underline).
        Println("红色粗体下划线")
    
    // 逐步添加属性
    c := color.New(color.FgBlue)
    c.Add(color.Bold)
    c.Println("蓝色粗体")
}
```

### 模板 3：全局实例

```go
package main

import "gitee.com/MM-Q/color"

func main() {
    // 获取全局实例（默认白色加粗）
    c := color.G()
    
    // 使用快捷方法
    c.Red("红色文字")
    c.Green("绿色文字")
    
    // 使用日志级别方法
    c.Info("信息日志")
    c.Success("成功日志")
    c.Warn("警告日志")
    c.Error("错误日志")
}
```

### 模板 4：RGB 颜色

```go
package main

import "gitee.com/MM-Q/color"

func main() {
    // RGB 前景色
    color.RGB(255, 128, 0).Println("橙色文字")
    
    // RGB 背景色
    color.BgRGB(0, 128, 255).Println("蓝色背景")
    
    // 组合前景和背景
    color.RGB(255, 255, 255).
        AddBgRGB(255, 0, 0).
        Println("白字红底")
}
```

### 模板 5：终端 UI 组件

```go
package main

import (
    "fmt"
    "gitee.com/MM-Q/color"
)

// 打印章节标题
func printSection(title string) {
    fmt.Println()
    color.New(color.FgHiWhite, color.BgBlue, color.Bold).
        Printf(" %s ", title)
    fmt.Println()
    fmt.Println()
}

// 打印日志
func printLog(level, message string) {
    switch level {
    case "INFO":
        color.Cyan("[INFO]  %s", message)
    case "WARN":
        color.Yellow("[WARN]  %s", message)
    case "ERROR":
        color.Red("[ERROR] %s", message)
    case "SUCCESS":
        color.Green("[OK]    %s", message)
    }
}

func main() {
    printSection("系统状态")
    printLog("INFO", "系统启动成功")
    printLog("SUCCESS", "数据库连接已建立")
    printLog("WARN", "内存使用率超过 80%")
    printLog("ERROR", "连接超时")
}
```

## 常见问题

### Q: 颜色为什么不显示？

A: 检查以下几点：
1. 是否设置了 `NO_COLOR` 环境变量
2. 输出是否被重定向到文件或管道
3. Windows 系统是否支持 ANSI（Windows 10 以上默认支持）
4. 是否手动禁用了颜色：`color.NoColor = true`

### Q: 如何禁用所有颜色？

A: 三种方式：
```go
// 方式 1：代码中设置
color.NoColor = true

// 方式 2：环境变量
export NO_COLOR=1

// 方式 3：针对单个对象
c := color.New(color.FgRed)
c.DisableColor()
```

### Q: 如何返回带颜色的字符串（不打印）？

A: 使用 S 前缀的函数：
```go
redStr := color.SRed("红色文字")
fmt.Printf("前缀 %s 后缀\n", redStr)
```

### Q: 从 github.com/fatih/color 迁移需要注意什么？

A: 主要变化：
- `RedString()` → `SRed()`
- `GreenString()` → `SGreen()`
- 其他类似
- 新增全局实例 `color.G()`

## 最佳实践

1. **简单场景用便捷函数**：`color.Red()`, `color.Green()` 等
2. **复杂样式用链式调用**：`color.New().Add().Print()`
3. **需要字符串用 S 前缀**：`color.SRed()`, `color.SGreen()`
4. **频繁使用用全局实例**：`c := color.G()`
5. **自定义颜色用 RGB**：`color.RGB(r, g, b)`

## 可用颜色常量

### 标准前景色
`FgBlack`, `FgRed`, `FgGreen`, `FgYellow`, `FgBlue`, `FgMagenta`, `FgCyan`, `FgWhite`

### 高亮前景色
`FgHiBlack`, `FgHiRed`, `FgHiGreen`, `FgHiYellow`, `FgHiBlue`, `FgHiMagenta`, `FgHiCyan`, `FgHiWhite`

### 标准背景色
`BgBlack`, `BgRed`, `BgGreen`, `BgYellow`, `BgBlue`, `BgMagenta`, `BgCyan`, `BgWhite`

### 文本样式
`Bold`, `Italic`, `Underline`, `BlinkSlow`, `ReverseVideo`, `CrossedOut`
