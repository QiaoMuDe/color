# color 包文档

```go
package color // import "gitee.com/MM-Q/color"
```

Package color 是一个 ANSI 颜色包，用于向标准输出输出彩色或 SGR 定义的文本。 API 可以通过多种方式使用，选择适合你的方式即可。

## 快速开始

使用简单且默认的辅助函数，配合预定义的前景色：

```go
color.Cyan("以青色打印文本。")

// 默认会自动追加换行符
color.Blue("以蓝色打印 %s。", "文本")

// 更多默认前景色..
color.Red("我们有红色")
color.Yellow("也有黄色！")
color.Magenta("还有很多其他颜色 ..")

// 高亮色
color.HiGreen("亮绿色。")
color.HiBlack("亮黑色就是灰色..")
color.HiWhite("闪亮的白色！")
```

## 自定义颜色组合

然而，有时需要自定义颜色组合。以下是一些创建自定义颜色对象 并使用每个独立颜色对象的打印函数的示例。

```go
// 创建一个新的颜色对象
c := color.New(color.FgCyan).Add(color.Underline)
c.Println("打印带下划线的青色文本。")

// 或者直接添加到 New() 中
d := color.New(color.FgCyan, color.Bold)
d.Printf("这也打印粗体青色 %s\n", "！")

// 混合前景色和背景色，创建新的组合！
red := color.New(color.FgRed)

boldRed := red.Add(color.Bold)
boldRed.Println("这将打印粗体红色文本。")

whiteBackground := red.Add(color.BgWhite)
whiteBackground.Println("红色文本配白色背景。")

// 使用你自己的 io.Writer 输出
color.New(color.FgBlue).Fprintln(myWriter, "蓝色！")

blue := color.New(color.FgBlue)
blue.Fprint(myWriter, "这将打印蓝色文本。")
```

## 创建 PrintXxx 函数

你可以创建 PrintXxx 函数来进一步简化：

```go
// 创建自定义打印函数以方便使用
red := color.New(color.FgRed).PrintfFunc()
red("警告")
red("错误：%s", err)

// 混合多个属性
notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
notice("不要忘记这个...")
```

## 使用 FprintXxx 函数

你也可以使用 FprintXxx 函数传入你自己的 io.Writer：

```go
blue := color.New(FgBlue).FprintfFunc()
blue(myWriter, "重要通知：%s", stars)

// 混合多个属性
success := color.New(color.Bold, color.FgGreen).FprintlnFunc()
success(myWriter, "不要忘记这个...")
```

## 创建 SprintXxx 函数

或者创建 SprintXxx 函数来将字符串与其他非彩色字符串混合：

```go
yellow := New(FgYellow).SprintFunc()
red := New(FgRed).SprintFunc()

fmt.Printf("这是一个 %s，这是一个 %s。\n", yellow("警告"), red("错误"))

info := New(FgWhite, BgGreen).SprintFunc()
fmt.Printf("这个 %s 太棒了！\n", info("包"))
```

## Windows 支持

Windows 支持默认启用。所有 Print 函数都能按预期工作。 但是，仅对于 color.SprintXXX 函数，用户应该使用 fmt.FprintXXX 并将输出设置为 color.Output：

```go
fmt.Fprintf(color.Output, "Windows 支持：%s", color.SGreen("通过"))

info := New(FgWhite, BgGreen).SprintFunc()
fmt.Fprintf(color.Output, "这个 %s 太棒了！\n", info("包"))
```

## 与现有代码集成

可以与现有代码一起使用。只需使用 Set() 方法将标准输出设置为给定参数。 这样就不需要重写现有代码。

```go
// 使用便捷的标准颜色。
color.Set(color.FgYellow)

fmt.Println("现有文本现在将显示为黄色")
fmt.Printf("这个也是 %s\n", "黄色")

color.Unset() // 不要忘记取消设置

// 你可以混合参数
color.Set(color.FgMagenta, color.Bold)
defer color.Unset() // 在你的函数中使用它

fmt.Println("所有文本现在将显示为粗体洋红色。")
```

## 禁用颜色输出

