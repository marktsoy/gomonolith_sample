package apiserver

// Config TODO
type Config struct {
	DatabaseURL string `toml:"store"`
}

//NewConfig TODO
func NewConfig() *Config {
	return &Config{
		DatabaseURL: "",
	}
}
