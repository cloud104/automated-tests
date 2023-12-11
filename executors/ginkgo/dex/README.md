# Name: Dex

## TestKube Type: Golang/Ginkgo

## Verifications:

```
Feature: Dex Addon Deployment Check

  Scenario: Dex Addon Deployment Status Check
    Given the Dex addon is installed
    When the Dex deployments are checked
    Then it is expected that all Dex deployments should be ready for usage
```

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable         | Default Value | Allowed Values    | Description                                           |
|------------------|---------------|-------------------|-------------------------------------------------------|
| PROFILE_ACTIVE   | kubernetes    | kubernetes, local | Specifies the active profile for test configuration.  |
| TEST_TIMEOUT     | 1m            |                   | Sets the timeout duration for the tests.              |
| NAMESPACE        | tks-system    |                   | The namespace where the addon was installed.          |

## Running the Tests

To run the Dex tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "dex-test" \
      --description "dex-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/dex-testing" \
      --git-path "executors/ginkgo/dex-operator" \
      --namespace "testkube"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "dex-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the Dex tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
