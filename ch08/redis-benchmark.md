```
redis-benchmark -t set|get -n 100000 -q -d 5
```





"test"|"rps"|"avg_latency_ms"|"min_latency_ms"|"p50_latency_ms"|"p95_latency_ms"|"p99_latency_ms"|"max_latency_ms"
｜-|:-:|:-:|:-:|:-:|:-:|:-:|-:
"SET","51493.30","0.510","0.216","0.487","0.751","0.999","4.607"
"GET","52328.62","0.499","0.240","0.479","0.703","0.975","8.263"