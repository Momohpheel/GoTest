package config

var App struct {
	Port   string `env:"Port" envDefault:"4500"`
	ENV    string `env:"ENV" envDefault:"local"`
	JWTKey string `env:"JWTKey" envDefault:"kiuru72h2ywn"`
}
