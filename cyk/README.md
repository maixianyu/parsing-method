For simplicity, the CYK program here require grammar to be Chomsky normal form(CNF),
and cannot transform context-free grammar to CNF automatically. 

step1:
build a recognition table;

step2:
use unger-style method to parse input string according to the recognition table;

time complexity: O(n^3)

Test in sample/number, with input "12.34+56":<br/>
    Number -><br/>
    N1Scale -><br/>
    IntegerFractionScale -><br/>
    IntegerDigitFractionScale -><br/>
    1DigitFractionScale -><br/>
    12FractionScale -><br/>
    12T1IntegerScale -><br/>
    12.IntegerScale -><br/>
    12.IntegerDigitScale -><br/>
    12.3DigitScale -><br/>
    12.34Scale -><br/>
    12.34N2Integer -><br/>
    12.34T2SignInteger -><br/>
    12.34eSignInteger -><br/>
    12.34e+Integer -><br/>
    12.34e+IntegerDigit -><br/>
    12.34e+5Digit -><br/>
    12.34e+56 -><br/>

Ref:<br/>
Parsing Techniques - A Practical Guide<br/>
Dick Grune and Ceriel J.H. Jacobs<br/>
VU University Amsterdam, Amsterdam, The Netherlands<br/>
Published (2008) by Springer US, ISBN 978-1-4419-1901-4<br/>
