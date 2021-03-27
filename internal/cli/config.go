package cli

type Config struct {
	Listen struct {
		Ip   string
		Port int
	}
	Auth struct {
		Token string
	}
	Verbose int
}

func NewConfig() *Config {
	return &Config{
		Verbose: 1,
	}
}
