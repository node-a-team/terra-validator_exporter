# terra-validator_exporter :satellite:
![CreatePlan](https://img.shields.io/badge/relase-v0.1.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.12.4%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)

Prometheus exporter for Terra Validators


## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter (localhost:26660), and other specific information monitoring purposes


## Collecting information list
> **Network**
- chainId: Name of the chain
- blockHeight: Height of the current block
- currentBlockTime: Time it takes to create & confirm block (current block time - previous block time)
- bondedTokens: Number of currently bonded Luna
- notBondedTokens: Number of unbonded Luna
- totalBondedTokens: Number of currently bonded & unbonded Luna
- bondedRate: Ratio of bonded tokens within the network
- validatorCount: Number of validators within the network
- precommitRate: Precommit Ratio of precommits collected in a round
- proposerWalletAccountNumber: Account number given on each validator’s wallet (Required to show Proposer in Grafana) 

> **Validator Info**
- moniker: Name of the validator
- accountAddress: Validator's Account address
- consHexAddress: Validator's Consensus Hex address
- operatorAddress: Validator's Operator address
- validatorPubKey: Validator's Validator pubkey(```terrad tendermint show-validator```)
- votingPower: Decimal truncated Total voting power of the validator
- delegatorShares: Validator's total delegated tokens
- delegatorCount: Number of each unique delegators for a validator
- delegationRatio: Ratio of validator's bonded tokens to the network's total bonded tokens
- selfDelegationAmount: Self-bonded amount of the validator
- proposerPriorityValue: Number which represents the priority of the validator proposing in the next round
- proposerPriority: Rank of the proposerPriorityValue
- proposingStatus: Shows if the validator is the proposer or not in the current round(true: 1, false: 0)
- validatorCommitStatus: Confirms if the validator has committed in this round(true: 1, false: 0)
- commissionMaxChangeRate: Max range of commission rate whic hthe validator can change
- commissionMaxRate: The highest commission rate which the validator can charge
- commissionRate: Commission rate of the validator charged on delegators' rewards
- balances(luna, krw, sdr, usd): Wallet information of the validator which shows the balance
- commission(luna, krw, sdr, usd): Accumulated commission fee of the validator
- rewards(luna, krw, sdr, usd): Accumulated rewards of the validator
- minSelfDelegation: The required minimum number of tokens which the validator must self-delegate
- jailed: Confirms if the validator is jailed or not(true: 1, false: 0)

![CreatePlan](./example/monitoring_example(prometheus).png)


## Quick Start
RPC and REST server information is required to run the program
- Download
```
wget https://github.com/node-a-team/terra-validator_exporter/releases/download/v0.1.0/terra-validator_exporter_v0.1.0.tar.gz
tar -xzvf terra-validator_exporter_v0.1.0.tar.gz &&  cd terra-validator_exporter
```

 - Config Setup
 1) Input RPC and Rest server information
 2) Input validator operator address(```terracli keys show [Key Name] --bech=val --address```)
 3) Set exporter port
 4) Set outPrint (if true: prints collected information from the exporter)
```
vi config.toml
```
```
# TOML Document for Terra-Validator Exporter(Pometheus & Grafana)

title = "Terra-Validator Exporter TOML"
network = "terra"

# RPC-Server
[rpc]
address = "localhost:26657"

[rest_server]
address = "localhost:1317"

[validator_info]
operatorAddress = ""

[option]
exporterListenPort = "26661"
outputPrint = true
```

![CreatePlan](./example/config.png)

 - Run
```
./terra-validator_exporter
```

![CreatePlan](./example/config_outputPrint(true).png)


## Grafana Example
Can set alarms using the functions on Grafana (ex. Alarms if the validator fails to precommit or gets jailed)
![CreatePlan](./example/monitoring_example(grafana).png)

