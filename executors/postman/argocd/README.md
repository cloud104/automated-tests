# Name: argocd

## TestKube Type: postman/collection

## Verifications:

- Verify login attempt using access credentials and get session token. ( Expected: "200" )

- Check if the configured repositories have a successful connection. ( Expected: "Connection: Successful" )

- Check the return of the apps. ( Expected: "Synced" | "Healthy" )
 

## Endpoints:

- http://argo-cd-argocd-server.tks-system.svc.cluster.local/api/v1/session (requires authentication)
- http://argo-cd-argocd-server.tks-system.svc.cluster.local/api/v1/repositories (requires authentication)
- http://argo-cd-argocd-server.tks-system.svc.cluster.local/api/v1/applications? (requires authentication)


## Variables:

- ARGOCD_PASS (--secret-variable)


## Create Test:

```
kubectl testkube create test --name argocd --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/argocd/argocd.json
```

## Run Test:

```
kubectl testkube run test argocd -s ARGOCD_PASS=""
```
