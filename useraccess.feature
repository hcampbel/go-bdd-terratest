Feature: User Object Storage
  Scenario: create object store
    When I create new bucket my-gobdd-terratest-bucket
    Then the my-gobdd-terratest-bucket creation succeeded