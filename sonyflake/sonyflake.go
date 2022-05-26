package sonyflake

import (
	log "github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

var SF *sonyflake.Sonyflake

func init() {
	SF = sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: nil,
	})
	if SF == nil {
		log.Fatal("SonyFlake not created")
	}
}
