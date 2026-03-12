package config

// Config 应用配置
type Config struct {
	Server          ServerConfig    `json:"server"`
	Database        DatabaseConfig  `json:"database"`
	RecommendConfig RecommendConfig `json:"recommend"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Charset  string `json:"charset"`
}

type RecommendConfig struct {
	CollaborativeCount int32 `json:"collaborativeCount"`
}

// GetDefaultConfig 获取默认配置
func GetDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "127.0.0.1",
			Port: 8888,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "123456",
			DBName:   "advert_recommend",
			Charset:  "utf8mb4",
		},
		RecommendConfig: RecommendConfig{
			CollaborativeCount: 5,
		},
	}
}

var Global *Config
