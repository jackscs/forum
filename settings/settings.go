package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量 用来保存程序的所有配置信息
var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name    string
	Mode    string
	Port    int
	Version string
}

type LogConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type MysqlConfig struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisConfig struct {
	Host     string
	Port     int
	Db       int
	Password string
	PoolSize int
}

func Init() (err error) {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()    // 查找并读取配置文件
	if err != nil {
		// 处理读取配置文件的错误
		fmt.Println("viper.ReadInConfig() failed ,err =", err)
		return
	}

	//把读取到的配置信息反序列到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}
	fmt.Println(Conf)
	fmt.Printf("%+v\n", Conf.AppConfig.Version)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
