package flyweight

import "fmt"

// 享元模式

// 具体享元

// ConcreteCharacter 是具体享元，存储字符的内在状态
type ConcreteCharacter struct {
	symbol string // 内在状态：字符本身
}

// NewConcreteCharacter 创建具体享元实例
func NewConcreteCharacter(symbol string) *ConcreteCharacter {
	return &ConcreteCharacter{symbol: symbol}
}

// Display 显示字符，接受外在状态（字体大小）
func (c *ConcreteCharacter) Display(fontSize int) string {
	return fmt.Sprintf("字符: %s, 字体大小: %d", c.symbol, fontSize)
}
