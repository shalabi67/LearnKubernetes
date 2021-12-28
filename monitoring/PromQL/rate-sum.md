# Rate and Sum

## Rate
Will return the rate of change per second over a range. This is used for counters.

### Examples
#### rate(worker_jobs_total[5m]) 
This query will return the rate of change per second over 5 min range.

| Metrics | result |
| --- | --- |
| {instance="promo:8280", job="batch", os="linux", runtime="docker", status="failed"} | 4.779166666666667 |
| {instance="promo:8280", job="batch", os="linux", runtime="docker", status="processed"} | 48.016666666666666 |
| {instance="promo:8281", job="batch", os="linux", runtime="docker", status="failed"} | 4.845833333333333 |
| {instance="promo:8281", job="batch", os="linux", runtime="docker", status="processed"} | 42.0125 |

#### rate(worker_jobs_total{status="failed"}[5m])
This query will return the rate of failed jobs per second over 5 min range.

| Metrics | result |
| --- | --- |
| {instance="promo:8280", job="batch", os="linux", runtime="docker", status="failed"} | 4.779166666666667 |
| {instance="promo:8281", job="batch", os="linux", runtime="docker", status="failed"} | 4.845833333333333 |


## Sum and Rate
The rate function will return a gauge or a victor based on labels, while the sum sums the values of labels.

### Examples
#### sum by (status)(rate(worker_jobs_total[5m]))
This will sum the values of status label excluding job, os and instance for the rate of `worker_jobs_total` 
in the range of 5 min. 

Metrics | result |
| --- | --- |
{status="failed"} | 9.625 |
{status="processed"} | 90.02916666666667 |

the value 9.625 is 4.845833333333333 + 4.779166666666667

the value 90.02916666666667 is 	42.0125 + 48.016666666666666

#### sum by (instance, status)(rate(worker_jobs_total[5m]))
This will sum the values of status label excluding job, and os for the rate of `worker_jobs_total`
in the range of 5 min. 

| Metrics | result |
| --- | --- |
| {instance="promo:8280", job="batch", os="linux", runtime="docker", status="failed"} | 4.779166666666667 |
| {instance="promo:8280", job="batch", os="linux", runtime="docker", status="processed"} | 48.016666666666666 |
| {instance="promo:8281", job="batch", os="linux", runtime="docker", status="failed"} | 4.845833333333333 |
| {instance="promo:8281", job="batch", os="linux", runtime="docker", status="processed"} | 42.0125 |

Notice that this is exactly the same result as rate(worker_jobs_total[5m])  because job and os are not changing.