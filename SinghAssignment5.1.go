/*

       CSS 410W
       Spring Semester, 2019

       Assignment 5.1:Program to request the number of hours spent at
		      each activity and then display the number of calories burned and the number of pounds 
	              worked off.It also ask number of athletes and print data for multiple athletes.

       Programmed by: Surya Partap Singh
       Due: Tuesday,January 28,2020

*/
package main

import (
	
        "fmt"
        "math"
       
        
        
)

func main() {
	var cyclingCalories float64
	var runningCalories float64
	var swimmingCalories float64
	var hoursCyc float64
	var hoursRun float64
	var hoursSwim float64
	var CycCaloriestotal float64
	var RunCaloriestotal float64
	var SwimCaloriestotal float64
	var TotalCalories float64
	var PoundsBurned float64
  	var Athlete int
 
	cyclingCalories  = 200
	runningCalories  = 475
	swimmingCalories = 275
  	fmt.Print("How many triathletes are there?")
  	fmt.Scanf("%d",&Athlete)
  
  	for i:=1;i<=Athlete;i++{
  		fmt.Print("\n")
  		fmt.Print("Athlete #",i,":","\n")
		fmt.Print("Enter number of hours spend on cycling: ")
	 	_, err := fmt.Scanf("%f", &hoursCyc)
		fmt.Print("Enter number of hours spend in running: ")
		fmt.Scanf("%f",&hoursRun)
		fmt.Print("Enter number of hours spend on swimming: ")
		fmt.Scanf("%f", &hoursSwim)
  
		if err != nil {
            		fmt.Println(err)
         	}
		CycCaloriestotal=hoursCyc*cyclingCalories
		RunCaloriestotal=hoursRun*runningCalories
		SwimCaloriestotal=hoursSwim*swimmingCalories

		TotalCalories=CycCaloriestotal+RunCaloriestotal+SwimCaloriestotal
		fmt.Print("Total Calories burned: ",TotalCalories,"\n")
  
		PoundsBurned=float64(TotalCalories)/3500

		fmt.Print("Pounds burned:",math.Floor(PoundsBurned*100)/100,"\n")}            //Used Math library to round the number to two decimal point
 
	
}