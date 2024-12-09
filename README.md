# Potential Integer Overflow in Go Function

This repository demonstrates a potential integer overflow bug in a simple Go function and its solution.

The function `myFunc` adds two integers. If the sum exceeds the maximum value representable by the `int` type, an overflow occurs, leading to unexpected results.

The solution introduces type checking and error handling to mitigate this issue.
