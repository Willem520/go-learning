package test

import (
	"github.com/robfig/cron"
	"log"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	cron := cron.New()
	cron.AddFunc("*/5 * * * * ?", func() {
		log.Printf("======test")
	})
	cron.Start()
	defer cron.Stop()
	time.Sleep(time.Minute)
}
