# PromQl summary and histogram

## setup
```
cd monitoring/PromQL/histogram/
chmod +x generate-requests.sh 

```

## Summary
Summary adds _count and _sum suffix to the metrics like this one, both are counters.
```
# TYPE web_delay_seconds summary
web_delay_seconds_sum 800
web_delay_seconds_count 149
```
- [complete queries](http://localhost:9090/graph?g0.expr=web_delay_seconds_count&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h&g1.expr=web_delay_seconds_sum&g1.tab=1&g1.stacked=0&g1.show_exemplars=0&g1.range_input=1h&g2.expr=rate(web_delay_seconds_count%20%5B5m%5D)&g2.tab=1&g2.stacked=0&g2.show_exemplars=0&g2.range_input=1h&g3.expr=rate(web_delay_seconds_sum%5B5m%5D)&g3.tab=1&g3.stacked=0&g3.show_exemplars=0&g3.range_input=1h&g4.expr=rate(web_delay_seconds_sum%5B5m%5D)%20%2F%20rate(web_delay_seconds_count%20%5B5m%5D)&g4.tab=1&g4.stacked=0&g4.show_exemplars=0&g4.range_input=1h)
- [complete queries as graph](http://localhost:9090/graph?g0.expr=web_delay_seconds_count&g0.tab=0&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h&g1.expr=web_delay_seconds_sum&g1.tab=0&g1.stacked=0&g1.show_exemplars=0&g1.range_input=1h&g2.expr=rate(web_delay_seconds_count%20%5B5m%5D)&g2.tab=0&g2.stacked=0&g2.show_exemplars=0&g2.range_input=1h&g3.expr=rate(web_delay_seconds_sum%5B5m%5D)&g3.tab=0&g3.stacked=0&g3.show_exemplars=0&g3.range_input=1h&g4.expr=rate(web_delay_seconds_sum%5B5m%5D)%20%2F%20rate(web_delay_seconds_count%20%5B5m%5D)&g4.tab=0&g4.stacked=0&g4.show_exemplars=0&g4.range_input=1h)
### rate(web_delay_seconds_count [5m])
Number of slow mode requests per seconds that had been received in the duration of 5 minutes.

| Metrics | result |
| --- | --- |

