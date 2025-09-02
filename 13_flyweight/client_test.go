package flyweight

import (
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestFlyweight 测试享元模式的场景
func TestFlyweight(t *testing.T) {
	factory := NewCharacterFactory()

	tests := []struct {
		name         string
		symbol       string
		fontSize     int
		expectedDesc string
	}{
		{
			name:         "Character A with size 12",
			symbol:       "A",
			fontSize:     12,
			expectedDesc: "字符: A, 字体大小: 12",
		},
		{
			name:         "Character A with size 16",
			symbol:       "A",
			fontSize:     16,
			expectedDesc: "字符: A, 字体大小: 16",
		},
		{
			name:         "Character B with size 14",
			symbol:       "B",
			fontSize:     14,
			expectedDesc: "字符: B, 字体大小: 14",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char := factory.GetCharacter(tt.symbol)
			result := char.Display(tt.fontSize)
			if result != tt.expectedDesc {
				t.Errorf("expected desc %q, got %q", tt.expectedDesc, result)
			}
			t.Logf("Rendered: %s", result)
		})
	}

	// 验证共享性：相同字符的享元对象应为同一个实例
	char1 := factory.GetCharacter("A")
	char2 := factory.GetCharacter("A")
	if char1 != char2 {
		t.Errorf("expected same Character instance for symbol 'A'")
	}
}
