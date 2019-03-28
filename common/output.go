package common

import (
	"log"
	"strings"
)

/* CombTraceWithTemplate combine trace and rhSide to produce a result []string */
/* rhSide = []string{"Term", "+", "Factor"} */
/* trace =  [][]string{
	[]string{"TermxFactor", "FactorxFactor", "ixFactor", "ixi"},
	[]string{"+"},
	[]string{"i"},
} */
/* res = []string{
	"Term+Factor"
	"TermxFactor+Factor"
	"FactorxFactor+Factor"
	"ixFactor+Factor"
	"ixi+Factor"
	"ixi+i"
} */
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

func AppendString2StrSlice(input []string, suffix string) []string {
	res := []string{}
	for _, t := range input {
		res = append(res, t + suffix)
	}
	return res
}