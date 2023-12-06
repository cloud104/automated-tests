# Name: Falco

## TestKube Type: Golang/Ginkgo

## Verifications:

```
Feature: Falco is Installed and Healthy

  Scenario: Falco Installation Check
    Given the Falco DaemonSet is installed
    When the health of Falco pods on each node is checked
    Then it is expected that all Falco pods on each node should be healthy

  Scenario: Falco Exporter Installation Check
    Given the Falco Exporter DaemonSet is installed
    When the health of Falco Exporter pods on each node is checked
    Then it is expected that all Falco Exporter pods on each node should be healthy
```

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable         | Default Value | Allowed Values    | Description                                           |
|------------------|---------------|-------------------|-------------------------------------------------------|
| PROFILE_ACTIVE   | kubernetes    | kubernetes, local | Specifies the active profile for test configuration.  |
| TEST_TIMEOUT     | 1m            |                   | Sets the timeout duration for the tests.              |
| NAMESPACE        | tks-system    |                   | The namespace where the addon was installed.          |

## Running the Tests

To run the Falco tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "falco-test" \
      --description "falco-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/falco-testing" \
      --git-path "executors/ginkgo/falco-operator" \
      --namespace "testkube"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "falco-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the Falco tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
