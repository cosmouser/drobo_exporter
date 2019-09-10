package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9045").String()

func handler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' parameter must be specified", 400)
		return
	}
	registry := prometheus.NewRegistry()
	collector := collector{target: target}
	registry.MustRegister(collector)

	// Delegate http serving to Prometheus client library, which will call collector.Collect.
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func main() {
	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("drobo_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log.Infoln("Starting drobo exporter")
	log.Infoln("Build context", version.BuildContext())

	http.Handle("/metrics", promhttp.Handler()) // Normal metrics endpoint for Drobo exporter itself.
	http.HandleFunc("/drobo", handler)          // Endpoint to do Drobo scrapes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head>
		<title>Drobo Exporter</title>
		<style>
		label{
		display:inline-block;
		width:75px;
		}
		form label {
		margin: 10px;
		}
		form input {
		margin: 10px;
		}
		</style>
		</head>
		<body>
		<h1>Drobo Exporter</h1>
		<form action="/drobo">
		<label>Target:</label> <input type="text" name="target" placeholder="X.X.X.X" value="1.2.3.4"><br>
		<input type="submit" value="Submit">
		</form>
		</body>
		</html>`))
	})
	log.Infof("Listening on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
