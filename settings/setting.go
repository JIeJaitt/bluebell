package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf ... 全局变量，用来保存程序的所有配置信息
var Conf = new(Config)

// Config ...
type Config struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig      `mapstructure:"log"`
	*MySQLConfig    `mapstructure:"mysql"`
	*RedisConfig    `mapstructure:"redis"`
	*PostListConfig `mapstructure:"postlist"`
}

// LogConfig ...
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MySQLConfig ...
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig ...
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// PostListConfig ...
type PostListConfig struct {
	Size int64 `mapstructure:"size"`
}

// Init ...
// 使用viper库管理配置
func Init() error {
	// 设置配置文件
	viper.SetConfigFile("./conf/config.json")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil { //读取配置文件失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil { //反序列化失败
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}

	viper.WatchConfig() // 监控配置文件的修改
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed：%v\n", e.Name)
		if err := viper.Unmarshal(Conf); err != nil { // 把变化后的配置文件重新反序列化到Conf变量中
			fmt.Printf("viper.Unmarshal failed again, err:%v\n", err)
			return
		}
	})

	return nil
}
