package executor

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// ExtractSemanticTree 从页面提取语义树
func ExtractSemanticTree(ctx context.Context, page *rod.Page) (*SemanticTree, error) {
	tree := &SemanticTree{
		Elements: make(map[string]*SemanticNode),
	}

	// 提取所有可交互元素
	elements, err := extractInteractiveElements(ctx, page)
	if err != nil {
		return nil, fmt.Errorf("failed to extract interactive elements: %w", err)
	}

	// 构建语义节点
	for _, elem := range elements {
		node, err := buildSemanticNode(ctx, elem)
		if err != nil {
			continue // 跳过无法处理的元素
		}
		if node != nil {
			tree.Elements[node.ID] = node
		}
	}

	// 构建树形结构（简化版，将所有节点作为根节点的子节点）
	tree.Root = &SemanticNode{
		ID:       "root",
		Type:     "root",
		Children: make([]*SemanticNode, 0, len(tree.Elements)),
	}
	for _, node := range tree.Elements {
		tree.Root.Children = append(tree.Root.Children, node)
	}

	return tree, nil
}

// extractInteractiveElements 提取所有可交互元素
func extractInteractiveElements(ctx context.Context, page *rod.Page) (rod.Elements, error) {
	// 定义可交互元素的选择器
	selectors := []string{
		// 表单元素
		"input:not([type='hidden'])",
		"textarea",
		"select",
		"button",
		
		// 链接
		"a[href]",
		
		// 可点击元素
		"[role='button']",
		"[role='link']",
		"[role='menuitem']",
		"[role='tab']",
		"[role='checkbox']",
		"[role='radio']",
		"[onclick]",
		
		// 其他交互元素
		"[contenteditable='true']",
	}

	allElements := rod.Elements{}
	for _, selector := range selectors {
		elements, err := page.Elements(selector)
		if err != nil {
			continue
		}
		allElements = append(allElements, elements...)
	}

	return allElements, nil
}

// buildSemanticNode 构建语义节点
func buildSemanticNode(ctx context.Context, elem *rod.Element) (*SemanticNode, error) {
	// 获取元素基本信息
	tagName, err := elem.Eval("() => this.tagName.toLowerCase()")
	if err != nil {
		return nil, err
	}

	node := &SemanticNode{
		Type:       tagName.Value.Str(),
		Attributes: make(map[string]string),
		Metadata:   make(map[string]interface{}),
	}

	// 获取常用属性
	attributes := []string{"id", "name", "class", "type", "role", "aria-label", "placeholder", "value", "href", "title"}
	for _, attr := range attributes {
		if val, err := elem.Attribute(attr); err == nil && val != nil {
			node.Attributes[attr] = *val
		}
	}

	// 设置节点属性
	if id, ok := node.Attributes["id"]; ok && id != "" {
		node.ID = id
	} else {
		// 生成唯一 ID
		node.ID = generateElementID(elem)
	}

	// 设置类型
	if elemType, ok := node.Attributes["type"]; ok {
		node.Type = elemType
	}

	// 设置 Role
	if role, ok := node.Attributes["role"]; ok {
		node.Role = role
	}

	// 设置标签
	if ariaLabel, ok := node.Attributes["aria-label"]; ok {
		node.Label = ariaLabel
	} else if title, ok := node.Attributes["title"]; ok {
		node.Label = title
	} else if name, ok := node.Attributes["name"]; ok {
		node.Label = name
	}

	// 设置 Placeholder
	if placeholder, ok := node.Attributes["placeholder"]; ok {
		node.Placeholder = placeholder
	}

	// 设置值
	if value, ok := node.Attributes["value"]; ok {
		node.Value = value
	}

	// 获取元素文本
	text, err := elem.Text()
	if err == nil && text != "" {
		node.Text = strings.TrimSpace(text)
		if node.Label == "" {
			node.Label = node.Text
		}
	}

	// 如果仍然没有标签，使用元素内的文本内容（限制长度）
	if node.Label == "" {
		innerText, err := elem.Eval("() => this.innerText || this.textContent")
		if err == nil && innerText.Value.Str() != "" {
			text := strings.TrimSpace(innerText.Value.Str())
			if len(text) > 50 {
				text = text[:50] + "..."
			}
			node.Label = text
		}
	}

	// 获取选择器
	node.Selector = buildCSSSelector(node)
	node.XPath = buildXPath(node)

	// 获取元素位置
	shape, err := elem.Shape()
	if err == nil && shape != nil && len(shape.Quads) > 0 {
		// 使用第一个 quad 来计算位置
		quad := shape.Quads[0]
		minX, maxX := quad[0], quad[0]
		minY, maxY := quad[1], quad[1]
		for i := 0; i < len(quad); i += 2 {
			if quad[i] < minX {
				minX = quad[i]
			}
			if quad[i] > maxX {
				maxX = quad[i]
			}
			if quad[i+1] < minY {
				minY = quad[i+1]
			}
			if quad[i+1] > maxY {
				maxY = quad[i+1]
			}
		}
		node.Position = &ElementPosition{
			X:      minX,
			Y:      minY,
			Width:  maxX - minX,
			Height: maxY - minY,
		}
	}

	// 检查可见性和可用性
	visible, err := elem.Visible()
	if err == nil {
		node.IsVisible = visible
	}

	// 检查是否启用
	disabled, err := elem.Property("disabled")
	if err == nil {
		isDisabled := disabled.Bool()
		node.IsEnabled = !isDisabled
	} else {
		node.IsEnabled = true
	}

	return node, nil
}

