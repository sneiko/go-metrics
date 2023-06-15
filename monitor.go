package metrics

import (
	"metrics/db"
	"metrics/entity"
)

func AddToMonitor(sql string, params any) *entity.MonitorRow {
	query := entity.NewMonitorRow(sql, params)
	db.AddToDb(query)
	return query
}

func GetCurrentMonitor() []*entity.MonitorRow {
	return db.InMemoryDb
}

func EstimateQuery(sql string, params any, runner func()) {
	query := AddToMonitor(sql, params)
	runner()
	query.Finish()
}
