package common

import (
	"log"
	"strings"
)

func CombTraceWithTemplate(traces [][]string, rhSide RightHandSide) []string {
	/* copy rhSide to tmpl, avoiding make a change to rhSide*/
	tmpl := make([]string, len(rhSide))
	if copy(tmpl, rhSide) == 0 {
		log.Fatalf("fail to copy from rhSide(%v) to tmpl", rhSide)
	}

	/* use template to generate output with traces */
	res := []string{}
	for idx, trace := range traces {
		var s string
		for _, s = range trace {
			subRes := make([]string, len(rhSide))
			copy(subRes, tmpl)
			subRes[idx] = s
			str := strings.Join(subRes, "")
			if len(res) == 0 || str != res[len(res)-1] {
				res = append(res, str)
			}
		}
		tmpl[idx] = s
	}

	return res
}