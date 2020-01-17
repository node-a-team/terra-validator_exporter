# terra-validator_exporter :satellite:
![CreatePlan](https://img.shields.io/badge/relase-v0.3.1-red)
![CreatePlan](https://img.shields.io/badge/go-1.13.1%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)

Prometheus exporter for Terra Validators

## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter (localhost:26660), and other specific information monitoring purposes


## Collecting information list
- blockHeight: Height of the current block
- bondedTokens: Number of currently bonded Luna
- notBondedTokens: Number of unbonded Luna
- bondedRatio: Ratio of bonded tokens within the network
- totalSupply: Total Supply of Luna

- precommitStatus: Check precommit for previous block
- commitVoteType: Check commit vote type for previous block(false: 0, prevote: 1, precommit: 2)
- jailStatus: Confirms if the validator is jailed or not(true: 1, false: 0)

- votingPower: Decimal truncated Total voting power of the validator
- delegatorCount: Number of each unique delegators for a validator
- delegatorShares: Validator's total delegated tokens
- delegationRatio: Ratio of validator's bonded tokens to the network's total bonded tokens
- selfDelegationAmount: Self-bonded amount of the validator
- minSelfDelegation: The required minimum number of tokens whic hthe validator must self-delegate

- balances(uluna, ukrw, usdr, uusd, umnt): Wallet information of the validator which shows the balance
- commission(uluna, ukrw, usdr, uusd, umnt): Accumulated commission fee of the validator
- rewards(uluna, ukrw, usdr, uusd, umnt): Accumulated rewards of the validator

- commissionMaxRate: The highest commission rate which the validator can charge
- commissionMaxChangeRate: Max range of commission rate whic hthe validator can change
- commissionRate: Commission rate of the validator charged on delegators' rewards

- proposerRanking: Ranking to become a proposer
- proposerStatus: Shows if the validator is the proposer or not in the current round(true: 1, false: 0)

- totalProposalCount: Total number of proposals
- votingProposalCount: Proposal that is voting period

- oracleMiss: Check the missed oracle count

- labels: moniker, chainId, accountAddress, operatorAddress, consHexAddress, 




## Install
```bash
mkdir exporter && cd exporter

wget https://github.com/node-a-team/terra-validator_exporter/releases/download/v0.3.1/terra-validator_exporter.tar.gz  && sha256sum terra-validator_exporter.tar.gz | fgrep 1221905e7983dfc317e768099eacbb319d9eff39117503deb917218011907477 && tar -zxvf terra-validator_exporter.tar.gz ||  echo "Bad Binary!"
```

## Config
1. Modify to the appropriate RPC and REST server address
2. Modify the value of ```operatorAddr``` to the operator address of the validator you want to monitor.
3. You can change the service port(default: 26661)
```bash
vi config.toml
```
```bash
# TOML Document for Terra-Validator Exporter(Pometheus & Grafana)

title = "TOML Document"

[Servers]
        [Servers.addr]
        rpc = "localhost:26657"
        rest = "localhost:1317"

[Validator]
operatorAddr = "terravaloper1tusfpgvjrplqg2fm7wacy4slzjmnzswcfufuvp"

[Options]
listenPort = "26661"

```

## Start
  
```bash
./terra-validator_exporter {path to config.toml}

// ex)
./terra-validator_exporter /data/terra/exporter
```

## Use systemd service
  
```sh
# Make log directory & file
sudo mkdir /var/log/userLog  
sudo touch /var/log/userLog/terra-validator_exporter.log  
# user: iris
sudo chown terra:terra /var/log/userLog/terra-validator_exporter.log

# $HOME: /data/terra
# Path to config.toml: /data/terra/exporter
sudo tee /etc/systemd/system/terra-validator_exporter.service > /dev/null <<EOF
[Unit]
Description=Terra Validator Exporter
After=network-online.target

[Service]
User=terra
WorkingDirectory=/data/terra
ExecStart=/data/terra/exporter/terra-validator_exporter \
        /data/terra/exporter
StandardOutput=file:/var/log/userLog/terra-validator_exporter.log
StandardError=file:/var/log/userLog/terra-validator_exporter.log
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable terra-validator_exporter.service
sudo systemctl start terra-validator_exporter.service


## log
tail -f /var/log/userLog/terra-validator_exporter.log
```

## Example
https://www.notion.so/wlsaud619/terra-validator_exporter-1e9c6cf1bdb0483180829676b533565b
