package flyweight

// 享元模式

// 享元接口

// Character 是享元接口，定义字符显示行为
type Character interface {
	Display(fontSize int) string
}
