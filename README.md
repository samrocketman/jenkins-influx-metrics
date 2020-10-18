# Jenkins server metrics for InfluxDB

> WARNING: this project is pre-alpha quality and not ready for use.  When I feel it is ready for use then I'll remove this banner.

A simple program to collect metrics from a Jenkins server and surface the
metrics as [InfluxDB line protocol][idb-line-protocol].  Written in Go.

Prereqisites:

- Install the [Jenkins Metrics Plugin][jenkins-metrics].

# Usage

Let's say you have a Jenkins server at localhost and you've set an environment
variable `METRICS_TOKEN`.  You can view your Jenkins metrics as influx line
protocol with the following command.

    ./jenkins-influx-metrics \
      -jenkins-url http://localhost:8080/ \
      -metrics-token $METRICS_TOKEN \
      -tag-set 'host=localhost'

Integrating with telegraf to push the Jenkins server metrics to InfluxDB could
look something like the following.

    # Collect Jenkins metrics every minute
    [[inputs.exec]]
      commands = [
        "./jenkins-influx-metrics -jenkins-url http://localhost:8080/ -metrics-token $METRICS_TOKEN -tag-set 'host=localhost'"
      ]
      timeout = "30s"
      data_format = "influx"
      interval = "1m"
      [inputs.exec.tags]
        tag1 = "foo"
        tag2 = "bar"

Learn more about [Telegraf configuration][telegraf-conf] or more specifically
about [Telegraf `inputs.exec`][telegraf-exec].

# Building

    go build

[idb-line-protocol]: https://docs.influxdata.com/influxdb/v1.2/write_protocols/line_protocol_tutorial/
[jenkins-metrics]: https://wiki.jenkins-ci.org/display/JENKINS/Metrics+Plugin
[telegraf-conf]: https://github.com/influxdata/telegraf/blob/master/docs/CONFIGURATION.md
[telegraf-exec]: https://github.com/influxdata/telegraf/blob/master/etc/telegraf.conf
