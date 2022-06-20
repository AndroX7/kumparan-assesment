package cron

import (
	"github.com/robfig/cron/v3"
)

type Usecase interface {
	Cron() *cron.Cron
}
