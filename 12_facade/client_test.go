package facade

import "testing"

// 单元测试
// 模拟客户端调用

// TestFacade 测试外观模式的场景
func TestFacade(t *testing.T) {
	facade := NewHomeTheaterFacade()

	tests := []struct {
		name         string
		action       func() string
		expectedDesc string
	}{
		{
			name:         "WatchMovie",
			action:       func() string { return facade.WatchMovie("Inception") },
			expectedDesc: "开始观看电影...\n",
		},
		{
			name:         "EndMovie",
			action:       func() string { return facade.EndMovie() },
			expectedDesc: "结束电影播放...\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action()
			if result != tt.expectedDesc {
				t.Errorf("expected desc %q, got %q", tt.expectedDesc, result)
			}
			t.Logf("Action result: %s", result)
		})
	}
}
