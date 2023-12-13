# Name: Dex

## TestKube Type: postman/collection

## Verifications:

- Checks if the health endpoints is as expected. (Expected: "OK")

## Endpoints:

- http://dex.tks-system.svc.cluster.local:5558/healthz/live
- http://dex.tks-system.svc.cluster.local:5558/healthz/ready
- http://tks-login.tks-system.svc.cluster.local:5555/healthz


## Create Test:

```
kubectl testkube create test --name dex--type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/dex/dex.json
```

## Run Test:

```
kubectl testkube run test dex
```
