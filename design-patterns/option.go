package main

import "fmt"

// Server 是我们要配置的对象
type Server struct {
	Host string
	Port int
}

// Option 是一个函数类型，用于配置 Server
type Option func(*Server)

// NewServer 是 Server 的构造函数，接收可变数量的 Option
func NewServer(opts ...Option) *Server {
	// 创建一个 Server 对象并设置默认值
	s := &Server{
		Host: "localhost",
		Port: 8080,
	}

	// 应用每个 Option 函数来配置 Server
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithHost 返回一个 Option，用于设置 Server 的 Host
func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

// WithPort 返回一个 Option，用于设置 Server 的 Port
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func main() {
	// 使用 Option 模式创建和配置 Server 对象
	server := NewServer(
		WithHost("example.com"),
		WithPort(9090),
	)

	fmt.Printf("Server is running on %s:%d\n", server.Host, server.Port)
}
