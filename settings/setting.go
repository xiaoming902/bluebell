package settings

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string        `mapstructure:"name"`
	Mode         string        `mapstructure:"mode"`
	Version      string        `mapstructure:"version"`
	StartTime    string        `mapstructure:"start_time"`
	MachineId    int64         `mapstructure:"machine_id"`
	Add          string        `mapstructure:"add"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*LogConfig   `mapstructure:"log"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

func Init() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()          // 读取配置信息
	if err != nil {                      // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
}
