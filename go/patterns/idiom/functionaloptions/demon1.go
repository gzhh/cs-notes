package functionaloptions

type (
	// Config holds the application's configuration.
	Config struct {
		// LogLevel sets the application's logger log level.
		LogLevel string
	}

	// ConfigOption allows to customise a *Config.
	ConfigOption func(*Config)
)

// ConfigWithLogLevel allows overriding the default log level.
func ConfigWithLogLevel(level string) ConfigOption {
	return func(conf *Config) {
		conf.LogLevel = level
	}
}

// New returns a new Config with overridable defaults.
func New(opts ...ConfigOption) *Config {
	// setting defaults
	conf := &Config{
		LogLevel: "error",
	}

	// applying options, if passed.
	for _, opt := range opts {
		opt(conf)
	}

	// returning a customised config.
	return conf
}
