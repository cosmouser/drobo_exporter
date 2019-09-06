package main

import (
	"reflect"
	"time"

	"encoding/xml"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"net"
)

type collector struct {
	target string
}

// Describe implements Prometheus.Collector.
func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}

// Collect implements Prometheus.Collector.
func (c collector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	samples, err := ScrapeTarget(c.target)
	if err != nil {
		log.Infof("Error scraping target %s: %s", c.target, err)
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("drobo_error", "Error scraping target", nil, nil), err)
		return
	}
	for _, sample := range samples {
		ch <- sample
	}
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("drobo_scrape_duration_seconds", "Time Drobo scrape took.", nil, nil),
		prometheus.GaugeValue,
		time.Since(start).Seconds())
}

// ScrapeTarget fetches xml from a drobo and returns metrics
func ScrapeTarget(target string) ([]prometheus.Metric, error) {
	addr := net.ParseIP(target)
	if addr == nil {
		return nil, errors.New("invalid target address")
	}
	out, err := dialUPNP(addr)
	if err != nil {
		return nil, err
	}
	update := ESATMUpdate{}
	err = xml.Unmarshal(out, &update)
	if err != nil {
		return nil, err
	}
	// update.Status (9) and update.FirmwareFeatureStates (3) are state sets
	result := make([]prometheus.Metric, 0, 64)
	updateVal := reflect.ValueOf(update)
	values := make([]interface{}, updateVal.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = updateVal.Field(i).Interface()
	}
	return nil, nil
}
