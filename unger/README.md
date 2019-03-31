unger method:
    top-down
    depth-first search

time complexity: O(C^N)

Test in sample/arithmetic, with input "(i+i)xi":<br/>
    Expr -><br/>
    Term -><br/>
    TermxFactor -><br/>
    FactorxFactor -><br/>
    (Expr)xFactor -><br/>
    (Expr+Term)xFactor -><br/>
    (Term+Term)xFactor -><br/>
    (Factor+Term)xFactor -><br/>
    (i+Term)xFactor -><br/>
    (i+Factor)xFactor -><br/>
    (i+i)xFactor -><br/>
    (i+i)xi -><br/>

Ref:<br/>
Parsing Techniques - A Practical Guide<br/>
Dick Grune and Ceriel J.H. Jacobs<br/>
VU University Amsterdam, Amsterdam, The Netherlands<br/>
Published (2008) by Springer US, ISBN 978-1-4419-1901-4<br/> 