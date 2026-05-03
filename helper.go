package color

// Black 以黑色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Black(format string, a ...interface{}) { colorPrint(format, FgBlack, a...) }

// Red 以红色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Red(format string, a ...interface{}) { colorPrint(format, FgRed, a...) }

// Green 以绿色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Green(format string, a ...interface{}) { colorPrint(format, FgGreen, a...) }

// Yellow 以黄色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Yellow(format string, a ...interface{}) { colorPrint(format, FgYellow, a...) }

// Blue 以蓝色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Blue(format string, a ...interface{}) { colorPrint(format, FgBlue, a...) }

// Magenta 以洋红色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Magenta(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }

// Cyan 以青色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func Cyan(format string, a ...interface{}) { colorPrint(format, FgCyan, a...) }

// White 以白色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func White(format string, a ...interface{}) { colorPrint(format, FgWhite, a...) }

// BlackString 返回带有黑色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有黑色前景的字符串
func BlackString(format string, a ...interface{}) string { return colorString(format, FgBlack, a...) }

// RedString 返回带有红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有红色前景的字符串
func RedString(format string, a ...interface{}) string { return colorString(format, FgRed, a...) }

// GreenString 返回带有绿色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有绿色前景的字符串
func GreenString(format string, a ...interface{}) string { return colorString(format, FgGreen, a...) }

// YellowString 返回带有黄色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有黄色前景的字符串
func YellowString(format string, a ...interface{}) string { return colorString(format, FgYellow, a...) }

// BlueString 返回带有蓝色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有蓝色前景的字符串
func BlueString(format string, a ...interface{}) string { return colorString(format, FgBlue, a...) }

// MagentaString 返回带有洋红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有洋红色前景的字符串
func MagentaString(format string, a ...interface{}) string {
	return colorString(format, FgMagenta, a...)
}

// CyanString 返回带有青色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有青色前景的字符串
func CyanString(format string, a ...interface{}) string { return colorString(format, FgCyan, a...) }

// WhiteString 返回带有白色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有白色前景的字符串
func WhiteString(format string, a ...interface{}) string { return colorString(format, FgWhite, a...) }

// HiBlack 以高亮黑色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiBlack(format string, a ...interface{}) { colorPrint(format, FgHiBlack, a...) }

// HiRed 以高亮红色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiRed(format string, a ...interface{}) { colorPrint(format, FgHiRed, a...) }

// HiGreen 以高亮绿色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiGreen(format string, a ...interface{}) { colorPrint(format, FgHiGreen, a...) }

// HiYellow 以高亮黄色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiYellow(format string, a ...interface{}) { colorPrint(format, FgHiYellow, a...) }

// HiBlue 以高亮蓝色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiBlue(format string, a ...interface{}) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta 以高亮洋红色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiMagenta(format string, a ...interface{}) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan 以高亮青色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiCyan(format string, a ...interface{}) { colorPrint(format, FgHiCyan, a...) }

// HiWhite 以高亮白色前景打印文本。
// 默认会在 format 末尾追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
func HiWhite(format string, a ...interface{}) { colorPrint(format, FgHiWhite, a...) }

// HiBlackString 返回带有高亮黑色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮黑色前景的字符串
func HiBlackString(format string, a ...interface{}) string {
	return colorString(format, FgHiBlack, a...)
}

// HiRedString 返回带有高亮红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮红色前景的字符串
func HiRedString(format string, a ...interface{}) string { return colorString(format, FgHiRed, a...) }

// HiGreenString 返回带有高亮绿色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮绿色前景的字符串
func HiGreenString(format string, a ...interface{}) string {
	return colorString(format, FgHiGreen, a...)
}

// HiYellowString 返回带有高亮黄色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮黄色前景的字符串
func HiYellowString(format string, a ...interface{}) string {
	return colorString(format, FgHiYellow, a...)
}

// HiBlueString 返回带有高亮蓝色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮蓝色前景的字符串
func HiBlueString(format string, a ...interface{}) string { return colorString(format, FgHiBlue, a...) }

// HiMagentaString 返回带有高亮洋红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮洋红色前景的字符串
func HiMagentaString(format string, a ...interface{}) string {
	return colorString(format, FgHiMagenta, a...)
}

// HiCyanString 返回带有高亮青色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮青色前景的字符串
func HiCyanString(format string, a ...interface{}) string { return colorString(format, FgHiCyan, a...) }

// HiWhiteString 返回带有高亮白色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮白色前景的字符串
func HiWhiteString(format string, a ...interface{}) string {
	return colorString(format, FgHiWhite, a...)
}
