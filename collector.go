package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"encoding/xml"
	"github.com/iancoleman/strcase"
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
	return update.extractMetrics(), nil
}

func (e ESATMUpdate) extractMetrics() []prometheus.Metric {
	// v.Type().Field(0).Name displays the name of a struct field
	// If the field name matches something in our state set map
	// process as state set according to the label name and
	// label values for that map key.
	// If the field type is int then process as gauge.
	// Else: process as string

	// update.Status (9) and update.FirmwareFeatureStates (3) are state sets
	result := make([]prometheus.Metric, 0, 64)
	updateVal := reflect.ValueOf(e)
	for i := 0; i < updateVal.NumField(); i++ {
		t := prometheus.UntypedValue
		log.Infof("%30s: %30v %30v", updateVal.Type().Field(i).Name, updateVal.Field(i).Interface(), updateVal.Field(i).Kind())
		field := updateVal.Type().Field(i).Name
		promField := "drobo_" + strcase.ToSnake(field)
		if _, ok := stateSets[field]; ok {
			samples := enumAsStateSet(int(updateVal.Field(i).Int()), field)
			result = append(result, samples...)
		} else {
			if updateVal.Field(i).Kind().String() == "int" {
				t = prometheus.GaugeValue
				sample, err := prometheus.NewConstMetric(prometheus.NewDesc(promField, helpInfo[field], nil, nil),
					t, float64(updateVal.Field(i).Int()))
				if err != nil {
					sample = prometheus.NewInvalidMetric(prometheus.NewDesc("drobo_error", "Error calling NewConstMetric", nil, nil),
						fmt.Errorf("error for metric %s", promField))
					log.Error(err)
				}
				result = append(result, sample)
			}
		}
	}
	return result
}
func enumAsStateSet(value int, field string) []prometheus.Metric {
	results := []prometheus.Metric{}
	state, ok := stateSets[field][value]
	if !ok {
		// Fallback to using the value.
		state = strconv.Itoa(value)
	}
	promField := "drobo_" + strcase.ToSnake(field)
	newMetric, err := prometheus.NewConstMetric(prometheus.NewDesc(promField, helpInfo[field], []string{promField}, nil),
		prometheus.GaugeValue, 1.0, state)
	if err != nil {
		newMetric = prometheus.NewInvalidMetric(prometheus.NewDesc("drobo_error", "Error calling NewConstMetric for EnumAsStateSet", nil, nil),
			fmt.Errorf("error for metric %s", promField))
	}
	results = append(results, newMetric)

	for k, v := range stateSets[field] {
		if k == value {
			continue
		}
		newMetric, err := prometheus.NewConstMetric(prometheus.NewDesc(promField, helpInfo[field], []string{promField}, nil),
			prometheus.GaugeValue, 0.0, v)
		if err != nil {
			newMetric = prometheus.NewInvalidMetric(prometheus.NewDesc("drobo_error", "Error calling NewConstMetric for EnumAsStateSet", nil, nil),
				fmt.Errorf("error for metric %s", promField))
		}
		results = append(results, newMetric)
	}
	return results
}

var stateSets = map[string]map[int]string{
	"Status": map[int]string{
		0x8000:  "droboOK",
		0x8004:  "overYellowThreshold",
		0x8006:  "overRedThreshold",
		0x8010:  "hasBadDrive",
		0x8046:  "aDriveHasBeenRemoved",
		0x8240:  "dataProtectionInProgress",
		0x18000: "dashboardIndicatesDroboIsOK",
		0x18006: "droboOverRedThreshold",
		0x18240: "dataProtectionInProgressDoNotRemoveDrives",
	},
	"FirmwareFeatureStates": map[int]string{
		0x4: "unknown",
		0x6: "singleRedundancy",
		0x7: "dualRedundancy",
	},
}
