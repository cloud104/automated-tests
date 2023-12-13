# Name: Metrics-Server

## TestKube Type: postman/collection

## Verifications:

- Check Metrics-server API is available. ( Expected: "Available" | "True" )
- Check if request metric endpoint return expected. ( Expected: "200")
 

## Endpoints:

- https://{{API_ADDRESS}}/apis/apiregistration.k8s.io/v1/apiservices/v1beta1.metrics.k8s.io
- https://{{API_ADDRESS}}/apis/metrics.k8s.io/v1beta1/nodes


## Environment Variables

The following variable is required to run the test on the application:

| Variable     | Description                                                      |
|--------------|------------------------------------------------------------------|
| API_ADDRESS  | Specify the cluster API address.                                 |
| SERVICE_ACCOUNT_TOKEN | Specify the service account token used to query on api. |


## Create Test:

```
kubectl testkube create test --name metrics-server --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/metrics-server/metrics-server.json
```

## Run Test:

```
kubectl testkube run test metrics-server -v API_ADDRESS="" -s SERVICE_ACCOUNT_TOKEN=""
```
