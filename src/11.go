package src

import (
	"fmt"
	"net/http"
	"time"
)

// Config represents a configuration.
type Config struct {
	Port int
}

// ConfigBuilder is a builder for Config.
type ConfigBuilder struct {
	port *int
}

// Port sets the port of the ConfigBuilder.
func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Build builds a Config from the ConfigBuilder.
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	defaultPort := 8080
	errorPort := 3110

	if b.port == nil {
		cfg.Port = defaultPort
	} else {
		cfg.Port = *b.port
	}

	if cfg.Port == 0 {
		cfg.Port = errorPort
	} else if cfg.Port < 1024 || cfg.Port > 65535 {
		return cfg, fmt.Errorf("invalid port number: %d", cfg.Port)
	}

	return cfg, nil
}

// NewServer creates a new server.
func NewServer(addr string, cfg Config) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, cfg.Port),
		Handler: nil,
	}
}

type options struct {
	port    *int
	timeout *int
}

// Option represents an option for a server.
type Option func(*options) error

// WithPort creates a new Option with the specified port.
func WithPort(port int) Option {
	return func(o *options) error {
		if port < 1024 || port > 65535 {
			return fmt.Errorf("invalid port number: %d", port)
		}

		o.port = &port
		return nil
	}
}

// WithTimeout creates a new Option with the specified timeout.
func WithTimeout(timeout int) Option {
	return func(o *options) error {
		if timeout < 0 {
			return fmt.Errorf("invalid timeout: %d", timeout)
		}

		o.timeout = &timeout
		return nil
	}
}

// NewServerWithOptions creates a new server with the specified options.
func NewServerWithOptions(addr string, opts ...Option) (*http.Server, error) {
	o := &options{}
	for _, opt := range opts {
		if err := opt(o); err != nil {
			return nil, err
		}
	}

	port := 8080
	if o.port != nil {
		port = *o.port
	}

	timeout := 10 * time.Second
	if o.timeout != nil {
		timeout = time.Duration(*o.timeout) * time.Second
	}

	return &http.Server{
		Addr:        fmt.Sprintf("%s:%d", addr, port),
		Handler:     nil,
		ReadTimeout: timeout,
	}, nil
}

func test() {

	cfg, err := (&ConfigBuilder{}).Port(2332).Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	server := NewServer("localhost", cfg)
	fmt.Println(server.Addr)

	server2, err := NewServerWithOptions("localhost", WithPort(2332), WithTimeout(5))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(server2.Addr, server2.ReadTimeout)
}
