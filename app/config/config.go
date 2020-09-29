package config

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}
