package metric

import (

//	"fmt"
//	"encoding/hex"
//	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/prometheus/client_golang/prometheus"
)

var (

)

func NewGauge(nameSpace string, name string, help string) prometheus.Gauge {
	result := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "" + nameSpace,
			Name:      "" + name,
			Help:      "" + help,
		},
	)

	return result
}
/*
func NewCounterVec(nameSpace string, name string, help string) prometheus.CounterVec {
        result := prometheus.NewCounterVec(
                prometheus.GaugeOpts{
                        Namespace: "" + nameSpace,
                        Name:      "" + name,
                        Help:      "" + help,
                },
		[]string("chainID", "moniker", "operAddr", "accAddr", "consHexAddr"),
	)

        return result
}
*/

