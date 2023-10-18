# Name: Stack-Logs

## TestKube Type: postman/collection

## Verifications:

- Verify login attempt using access credentials. ( Expected: "200" )

- Elastic_cluster_status: Verify whether the cluster status returned by the api is as expected ( Expected: "green" )
 
- Kibana_status: Verify whether the cluster status returned by the api is as expected ( Expected: "green" )

## URL:

- Service Elastic: https://tks-logs-es-http.tks-logs.svc.cluster.local:9200
- Service Kibana: http://tks-logs-kb-http.tks-logs.svc.cluster.local:5601

## Endpoints:

- /_cluster/health/ (requires authentication) - ElasticSearch
- /api/stats (requires authentication) - Kibana

## Variables:

- USER string
- PASS string

## Create Test:

```
kubectl testkube create test --name stack-logs --type postman/collection --test-content-type git-file --git-uri https://github.com/cloud104/automated-tests.git --git-branch master --git-path executors/postman/logs/logs.json
```

## Run Test:

```
kubectl testkube run test stack-logs --secret-variable USER=$USERNAME --secret-variable PASS=$PASSWORD 
```
