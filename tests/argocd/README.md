# Name: check-argocd

## TestKube Type: postman/collection

### CRD: 

- ../../executors/postman/crd-argocd.yaml

## Verifications:

- Verify login attempt using access credentials and get session token. ( Expected: "200" )

- Check if the configured repositories have a successful connection. ( Expected: "Connection: Successful" )
 
## URL:

- Service: http://argo-cd-argocd-server.tks-system.svc.cluster.local

## Endpoints:

- /api/v1/session (requires authentication)
- /api/v1/repositories (requires authentication)

## Variables:

- USER string
- PASS string

## Create Test:

```
kubectl testkube create test --name check-argocd --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path tests/argocd/argocd.json
```
## Run Test:

```
kubectl testkube run test check-argocd --secret-variable USER=$USERNAME --secret-variable PASS=$PASSWORD
```