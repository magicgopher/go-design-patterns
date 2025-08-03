package proxy

import (
	"io"
	"os"
	"strings"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestImageProxy 测试图像代理的延迟加载和显示功能。
func TestImageProxy(t *testing.T) {
	// 定义测试用例，包含图像文件名和期望的显示结果。
	tests := []struct {
		name            string
		filename        string
		expectedDisplay string
	}{
		{
			name:            "TestImage1",
			filename:        "image1.jpg",
			expectedDisplay: "显示图像: image1.jpg",
		},
		{
			name:            "TestImage2",
			filename:        "image2.png",
			expectedDisplay: "显示图像: image2.png",
		},
	}

	// 遍历测试用例，验证代理的显示功能。
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建代理对象。
			proxy := NewImageProxy(tt.filename)

			// 第一次调用 Display，触发加载。
			output := proxy.Display()
			if output != tt.expectedDisplay {
				t.Errorf("期望显示 '%s'，但得到 '%s'", tt.expectedDisplay, output)
			}

			// 第二次调用 Display，验证不重复加载。
			output = proxy.Display()
			if output != tt.expectedDisplay {
				t.Errorf("期望显示 '%s'，但得到 '%s'", tt.expectedDisplay, output)
			}

			// 日志输出，供调试。
			t.Logf("成功显示: %s", output)
		})
	}
}

// TestLazyLoading 测试代理的延迟加载功能。
func TestLazyLoading(t *testing.T) {
	// 创建代理对象。
	proxy := NewImageProxy("test.jpg")

	// 检查 realImage 是否未加载。
	if proxy.realImage != nil {
		t.Errorf("期望 realImage 未加载，但已加载")
	}

	// 调用 Display，触发加载。
	proxy.Display()

	// 检查 realImage 是否已加载。
	if proxy.realImage == nil {
		t.Errorf("期望 realImage 已加载，但仍为 nil")
	}
	if !proxy.realImage.loaded {
		t.Errorf("期望 realImage.loaded 为 true，但得到 false")
	}
}

// TestMultipleDisplays 测试多次调用 Display 是否只加载一次。
func TestMultipleDisplays(t *testing.T) {
	// 创建管道以捕获标准输出。
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("创建管道失败: %v", err)
	}
	// 保存原始 Stdout 并重定向。
	origStdout := os.Stdout
	os.Stdout = w
	defer func() {
		// 恢复原始 Stdout。
		os.Stdout = origStdout
		w.Close()
	}()

	// 创建代理对象。
	proxy := NewImageProxy("multi.jpg")

	// 多次调用 Display。
	proxy.Display()
	proxy.Display()

	// 关闭写入端并读取输出。
	w.Close()
	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("读取管道输出失败: %v", err)
	}

	// 检查加载日志只出现一次。
	loadCount := strings.Count(string(output), "加载图像: multi.jpg")
	if loadCount != 1 {
		t.Errorf("期望加载图像 1 次，但得到 %d 次", loadCount)
	}
}
