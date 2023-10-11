/* A Displacement function which take input from use and calculates the displacement over a given time . s = displacement , t= time, a = acceleration, v0 = initial velocity, s0 = initial displacement.  S =1/2 at^2 + v0t+s0
*/


package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a , v0 ,s0 float64)func (float64) float64{
		return func(t float64) float64{
			return 0.5 * a * math.Pow(t,2) + v0 *t + s0
		}
}	

func main(){
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&acceleration)

	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&initialVelocity)

	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&initialDisplacement)

	fmt.Print("Enter time (t): ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	displacement := fn(time)
	fmt.Printf("The displacement after %.2f seconds is %.2f\n", time, displacement)
}