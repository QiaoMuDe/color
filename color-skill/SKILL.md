---
name: go-color
description: 为 Go 语言的 color 库（gitee.com/MM-Q/color）生成彩色输出代码。当用户需要给终端输出添加颜色、使用 color 库美化输出、生成彩色日志/表格/进度条等终端 UI 组件、询问 color 库的 API 用法或调试颜色相关问题时触发此 skill。支持便捷函数、链式调用、RGB 真彩色、全局实例、字符串返回方法等功能。
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
| 格式化打印 | 带 f 后缀函数 | `color.Redf("错误: %s", err)` |
| 组合多种样式 | 链式调用 | `color.New(color.FgRed, color.Bold).Println(...)` |
| 需要返回字符串 | S 前缀函数 | `str := color.SRed("红色文字")` |
| 格式化返回字符串 | 带 f 后缀 S 函数 | `str := color.SRedf("错误: %s", err)` |
| 频繁使用相同样式 | 全局实例 | `c := color.G(); c.Red(...)` |
| 自定义 RGB 颜色 | RGB 函数 | `color.RGB(255, 128, 0).Println(...)` |

**重要区别**：
- 不带 `f` 后缀的方法（如 `Red()`）：只接受字符串，自动追加换行符
- 带 `f` 后缀的方法（如 `Redf()`）：支持格式化字符串，不自动追加换行符

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
    // 便捷函数 - 最简单的方式（自动追加换行符）
    color.Red("这是红色文字")
    color.Green("这是绿色文字")
    color.Blue("这是蓝色文字")
    
    // 格式化输出（不自动追加换行符，需手动添加）
    color.Yellowf("警告: %s\n", "磁盘空间不足")
    
    // 返回字符串（不打印）
    redStr := color.SRed("红色文字")
    fmt.Println(redStr)
    
    // 格式化返回字符串
    errStr := color.SRedf("错误: %s", "连接失败")
    fmt.Println(errStr)
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
    // 获取全局实例（默认启用加粗）
    c := color.G()
    
    // 使用快捷方法（自动追加换行符）
    c.Red("红色文字")
    c.Green("绿色文字")
    
    // 使用格式化方法（不自动追加换行符）
    c.Bluef("蓝色: %s\n", "信息")
    
    // 返回字符串
    str := c.SRed("红色字符串")
    fmt.Println(str)
    
    // 链式配置
    c.SetBold(true).SetUnderline(true).Red("粗体下划线红色")
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
        color.Cyanf("[INFO]  %s\n", message)
    case "WARN":
        color.Yellowf("[WARN]  %s\n", message)
    case "ERROR":
        color.Redf("[ERROR] %s\n", message)
    case "SUCCESS":
        color.Greenf("[OK]    %s\n", message)
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
- **重要**：API 使用方式变化
  - 原库：`Red("text %s", arg)` 支持格式化
  - 本库：`Red("text")` 只接受字符串，`Redf("text %s", arg)` 支持格式化

## 最佳实践

1. **简单场景用便捷函数**：`color.Red()`, `color.Green()` 等（自动追加换行符）
2. **格式化输出用 f 后缀**：`color.Redf()`, `color.Greenf()` 等（不自动追加换行符）
3. **复杂样式用链式调用**：`color.New().Add().Print()`
4. **需要字符串用 S 前缀**：`color.SRed()`, `color.SGreen()`（不带f后缀）或 `color.SRedf()`, `color.SGreenf()`（带f后缀）
5. **频繁使用用全局实例**：`c := color.G()`
6. **自定义颜色用 RGB**：`color.RGB(r, g, b)`
7. **灰色文字**：`color.Gray()` 或 `color.SGray()`（`FgHiBlack` 的别名）

## 可用颜色常量

### 标准前景色
`FgBlack`, `FgRed`, `FgGreen`, `FgYellow`, `FgBlue`, `FgMagenta`, `FgCyan`, `FgWhite`

### 高亮前景色
`FgHiBlack`, `FgHiRed`, `FgHiGreen`, `FgHiYellow`, `FgHiBlue`, `FgHiMagenta`, `FgHiCyan`, `FgHiWhite`

### 灰色（别名）
`FgGray`（`FgHiBlack` 的别名）, `BgGray`（`BgHiBlack` 的别名）

### 标准背景色
`BgBlack`, `BgRed`, `BgGreen`, `BgYellow`, `BgBlue`, `BgMagenta`, `BgCyan`, `BgWhite`

### 文本样式
`Bold`, `Italic`, `Underline`, `BlinkSlow`, `ReverseVideo`, `CrossedOut`
