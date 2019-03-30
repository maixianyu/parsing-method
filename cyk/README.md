For simplicity, the CYK program here require grammar to be Chomsky normal form(CNF),
and cannot transform context-free grammar to CNF automatically. 

step1:
build a recognition table;

step2:
use unger-style method to parse input string according to the recognition table;

time complexity: O(n^3)

Test in sample/number, with input "12.34+56":
    Number ->
    N1Scale ->
    IntegerFractionScale ->
    IntegerDigitFractionScale ->
    1DigitFractionScale ->
    12FractionScale ->
    12T1IntegerScale ->
    12.IntegerScale ->
    12.IntegerDigitScale ->
    12.3DigitScale ->
    12.34Scale ->
    12.34N2Integer ->
    12.34T2SignInteger ->
    12.34eSignInteger ->
    12.34e+Integer ->
    12.34e+IntegerDigit ->
    12.34e+5Digit ->
    12.34e+56 ->

Ref:
Parsing Techniques - A Practical Guide
Dick Grune and Ceriel J.H. Jacobs
VU University Amsterdam, Amsterdam, The Netherlands
Published (2008) by Springer US, ISBN 978-1-4419-1901-4 
