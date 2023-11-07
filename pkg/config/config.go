package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

// Conf 全局配置变量
var Conf = new(config)

type config struct {
	System    *SystemConfig    `mapstructure:"system" json:"system"`
	Logs      *LogsConfig      `mapstructure:"logs" json:"logs"`
	Mysql     *MysqlConfig     `mapstructure:"mysql" json:"mysql"`
	Redis     *RedisConfig     `mapstructure:"redis" json:"redis"`
	Casbin    *CasbinConfig    `mapstructure:"casbin" json:"casbin"`
	Jwt       *JwtConfig       `mapstructure:"jwt" json:"jwt"`
	RateLimit *RateLimitConfig `mapstructure:"ratelimit" json:"rateLimit"`
	SnowFlake *SnowFlakeConfig `mapstructure:"snowFlake" json:"snowFlake"`
}

type SystemConfig struct {
	Mode            string `mapstructure:"mode" json:"mode"`
	Host            string `mapstructure:"host" json:"host"`
	Port            int    `mapstructure:"port" json:"port"`
	RSAPublicKey    string `mapstructure:"rsa-public-key" json:"rsaPublicKey"`
	RSAPrivateKey   string `mapstructure:"rsa-private-key" json:"rsaPrivateKey"`
	RSAPublicBytes  []byte `mapstructure:"-" json:"-"`
	RSAPrivateBytes []byte `mapstructure:"-" json:"-"`
	I18nLanguage    string `mapstructure:"i18n-language" json:"i18nLanguage"`
}

type LogsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type MysqlConfig struct {
	Username  string `mapstructure:"username" json:"username"`
	Password  string `mapstructure:"password" json:"password"`
	Database  string `mapstructure:"database" json:"database"`
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Query     string `mapstructure:"query" json:"query"`
	LogMode   bool   `mapstructure:"log-mode" json:"logMode"`
	Charset   string `mapstructure:"charset" json:"charset"`
	Collation string `mapstructure:"collation" json:"collation"`
}

type RedisConfig struct {
	Password string `mapstructure:"password" json:"password"`
	Database int    `mapstructure:"database" json:"database"`
	Addr     string `mapstructure:"addr" json:"addr"`
}

type CasbinConfig struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type JwtConfig struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type RateLimitConfig struct {
	FillInterval int64 `mapstructure:"fill-interval" json:"fillInterval"`
	Capacity     int64 `mapstructure:"capacity" json:"capacity"`
	Quantum      int64 `mapstructure:"quantum" json:"quantum"`
}

type SnowFlakeConfig struct {
	WorkerID     int64 `mapstructure:"workerId" json:"workerId"`
	DatacenterID int64 `mapstructure:"datacenterId" json:"datacenterId"`
}

// InitConfig 加载配置文件
func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println("读取应用目录失败:", err)
		os.Exit(1)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	// 读取配置信息
	err = viper.ReadInConfig()

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("初始化配置文件失败:", err)
			os.Exit(1)
		}
		fmt.Println("配置文件已修改")
		// 读取rsa key
		Conf.System.RSAPublicBytes = RSAReadKeyFromFile(Conf.System.RSAPublicKey)
		Conf.System.RSAPrivateBytes = RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
	})

	if err != nil {
		fmt.Println("读取配置文件失败:", err)
		os.Exit(1)
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("初始化配置文件失败:", err)
		os.Exit(1)
	}
	// 读取rsa key
	Conf.System.RSAPublicBytes = RSAReadKeyFromFile(Conf.System.RSAPublicKey)
	Conf.System.RSAPrivateBytes = RSAReadKeyFromFile(Conf.System.RSAPrivateKey)
}

// 从文件中读取RSA key
func RSAReadKeyFromFile(filename string) (b []byte) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b = make([]byte, fileInfo.Size())
	f.Read(b)
	return b
}
