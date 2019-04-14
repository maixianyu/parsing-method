Pushdown automato require grammar to be Greibach Normal Form (GNF) :
<br/>A -> a
<br/>or
<br/>A -> aB1B2...Bn

<br/>a sample of GNF here is :
<br/>S -> A B | D C
<br/>A -> a | a A
<br/>B -> b c | b B c
<br/>D -> a b | a D b
<br/>C -> c | c C

<br/>an example of input:
<br/>aabc

###breadthFirst method:
<br/>λ go test parse.go parse_test.go -v
<br/> parse_test.go:133: [S A a A a B b c #]
<br/>
<br/>λ go test parse.go performance_test.go -v
<br/>=== RUN   TestPerformance
<br/>--- PASS: TestPerformance (1.81s)
<br/>    performance_test.go:23:   100000         16337 ns/op
<br/>PASS
<br/>ok      command-line-arguments  2.260s

###backtracking method:
<br/>=== RUN   TestParse
<br/>--- PASS: TestParse (0.00s)
<br/>    parse_test.go:133: [S A a A a B b c #]
<br/>PASS
<br/>
<br/>λ go test parse.go performance_test.go -cpuprofile cpu.prof -v
<br/>=== RUN TestPerformance
<br/>--- PASS: TestPerformance (1.67s)
<br/>performance_test.go:23: 300000 5372 ns/op
<br/>PASS
<br/>ok command-line-arguments 2.082s