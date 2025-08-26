package composite

import (
	"fmt"
	"testing"
)

// 模拟客户端调用

func TestCompositePattern(t *testing.T) {
	// --- 1. 创建叶子节点（文件） ---
	file1 := NewFile("Go设计模式.pdf")
	file2 := NewFile("代码大全.pdf")
	file3 := NewFile("个人简历.docx")

	// --- 2. 创建组合节点（文件夹） ---
	folder1 := NewFolder("技术书籍")
	folder2 := NewFolder("个人文档")
	rootFolder := NewFolder("根目录")

	// --- 3. 构建树形结构 ---
	folder1.Add(file1)
	folder1.Add(file2)

	folder2.Add(file3)

	rootFolder.Add(folder1)
	rootFolder.Add(folder2)

	// --- 4. 客户端统一操作 ---
	// 客户端不需要区分操作的是文件还是文件夹，统一调用Search方法

	fmt.Println("--- 场景1: 在整个文件系统中搜索 '设计' ---")
	// 在根目录上执行搜索，会递归地在所有子节点上执行
	rootFolder.Search("设计")

	fmt.Println("\n--- 场景2: 在 '技术书籍' 文件夹中搜索 '代码' ---")
	// 也可以在子树上执行搜索
	folder1.Search("代码")

	fmt.Println("\n--- 场景3: 在单个文件上搜索 '简历' ---")
	// 也可以直接在叶子节点上执行操作
	file3.Search("简历")
}
