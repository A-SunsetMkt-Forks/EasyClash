package config

var DefaultBase = BaseConf{
	Port:               7890,
	SocksPort:          7891,
	RedirPort:          7892,
	AllowLan:           false,
	ExternalController: "127.0.0.1:9090",
}
