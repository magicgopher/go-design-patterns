package proxy

import "fmt"

// 代理模式

// Image 抽象主题

// Image 定义图像的行为。
type Image interface {
	// Display 显示图像。
	Display() string
}

// RealImage 真实主题

// RealImage 真实图像，模拟加载和显示大图像文件。
type RealImage struct {
	filename string
	loaded   bool
}

// NewRealImage 创建一个新的真实图像。
func NewRealImage(filename string) *RealImage {
	return &RealImage{
		filename: filename,
		loaded:   false,
	}
}

// load 模拟从磁盘加载图像（耗时操作）。
func (r *RealImage) load() {
	if !r.loaded {
		fmt.Printf("加载图像: %s\n", r.filename)
		r.loaded = true
	}
}

// Display 显示图像。
func (r *RealImage) Display() string {
	r.load() // 在显示前确保图像已加载
	return fmt.Sprintf("显示图像: %s", r.filename)
}
