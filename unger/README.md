unger method:
    top-down
    depth-first search

time complexity: O(C^N)

Test in sample/arithmetic, with input "(i+i)xi":
    Expr ->
    Term ->
    TermxFactor ->
    FactorxFactor ->
    (Expr)xFactor ->
    (Expr+Term)xFactor ->
    (Term+Term)xFactor ->
    (Factor+Term)xFactor ->
    (i+Term)xFactor ->
    (i+Factor)xFactor ->
    (i+i)xFactor ->
    (i+i)xi ->

Ref:
Parsing Techniques - A Practical Guide
Dick Grune and Ceriel J.H. Jacobs
VU University Amsterdam, Amsterdam, The Netherlands
Published (2008) by Springer US, ISBN 978-1-4419-1901-4 