可能会有需要禁用颜色输出的情况（例如将应用程序的标准输出管道传输到其他地方）。 `Color` 支持全局和单个颜色定义禁用颜色。例如，假设你有一个 CLI 应用程序 和一个 `--no-color` 布尔标志。你可以轻松禁用颜色输出：

```go
var flagNoColor = flag.Bool("no-color", false, "禁用颜色输出")

if *flagNoColor {
    color.NoColor = true // 禁用彩色输出
}
```

你也可以通过将 NO_COLOR 环境变量设置为任何值来禁用颜色。

它还支持单个颜色定义（本地）。你可以随时禁用/启用颜色输出：

```go
c := color.New(color.FgCyan)
c.Println("打印青色文本")

c.DisableColor()
c.Println("这将不打印任何颜色")

c.EnableColor()
c.Println("这又打印青色了...")
```

---

## 变量

```go
var (
    // NoColor 定义输出是否着色。它根据 stdout 的文件描述符是否指向终端
    // 动态设置为 false 或 true。如果设置了 NO_COLOR 环境变量（无论其值是什么），
    // 它也会被设置为 true。这是一个全局选项, 影响所有颜色。
    // 如需对每个颜色块进行更多控制，请单独使用 DisableColor() 方法。
    NoColor = noColorIsSet() || os.Getenv("TERM") == "dumb" || !stdoutIsTerminal()

    // Output 定义打印函数的标准输出。默认使用 stdOut()。
    Output = stdOut()

    // Error 定义打印函数的标准错误输出。默认使用 stdErr()。
    Error = stdErr()
)
```

---

## 函数

### Black

```go
func Black(format string, a ...interface{})
```

Black 以黑色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SBlack

```go
func SBlack(format string, a ...interface{}) string
```

SBlack 返回带有黑色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有黑色前景的字符串

---

### Blue

```go
func Blue(format string, a ...interface{})
```

Blue 以蓝色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SBlue

```go
func SBlue(format string, a ...interface{}) string
```

SBlue 返回带有蓝色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有蓝色前景的字符串

---

### Cyan

```go
func Cyan(format string, a ...interface{})
```

Cyan 以青色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SCyan

```go
func SCyan(format string, a ...interface{}) string
```

SCyan 返回带有青色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有青色前景的字符串

---

### Green

```go
func Green(format string, a ...interface{})
```

Green 以绿色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SGreen

```go
func SGreen(format string, a ...interface{}) string
```

SGreen 返回带有绿色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有绿色前景的字符串

---

### HiBlack

```go
func HiBlack(format string, a ...interface{})
```

HiBlack 以高亮黑色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiBlack

```go
func SHiBlack(format string, a ...interface{}) string
```

SHiBlack 返回带有高亮黑色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮黑色前景的字符串

---

### HiBlue

```go
func HiBlue(format string, a ...interface{})
```

HiBlue 以高亮蓝色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiBlue

```go
func SHiBlue(format string, a ...interface{}) string
```

SHiBlue 返回带有高亮蓝色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮蓝色前景的字符串

---

### HiCyan

```go
func HiCyan(format string, a ...interface{})
```

HiCyan 以高亮青色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiCyan

```go
func SHiCyan(format string, a ...interface{}) string
```

SHiCyan 返回带有高亮青色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮青色前景的字符串

---

### HiGreen

```go
func HiGreen(format string, a ...interface{})
```

HiGreen 以高亮绿色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiGreen

```go
func SHiGreen(format string, a ...interface{}) string
```

SHiGreen 返回带有高亮绿色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮绿色前景的字符串

---

### HiMagenta

```go
func HiMagenta(format string, a ...interface{})
```

HiMagenta 以高亮洋红色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiMagenta

```go
func SHiMagenta(format string, a ...interface{}) string
```

SHiMagenta 返回带有高亮洋红色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮洋红色前景的字符串

---

### HiRed

```go
func HiRed(format string, a ...interface{})
```

HiRed 以高亮红色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiRed

```go
func SHiRed(format string, a ...interface{}) string
```

SHiRed 返回带有高亮红色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮红色前景的字符串

---

### HiWhite

```go
func HiWhite(format string, a ...interface{})
```

HiWhite 以高亮白色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiWhite

