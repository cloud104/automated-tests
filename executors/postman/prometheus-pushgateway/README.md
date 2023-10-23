# Name: Prometheus-Pushgateway

## TestKube Type: postman/collection

## Verifications:

- Checks if the health endpoint is as expected. (Expected: "OK")
- checks whether the application is ready to serve traffic. (Expected: "OK")
- Checks if it is possible to send a random metric through an http request. (Expected status code: "200")
 

## Endpoints:

- http://prometheus-pushgateway.tks-system.svc.cluster.local:9091/-/healthy
- http://prometheus-pushgateway.tks-system.svc.cluster.local:9091/-/ready
- http://prometheus-pushgateway.tks-system.svc.cluster.local:9091/metrics/job/testkube_test_metric


## Create Test:

```
kubectl testkube create test --name prometheus-pushgateway --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/prometheus-pushgateway/prometheus-pushgateway.json
```

## Run Test:

```
kubectl testkube run test prometheus-pushgateway
```
