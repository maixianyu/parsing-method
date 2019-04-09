package common

import(
	"fmt"
	"runtime"
	"reflect"
)

/*get name of funtion i*/
func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

/* a function to compare []string */
func StringSliceEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		fmt.Printf("len(s1)=%d, len(s2)=%d\n", len(s1), len(s2))
		return false
	} else if (s1 == nil) != (s2 == nil) {
		fmt.Printf("s1=%v, s2=%v\n", s1, s2)
		return false
	}

	for idx, _ := range s1 {
		if s1[idx] != s2[idx] {
			fmt.Printf("idx=%d, s1[idx]=%v, s2[idx]=%v\n", idx, s1[idx], s2[idx])
			return false
		}
	}

	return true
}


/* a function to compare [][]string */
func StringSSEqual(res [][]string, expect [][]string) bool {
	if len(res) != len(expect) {
		return false
	}
	for idx, _ := range expect {
		if StringSliceEqual(res[idx], expect[idx]) == false {
			fmt.Printf("idx=%d, res[idx]=%v, res[idx]=%v\n", idx, res[idx], expect[idx])
			return false
		}
	}
	return true
}


/* a function to compare [][][]string */
func StringSSSEqual(res [][][]string, expect [][][]string) bool {
	if len(res) != len(expect) {
		return false
	}
	for idx, _ := range expect {
		if StringSSEqual(res[idx], expect[idx]) == false {
			fmt.Printf("idx=%d, res[idx]=%v, res[idx]=%v\n", idx, res[idx], expect[idx])
			return false
		}
	}
	return true
}