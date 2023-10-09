# Name: Certmanager

## TestKube Type: Golang/Ginkgo

## Verifications:

- The check validates that you can create a certificate using your CRD via a yaml manifest on a kubernetes cluster.

## Variables: ( on Manifest )

- CLUSTERID string
- REGIONDNS string

## Create Test:

```
kubectl testkube create test --name check-certmanager --type ginkgo/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/ginkgo/certmanager
```

### Or

```
kubectl create -f https://raw.githubusercontent.com/cloud104/automated-tests/master/executors/ginkgo/certmanager/certmanager.yaml
```

## Run Test:

```
kubectl testkube run test check-certmanager --secret-variable REGIONDNS=$region_dns --secret-variable CLUSTERID=$clusterid
```
