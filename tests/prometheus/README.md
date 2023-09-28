# Name: check-prometheus

## Verifications:

- Checks if the health endpoint is as expected. (Expected: "Prometheus Server is Healthy")
 
## URL:

- Service: http://prometheus-tks-prometheus.tks-system.svc.cluster.local:9090

## Endpoints:

- /-/healthy (not requires authentication)

## TestKube Type: curl/test

## Create Test:

```
kubectl testkube create test --name check-prometheus --type curl/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path tests/prometheus/prometheus.json
```
## Run Test:

```
kubectl testkube run test check-prometheus
```