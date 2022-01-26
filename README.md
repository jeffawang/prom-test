
run http server with metrics

```
go run .
```

run prometheus

```
docker run -p 9090:9090 -v $PWD/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```