// generateElementID 生成元素唯一 ID
func generateElementID(elem *rod.Element) string {
	// 使用元素的 describe 信息生成 ID
	desc, err := elem.Describe(1, false)
	if err != nil {
		return fmt.Sprintf("elem_%s", elem.Object.ObjectID)
	}
	
	data := fmt.Sprintf("%d_%s", desc.NodeID, desc.LocalName)
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("elem_%x", hash[:8])
}

// buildCSSSelector 构建 CSS 选择器
func buildCSSSelector(node *SemanticNode) string {
	// 优先使用 ID
	if id, ok := node.Attributes["id"]; ok && id != "" {
		return "#" + id
	}

	// 使用 name 属性
	if name, ok := node.Attributes["name"]; ok && name != "" {
		return fmt.Sprintf("%s[name='%s']", node.Type, name)
	}

	// 使用类型和其他属性组合
	selector := node.Type
	if class, ok := node.Attributes["class"]; ok && class != "" {
		// 只使用第一个类名
		classes := strings.Split(class, " ")
		if len(classes) > 0 && classes[0] != "" {
			selector += "." + classes[0]
		}
	}

	return selector
}

// buildXPath 构建 XPath
func buildXPath(node *SemanticNode) string {
	// 如果有 ID，使用 ID
	if id, ok := node.Attributes["id"]; ok && id != "" {
		return fmt.Sprintf("//%s[@id='%s']", node.Type, id)
	}

	// 如果有 name，使用 name
	if name, ok := node.Attributes["name"]; ok && name != "" {
		return fmt.Sprintf("//%s[@name='%s']", node.Type, name)
	}

	// 使用文本内容
	if node.Text != "" {
		return fmt.Sprintf("//%s[contains(text(), '%s')]", node.Type, node.Text)
	}

	return fmt.Sprintf("//%s", node.Type)
}

// FindElementByLabel 通过标签查找元素
func (tree *SemanticTree) FindElementByLabel(label string) *SemanticNode {
	label = strings.ToLower(strings.TrimSpace(label))
	
	for _, node := range tree.Elements {
		if strings.Contains(strings.ToLower(node.Label), label) {
			return node
		}
		if strings.Contains(strings.ToLower(node.Text), label) {
			return node
		}
		if strings.Contains(strings.ToLower(node.Placeholder), label) {
			return node
		}
	}
	
	return nil
}

// FindElementByType 通过类型查找元素
func (tree *SemanticTree) FindElementsByType(elemType string) []*SemanticNode {
	result := make([]*SemanticNode, 0)
	
	for _, node := range tree.Elements {
		if node.Type == elemType {
			result = append(result, node)
		}
	}
	
	return result
}

// FindElementByID 通过 ID 查找元素
func (tree *SemanticTree) FindElementByID(id string) *SemanticNode {
	return tree.Elements[id]
}

// GetVisibleElements 获取所有可见元素
func (tree *SemanticTree) GetVisibleElements() []*SemanticNode {
	result := make([]*SemanticNode, 0)
	
	for _, node := range tree.Elements {
		if node.IsVisible {
			result = append(result, node)
		}
	}
	
	return result
}

