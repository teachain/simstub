package commons

import (
	"github.com/jinzhu/configor"
)

//ProxyContractAddress string //代理合约地址
//StubContractAddress  string //桥接合约地址
type Config struct {
	Common       CommonConfig   `yaml:"common"`
	Node         NodeConfig     `yaml:"node"`
	Account      *AccountConfig `yaml:"account"`
	ProxyAddress string         `yaml:"proxy_addr"`
	DBConfig     *DBConfig      `yaml:"db"`
	DefaultChain *DefaultChain   `yaml:"default_chain"`
}
type DBConfig struct {
	Username string `default:"root" yaml:"username"`
	Password string `default:"" yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `default:"3306" yaml:"port"`
	DbName   string `required:"true" yaml:"db_name"`
	Charset  string `default:"utf8" yaml:"charset"`
	MaxIdle  int    `default:"1000" yaml:"max_idle"`
	MaxOpen  int    `default:"2000" yaml:"max_open"`
	LogMode  bool   `yaml:"log_mode"`
	Loc      string `required:"true" yaml:"loc"`
}

type CommonConfig struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Zone     string `yaml:"zone"`
	Chain    string `yaml:"chain"`
	Resource string `yaml:"resource"` //要跟anchor合约中的一致
}

type NodeConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
}
type AccountConfig struct {
	Keystore string `yaml:"keystore"`
	Password string `yaml:"password"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	err := configor.Load(&config, path)
	if err != nil {
		return nil, err
	}
	return &config, err
}

type DefaultChain struct {
	Path           string `yaml:"path"`
	Method         string `yaml:"method"`
	CallbackPath   string `yaml:"callback_path"`
	CallbackMethod string `yaml:"callback_method"`
	CrossChainType string `yaml:"cross_chain_type"`
}
