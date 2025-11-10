package utils

import (
	"fmt"
	"time"
)

type Timing struct {
	start time.Time
	name  string
}

func StartTiming(name string) *Timing {
	return &Timing{
		start: time.Now(),
		name:  name,
	}
}

func (t *Timing) End() {
	duration := time.Since(t.start)
	fmt.Printf("⏱ [%s] 耗时: %.2f ms\n", t.name, float64(duration.Microseconds())/1000.0)
}
