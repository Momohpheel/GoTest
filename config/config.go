package config

import "github.com/caarlos0/env/v6"

func init() {
	_ = env.Parse(&App)
}
