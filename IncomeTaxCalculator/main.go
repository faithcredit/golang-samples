package main

import (
	"fmt"
)

func main() {

	var max10 float64 = 9875
	var max12 float64 = 40125
	var max22 float64 = 85525
	var max24 float64 = 163300
	var max32 float64 = 207350
	var max35 float64 = 518400

	var tier10_tax float64 = max10 * .1
	var tier12_tax float64 = tier10_tax + ((max12 - max10) * .12)
	var tier22_tax float64 = tier12_tax + ((max22 - max12) * .22)
	var tier24_tax float64 = tier22_tax + ((max24 - max22) * .24)
	var tier32_tax float64 = tier24_tax + ((max32 - max24) * .32)
	var tier35_tax float64 = tier32_tax + ((max35 - max32) * .35)

	// ask user for the gross income
	fmt.Print("Enter your gross income from your W-2 for 2020:")

	var grossIncome float64
	fmt.Scanln(&grossIncome) // take gross income input from user

	fmt.Print("How many dependents are you claiming? ")
	var numDep int
	fmt.Scanln(&numDep) // take number of dependents input from user

	//calculate taxable income

	taxableIncome := grossIncome - 12200 - (2000 * float64(numDep))

	// calculate tax due

	var taxDue float64

	if taxableIncome <= 0 {
		taxDue = 0
	} else if taxableIncome <= max10 {
		taxDue = taxableIncome * 0.1
	} else if taxableIncome <= max12 {
		taxDue = tier10_tax + ((taxableIncome - max10) * .12)
	} else if taxableIncome <= max22 {
		taxDue = tier12_tax + ((taxableIncome - max12) * .22)
	} else if taxableIncome <= max24 {
		taxDue = tier22_tax + ((taxableIncome - max22) * .24)
	} else if taxableIncome <= max32 {
		taxDue = tier24_tax + ((taxableIncome - max24) * .32)
	} else if taxableIncome <= max35 {
		taxDue = tier32_tax + ((taxableIncome - max32) * .35)
	} else if taxableIncome > max35 {
		taxDue = tier35_tax + ((taxableIncome - max35) * .37)
	}
	fmt.Print("Your gross income is: $")
	fmt.Println(grossIncome)
	fmt.Print("Your claimed number of dependents is: ")
	fmt.Println(numDep)
	fmt.Print("Your taxable income is: $")
	fmt.Println(taxableIncome)
	fmt.Print("Your tax due is: $")
	fmt.Println(int(taxDue))
}
