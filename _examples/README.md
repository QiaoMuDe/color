# color 包使用示例

本目录包含 color 包的各种使用示例，每个示例都是独立的可运行程序。

## 目录结构

```
_examples/
├── README.md          # 本文件
├── basic/             # 基础用法示例
│   └── main.go
├── global/            # 全局实例用法示例
│   └── main.go
├── utils/             # 终端 UI 辅助函数示例
│   ├── utils.go       # 辅助函数定义
│   └── main.go        # 演示调用
└── advanced/          # 高级用法示例（待添加）
    └── main.go
```

## 运行示例

### 基础用法示例

演示 color 包的基本功能：便捷函数、链式调用、RGB颜色等。

```bash
cd basic
go run main.go
```

### 全局实例示例

演示如何使用全局实例进行便捷操作，默认启用白色加粗样式。

```bash
cd global
go run main.go
```

### 终端 UI 辅助函数示例

演示如何使用 color 库构建常用的终端 UI 组件（表格、日志、进度条等）。

```bash
cd utils
go run .
```

> **注意**: 由于 utils 示例包含多个文件（utils.go 和 main.go），需要使用 `go run .` 而不是 `go run main.go`。

## 示例内容概览

### basic/main.go

- 便捷函数（Red/Green/Blue 等）
- 格式化输出
- 链式调用组合样式
- RGB 真彩色
- 输出到不同目标（stderr、文件）
- 函数生成器
- 颜色控制（禁用/启用）
- 实际应用场景（日志级别、表格输出）

### global/main.go

- 基本使用（默认白色加粗）
- 颜色快捷方法（GRed/GGreen 等）
- 高亮颜色快捷方法
- 日志级别快捷方法（GInfo/GSuccess/GWarn/GError/GDebug）
- 样式组合
- 动态添加样式
- 禁用/启用颜色
- 获取全局实例进行高级操作
- 实际应用场景（命令行工具输出）

### utils/main.go & utils.go

- 章节标题和分隔线
- 日志级别打印（Info/Success/Warn/Error/Debug）
- 表格输出（带表头、数据行、分隔线）
- 列表输出（无序列表、有序列表）
- 状态标签（运行中/已停止/警告/待处理）
- 进度条（带动画效果）
- 代码块和引用文本
- 提示框（提示/警告/错误/成功）
- 综合示例：系统监控面板

## 添加新示例

如需添加新的示例，请：

1. 在 `_examples/` 下创建新的子目录
2. 在该目录下创建 `main.go` 文件
3. 确保 `package main` 和 `func main()` 定义正确
4. 在本 README 中添加说明
