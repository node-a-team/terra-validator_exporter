package consensus

import (
	t "terra-validator_exporter/types"

	"encoding/json"
	"os/exec"
)

var (
	consensusStatus t.ConsensusStatus
)

func ConsensusStatus() t.ConsensusStatus {
	cmd := "curl -s -XGET " + t.RpcServer + "/consensus_state" + " -H \"accept:application/json\""
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &consensusStatus)

	return consensusStatus
}
