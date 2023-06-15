package db

import "metrics/entity"

var InMemoryDb = make([]*entity.MonitorRow, 0, 100)

func AddToDb(s *entity.MonitorRow) {
	InMemoryDb = append(InMemoryDb, s)
}

func ClearDb() {
	InMemoryDb = make([]*entity.MonitorRow, 0, 100)
}
