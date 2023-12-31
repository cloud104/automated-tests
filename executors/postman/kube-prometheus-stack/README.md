# Name: kube-prometheus-stack

## TestKube Type: postman/collection

## Verifications:

### Grafana

- Verify login attempt using access credentials. 

- Verify database return at health check endpoint. ( Expected: "ok" )
 
- Check if the default data source matches the value configured in the variable. ( Expected: "isDefault = True" )

### Prometheus

- Checks if the health endpoint is as expected. (Expected: "Prometheus Server is Healthy")

### Alertmanager

- Checks if the health endpoint is as expected. (Expected: "OK")
- Check have active alertmanager on prometheus.

### Node-exporter

- Query the Prometheus API using the Node-exporter target to check if a metric is valid.


## Endpoints:

Grafana:
- http://kube-prometheus-stack-grafana.tks-system.svc.cluster.local/api/health 
- http://kube-prometheus-stack-grafana.tks-system.svc.cluster.local/api/datasources

Prometheus:
- http://prometheus-tks-prometheus.tks-system.svc.cluster.local:9090/-/healthy

Alertmanager:
- http://prometheus-tks-alertmanager.tks-system.svc.cluster.local:9093/-/healthy
- http://prometheus-tks-prometheus.tks-system.svc.cluster.local:9090/api/v1/alertmanagers

Node-exporter:
- http://prometheus-tks-prometheus.tks-system.svc.cluster.local:9090/api/v1/targets/metadata?match_target={job="node-exporter"}&metric=promhttp_metric_handler_requests_total


## Environment Variables

The following variable is required to run the test on the application:

| Variable     | Description                                            |
|--------------|--------------------------------------------------------|
| GRAFANA_USER  | Specifies the username for login on Grafana Api.      |
| GRAFANA_PASS  | Specifies the password for login on Grafana Api.      |
| GRAFANA_DATASOURCE | Specifies the Datasource validate on Grafana Api.|


## Create Test:

```
kubectl testkube create test --name kube-prometheus-stack --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/kube-prometheus-stack/kube-prometheus-stack.json
```

## Run Test:

```
kubectl testkube run test kube-prometheus-stack -s GRAFANA_USER="" -s GRAFANA_PASS="" -v GRAFANA_DATASOURCE=""
```
