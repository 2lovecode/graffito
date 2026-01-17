package config

import (
	"fmt"
	"os"
)

// Config 应用配置
type Config struct {
	// Server配置
	Server ServerConfig
	
	// Sandbox配置
	Sandbox SandboxConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string
	Port string
}

// SandboxConfig 沙箱配置
type SandboxConfig struct {
	Timeout int // 超时时间（秒）
}

var globalConfig *Config

// Load 加载配置
func Load() (*Config, error) {
	if globalConfig != nil {
		return globalConfig, nil
	}

	config := &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Sandbox: SandboxConfig{
			Timeout: getEnvInt("SANDBOX_TIMEOUT", 30),
		},
	}

	globalConfig = config
	return config, nil
}

// Get 获取全局配置
func Get() *Config {
	if globalConfig == nil {
		config, err := Load()
		if err != nil {
			// 返回默认配置
			return &Config{
				Server: ServerConfig{
					Host: "0.0.0.0",
					Port: "8080",
				},
				Sandbox: SandboxConfig{
					Timeout: 30,
				},
			}
		}
		return config
	}
	return globalConfig
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt 获取环境变量（整数类型），如果不存在则返回默认值
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var result int
		if _, err := fmt.Sscanf(value, "%d", &result); err == nil {
			return result
		}
	}
	return defaultValue
}
