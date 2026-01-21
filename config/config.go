package config

import (
	"flag"
	"os"
	"sync"
)

type Config struct {
	Port         int
	Env          string
	ResendAPIKey string
	ContactEmail string
}

var (
	once     sync.Once
	instance *Config
)

func LoadConfigOnce() *Config {
	once.Do(func() {
		instance = &Config{}
		flag.IntVar(&instance.Port, "port", 3000, "Server port")
		flag.StringVar(&instance.Env, "env", "development", "Environment (development|production)")
		flag.StringVar(&instance.ResendAPIKey, "resend-api-key", os.Getenv("RESEND_API_KEY"), "Resend API Key")
		flag.StringVar(&instance.ContactEmail, "contact-email", os.Getenv("CONTACT_EMAIL"), "Contact Email recipient")
		flag.Parse()

		if instance.ContactEmail == "" {
			instance.ContactEmail = "nelsonmarro@gmail.com" // Fallback
		}
	})
	return instance
}
