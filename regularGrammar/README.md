Chp 6. Regular grammars and finite-state automata

DFA - Deterministic Finite-state Automaton

The implementation is almost a copy version in Go from Ross Cox's blog:<br/>
https://swtch.com/~rsc/regexp/regexp1.html<br/>
This blog makes a very explict comment on the theory and implementation in C.<br/>

<br/>
###testcase1
<br/>regexp: (a?)^na^n
<br/>input: a^n
<br/>command: go test DFA.go NFA.go performance_test.go -v

<br/>DFA:
<br/>performance_test.go:49: textSize=10
<br/>performance_test.go:50: 100000 13069 ns/op
<br/>performance_test.go:49: textSize=20
<br/>performance_test.go:50: 50000 27894 ns/op
<br/>performance_test.go:49: textSize=30
<br/>performance_test.go:50: 30000 45780 ns/op
<br/>performance_test.go:49: textSize=40
<br/>performance_test.go:50: 20000 67933 ns/op
<br/>performance_test.go:49: textSize=50
<br/>performance_test.go:50: 20000 94010 ns/op

<br/>NFA:
<br/>performance_test.go:49: textSize=10
<br/>performance_test.go:50: 300000 5749 ns/op
<br/>performance_test.go:49: textSize=20
<br/>performance_test.go:50: 100000 12464 ns/op
<br/>performance_test.go:49: textSize=30
<br/>performance_test.go:50: 100000 21595 ns/op
<br/>performance_test.go:49: textSize=40
<br/>performance_test.go:50: 50000 31835 ns/op
<br/>performance_test.go:49: textSize=50
<br/>performance_test.go:50: 30000 43636 ns/op

<br/>
###testcase2
<br/>regexp: 0(a|b|c|d|...|z)1
<br/>input: 0abcdabcdabcda...abcd1
<br/>command: go test DFA.go NFA.go performance_test.go -v
<br/>
<br/>performance_test.go:55: command-line-arguments.ParseDFA
<br/>performance_test.go:60: textSize=10
<br/>performance_test.go:61:   100000         12066 ns/op
<br/>performance_test.go:60: textSize=20
<br/>performance_test.go:61:   100000         11248 ns/op
<br/>performance_test.go:60: textSize=30
<br/>performance_test.go:61:   200000         11139 ns/op
<br/>performance_test.go:60: textSize=40
<br/>performance_test.go:61:   200000         12739 ns/op
<br/>performance_test.go:60: textSize=50
<br/>performance_test.go:61:   200000         12084 ns/op
<br/>performance_test.go:55: command-line-arguments.ParseNFA
<br/>performance_test.go:60: textSize=10
<br/>performance_test.go:61:   200000         10549 ns/op
<br/>performance_test.go:60: textSize=20
<br/>performance_test.go:61:   100000         12934 ns/op
<br/>performance_test.go:60: textSize=30
<br/>performance_test.go:61:   100000         15758 ns/op
<br/>performance_test.go:60: textSize=40
<br/>performance_test.go:61:   100000         20592 ns/op
<br/>performance_test.go:60: textSize=50
<br/>performance_test.go:61:   100000         22059 ns/op