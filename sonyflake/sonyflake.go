package sonyflake

import (
	"github.com/sony/sonyflake"
	"log"
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
