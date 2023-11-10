# VaultOperator Tests

This repository contains automated tests for the VaultOperator using TestKube.

## Running the Tests

To run the VaultOperator tests using TestKube, follow the steps below:

1. **Create the Test:**
    ```bash
    kubectl testkube create test \
      --name "vault-operator-test" \
      --description "vault-operator-test" \
      --type "ginkgo/test" \
      --test-content-type "git-file" \
      --git-uri "https://github.com/cloud104/automated-tests" \
      --git-branch "feature/vault-operator-testing" \
      --git-path "executors/ginkgo/vault-operator" \
      --namespace "testkube"
    ```

2. **Run the Test:**
    ```bash
    kubectl testkube run test "vault-operator-test" --image "kurtis/testkube-executor-ginkgo:1.15.16" --namespace "testkube"
    ```

This will initiate the VaultOperator tests within the specified namespace (testkube) using the TestKube framework.

Make sure to replace the values for the Git URI, Git Branch, and other parameters according to your specific test
environment. Additionally, ensure that the required images and dependencies are available in your environment.
