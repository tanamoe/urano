package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Fahasa   FahasaConfig   `mapstructure:"fahasa"`
	Jobs     JobsConfig     `mapstructure:"jobs"`
}

type AppConfig struct {
	ListenAddress  string   `mapstructure:"listenAddress"`
	AllowedOrigins []string `mapstructure:"allowedOrigins"`
}

type DatabaseConfig struct {
	Conn string `mapstructure:"conn"`
}

type FahasaConfig struct {
	SearchToken string `mapstructure:"searchToken"`
}

type JobsConfig struct {
	DailyRegistry ScheduleConfig `mapstructure:"dailyRegistry"`
}

type ScheduleConfig struct {
	Enable  bool   `mapstructure:"enable"`
	Crontab string `mapstructure:"crontab"`
}

func Load() (*Config, error) {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
