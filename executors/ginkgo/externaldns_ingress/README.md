# Name: Externaldns-Ingress

## TestKube Type: Golang/Ginkgo

## Verifications:

- Check if it is possible to create resources via yaml manifest (Deployment, service and ingress (nginx web server)).
- Expect status code 200 return from request for the created ingress.

## Variables:

- CLUSTER_ID string
- REGION_DNS string

## Create Test:

```
kubectl testkube create test --name externaldns-ingress --type ginkgo/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/ginkgo/externaldns_ingress
```

## Run Test:

```
kubectl testkube run test externaldns-ingress --secret-variable CLUSTER_ID=$CLUSTER_ID --secret-variable REGION_DNS=$REGION_DNS
```
