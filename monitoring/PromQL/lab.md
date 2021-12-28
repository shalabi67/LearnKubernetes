# Lab

## sum by(instance) (node_cpu_seconds_total{mode=~"system|user"})
Write a Query to Determine Which Instances Have High CPU Utilization
another solution is
sum by(instance) (node_cpu_seconds_total{mode=~"system"} + ignoring(mode) node_cpu_seconds_total{mode=~"user"})

## rate(http_request_duration_seconds_sum[5m])
Get the Per-Second Rate of Increase in HTTP Request Duration for All Instances