# Name: External-dns

## TestKube Type: postman/collection

## Verifications:

- Checks if the health endpoint is as expected. (Expected: "OK")


## Endpoints:

- http://external-dns.tks-system.svc.cluster.local:7979/healthz


## Create Test:

```
kubectl testkube create test --name external-dns --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/external-dns/external-dns.json
```

## Run Test:

```
kubectl testkube run test external-dns
```
