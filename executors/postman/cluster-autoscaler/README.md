# Name: cluster-autoscaler

## TestKube Type: postman/collection

## Verifications:

- Verify the health of the Cluster Autoscaler

## Create Test:

```
kubectl testkube create test \
  --name cluster-autoscaler \
  --type postman/collection \
  --test-content-type git-file \
  --job-template templates/worker-system-node-pool-job.yaml \
  --git-uri https://github.com/cloud104/automated-tests.git \
  --git-branch feature/cluster-autoscaler-testing \
  --git-path executors/postman/cluster-autoscaler/cluster-autoscaler.json \
```

## Run Test:

```
kubectl testkube run test cluster-autoscaler
```