// GetClickableElements 获取所有可点击元素
func (tree *SemanticTree) GetClickableElements() []*SemanticNode {
	result := make([]*SemanticNode, 0)
	clickableTypes := map[string]bool{
		"button": true,
		"a":      true,
		"submit": true,
	}
	
	for _, node := range tree.Elements {
		if node.IsVisible && node.IsEnabled {
			if clickableTypes[node.Type] || node.Role == "button" || node.Role == "link" {
				result = append(result, node)
			}
		}
	}
	
	return result
}

// GetInputElements 获取所有输入元素
func (tree *SemanticTree) GetInputElements() []*SemanticNode {
	result := make([]*SemanticNode, 0)
	inputTypes := map[string]bool{
		"text":     true,
		"email":    true,
		"password": true,
		"search":   true,
		"tel":      true,
		"url":      true,
		"number":   true,
		"textarea": true,
	}
	
	for _, node := range tree.Elements {
		if node.IsVisible && node.IsEnabled {
			if inputTypes[node.Type] {
				result = append(result, node)
			}
		}
	}
	
	return result
}

// SerializeToSimpleText 将语义树序列化为简单文本（用于 LLM）
func (tree *SemanticTree) SerializeToSimpleText() string {
	var builder strings.Builder
	builder.WriteString("Page Interactive Elements:\n\n")
	
	// 按类型分组
	clickable := tree.GetClickableElements()
	inputs := tree.GetInputElements()
	
	if len(clickable) > 0 {
		builder.WriteString("Clickable Elements:\n")
		for i, node := range clickable {
			// 生成更明确的标识
			label := node.Label
			if label == "" {
				label = node.Text
			}
			if label == "" && node.Attributes["id"] != "" {
				label = fmt.Sprintf("id:%s", node.Attributes["id"])
			}
			if label == "" && node.Attributes["name"] != "" {
				label = fmt.Sprintf("name:%s", node.Attributes["name"])
			}
			if label == "" {
				label = fmt.Sprintf("<%s>", node.Type)
			}
			
			builder.WriteString(fmt.Sprintf("  Clickable Element [%d]: %s", i+1, label))
			if node.Type != "" {
				builder.WriteString(fmt.Sprintf(" (type: %s)", node.Type))
			}
			if node.Text != "" && node.Text != node.Label {
				builder.WriteString(fmt.Sprintf(" - %s", node.Text))
			}
			builder.WriteString("\n")
		}
		builder.WriteString("\n")
	}
	
	if len(inputs) > 0 {
		builder.WriteString("Input Elements:\n")
		for i, node := range inputs {
			// 生成更明确的标识
			label := node.Label
			if label == "" {
				label = node.Placeholder
			}
			if label == "" && node.Attributes["id"] != "" {
				label = fmt.Sprintf("id:%s", node.Attributes["id"])
			}
			if label == "" && node.Attributes["name"] != "" {
				label = fmt.Sprintf("name:%s", node.Attributes["name"])
			}
			if label == "" {
				label = fmt.Sprintf("<%s>", node.Type)
			}
			
			builder.WriteString(fmt.Sprintf("  Input Element [%d]: %s", i+1, label))
			if node.Type != "" && node.Type != "text" {
				builder.WriteString(fmt.Sprintf(" (type: %s)", node.Type))
			}
			if node.Placeholder != "" && node.Placeholder != label {
				builder.WriteString(fmt.Sprintf(" [placeholder: %s]", node.Placeholder))
			}
			if node.Value != "" {
				builder.WriteString(fmt.Sprintf(" [value: %s]", node.Value))
			}
			builder.WriteString("\n")
		}
	}
	
	return builder.String()
}

// HighlightElement 在页面上高亮显示元素（用于调试）
func HighlightElement(ctx context.Context, page *rod.Page, selector string) error {
	elem, err := page.Element(selector)
	if err != nil {
		return err
	}
	
	// 添加高亮样式
	_, err = elem.Eval(`function() {
		this.style.outline = '3px solid red';
		this.style.outlineOffset = '2px';
		setTimeout(() => {
			this.style.outline = '';
			this.style.outlineOffset = '';
		}, 2000);
	}`)
	
	return err
}

