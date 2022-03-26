package domain

type Config struct {
	Server struct {
		Enabled bool   `yaml:"enabled", envconfig:"SERVER_ENABLED"`
		Port    string `yaml:"port", envconfig:"SERVER_PORT"`
		Host    string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Client struct {
		Enabled bool `yaml:"enabled", envconfig:"CLIENT_ENABLED"`
		Dos     struct {
			Url          string `yaml:"url", envconfig:"CLIENT_DOS_URL"`
			Method       string `yaml:"method", envconfig:"CLIENT_DOS_METHOD"`
			RequestCount int    `yaml:"requestCount", envconfig:"CLIENT_DOS_REQUEST_COUNT"`
		} `yaml:"dos"`
		MaxIdleConns        string `yaml:"maxIdleConns", envconfig:"CLIENT_MAX_ID_CONNS"`
		MaxConnsPerHost     string `yaml:"maxConnsPerHost", envconfig:"CLIENT_MAX_CONNS_PER_HOST"`
		MaxIdleConnsPerHost string `yaml:"maxIdleConnsPerHost", envconfig:"CLIENT_MAX_IDLE_CONNS_PER_HOST"`
	} `yaml:"client"`
}
