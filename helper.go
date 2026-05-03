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

// SBlack 返回带有黑色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有黑色前景的字符串
func SBlack(format string, a ...interface{}) string { return colorString(format, FgBlack, a...) }

// SRed 返回带有红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有红色前景的字符串
func SRed(format string, a ...interface{}) string { return colorString(format, FgRed, a...) }

// SGreen 返回带有绿色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有绿色前景的字符串
func SGreen(format string, a ...interface{}) string { return colorString(format, FgGreen, a...) }

// SYellow 返回带有黄色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有黄色前景的字符串
func SYellow(format string, a ...interface{}) string { return colorString(format, FgYellow, a...) }

// SBlue 返回带有蓝色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有蓝色前景的字符串
func SBlue(format string, a ...interface{}) string { return colorString(format, FgBlue, a...) }

// SMagenta 返回带有洋红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有洋红色前景的字符串
func SMagenta(format string, a ...interface{}) string {
	return colorString(format, FgMagenta, a...)
}

// SCyan 返回带有青色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有青色前景的字符串
func SCyan(format string, a ...interface{}) string { return colorString(format, FgCyan, a...) }

// SWhite 返回带有白色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有白色前景的字符串
func SWhite(format string, a ...interface{}) string { return colorString(format, FgWhite, a...) }

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

// SHiBlack 返回带有高亮黑色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮黑色前景的字符串
func SHiBlack(format string, a ...interface{}) string {
	return colorString(format, FgHiBlack, a...)
}

// SHiRed 返回带有高亮红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮红色前景的字符串
func SHiRed(format string, a ...interface{}) string { return colorString(format, FgHiRed, a...) }

// SHiGreen 返回带有高亮绿色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮绿色前景的字符串
func SHiGreen(format string, a ...interface{}) string {
	return colorString(format, FgHiGreen, a...)
}

// SHiYellow 返回带有高亮黄色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮黄色前景的字符串
func SHiYellow(format string, a ...interface{}) string {
	return colorString(format, FgHiYellow, a...)
}

// SHiBlue 返回带有高亮蓝色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮蓝色前景的字符串
func SHiBlue(format string, a ...interface{}) string { return colorString(format, FgHiBlue, a...) }

// SHiMagenta 返回带有高亮洋红色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮洋红色前景的字符串
func SHiMagenta(format string, a ...interface{}) string {
	return colorString(format, FgHiMagenta, a...)
}

// SHiCyan 返回带有高亮青色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮青色前景的字符串
func SHiCyan(format string, a ...interface{}) string { return colorString(format, FgHiCyan, a...) }

// SHiWhite 返回带有高亮白色前景的字符串。
//
// 参数:
//   - format: 格式字符串
//   - a: 格式化参数
//
// 返回值:
//   - string: 带有高亮白色前景的字符串
func SHiWhite(format string, a ...interface{}) string {
	return colorString(format, FgHiWhite, a...)
}
