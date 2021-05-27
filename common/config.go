package common

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	ProxyContractAddress string //代理合约地址
	StubContractAddress  string //桥接合约地址
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	err := configor.Load(&config, path)
	if err != nil {
		return nil, err
	}
	return &config, err
}
