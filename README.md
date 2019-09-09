# drobo_exporter
prometheus exporter for drobo devices

Modeled after [snmp_exporter](https://github.com/prometheus/snmp_exporter) this tool scrapes Drobo products and exposes their metrics for prometheus. Help data is from [droboports](https://github.com/droboports/droboports.github.io/wiki/NASD-XML-format#mstatus). Currently supports Drobo 5N and Drobo 800FS devices.

## Usage

```sh
./drobo_exporter
```

## Prometheus Configuration

The drobo exporter needs to be passed the address as a parameter, this can be
done with relabelling.

Example config:
```YAML
scrape_configs:
  - job_name: 'drobo'
    static_configs:
      - targets:
        - 192.168.1.2:5000  # drobo device.
    metrics_path: /drobo
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:9045  # The drobo exporter's real hostname:port.
```

## to do
- expose metrics for each individual drive in mSlotsExp
