package conf

// Config 应用配置
type Config struct {
	System    `mapstructure:"system"`
	LogConfig `mapstructure:"log"`
	Database  `mapstructure:"database"`
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

type Database struct {
	Source      string `mapstructure:"source"`
	LogMode     bool   `mapstructure:"log_mode"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
}
