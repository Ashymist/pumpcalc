package pumpcalc

import "errors"

var DivideByZeroError = errors.New("Can't divide by zero!")
var NegativeMassFlowKgSError = errors.New("MassFlowKg can't be negative")
