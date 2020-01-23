/*

       CSS 410W
       Spring Semester, 2019

       Assignment 4.1:Program to request the number of hours spent at
					  each activity and then display the number of calories burned and the number of pounds 
	                  worked off.

       Programmed by: Surya Partap Singh
       Due: Wednesday, 2020

*/
package main

import (
	
        "fmt"
        "math"
       
        
        
)

func main() {
	var cyclingCalories int
	var runningCalories int
	var swimmingCalories int
	var hoursCyc int
	var hoursRun int
	var hoursSwim int
	var CycCaloriestotal int
	var RunCaloriestotal int
	var SwimCaloriestotal int
	var TotalCalories int
	var PoundsBurned float64
 
	cyclingCalories  = 200
	runningCalories  = 475
	swimmingCalories = 275
	fmt.Print("Enter number of hours spend on cycling:")
	 _, err := fmt.Scanf("%d", &hoursCyc)
	fmt.Print("Enter number of hours spend in running:")
	fmt.Scanf("%d",&hoursRun)
	fmt.Print("Enter number of hours spend on swimming:")
	fmt.Scanf("%d", &hoursSwim)
  
	if err != nil {
            fmt.Println(err)
         }
	CycCaloriestotal=hoursCyc*cyclingCalories
	RunCaloriestotal=hoursRun*runningCalories
	SwimCaloriestotal=hoursSwim*swimmingCalories

	fmt.Print("Calories burned in cycling:",CycCaloriestotal,"\n")
	fmt.Print("Calories burned in Running:",RunCaloriestotal,"\n")
	fmt.Print("Calories burned in Swimming:",SwimCaloriestotal,"\n")
	TotalCalories=CycCaloriestotal+RunCaloriestotal+SwimCaloriestotal
	fmt.Print("Total Calories burned:",TotalCalories,"\n")
  
	PoundsBurned=float64(TotalCalories)/3500

	fmt.Print("Pounds burned:",math.Floor(PoundsBurned*100)/100,"\n")            //Used Math library to round the number to two decimal point
	
}