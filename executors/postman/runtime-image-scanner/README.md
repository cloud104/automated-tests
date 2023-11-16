# Name: Runtime-image-scanner

## TestKube Type: postman/collection

## Verifications:

- Checks if the metrics endpoint is as expected. (Expected status code: "200") 

## Endpoints:

- http://runtime-image-scanner.tks-system.svc.cluster.local:8080/metrics


## Create Test:

```
kubectl testkube create test --name runtime-image-scanner --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/runtime-image-scanner/runtime-image-scanner.json
```

## Run Test:

```
kubectl testkube run test runtime-image-scanner
```
