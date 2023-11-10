# Name: VaultOperator

## TestKube Type: Golang/Ginkgo

## Verifications:

Feature: Installing Vault using the Operator

  Scenario: Applying the Operator CRD
    Given the Operator CRD is applied
    Then eventually, there should be 1 ready pod

  Scenario: Completing the installation process
    Given the installation process is complete
    When creating a new secret to Vault
    Then the new secret should be accessible

### Environment Variables

The following environment variables are available for configuring the tests:

| Variable         | Default Value                      | Allowed Values    |
|------------------|------------------------------------|-------------------|
| PROFILE_ACTIVE   | kubernetes                         | kubernetes, local |
| TEST_SKIP_DELETE | false                              | true, false       |
| TEST_TIMEOUT     | 1m                                 |                   |
| VAULT_ADDRESS    | http://vault-test.<namespace>:8200 |                   |
| VAULT_USERNAME   | <generated>                        |                   |
| VAULT_PASSWORD   | <generated>                        |                   |

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
