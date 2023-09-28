# Name: check-alertmanager

## TestKube Type: curl/test

### CRD: 

- ../../executors/curl/crd-alertmanager.yaml

## Verifications:

- Checks if the health endpoint is as expected. (Expected: "OK")
 
## URL:

- Service: http://prometheus-tks-alertmanager.tks-system.svc.cluster.local:9093

## Endpoints:

- /-/healthy (not requires authentication)

## Create Test:

```
kubectl testkube create test --name check-alertmanager --type curl/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path tests/alertmanager/alertmanager.json
```
## Run Test:

```
kubectl testkube run test check-alertmanager
```