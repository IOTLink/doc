# Copyright IBM Corp. 2017 All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

Feature: Peer Service
    As a user I want to be able have channels and chaincodes to execute


Scenario Outline: FAB-3505: Test chaincode example02 deploy, invoke, and query
  Given I have a bootstrapped fabric network of type <type>
  When a user deploys chaincode at path "github.com/hyperledger/fabric/chaincode_example02" with ["init", "a", "1000" , "b", "2000"] with name "mycc"
  Then the chaincode is deployed
  When a user queries on the chaincode named "mycc" with args ["query", "a"]
  Then a user receives expected response of 1000
  When a user invokes on the chaincode named "mycc" with args ["txId1", "invoke", "a", 10]
  And a user queries on the chaincode named "mycc" with args ["query", "a"]
  Then a user receives expected response of 990

  Given "peer1" is taken down
  When a user invokes on the chaincode named "mycc" with args ["txId1", "invoke", "a", 10]
  And I wait "15" seconds
  And "peer1" comes back up
  When a user queries on the chaincode named "mycc" with args ["query", "a"] on "peer1"
  Then a user receives expected response is 980
  Examples:
    | type  |
    | solo  |
    | kafka |


Scenario Outline: FAB-1440: Test basic chaincode deploy, invoke, query
  Given I have a bootstrapped fabric network of type <type>
  When a user deploys chaincode
  Then the chaincode is deployed
  When a user queries on the chaincode
  Then a user receives expected response
  Examples:
    | type  |
    | solo  |
    | kafka |


@skip
Scenario: FAB-3861: Basic Chaincode Execution (example02)
    Given I have a bootstrapped fabric network
    When a user deploys chaincode
    Then the chaincode is deployed


@skip
Scenario: FAB-3865: Multiple Channels Per Peer
    Given I have a bootstrapped fabric network
    When a user deploys chaincode
    Then the chaincode is deployed


@skip
Scenario: FAB-3866: Multiple Chaincodes Per Peer
    Given I have a bootstrapped fabric network
    When a user deploys chaincode
    Then the chaincode is deployed
