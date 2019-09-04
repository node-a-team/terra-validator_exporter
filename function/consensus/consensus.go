package consensus

import (
        t "github.com/node-a-team/terra-validator_exporter/types"

        "os/exec"
        "encoding/json"
)


var (
	consensusStatus t.ConsensusStatus
)

func ConsensusStatus() t.ConsensusStatus {
        cmd := "curl -s -XGET " +t.RpcServer +"/consensus_state"  +" -H \"accept:application/json\""
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
        json.Unmarshal(out, &consensusStatus)

	return consensusStatus
}
