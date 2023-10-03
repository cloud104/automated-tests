# Name: Check-Stack-Logs

## TestKube Type: postman/collection

## Verifications:

- Verify login attempt using access credentials. ( Expected: "200" )


 
## URL:

- Service: http://kube-prometheus-stack-grafana.tks-system.svc.cluster.local

## Endpoints:

- /api/datasources ( requires authentication)
- /api/health ( does not require authentication )

## Variables:

- USER string
- PASS string
- DATASOURCE string

## Create Test:

```
kubectl testkube create test --name check-grafana --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/grafana/grafana.json
```
## Run Test:

```
kubectl testkube run test check-stack-logs --secret-variable USER=$USERNAME --secret-variable PASS=$PASSWORD 
```