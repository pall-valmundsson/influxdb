package httpd

type Config struct {
	Enabled      bool   `toml:"enabled"`
	BindAddress  string `toml:"bind-address"`
	AuthEnabled  bool   `toml:"auth-enabled"`
	LogEnabled   bool   `toml:"log-enabled"`
	WriteTracing bool   `toml:"write-tracing"`
	PprofEnabled bool   `toml:"pprof-enabled"`
}

func NewConfig() Config {
	return Config{
		Enabled:    true,
		LogEnabled: true,
	}
}
