package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("config") // Nome do arquivo
	v.SetConfigType("yaml")   // Tipo do arquivo
	v.AddConfigPath(".")      // Caminho padrão

	// Definir variáveis padrão
	v.SetDefault("server.port", "8080")
	v.SetDefault("auth.secret", "supersecretapikey")
	v.SetDefault("database.url", "postgresql://user:password@localhost:5432/app")

	// Ler configurações do arquivo
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using defaults: %v", err)
	}

	Config = v
	return nil
}
