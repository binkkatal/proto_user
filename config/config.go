package config

// DBConfig for database
type DBConfig struct {
	Username        string `yaml:"user"`
	Password        string `yaml:"pass"`
	DBName          string `yaml:"dbname"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int64  `yaml:"connMaxLifetime"`
}

