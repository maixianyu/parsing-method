Chp 6. Regular grammars and finite-state automata

DFA - Deterministic Finite-state Automaton

The implementation is almost a copy version in Go from Ross Cox's blog:<br/>
https://swtch.com/~rsc/regexp/regexp1.html<br/>
This blog makes a very explict comment on the theory and implementation in C.<br/>

<br/>λ go test performance_test.go parse.go -v
<br/>=== RUN   TestPerformance
<br/>--- PASS: TestPerformance (18.52s)
<br/>    performance_test.go:34: textSize=10
<br/>    performance_test.go:35:   300000          5282 ns/op
<br/>    performance_test.go:34: textSize=20
<br/>    performance_test.go:35:   200000         11528 ns/op
<br/>    performance_test.go:34: textSize=30
<br/>    performance_test.go:35:   100000         19817 ns/op
<br/>    performance_test.go:34: textSize=40
<br/>    performance_test.go:35:    50000         29381 ns/op
<br/>    performance_test.go:34: textSize=50
<br/>    performance_test.go:35:    30000         40040 ns/op
<br/>    performance_test.go:34: textSize=60
<br/>    performance_test.go:35:    30000         53188 ns/op
<br/>    performance_test.go:34: textSize=70
<br/>    performance_test.go:35:    20000         67491 ns/op
<br/>    performance_test.go:34: textSize=80
<br/>    performance_test.go:35:    20000         83022 ns/op
<br/>    performance_test.go:34: textSize=90
<br/>    performance_test.go:35:    10000        100520 ns/op
<br/>    performance_test.go:34: textSize=100
<br/>    performance_test.go:35:    10000        119081 ns/op
<br/>PASS
<br/>ok      command-line-arguments  19.231s