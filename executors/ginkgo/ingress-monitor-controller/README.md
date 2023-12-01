# Name: IngressMonitorController

## TestKube Type: Golang/Ginkgo

## Verifications:

- It checks whether the deployment is ready for usage.

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable       | Default Value | Allowed Values    | Description                                          |
|----------------|---------------|-------------------|------------------------------------------------------|
| PROFILE_ACTIVE | kubernetes    | kubernetes, local | Specifies the active profile for test configuration. |
| TEST_TIMEOUT   | 1m            |                   | Sets the timeout duration for the tests.             |
| NAMESPACE      | tks-system    |                   | The namespace where the addon was installed.         |

## Running the Tests

To run the IngressMonitorController tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "ingress-monitor-controller-test" \
      --description "ingress-monitor-controller-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/ingress-monitor-controller-testing" \
      --git-path "executors/ginkgo/ingress-monitor-controller" \
      --namespace "testkube"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "ingress-monitor-controller-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the IngressMonitorController tests within the specified namespace (testkube) using the TestKube
framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
