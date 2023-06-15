package db

import "metrics/internal/entity"

var InMemoryDb = make([]*entity.MonitorRow, 0, 100)

func AddToDb(s *entity.MonitorRow) {
	InMemoryDb = append(InMemoryDb, s)
}

func ClearDb() {
	InMemoryDb = make([]*entity.MonitorRow, 0, 100)
}
