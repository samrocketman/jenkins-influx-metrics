# How metrics are displayed

Tags are optional.  Pass `-tag-set` option to CLI e.g. `-tag-set host=myhost`.

In the metrics API, if a metric field ends with `.value` then the `.value` is
ignored and the metric name is derived without `_value`.  If a metric value is a
`String` then it is completely ignored as a metric entirely.  Only integers and
floats are surfaced as metrics.

# `metrics` endpoint

### Gauges

The last CSV field in they key is the key of the fieldset.  All preceding
periods and hyphens are turned into underscores and used as the measurement.
For example:

    "jenkins.executor.count.value" : {
      "value" : 2
    },
    "jenkins.executor.free.value" : {
      "value" : 2
    },
    "jenkins.executor.in-use.value" : {
      "value" : 0
    }

Following the rules so far: `.value` is ignored, the last CSV field is a part
of the line protocol fieldset, and all preceding fields have periods replaced
with underscores.  So the resulting influx line protocol would be:

    jenkins_executor count=2i,free=2i,in-use=0i

One more example,

    "vm.memory.heap.committed" : {
      "value" : 2466775040
    },
    "vm.memory.heap.init" : {
      "value" : 526385152
    },
    "vm.memory.heap.max" : {
      "value" : 7486832640
    },
    "vm.memory.heap.usage" : {
      "value" : 0.1394093745896796
    },
    "vm.memory.heap.used" : {
      "value" : 1043734656
    },

The above example turns into the following influx line protocol.

    vm_memory_heap committed=2466775040i,init=526385152i,max=7486832640i,usage=0.1394093745896796,used=1043734656i

### Histograms

Histograms are currently ignored from the metrics plugin.

### Timers

Each key is treated like an influx metric.  The `values` subkey is ignored.  The
subkey/value pairs are part of the line protocol fieldset.  The metric is named
after the parent key but with the periods and hyphens replaced with underscores.
For example,

    "jenkins.health-check.duration" : {
      "count" : 7,
      "max" : 0.0014854870000000002,
      "mean" : 8.38357636372125E-4,
      "min" : 3.24028E-4,
      "p50" : 8.21298E-4,
      "p75" : 8.21298E-4,
      "p95" : 0.001079536,
      "p98" : 0.001079536,
      "p99" : 0.001079536,
      "p999" : 0.0014854870000000002,
      "values" : [ 3.24028E-4, 4.93388E-4, 5.534600000000001E-4, 6.15958E-4, 8.21298E-4, 0.001079536, 0.0014854870000000002 ],
      "stddev" : 1.8236473518573368E-4,
      "m15_rate" : 0.3636368352669909,
      "m1_rate" : 0.6590553102091476,
      "m5_rate" : 0.6978092521149444,
      "mean_rate" : 0.983123130313119,
      "duration_units" : "seconds",
      "rate_units" : "calls/minute"
    }

The above example would turn into the following influx line protocol.

    jenkins_health_check_duration count=7i,max=0.0014854870000000002,mean=8.38357636372125E-4,min=3.24028E-4...  etc

# `healthcheck` endpoint

Sample metrics from health check endpoint.

	{
	  "disk-space" : {
		"healthy" : true
	  },
	  "plugins" : {
		"healthy" : true,
		"message" : "No failed plugins"
	  },
	  "temporary-space" : {
		"healthy" : true
	  },
	  "thread-deadlock" : {
		"healthy" : true
	  }
	}

Resulting influx line protocol.

	healthcheck_endpoint disk_space_healthy=true,plugins_healthy=true,temporary_space_healthy=true,thread_deadlock_healthy=true
