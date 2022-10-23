package config

type AppConfig struct {
	Api Api `json:"api"`
}

type Api struct {
	Port string `json:"port"`
}
