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

### backtracking
Backtracking is a property of the parser, not of the implementation.
Both depthFirst method and recursive descent method are of backtracking.

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

###depthFirst method:
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

###recursiveDescent method:
<br/>λ go test parse.go parse_test.go  -v
<br/>=== RUN   TestParse
<br/>--- PASS: TestParse (0.00s)
<br/>    parse_test.go:16: a a b c [[S -> AB A -> aA A -> a B -> bc]]
<br/>    parse_test.go:25: a b c c [[S -> DC D -> ab C -> cC C -> c]]
<br/>PASS
<br/>ok      command-line-arguments  0.380s
<br/>
<br/>λ go test parse.go performance_test.go -v -cpuprofile=cpu.prof
<br/>=== RUN   TestPerformance
<br/>--- PASS: TestPerformance (2.14s)
<br/>    performance_test.go:16:  2000000           733 ns/op
<br/>PASS
<br/>ok      command-line-arguments  3.009s