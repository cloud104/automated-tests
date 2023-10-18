# Name: Tks-logs

## TestKube Type: postman/collection

## Verifications:

- Verify login attempt using access credentials.

- Elastic_cluster_status: Verify whether the cluster status returned by the api is as expected ( Expected: "green" )
 
- Kibana_status: Verify whether the cluster status returned by the api is as expected ( Expected: "green" )


## Endpoints:

ElasticSearch: https://tks-logs-es-http.tks-logs.svc.cluster.local:9200/_cluster/health/ (requires authentication)
Kibana: http://tks-logs-kb-http.tks-logs.svc.cluster.local:5601/api/stats (requires authentication)

## Variables:

- USER (--sercret-variable)
- PASS (--sercret-variable)


## Create Test:

```
kubectl testkube create test --name tks-logs --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/tks-logs/tks-logs.json
```

## Run Test:

```
kubectl testkube run test tks-logs --secret-variable USER="" --secret-variable PASS="" 
```
