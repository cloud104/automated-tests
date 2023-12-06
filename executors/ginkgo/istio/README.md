# Name: Istio

## TestKube Type: Golang/Ginkgo

## Verifications:

Feature: Istio is up and running

  Scenario: IstioOperator CRD Health Check
    Given the IstioOperator CRD is applied
    When the Istio components are checked for health
    Then it is expected that all components should be healthy

  Scenario: Addon Installation Check
    Given the addon is installed
    When the egressgateway deployment is checked
    Then it should have an available egressgateway deployment

    When the ingressgateway deployment is checked
    Then it should have an available ingressgateway deployment

    When the istiod deployment is checked
    Then it should have an available istiod deployment

    When the kiali deployment is checked
    Then it should have an available kiali deployment

    When the kiali-oauth2-proxy deployment is checked
    Then it should have an available kiali-oauth2-proxy deployment

    When the prometheus deployment is checked
    Then it should have an available prometheus deployment

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable         | Default Value | Allowed Values    | Description                                           |
|------------------|---------------|-------------------|-------------------------------------------------------|
| PROFILE_ACTIVE   | kubernetes    | kubernetes, local | Specifies the active profile for test configuration.  |
| TEST_SKIP_DELETE | false         | true, false       | Indicates whether to skip deletion of test resources. |
| TEST_TIMEOUT     | 1m            |                   | Sets the timeout duration for the tests.              |
| NAMESPACE        | istio-system  |                   | The namespace where the addon was installed.          |

## Running the Tests

To run the Istio tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "istio-test" \
      --description "istio-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/istio-testing" \
      --git-path "executors/ginkgo/istio-operator" \
      --namespace "testkube"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "istio-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the Istio tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
