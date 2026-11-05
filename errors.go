package pumpcalc

import "errors"

var DivideByZeroError = errors.New("Can't divide by zero!")
var NegativeDensityError = errors.New("Density can't be negative")
var NegativeMassFlowKgSError = errors.New("MassFlowKg can't be negative")
var NegativeVelocityError = errors.New("Velocity can't be negative in this context")
var NegativeKinematicViscosityError = errors.New("Kinematic viscosity can't be negative")
var NotPositiveFrictionCoefficientError = errors.New("Friction coefficient can't be zero or negative")
var NotPositiveGravityError = errors.New("Gravity can't be zero or negative")
