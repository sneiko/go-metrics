# go-metrics
Simple pkg for monitoring estimation of SQL queries 

# Usage:
## Init 
```
metrics.RunServer(metrics.Options{
		Host: "localhost",
		Port: 8080,
})
```

and open in browser page by: http://localhost:8080 for control your queries


## Methods
### Add query to monitor
```
func AddToMonitor(sql string, params any) *entity.MonitorRow
```

### Get last saved queries
```
func GetCurrentMonitor() []*entity.MonitorRow
```

### For estimate run your queries in "runner" func 
```
func EstimateQuery(sql string, params any, runner func())
```