```go
func SHiWhite(format string, a ...interface{}) string
```

SHiWhite 返回带有高亮白色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮白色前景的字符串

---

### HiYellow

```go
func HiYellow(format string, a ...interface{})
```

HiYellow 以高亮黄色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SHiYellow

```go
func SHiYellow(format string, a ...interface{}) string
```

SHiYellow 返回带有高亮黄色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有高亮黄色前景的字符串

---

### Magenta

```go
func Magenta(format string, a ...interface{})
```

Magenta 以洋红色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SMagenta

```go
func SMagenta(format string, a ...interface{}) string
```

SMagenta 返回带有洋红色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有洋红色前景的字符串

---

### Red

```go
func Red(format string, a ...interface{})
```

Red 以红色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SRed

```go
func SRed(format string, a ...interface{}) string
```

SRed 返回带有红色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有红色前景的字符串

---

### Unset

```go
func Unset()
```

Unset 重置所有转义属性并清除输出。 通常在 Set() 之后调用。

---

### White

```go
func White(format string, a ...interface{})
```

White 以白色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SWhite

```go
func SWhite(format string, a ...interface{}) string
```

SWhite 返回带有白色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有白色前景的字符串

---

### Yellow

```go
func Yellow(format string, a ...interface{})
```

Yellow 以黄色前景打印文本。 默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### SYellow

```go
func SYellow(format string, a ...interface{}) string
```

SYellow 返回带有黄色前景的字符串。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 带有黄色前景的字符串

---

## 类型

### Attribute

```go
type Attribute int
```

Attribute 定义单个 SGR（Select Graphic Rendition）代码， 用于控制终端文本的显示属性，如颜色、样式等。

#### 基础文本样式属性

这些属性控制文本的显示样式，如加粗、斜体、下划线等。

```go
const (
    Reset        Attribute = iota // 重置所有属性为默认值
    Bold                          // 加粗文本
    Faint                         // 轻淡/暗淡文本（降低亮度）
    Italic                        // 斜体文本
    Underline                     // 下划线文本
    BlinkSlow                     // 慢速闪烁（每秒少于 150 次）
    BlinkRapid                    // 快速闪烁（每分钟 150 次以上）
    ReverseVideo                  // 反显（前景色和背景色互换）
    Concealed                     // 隐蔽/隐藏文本（与背景色相同）
    CrossedOut                    // 删除线文本
)
```

#### 重置属性

这些属性用于取消对应的文本样式。

```go
const (
    ResetBold      Attribute = iota + 22 // 重置加粗
    ResetItalic                          // 重置斜体
    ResetUnderline                       // 重置下划线
    ResetBlinking                        // 重置闪烁

    ResetReversed   // 重置反显
    ResetConcealed  // 重置蔽蔽
    ResetCrossedOut // 重置删除线
)
```

#### 前景文本颜色（标准 8 色）

这些颜色在大多数终端中都受支持。

```go
const (
    FgBlack   Attribute = iota + 30 // 黑色前景
    FgRed                           // 红色前景
    FgGreen                         // 绿色前景
    FgYellow                        // 黄色前景
    FgBlue                          // 蓝色前景
    FgMagenta                       // 洋红色/品红色前景
    FgCyan                          // 青色前景
    FgWhite                         // 白色前景
)
```

#### 前景高亮文本颜色（高亮 8 色）

这些是高亮版本的标准前景色，比标准色更明亮。

```go
const (
    FgHiBlack   Attribute = iota + 90 // 高亮黑色前景（灰色）
    FgHiRed                           // 高亮红色前景
    FgHiGreen                         // 高亮绿色前景
    FgHiYellow                        // 高亮黄色前景
    FgHiBlue                          // 高亮蓝色前景
    FgHiMagenta                       // 高亮洋红色前景
    FgHiCyan                          // 高亮青色前景
    FgHiWhite                         // 高亮白色前景
)
```

#### 背景文本颜色（标准 8 色）

这些颜色用于设置文本的背景色。

```go
const (
    BgBlack   Attribute = iota + 40 // 黑色背景
    BgRed                           // 红色背景
    BgGreen                         // 绿色背景
    BgYellow                        // 黄色背景
    BgBlue                          // 蓝色背景
    BgMagenta                       // 洋红色/品红色背景
    BgCyan                          // 青色背景
    BgWhite                         // 白色背景
)
```

