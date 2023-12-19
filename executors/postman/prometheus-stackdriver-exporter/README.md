# Name: Stackdriver Prometheus Exporter

## TestKube Type: postman/collection

## Verifications:

- Checks if the metrics endpoint is as expected. (Expected status code: "200")

## Endpoints:

- http://prometheus-stackdriver-exporter.tks-system:9255/metrics

### Environment Variables

The following variables are available for configuring the service:

| Variable     | Default Value                   | Description                                  |
|--------------|---------------------------------|----------------------------------------------|
| SERVICE_NAME | prometheus-stackdriver-exporter | Specifies the name of the service.           |
| NAMESPACE    | tks-system                      | The namespace where the service is deployed. |
| PORT         | 9255                            | The port on which the service is accessible. |

## Create Test:

```
kubectl testkube create test --name prometheus-stackdriver-exporter --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/prometheus-stackdriver-exporter/prometheus-stackdriver-exporter.json
```

## Run Test:

```
kubectl testkube run test prometheus-stackdriver-exporter
```
