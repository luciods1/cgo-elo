This small project binds C code to Golang to calculate elo related operations.

Available operations:

* Probability of winning is calculated by the following logistic function:

<img src="https://latex.codecogs.com/svg.latex?\Large&space;p(w)=\frac{1}{10^{-(a-b)/400}+1}" title="\Large p(w)=\frac{1}{10^{-(a-b)/400}+1}" />
<br>

* Elo differential calculation is made using the current player score, the probability of wining against a player B, growth factor (a value that determine how easily our population will growth):

<img src="https://latex.codecogs.com/svg.latex?\Large&space;S=c-p(w)k" title="\Large S=c-p(w)k" />
<br>

Added benchmarks to check the performance between C and Go in this arithmetic operations.
You can run them by executing:
```
BIN_NAME=elo make benchmark
``` 
