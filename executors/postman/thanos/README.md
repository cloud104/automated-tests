# Name: thanos

## TestKube Type: postman/collection

## Verifications:

- Check the status of the Thanos Compactor
- Verify the health of the Thanos Query

## Create Test:

```
kubectl testkube create test \
  --name thanos \
  --type postman/collection \
  --test-content-type git-file \
  --job-template templates/worker-system-node-pool-job.yaml \
  --git-uri https://github.com/cloud104/automated-tests.git \
  --git-branch feature/thanos-testing \
  --git-path executors/postman/thanos/thanos.json
```

## Run Test:

```
kubectl testkube run test thanos
```
