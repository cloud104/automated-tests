# Name: Certmanager

## TestKube Type: Golang/Ginkgo

## Verifications:

- The check validates if  you can create a certificate using your CRD a yaml manifest on a kubernetes cluster.

## Variables:

- CLUSTER_ID string
- REGION_DNS string

## Create Test:

```
kubectl testkube create test --name certmanager --type ginkgo/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/ginkgo/certmanager
```

## Run Test:

```
kubectl testkube run test certmanager --secret-variable CLUSTER_ID=$CLUSTER_ID --secret-variable REGION_DNS=$REGION_DNS 
```
