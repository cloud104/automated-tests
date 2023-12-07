# Name: Spot Price Calculator

## TestKube Type: postman/collection

## Verifications:

- Checks if the metrics endpoint is as expected. (Expected status code: "200") 

## Endpoints:

- http://spot-price-calculator.tks-system:8080/metrics

### Environment Variables

The following variables are available for configuring the service:

| Variable     | Default Value         | Description                                  |
|--------------|-----------------------|----------------------------------------------|
| SERVICE_NAME | spot-price-calculator | Specifies the name of the service.           |
| NAMESPACE    | tks-system            | The namespace where the service is deployed. |
| PORT         | 8080                  | The port on which the service is accessible. |

## Create Test:

```
kubectl testkube create test --name spot-price-calculator --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/spot-price-calculator/spot-price-calculator.json
```

## Run Test:

```
kubectl testkube run test spot-price-calculator
```
