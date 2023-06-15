package entity

import (
	"fmt"
	"time"
)

type MonitorRow struct {
	SQL      string
	Params   any
	Duration *time.Duration

	startTime time.Time
}

func NewMonitorRow(sql string, params any) *MonitorRow {
	return &MonitorRow{
		SQL:       sql,
		Params:    params,
		startTime: time.Now(),
	}
}

func (m *MonitorRow) Finish() {
	estimation := time.Since(m.startTime)
	m.Duration = &estimation
}

func (m *MonitorRow) String() string {
	return fmt.Sprintf("%s | %v | %v", m.SQL, m.Params, m.Duration)
}
