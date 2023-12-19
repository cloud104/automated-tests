# Name: NodeLocalDNS

## TestKube Type: Golang/Ginkgo

## Verifications:

```
Feature: NodeLocalDNS Addon Installation Check

  Scenario: NodeLocalDNS DaemonSet Check
    Given the NodeLocalDNS addon is installed
    When the NodeLocalDNS DaemonSet is checked
    Then it should have exactly one DaemonSet for NodeLocalDNS

  Scenario: NodeLocalDNS Pod Health Check
    Given the NodeLocalDNS addon is installed
    When the health of NodeLocalDNS pods on each node is checked
    Then it is expected that all pods on each node should be healthy
```

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable       | Default Value | Allowed Values    | Description                                          |
|----------------|---------------|-------------------|------------------------------------------------------|
| PROFILE_ACTIVE | kubernetes    | kubernetes, local | Specifies the active profile for test configuration. |
| TEST_TIMEOUT   | 1m            |                   | Sets the timeout duration for the tests.             |
| NAMESPACE      | kube-system   |                   | The namespace where the addon was installed.         |

## Running the Tests

To run the NodeLocalDNS tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "node-local-dns-test" \
      --description "node-local-dns-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/node-local-dns-testing" \
      --git-path "executors/ginkgo/node-local-dns-operator" \
      --namespace "tks-system"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "node-local-dns-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the NodeLocalDNS tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
