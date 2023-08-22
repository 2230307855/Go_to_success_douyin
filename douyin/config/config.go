package config

import (
	"github.com/spf13/viper"
)

// 获取配置
// 只需要文件名 Config前面的就行，后面的就行拼接，不用拓展名
func GetConfig(filename string) (*viper.Viper, error) {

	v := viper.New()

	// 设置配置文件名（不含扩展名）
	v.SetConfigName(filename + "Config")
	// 设置配置文件路径（可根据实际路径进行调整）
	v.AddConfigPath("config")

	// 设置配置文件类型为 YAML
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil

}
