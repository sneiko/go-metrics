package metrics

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"metrics/entity"
)

type Options struct {
	Host string
	Port int
}

func RunServer(opt Options) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.New("index")
		tmplt, err := tmplt.Parse(`
		<html>
			<head>
			<title>SQL Monitor</title>
			</head>
			<body>
			<h2>Last pack of queries</h2>
		
			<table style="width: 100%">
			<thead>
			<td><b>SQL</b></td>
			<td><b>Params</b></td>
			<td><b>Estimate sec.</b></td>
			</thead>
				{{range .List}}
					<tr>
						<td>{{.SQL}}</td>
						<td>{{.Params}}</td>
						{{if .Duration }}
							<td>{{.Duration}}</td>
						{{else}}
							<td>N/A</td>
						{{end}}
					</tr>
				{{end}}
			</table>
			</body>
		</html>`)
		if err != nil {
			log.Fatalln(err)
		}

		result := struct {
			List []*entity.MonitorRow
		}{GetCurrentMonitor()}

		tmplt.Execute(w, result)
	})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", opt.Host, opt.Port), nil); err != nil {
		log.Fatal(err)
	}
}
