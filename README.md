# jenkins-metrics

A simple program to collect metrics from Jenkins and surface the metrics as
[InfluxDB line protocol][idb-line-protocol].  Written in Go.

Prereqisites:

- Install the [Jenkins Metrics Plugin][jenkins-metrics].

[idb-line-protocol]: https://docs.influxdata.com/influxdb/v1.2/write_protocols/line_protocol_tutorial/
[jenkins-metrics]: https://wiki.jenkins-ci.org/display/JENKINS/Metrics+Plugin
