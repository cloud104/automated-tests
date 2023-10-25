# Name: Certmanager

## TestKube Type: Golang/Ginkgo

## Verifications:

- The check validates if you can create a certificate through yaml manifest.

## Variables:

- CLUSTER_ID (--variable)
- REGION_DNS (--variable)

## Create Test:

```
kubectl testkube create test --name certmanager --type ginkgo/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/ginkgo/certmanager
```

## Run Test:

```
kubectl testkube run test certmanager -v CLUSTER_ID="" -v REGION_DNS="" 
```
