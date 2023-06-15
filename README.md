# go-metrics
Simple pkg for monitoring estimation of SQL queries 

# Usage:
## Init 
metrics.RunServer(metrics.Options{
		Host: "localhost",
		Port: 8080,
})

and open in browser page by: http://localhost:8080 for control your queries


## Methods
func AddToMonitor(sql string, params any) *entity.MonitorRow - add query to monitor
func GetCurrentMonitor() []*entity.MonitorRow - get last saved queries
func EstimateQuery(sql string, params any, runner func()) - for estimate run your queries in "runner" func 

