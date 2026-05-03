package color

import (
	"fmt"
	"strings"
	"sync"
)

var (
	// colorsCache 用于减少创建的 Color 对象数量，并允许重用已创建的具有所需 Attribute 的对象。
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // 保护 colorsCache
)

// boolPtr 返回指向布尔值的指针。
//
// 参数:
//   - v: 布尔值
//
// 返回值:
//   - *bool: 指向该布尔值的指针
func boolPtr(v bool) *bool {
	return &v
}

// getCachedColor 从缓存中获取或创建 Color 对象。
// 如果缓存中不存在，则创建新对象并存入缓存。
//
// 参数:
//   - p: SGR 属性
//
// 返回值:
//   - *Color: 对应属性的颜色对象
func getCachedColor(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		c = New(p)
		colorsCache[p] = c
	}

	return c
}

// colorPrint 是便捷函数的底层实现，用于打印彩色文本。
// 如果 format 不以换行符结尾，会自动追加换行符。
//
// 参数:
//   - format: 格式字符串
//   - p: 颜色属性
//   - a: 格式化参数
func colorPrint(format string, p Attribute, a ...interface{}) {
	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	if len(a) == 0 {
		_, _ = c.Print(format)
	} else {
		_, _ = c.Printf(format, a...)
	}
}

// colorString 是便捷函数的底层实现，用于返回彩色字符串。
//
// 参数:
//   - format: 格式字符串
//   - p: 颜色属性
//   - a: 格式化参数
//
// 返回值:
//   - string: 格式化后的彩色字符串
func colorString(format string, p Attribute, a ...interface{}) string {
	c := getCachedColor(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

// sprintln 是一个辅助函数，用于使用 fmt.Sprintln 格式化字符串并去除末尾的换行符。
//
// 参数:
//   - a: 要格式化的操作数
//
// 返回值:
//   - string: 格式化后的字符串（不含末尾换行符）
func sprintln(a ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}
