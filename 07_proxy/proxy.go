package proxy

// 代理模式

// 代理

// ImageProxy 图像代理，控制对真实图像的访问。
type ImageProxy struct {
	filename  string
	realImage *RealImage
}

// NewImageProxy 创建一个新的图像代理。
func NewImageProxy(filename string) *ImageProxy {
	return &ImageProxy{
		filename:  filename,
		realImage: nil, // 延迟初始化
	}
}

// Display 显示图像，通过代理控制加载。
func (p *ImageProxy) Display() string {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename) // 延迟加载
	}
	return p.realImage.Display()
}
