For simplicity, the CYK program here require grammar to be Chomsky normal form(CNF),
and cannot transform context-free grammar to CNF automatically. 

step1:
build a recognition table;

step2:
use unger-style method to parse input string according to the recognition table;

time complexity: O(n^3)

Ref:
Parsing Techniques - A Practical Guide
Dick Grune and Ceriel J.H. Jacobs
VU University Amsterdam, Amsterdam, The Netherlands
Published (2008) by Springer US, ISBN 978-1-4419-1901-4 
