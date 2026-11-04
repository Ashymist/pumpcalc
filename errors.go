package pumpcalc

import "errors"

var DivideByZeroError = errors.New("Can't divide by zero!")
var NegativeMassFlowKgSError = errors.New("MassFlowKg can't be negative")
var NegativeVelocityError = errors.New("Velocity can't be negative in this context")
