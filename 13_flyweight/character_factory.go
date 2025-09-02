package flyweight

// 享元模式

// 享元工厂

// CharacterFactory 是享元工厂，管理共享的字符对象
type CharacterFactory struct {
	characters map[string]Character
}

// NewCharacterFactory 创建享元工厂实例
func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		characters: make(map[string]Character),
	}
}

// GetCharacter 获取或创建共享的字符对象
func (f *CharacterFactory) GetCharacter(symbol string) Character {
	if char, exists := f.characters[symbol]; exists {
		return char
	}
	char := NewConcreteCharacter(symbol)
	f.characters[symbol] = char
	return char
}