#### 背景高亮文本颜色（高亮 8 色）

这些是高亮版本的背景色，比标准色更明亮。

```go
const (
    BgHiBlack   Attribute = iota + 100 // 高亮黑色背景（灰色）
    BgHiRed                            // 高亮红色背景
    BgHiGreen                          // 高亮绿色背景
    BgHiYellow                         // 高亮黄色背景
    BgHiBlue                           // 高亮蓝色背景
    BgHiMagenta                        // 高亮洋红色背景
    BgHiCyan                           // 高亮青色背景
    BgHiWhite                          // 高亮白色背景
)
```

---

### Color

```go
type Color struct {
    // Has unexported fields.
}
```

Color 定义一个由 SGR 参数定义的自定义颜色对象。

#### BgRGB

```go
func BgRGB(r, g, b int) *Color
```

BgRGB 返回一个新的 24 位 RGB 背景色。

**参数:**
- `r`: 红色分量 (0-255)
- `g`: 绿色分量 (0-255)
- `b`: 蓝色分量 (0-255)

**返回值:**
- `*Color`: 配置好的颜色对象

---

#### New

```go
func New(value ...Attribute) *Color
```

New 返回一个新创建的颜色对象。

**参数:**
- `value`: 任意数量的 SGR 参数。

**返回值:**
- `*Color`: 新创建的颜色对象。

---

#### RGB

```go
func RGB(r, g, b int) *Color
```

RGB 返回一个新的 24 位 RGB 前景色。

**参数:**
- `r`: 红色分量 (0-255)
- `g`: 绿色分量 (0-255)
- `b`: 蓝色分量 (0-255)

**返回值:**
- `*Color`: 配置好的颜色对象

---

#### Set

```go
func Set(p ...Attribute) *Color
```

Set 立即设置给定的 SGR 参数。 将使用给定的 SGR 参数更改输出颜色，直到调用 color.Unset() 为止。

**参数:**
- `p`: 任意数量的 SGR 参数

**返回值:**
- `*Color`: 配置好的颜色对象

---

#### Add

```go
func (c *Color) Add(value ...Attribute) *Color
```

Add 用于链式添加 SGR 参数。 可以使用任意数量的参数进行组合并创建自定义颜色对象。

**参数:**
- `value`: 任意数量的 SGR 参数

**返回值:**
- `*Color`: 当前颜色对象，支持链式调用

**示例:**

```go
c.Add(FgRed, Underline).Println("红色下划线文本")
```

---

#### AddBgRGB

```go
func (c *Color) AddBgRGB(r, g, b int) *Color
```

AddBgRGB 用于链式添加背景 RGB SGR 参数。 可以使用任意数量的参数进行组合并创建自定义颜色对象。

**参数:**
- `r`: 红色分量 (0-255)
- `g`: 绿色分量 (0-255)
- `b`: 蓝色分量 (0-255)

**返回值:**
- `*Color`: 当前颜色对象，支持链式调用

**示例:**

```go
c.AddBgRGB(255, 128, 0).Println("橙色背景")
```

---

#### AddRGB

```go
func (c *Color) AddRGB(r, g, b int) *Color
```

AddRGB 用于链式添加前景 RGB SGR 参数。 可以使用任意数量的参数进行组合并创建自定义颜色对象。

**参数:**
- `r`: 红色分量 (0-255)
- `g`: 绿色分量 (0-255)
- `b`: 蓝色分量 (0-255)

**返回值:**
- `*Color`: 当前颜色对象，支持链式调用

**示例:**

```go
c.AddRGB(255, 128, 0).Println("橙色文本")
```

---

#### DisableColor

```go
func (c *Color) DisableColor()
```

DisableColor 禁用颜色输出。 可用于在不更改任何现有代码的情况下禁用颜色输出，例如配合 "--no-color" 标志使用。 要重新启用，请使用 EnableColor() 方法。

---

#### EnableColor

```go
func (c *Color) EnableColor()
```

EnableColor 启用颜色输出。 与 DisableColor() 一起使用。如果颜色未被禁用，此方法没有副作用。

