package config

import cfg "github.com/Yalantis/go-config"

var config Config

func Get() *Config {
	return &config
}

func Load(fileName string) error {
	return cfg.Init(&config, fileName)
}

type (
	Config struct {
		AppName            string `json:"app_name" envconfig:"API_APP_NAME" default:"api"`
		LogPreset          string `json:"log_preset" envconfig:"API_LOG_PRESET" default:"development"`
		ListenURL          string `json:"listen_url" envconfig:"API_LISTEN_URL" default:":8080"`
		PaginationMaxLimit int64  `json:"pagination_max_limit" envconfig:"API_PAGINATION_MAX_LIMIT" default:"1000"`

		Postgres Postgres `json:"postgres"`
	}

	Postgres struct {
		Host     string `json:"host"          envconfig:"POSTGRES_HOST"              default:"localhost"`
		Port     string `json:"port"          envconfig:"API_POSTGRES_PORT"          default:"5432"`
		Database string `json:"database"      envconfig:"API_POSTGRES_DATABASE"      default:"manager"`
		User     string `json:"user"          envconfig:"API_POSTGRES_USER"          default:"postgres"`
		Password string `json:"password"      envconfig:"API_POSTGRES_PASSWORD"      default:"12345"`
	}
)
