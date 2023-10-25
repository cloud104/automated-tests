# Name: Ingress-Nginx

## TestKube Type: Golang/Ginkgo

## Verifications:

- Check if it is possible to create resources via yaml manifest (Deployment, service and ingress (nginx web server)).
- Expect status code 200 return from ingress request created.

## Variables:

- CLUSTER_ID (--variable)
- REGION_DNS (--variable)

## Create Test:

```
kubectl testkube create test --name ingress-nginx --type ginkgo/test --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/ginkgo/ingress-nginx
```

## Run Test:

```
kubectl testkube run test ingress-nginx -v CLUSTER_ID="" -v REGION_DNS=""
```
