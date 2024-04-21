package config

import (
	"hackathon/types"
)

var Settings SettingsType

type SettingsType struct {
	Specific struct {
		GeoZone struct {
			DefaultRadiusMeters uint64
		}
		Retrain struct {
			ScheduleSeconds uint64
		}
	}
	Integration struct {
		LocalDatabase struct {
			Type     types.DatabaseType
			Host     string
			Database string
			Username string
			Password string
		}
		BankDatabase struct {
			Type     types.DatabaseType
			Host     string
			Database string
			Username string
			Password string
		}
		Queue struct {
			Enabled bool
			Type    types.QueueType
			Host    string
			Topic   string
		}
		HTTP struct {
			Enabled     bool
			BindAddress string
		}
	}
}
