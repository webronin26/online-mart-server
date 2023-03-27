package config

type DatabaseConfig struct {
	Username     string
	Password     string
	Name         string
	Host         string
	Port         uint64
	Encoding     string
	Maxconns     uint64
	Maxidleconns uint64
	Timeout      uint64
	Logmode      bool
}
