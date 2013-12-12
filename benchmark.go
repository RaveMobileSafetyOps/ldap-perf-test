package main

import (
	"fmt"
	"time"
)

type BTimer struct {
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Ended     bool
}

func NewTimer() *BTimer {
	t := &BTimer{StartTime: time.Now(), Ended: false}
	return t
}

func (self *BTimer) End() {
	self.EndTime = time.Now()
	self.Ended = true
	self.Duration = self.EndTime.Sub(self.StartTime)
}

func (self *BTimer) String() string {
	if self.Ended {
		return fmt.Sprintf("dur: %s (start: %s / end: %s)", self.Duration.String(), self.StartTime.String(), self.EndTime.String())
	} else {
		return fmt.Sprintf("dur: %s (start: %s / IN PROGRESS)", time.Since(self.StartTime).String(), self.StartTime.String())
	}
}