---

#### Equals

```go
func (c *Color) Equals(c2 *Color) bool
```

Equals 比较两种颜色是否相等。

**参数:**
- `c2`: 要比较的另一个颜色对象

**返回值:**
- `bool`: 如果两种颜色相等则返回 true

---

#### Fprint

```go
func (c *Color) Fprint(w io.Writer, a ...interface{}) (n int, err error)
```

Fprint 使用其操作数的默认格式进行格式化并写入 w。 当操作数都不是字符串时，在它们之间添加空格。

**参数:**
- `w`: 目标写入器
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

> **注意:** 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。

---

#### FprintFunc

```go
func (c *Color) FprintFunc() func(w io.Writer, a ...interface{})
```

FprintFunc 返回一个新函数，该函数使用 color.Fprint() 将传入的参数打印为彩色。

**返回值:**
- `func(w io.Writer, a ...interface{})`: 打印函数

---

#### Fprintf

```go
func (c *Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

Fprintf 根据格式说明符进行格式化并写入 w。

**参数:**
- `w`: 目标写入器
- `format`: 格式字符串
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

> **注意:** 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。

---

#### FprintfFunc

```go
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...interface{})
```

FprintfFunc 返回一个新函数，该函数使用 color.Fprintf() 将传入的参数打印为彩色。

**返回值:**
- `func(w io.Writer, format string, a ...interface{})`: 打印函数

---

#### Fprintln

```go
func (c *Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

Fprintln 使用其操作数的默认格式进行格式化并写入 w，并在末尾添加换行符。

**参数:**
- `w`: 目标写入器
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

> **注意:** 在 Windows 上，如果 w 是 *os.File 类型，用户应该用 colorable.NewColorable() 包装 w。

---

#### FprintlnFunc

```go
func (c *Color) FprintlnFunc() func(w io.Writer, a ...interface{})
```

FprintlnFunc 返回一个新函数，该函数使用 color.Fprintln() 将传入的参数打印为彩色。

**返回值:**
- `func(w io.Writer, a ...interface{})`: 打印函数

---

#### Print

```go
func (c *Color) Print(a ...interface{}) (n int, err error)
```

Print 使用其操作数的默认格式进行格式化并写入标准输出。 当操作数都不是字符串时，在它们之间添加空格。

**参数:**
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

---

#### PrintFunc

```go
func (c *Color) PrintFunc() func(a ...interface{})
```

PrintFunc 返回一个新函数，该函数使用 color.Print() 将传入的参数打印为彩色。

**返回值:**
- `func(a ...interface{})`: 打印函数

---

#### Printf

```go
func (c *Color) Printf(format string, a ...interface{}) (n int, err error)
```

Printf 根据格式说明符进行格式化并写入标准输出。

**参数:**
- `format`: 格式字符串
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

---

#### PrintfFunc

```go
func (c *Color) PrintfFunc() func(format string, a ...interface{})
```

PrintfFunc 返回一个新函数，该函数使用 color.Printf() 将传入的参数打印为彩色。

**返回值:**
- `func(format string, a ...interface{})`: 打印函数

---

#### Println

```go
func (c *Color) Println(a ...interface{}) (n int, err error)
```

Println 使用其操作数的默认格式进行格式化并写入标准输出，并在末尾添加换行符。

**参数:**
- `a`: 要格式化的操作数

**返回值:**
- `int`: 写入的字节数
- `error`: 写入过程中的错误（如果有）

---

#### PrintlnFunc

```go
func (c *Color) PrintlnFunc() func(a ...interface{})
```

PrintlnFunc 返回一个新函数，该函数使用 color.Println() 将传入的参数打印为彩色。

**返回值:**
- `func(a ...interface{})`: 打印函数

---

#### Sprint

```go
func (c *Color) Sprint(a ...interface{}) string
```

Sprint 使用其操作数的默认格式进行格式化并返回结果字符串。 当操作数都不是字符串时，在它们之间添加空格。

**参数:**
- `a`: 要格式化的操作数

**返回值:**
- `string`: 格式化后的字符串

---

#### SprintFunc

```go
func (c *Color) SprintFunc() func(a ...interface{}) string
```

SprintFunc 返回一个新函数，该函数使用 color.Sprint() 将传入的参数格式化为彩色字符串。

**返回值:**
- `func(a ...interface{}) string`: 格式化函数

---

#### Sprintf

```go
func (c *Color) Sprintf(format string, a ...interface{}) string
```

Sprintf 根据格式说明符进行格式化并返回结果字符串。

**参数:**
- `format`: 格式字符串
- `a`: 要格式化的操作数

**返回值:**
- `string`: 格式化后的字符串

---

#### SprintfFunc

```go
func (c *Color) SprintfFunc() func(format string, a ...interface{}) string
```

SprintfFunc 返回一个新函数，该函数使用 color.Sprintf() 将传入的参数格式化为彩色字符串。

**返回值:**
- `func(format string, a ...interface{}) string`: 格式化函数

---

#### Sprintln

```go
func (c *Color) Sprintln(a ...interface{}) string
```

Sprintln 使用其操作数的默认格式进行格式化并返回结果字符串，并在末尾添加换行符。

**参数:**
- `a`: 要格式化的操作数

**返回值:**
- `string`: 格式化后的字符串

---

#### SprintlnFunc

```go
func (c *Color) SprintlnFunc() func(a ...interface{}) string
```

SprintlnFunc 返回一个新函数，该函数使用 color.Sprintln() 将传入的参数格式化为彩色字符串。

**返回值:**
- `func(a ...interface{}) string`: 格式化函数

---

#### UnsetWriter

```go
func (c *Color) UnsetWriter(w io.Writer)
```

UnsetWriter 使用给定的 io.Writer 重置所有转义属性并清除输出。 通常在 SetWriter() 之后调用。

**参数:**
- `w`: 目标写入器

---

## 全局实例

全局实例提供了一种便捷的方式来使用颜色输出，无需手动创建 Color 对象。全局实例具有独立的配置管理，支持样式设置（加粗、下划线等），并且是线程安全的。

### 使用示例

```go
// 使用全局实例直接打印
color.GetGlobal().Red("错误信息")
color.GetGlobal().Green("成功信息")

// 设置样式配置
color.GetGlobal().SetBold(true).SetUnderline(true)
color.GetGlobal().Blue("带下划线的蓝色粗体文本")

// 使用日志级别方法
color.GetGlobal().Info("这是一条信息")
color.GetGlobal().Success("操作成功")
color.GetGlobal().Warning("警告信息")
color.GetGlobal().Error("错误信息")

// 使用终端前缀方法
color.GetGlobal().Ok("操作完成")
color.GetGlobal().Warn("需要注意")
color.GetGlobal().Err("发生错误")

// 获取带颜色的字符串（不打印）
redText := color.GetGlobal().SRed("红色文本")
```

---

### GlobalColor

```go
type GlobalColor struct {
    // Has unexported fields.
}
```

GlobalColor 是全局颜色实例的类型，包含独立的配置和输出设置，与普通的 Color 实例完全分离。

---

### StyleConfig

```go
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
```

StyleConfig 定义颜色样式配置。

#### Clone

```go
func (s *StyleConfig) Clone() *StyleConfig
```

Clone 创建样式配置的深拷贝，返回一个新的 StyleConfig 实例，复制当前配置的所有字段。

**注意:** Output 字段是接口类型，只会复制引用，不会深拷贝底层的写入器。

**返回值:**
- `*StyleConfig`: 新的样式配置实例

---

### 全局实例函数

#### GetGlobal

```go
func GetGlobal() *GlobalColor
```

GetGlobal 返回全局颜色实例。首次调用时会自动初始化全局实例。

**返回值:**
- `*GlobalColor`: 全局颜色实例

---

#### ResetGlobal

```go
func ResetGlobal()
```

ResetGlobal 重置全局实例到默认状态。这会重新创建全局实例并应用默认配置。

---

### GlobalColor 方法

#### SetConfig

```go
func (g *GlobalColor) SetConfig(config *StyleConfig) *GlobalColor
```

SetConfig 设置样式配置。会自动克隆传入的配置，防止外部修改影响全局实例。如果传入 nil，则使用默认配置。

**参数:**
- `config`: 要设置的样式配置

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### GetConfig

```go
func (g *GlobalColor) GetConfig() *StyleConfig
```

GetConfig 获取当前样式配置。返回的是配置对象的指针，直接修改会影响全局实例。如果需要修改配置，建议使用 GetConfigClone() 或 SetConfig() 方法。

**返回值:**
- `*StyleConfig`: 当前样式配置

---

#### GetConfigClone

```go
func (g *GlobalColor) GetConfigClone() *StyleConfig
```

GetConfigClone 获取当前样式配置的克隆。返回一个新的配置对象，修改它不会影响全局实例。适合用于基于当前配置创建新配置的场景。

**返回值:**
- `*StyleConfig`: 当前样式配置的克隆

---

#### SetOutput

```go
func (g *GlobalColor) SetOutput(w io.Writer) *GlobalColor
```

SetOutput 设置输出目标。

**参数:**
- `w`: 输出写入器

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetNoColor

```go
func (g *GlobalColor) SetNoColor(noColor bool) *GlobalColor
```

SetNoColor 设置是否禁用颜色。当禁用颜色时，所有样式效果（加粗、下划线等）也会被禁用。

**参数:**
- `noColor`: true 表示禁用颜色，false 表示启用颜色

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetBold

```go
func (g *GlobalColor) SetBold(bold bool) *GlobalColor
```

SetBold 设置是否启用加粗。

**参数:**
- `bold`: true 表示启用加粗，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetUnderline

```go
func (g *GlobalColor) SetUnderline(underline bool) *GlobalColor
```

SetUnderline 设置是否启用下划线。

**参数:**
- `underline`: true 表示启用下划线，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetItalic

```go
func (g *GlobalColor) SetItalic(italic bool) *GlobalColor
```

SetItalic 设置是否启用斜体。

**参数:**
- `italic`: true 表示启用斜体，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetBlink

```go
func (g *GlobalColor) SetBlink(blink bool) *GlobalColor
```

SetBlink 设置是否启用闪烁。

**参数:**
- `blink`: true 表示启用闪烁，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetFaint

```go
func (g *GlobalColor) SetFaint(faint bool) *GlobalColor
```

SetFaint 设置是否启用暗淡效果。

**参数:**
- `faint`: true 表示启用暗淡效果，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

#### SetCrossedOut

```go
func (g *GlobalColor) SetCrossedOut(crossedOut bool) *GlobalColor
```

SetCrossedOut 设置是否启用删除线。

**参数:**
- `crossedOut`: true 表示启用删除线，false 表示禁用

**返回值:**
- `*GlobalColor`: 当前 GlobalColor 对象，支持链式调用

---

### 颜色打印方法

#### Black

```go
func (g *GlobalColor) Black(format string, a ...interface{})
```

Black 使用黑色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Red

```go
func (g *GlobalColor) Red(format string, a ...interface{})
```

Red 使用红色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Green

```go
func (g *GlobalColor) Green(format string, a ...interface{})
```

Green 使用绿色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Yellow

```go
func (g *GlobalColor) Yellow(format string, a ...interface{})
```

Yellow 使用黄色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Blue

```go
func (g *GlobalColor) Blue(format string, a ...interface{})
```

Blue 使用蓝色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Magenta

```go
func (g *GlobalColor) Magenta(format string, a ...interface{})
```

Magenta 使用洋红色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Cyan

```go
func (g *GlobalColor) Cyan(format string, a ...interface{})
```

Cyan 使用青色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### White

```go
func (g *GlobalColor) White(format string, a ...interface{})
```

White 使用白色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### 返回字符串的颜色方法

#### SBlack

```go
func (g *GlobalColor) SBlack(format string, a ...interface{}) string
```

SBlack 返回黑色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SRed

```go
func (g *GlobalColor) SRed(format string, a ...interface{}) string
```

SRed 返回红色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SGreen

```go
func (g *GlobalColor) SGreen(format string, a ...interface{}) string
```

SGreen 返回绿色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SYellow

```go
func (g *GlobalColor) SYellow(format string, a ...interface{}) string
```

SYellow 返回黄色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SBlue

```go
func (g *GlobalColor) SBlue(format string, a ...interface{}) string
```

SBlue 返回蓝色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SMagenta

```go
func (g *GlobalColor) SMagenta(format string, a ...interface{}) string
```

SMagenta 返回洋红色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SCyan

```go
func (g *GlobalColor) SCyan(format string, a ...interface{}) string
```

SCyan 返回青色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SWhite

```go
func (g *GlobalColor) SWhite(format string, a ...interface{}) string
```

SWhite 返回白色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

### 高亮颜色打印方法

#### HiRed

```go
func (g *GlobalColor) HiRed(format string, a ...interface{})
```

HiRed 使用高亮红色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiGreen

```go
func (g *GlobalColor) HiGreen(format string, a ...interface{})
```

HiGreen 使用高亮绿色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiYellow

```go
func (g *GlobalColor) HiYellow(format string, a ...interface{})
```

HiYellow 使用高亮黄色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiBlue

```go
func (g *GlobalColor) HiBlue(format string, a ...interface{})
```

HiBlue 使用高亮蓝色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiMagenta

```go
func (g *GlobalColor) HiMagenta(format string, a ...interface{})
```

HiMagenta 使用高亮洋红色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiCyan

```go
func (g *GlobalColor) HiCyan(format string, a ...interface{})
```

HiCyan 使用高亮青色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### HiWhite

```go
func (g *GlobalColor) HiWhite(format string, a ...interface{})
```

HiWhite 使用高亮白色样式打印文本。默认会在 format 末尾追加换行符。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### 返回字符串的高亮颜色方法

#### SHiRed

```go
func (g *GlobalColor) SHiRed(format string, a ...interface{}) string
```

SHiRed 返回高亮红色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiGreen

```go
func (g *GlobalColor) SHiGreen(format string, a ...interface{}) string
```

SHiGreen 返回高亮绿色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiYellow

```go
func (g *GlobalColor) SHiYellow(format string, a ...interface{}) string
```

SHiYellow 返回高亮黄色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiBlue

```go
func (g *GlobalColor) SHiBlue(format string, a ...interface{}) string
```

SHiBlue 返回高亮蓝色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiMagenta

```go
func (g *GlobalColor) SHiMagenta(format string, a ...interface{}) string
```

SHiMagenta 返回高亮洋红色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiCyan

```go
func (g *GlobalColor) SHiCyan(format string, a ...interface{}) string
```

SHiCyan 返回高亮青色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

#### SHiWhite

```go
func (g *GlobalColor) SHiWhite(format string, a ...interface{}) string
```

SHiWhite 返回高亮白色样式的字符串（不打印）。

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

**返回值:**
- `string`: 格式化后的字符串

---

### 日志级别方法

#### Info

```go
func (g *GlobalColor) Info(format string, a ...interface{})
```

Info 以信息级别（青色）打印日志。格式: `[INFO] xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Success

```go
func (g *GlobalColor) Success(format string, a ...interface{})
```

Success 以成功级别（绿色）打印日志。格式: `[SUCCESS] xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Warning

```go
func (g *GlobalColor) Warning(format string, a ...interface{})
```

Warning 以警告级别（黄色）打印日志。格式: `[WARN] xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Error

```go
func (g *GlobalColor) Error(format string, a ...interface{})
```

Error 以错误级别（红色）打印日志。格式: `[ERROR] xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Debug

```go
func (g *GlobalColor) Debug(format string, a ...interface{})
```

Debug 以调试级别（洋红色）打印日志。格式: `[DEBUG] xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

### 终端提示信息前缀方法

#### Ok

```go
func (g *GlobalColor) Ok(format string, a ...interface{})
```

Ok 以 OK 前缀（绿色）打印成功信息。格式: `ok: xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Warn

```go
func (g *GlobalColor) Warn(format string, a ...interface{})
```

Warn 以 WARN 前缀（黄色）打印警告信息。格式: `warn: xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

---

#### Err

```go
func (g *GlobalColor) Err(format string, a ...interface{})
```

Err 以 ERR 前缀（红色）打印错误信息。格式: `err: xxx`

**参数:**
- `format`: 格式字符串
- `a`: 格式化参数

