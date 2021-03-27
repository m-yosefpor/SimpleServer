package cli

var Cfg Config

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
