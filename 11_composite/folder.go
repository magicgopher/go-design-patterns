package composite

import "fmt"

// 组合模式 - 组合节点

// Folder 代表组合节点 (Composite)
type Folder struct {
	name     string
	children []Node // 可以包含其他Node（文件或文件夹）
}

// NewFolder 创建一个文件夹实例
func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		children: make([]Node, 0),
	}
}

// Name 返回文件夹名
func (f *Folder) Name() string {
	return f.name
}

// Add 向文件夹中添加一个节点（文件或子文件夹）
func (f *Folder) Add(node Node) {
	f.children = append(f.children, node)
}

// Remove 从文件夹中移除一个节点
func (f *Folder) Remove(node Node) {
	for i, child := range f.children {
		if child.Name() == node.Name() {
			f.children = append(f.children[:i], f.children[i+1:]...)
			return
		}
	}
}

// Search 在文件夹中搜索关键字
// 对于文件夹，它会递归地在其所有子节点上执行搜索操作
func (f *Folder) Search(keyword string) {
	fmt.Printf("在文件夹 '%s' 中递归搜索关键字...\n", f.Name())
	// 组合模式的核心：遍历所有子节点，并调用它们自己的Search方法
	for _, child := range f.children {
		child.Search(keyword)
	}
}
