package biteship

const (
	DefaultUrl = "https://api.biteship.com"
)

type ConfigOption struct {
	SecretKey   string
	BiteshipUrl string
}

func DefaultConfig(secretKey string) *ConfigOption {
	return &ConfigOption{
		SecretKey:   secretKey,
		BiteshipUrl: DefaultUrl,
	}
}
