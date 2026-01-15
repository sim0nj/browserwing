package executor

import (
	"time"

	"github.com/go-rod/rod"
)

// Page 表示一个浏览器页面及其上下文
type Page struct {
	RodPage    *rod.Page
	URL        string
	Title      string
	SemanticTree *SemanticTree
	LastUpdated time.Time
}

// SemanticTree 表示页面的语义树结构
type SemanticTree struct {
	Root     *SemanticNode
	Elements map[string]*SemanticNode // ID -> Node 映射，用于快速查找
}

// SemanticNode 表示页面中的一个可交互元素
type SemanticNode struct {
	ID          string                 // 唯一标识符
	Type        string                 // 元素类型：button, input, link, select, textarea 等
	Role        string                 // ARIA role
	Label       string                 // 元素标签/名称
	Placeholder string                 // 输入框的 placeholder
	Value       string                 // 当前值
	Text        string                 // 元素文本内容
	Selector    string                 // CSS 选择器
	XPath       string                 // XPath 路径
	Attributes  map[string]string      // 所有属性
	Position    *ElementPosition       // 元素位置信息
	IsVisible   bool                   // 是否可见
	IsEnabled   bool                   // 是否可用
	Children    []*SemanticNode        // 子节点
	Metadata    map[string]interface{} // 其他元数据
}

// ElementPosition 元素位置信息
type ElementPosition struct {
	X      float64 // X 坐标
	Y      float64 // Y 坐标
	Width  float64 // 宽度
	Height float64 // 高度
}

// OperationResult 操作结果
type OperationResult struct {
	Success   bool                   `json:"success"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Error     string                 `json:"error,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
}

// NavigateOptions 导航选项
type NavigateOptions struct {
	WaitUntil string        // 等待条件：load, domcontentloaded, networkidle
	Timeout   time.Duration // 超时时间
}

// ClickOptions 点击选项
type ClickOptions struct {
	WaitVisible bool          // 等待元素可见
	WaitEnabled bool          // 等待元素可用
	Timeout     time.Duration // 超时时间
	Button      string        // 鼠标按钮：left, right, middle
	ClickCount  int           // 点击次数
}

// TypeOptions 输入选项
type TypeOptions struct {
	Clear       bool          // 是否先清空
	WaitVisible bool          // 等待元素可见
	Timeout     time.Duration // 超时时间
	Delay       time.Duration // 每个字符之间的延迟
}

// SelectOptions 选择选项
type SelectOptions struct {
	WaitVisible bool          // 等待元素可见
	Timeout     time.Duration // 超时时间
}

// WaitForOptions 等待选项
type WaitForOptions struct {
	Timeout time.Duration // 超时时间
	State   string        // 等待状态：visible, hidden, attached, detached
}

// ScreenshotOptions 截图选项
type ScreenshotOptions struct {
	FullPage bool   // 是否截取完整页面
	Quality  int    // 质量 (0-100)
	Format   string // 格式：png, jpeg
}

// ExtractOptions 提取选项
type ExtractOptions struct {
	Selector string   // CSS 选择器
	Type     string   // 提取类型：text, html, attribute, property
	Attr     string   // 属性名（type=attribute 时使用）
	Multiple bool     // 是否提取多个元素
	Fields   []string // 要提取的字段列表
}

