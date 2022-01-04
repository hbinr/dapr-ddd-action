package conf

// Config 应用配置
type Config struct {
	System    `mapstructure:"system"`
	LogConfig `mapstructure:"log"`
}

type System struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

// LogConfig zap log配置
type LogConfig struct {
	Prefix     string `mapstructure:"prefix"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
