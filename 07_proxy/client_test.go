package proxy

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// 单元测试
// 模拟客户端调用

// TestProxySaleComputer 测试代理模式的 SaleComputer 方法
func TestProxySaleComputer(t *testing.T) {
	// 定义测试用例结构，包含测试名称、输入价格和预期输出
	tests := []struct {
		name     string  // 测试用例名称
		money    float64 // 输入价格
		expected string  // 预期输出
	}{
		{
			name:     "Test with price 10000", // 测试用例1：输入价格10000
			money:    10000.0,
			expected: "经销商代理销售电脑，原始价格: 10000.00，增加20%利润: 2000.00，总价: 12000.00\n销售电脑, 并拿到钱:10000.00",
		},
		{
			name:     "Test with price 5000", // 测试用例2：输入价格5000
			money:    5000.0,
			expected: "经销商代理销售电脑，原始价格: 5000.00，增加20%利润: 1000.00，总价: 6000.00\n销售电脑, 并拿到钱:5000.00",
		},
		{
			name:     "Test with price 0", // 测试用例3：输入价格0
			money:    0.0,
			expected: "经销商代理销售电脑，原始价格: 0.00，增加20%利润: 0.00，总价: 0.00\n销售电脑, 并拿到钱:0.00",
		},
	}

	// 创建代理实例
	proxy := NewProxy()

	// 遍历所有测试用例
	for _, tt := range tests {
		// 使用 t.Run 运行子测试，tt.name 为子测试名称
		t.Run(tt.name, func(t *testing.T) {
			// 创建管道以捕获标准输出
			r, w, _ := os.Pipe()
			// 保存原始标准输出
			old := os.Stdout
			// 将标准输出重定向到管道
			os.Stdout = w
			// 确保测试结束后恢复标准输出
			defer func() {
				os.Stdout = old
			}()

			// 调用代理的 SaleComputer 方法
			proxy.SaleComputer(tt.money)

			// 关闭管道写入端
			w.Close()
			// 创建缓冲区以读取管道输出
			var buf bytes.Buffer
			// 将管道内容复制到缓冲区
			io.Copy(&buf, r)
			// 获取实际输出
			output := buf.String()

			// 格式化预期输出
			expected := tt.expected
			// 比较实际输出与预期输出
			if output != expected {
				t.Errorf("SaleComputer(%.2f) output = %q; want %q", tt.money, output, expected)
			}
		})
	}
}
