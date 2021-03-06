package conf

type Config struct {
	Domain   string
	RealAddr []string `mapstructure:"real_addr"`
	Cookie   Cookie   `mapstructure:"cookie"`
	Feishu   Feishu   `mapstructure:"feishu"`
	Dingtalk Dingtalk `mapstructure:"dingtalk"`
}

type Cookie struct {
	Name   string `mapstructure:"name"`
	Domain string `mapstructure:"domain"`
}

type Feishu struct {
	BaseUrl      string `mapstructure:"base_url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}

type Dingtalk struct {
	ApiUrl       string `mapstructure:"api_url"`
	LoginUrl     string `mapstructure:"login_url"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUri  string `mapstructure:"redirect_uri"`
}

func (c *Config) Init() error {
	return nil
}
