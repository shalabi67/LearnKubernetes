# Basic PromQL
You can find metrics of every source, by going to prometheus targets and from there use the url of the required target.

Have a look at the metrics [metrics](http://promo:8280/metrics) and notice that `worker_jobs_active` is a gauge metrics. 
while `worker_jobs_total` is a counter metrics. 

A `gauge` is a metric that represents a single numerical value that can arbitrarily go up and down. 

## Example
### worker_jobs_active
| Metrics | result |
| --- | --- |
| worker_jobs_active{instance="promo:8280", job="batch", os="linux", runtime="docker"}  |  83 |
| worker_jobs_active{instance="promo:8281", job="batch", os="linux", runtime="docker"} |   45 |
### worker_jobs_active{instance="promo:8280"}
- worker_jobs_active{instance="promo:8280", job="batch", os="linux", runtime="docker"}    83
  
### worker_jobs_active[5m]
| Metrics | result |
| --- | --- |
| worker_jobs_active{instance="promo:8280", job="batch", os="linux", runtime="docker"} | 28 @1638389714.684 <br> 78 @1638389774.684 <br> 21 @1638389834.684 <br> 86 @1638389894.684 <br> 83 @1638389954.684 |
| worker_jobs_active{instance="promo:8281", job="batch", os="linux", runtime="docker"} | 60 @1638389751.695 <br> 36 @1638389811.695 <br> 42 @1638389871.695 <br> 98 @1638389931.695 <br> 45 @1638389991.695 |

 