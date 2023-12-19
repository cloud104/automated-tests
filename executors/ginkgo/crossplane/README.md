# Name: Crossplane

## TestKube Type: Golang/Ginkgo

## Verifications:

Feature: Installation of Kubernetes Provider

Scenario: Applying Provider Manifests
    Given the Kubernetes Provider manifests
    When they are applied
    Then it should eventually be in a healthy state
    And RBAC manifests should be applied
    And ProviderConfig manifests should be applied

Scenario: Successful Installation of Kubernetes Provider
    Given that the Kubernetes Provider is installed
    When Crossplane is used to create Kubernetes resources
    Then it should eventually have 1 ready pod

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable         | Default Value                      | Allowed Values    | Description                                               |
|------------------|------------------------------------|-------------------|-----------------------------------------------------------|
| PROFILE_ACTIVE   | kubernetes                         | kubernetes, local | Specifies the active profile for test configuration.       |
| TEST_SKIP_DELETE | false                              | true, false       | Indicates whether to skip deletion of test resources.      |
| TEST_TIMEOUT     | 1m                                 |                   | Sets the timeout duration for the tests.                   |

## Running the Tests

To run the Crossplane tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "crossplane-test" \
      --description "crossplane-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/crossplane-testing" \
      --git-path "executors/ginkgo/crossplane" \
      --namespace "tks-system"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "crossplane-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the Crossplane tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
