package composite

import (
	"fmt"
	"strings"
)

// 组合模式 - 组件接口与叶子节点

// Node 是文件和文件夹的通用接口 (Component)
type Node interface {
	Search(keyword string)
	Name() string
}

// File 代表叶子节点 (Leaf)
type File struct {
	name string
}

// NewFile 创建一个文件实例
func NewFile(name string) *File {
	return &File{name: name}
}

// Name 返回文件名
func (f *File) Name() string {
	return f.name
}

// Search 在文件中搜索关键字
// 对于文件，只检查它自己的名字
func (f *File) Search(keyword string) {
	fmt.Printf("在文件 '%s' 中搜索关键字...\n", f.Name())
	if strings.Contains(f.Name(), keyword) {
		fmt.Printf(">>> 在文件 '%s' 中找到关键字 '%s'\n", f.Name(), keyword)
	}
}
