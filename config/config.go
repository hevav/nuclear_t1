package config

import (
	"hackathon/db"
	"hackathon/queue"
)

var Config ConfigType

type ConfigType struct {
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
			Type     db.DatabaseType
			Host     string
			Database string
			Username string
			Password string
		}
		BankDatabase struct {
			Type     db.DatabaseType
			Host     string
			Database string
			Username string
			Password string
		}
		Queue struct {
			Enabled bool
			Type    queue.QueueType
			Host    string
			Topic   string
		}
		HTTP struct {
			Enabled     bool
			BindAddress string
		}
	}
}
