# Name: check-argocd

## Verifications:

- Verify login attempt using access credentials. ( Expected: "200" )

- Verify database return at health check endpoint. ( Expected: "ok" )
 
- Check if the default data source matches the value configured in the variable. ( Expected: "isDefault = True" )
 

## URL:

- Service: http://kube-prometheus-stack-grafana.tks-system.svc.cluster.local


## Endpoints:

- /api/datasources ( requires authentication)
- /api/health ( does not require authentication )


## Variables:

- USER string
- PASS string
- DATASOURCE string


## TestKube Type: postman/collection


## Create Test:

```
kubectl testkube create test --name check-grafana --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path tests/grafana/grafana.json
```

## Run Test:

```
kubectl testkube run test check-grafana --secret-variable USER=$USERNAME --secret-variable PASS=$PASSWORD --secret-variable DATASOURCE=$DATA
```
