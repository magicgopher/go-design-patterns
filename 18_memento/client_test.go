package memento

import (
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestMemento 测试备忘录模式的场景
func TestMemento(t *testing.T) {
	editor := &Editor{}
	history := &History{}

	tests := []struct {
		name         string
		content      string
		restoreIndex int
		expected     string
	}{
		{
			name:         "保存并恢复第一个状态",
			content:      "第一个内容",
			restoreIndex: 0,
			expected:     "第一个内容",
		},
		{
			name:         "保存并恢复第二个状态",
			content:      "第二个内容",
			restoreIndex: 1,
			expected:     "第二个内容",
		},
		{
			name:         "恢复到第一个状态",
			content:      "第三个内容",
			restoreIndex: 0,
			expected:     "第一个内容",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置内容并保存状态
			editor.SetContent(tt.content)
			history.AddMemento(editor.CreateMemento())

			// 恢复到指定状态
			memento := history.GetMemento(tt.restoreIndex)
			if memento == nil {
				t.Fatalf("无法获取索引 %d 处的备忘录", tt.restoreIndex)
			}
			editor.Restore(memento)

			// 验证恢复结果
			result := editor.GetContent()
			if result != tt.expected {
				t.Errorf("期望内容 %q，实际得到 %q", tt.expected, result)
			}
			t.Logf("恢复的内容: %s", result)
		})
	}
}