// WaitForElement 等待元素出现
func WaitForElement(ctx context.Context, page *rod.Page, selector string, opts *WaitForOptions) error {
	if opts == nil {
		opts = &WaitForOptions{}
	}
	
	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 30 * 1000000000 // 30秒
	}
	
	// 等待元素
	elem, err := page.Timeout(timeout).Element(selector)
	if err != nil {
		return fmt.Errorf("element not found: %s", selector)
	}
	
	// 根据状态等待
	switch opts.State {
	case "visible":
		return elem.WaitVisible()
	case "hidden":
		return elem.WaitInvisible()
	default:
		// 默认只等待元素存在
		return nil
	}
}

// EvaluateAccessibility 评估页面可访问性
func EvaluateAccessibility(ctx context.Context, page *rod.Page) (*AccessibilityReport, error) {
	report := &AccessibilityReport{
		Issues: make([]AccessibilityIssue, 0),
	}
	
	// 检查没有 alt 属性的图片
	images, _ := page.Elements("img:not([alt])")
	for _, img := range images {
		src, _ := img.Attribute("src")
		report.Issues = append(report.Issues, AccessibilityIssue{
			Type:     "missing-alt",
			Severity: "warning",
			Message:  "Image missing alt attribute",
			Element:  fmt.Sprintf("<img src='%s'>", *src),
		})
	}
	
	// 检查没有 label 的输入框
	inputs, _ := page.Elements("input:not([type='hidden']):not([aria-label]):not([id])")
	for range inputs {
		report.Issues = append(report.Issues, AccessibilityIssue{
			Type:     "missing-label",
			Severity: "error",
			Message:  "Input field missing label or aria-label",
		})
	}
	
	report.TotalIssues = len(report.Issues)
	return report, nil
}

// AccessibilityReport 可访问性报告
type AccessibilityReport struct {
	TotalIssues int                   `json:"total_issues"`
	Issues      []AccessibilityIssue  `json:"issues"`
}

// AccessibilityIssue 可访问性问题
type AccessibilityIssue struct {
	Type     string `json:"type"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Element  string `json:"element,omitempty"`
}

// GetElementFromPage 从页面获取 Rod Element
func GetElementFromPage(ctx context.Context, page *rod.Page, node *SemanticNode) (*rod.Element, error) {
	// 优先使用 ID
	if id, ok := node.Attributes["id"]; ok && id != "" {
		elem, err := page.Element("#" + id)
		if err == nil {
			return elem, nil
		}
	}
	
	// 使用构建的选择器
	if node.Selector != "" {
		elem, err := page.Element(node.Selector)
		if err == nil {
			return elem, nil
		}
	}
	
	// 使用 XPath
	if node.XPath != "" {
		elem, err := page.ElementX(node.XPath)
		if err == nil {
			return elem, nil
		}
	}
	
	return nil, fmt.Errorf("element not found: %s", node.Label)
}

// InjectAccessibilityHelpers 注入辅助脚本
func InjectAccessibilityHelpers(ctx context.Context, page *rod.Page) error {
	// 注入辅助函数到页面
	_, err := page.Eval(`() => {
		// 添加用于元素高亮的辅助函数
		window.__highlightElement = function(selector) {
			const elem = document.querySelector(selector);
			if (elem) {
				elem.style.outline = '3px solid blue';
				elem.style.outlineOffset = '2px';
				setTimeout(() => {
					elem.style.outline = '';
					elem.style.outlineOffset = '';
				}, 1000);
			}
		};
		
		// 添加获取元素语义信息的辅助函数
		window.__getElementInfo = function(elem) {
			return {
				tag: elem.tagName.toLowerCase(),
				text: elem.innerText || elem.textContent,
				value: elem.value,
				visible: elem.offsetParent !== null,
				rect: elem.getBoundingClientRect()
			};
		};
	}`)
	
	return err
}

// ScrollToElement 滚动到元素位置
func ScrollToElement(ctx context.Context, elem *rod.Element) error {
	return elem.ScrollIntoView()
}

// GetElementScreenshot 获取元素截图
func GetElementScreenshot(ctx context.Context, elem *rod.Element) ([]byte, error) {
	return elem.Screenshot(proto.PageCaptureScreenshotFormatPng, 100)
}

