/*

       CSS 410W
       Spring Semester, 2019

      Assignment 7:Go program that read monthly rainfall levels for a city, prints those values out in a table, and
                   then prints some statistics.Program includes three functions in addition to main.

       Programmed by: Surya Partap Singh
       Due: Tuesday,February 2, 2020

*/



package main

import (
  "fmt"
  "os"
  "text/tabwriter"
  "math"

)
const TWELVE int = 12


//**********************************************************************************************************
//Function: Reading()([]float64,error)
//Purpose:  This function reads values from file and returns it and also returns error
//Incoming: None
//Outgoing: None
//Return:  RainAmount,err
//********************************************************************************************************
func Reading()([]float64,error){
   f, err := os.Open("rainfall.txt")
   if err != nil {
      fmt.Println(err)
   }
   var RainAmount = make([]float64,TWELVE)
   for i:=0;i<TWELVE;i++{
      var flt float64
      var n int
      n, err = fmt.Fscanln(f,&flt)
      if n == 0 || err != nil {
        break
      if i!=TWELVE{
        fmt.Println(err)
      }
      }
      RainAmount[i]=flt
   }
   if err := f.Close(); err != nil {
        fmt.Println(err)
    }
  return RainAmount,err


}
//**********************************************************************************************************
//Function: RainFallTable(RainFall [TWELVE]float64,Month [TWELVE]string)([TWELVE]float64,[TWELVE]string)
//Purpose:  This function prints the amount of rainfall in a particular month
//Incoming: RainFall [TWELVE]float64,Month [TWELVE]string
//Outgoing: None
//Return:  RainFall,Month
//********************************************************************************************************
func RainFallTable(RainFall [TWELVE]float64,Month [TWELVE]string)([TWELVE]float64,[TWELVE]string){
  f, err := os.Open("rainfall.txt")
  if err != nil {
     fmt.Println(err)
  }
  for i:=0;i<TWELVE;i++{
     var flt float64
     var n int
     n, err = fmt.Fscanln(f,&flt)
     if n == 0 || err != nil {
       break
     }
     RainFall[i]=flt

  }
  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 2, 1, 1, ' ', tabwriter.AlignRight)
  Month =[12]string{"January","February","March","April","May","June","July","August","September","October","November","December"}
  fmt.Fprintln(w, "Rainfall\tMonths\t")
  fmt.Println("******************************")
  for i:=0;i<TWELVE;i++{
    fmt.Fprintln(w,Month[i],"\t",math.Floor(RainFall[i]*10)/10)
  }
  fmt.Fprintln(w)
	w.Flush()

  return RainFall,Month

}
//**********************************************************************************************************
//Function: Computing(RainFall [TWELVE]float64)
//Purpose:  This function check the index of largest number and smallest in the file, also calculates average rainfall
//Incoming: RainFall [TWELVE]float64
//Outgoing: None
//Return:  max1,min1,Average
//********************************************************************************************************
func Computing(RainFall [TWELVE]float64) (int, int,float64){
  f, err := os.Open("rainfall.txt")
  if err != nil {
     fmt.Println(err)
  }
  for i:=0;i<TWELVE;i++{
     var flt float64
     var n int
     n, err = fmt.Fscanln(f,&flt)
     if n == 0 || err != nil {
       break
     }
     RainFall[i]=flt

  }
  max := RainFall[0]
  min := RainFall[0]
  var max1 int
  var min1 int
  var sum float64
  var Average float64
  for _, value := range RainFall{
                if value > max {
                        max = value
                }
        }
  for _, value := range RainFall {
      if value < min {
              min = value
                      }
              }
  for i:=0;i<TWELVE;i++{
    if(RainFall[i]==max){
      max1 = i
    }
  }
  for i:=0;i<TWELVE;i++{
    if(RainFall[i]==min){
      min1 = i
    }
  }
  for i:=0;i<TWELVE;i++{
    sum+=RainFall[i]
  }
  Average=sum/float64(TWELVE)
  return max1,min1,Average

}
//This is the main driver which runs the program
func main() {
   fmt.Println("******************************")
   fmt.Println("Rain in the year 2019")

   var arr=[TWELVE]float64{}
   var mon=[TWELVE]string{}

   RainAmount,months:=RainFallTable(arr,mon)
   max2,min2,average:=Computing(RainAmount)
   array,err:=Reading()
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println("Largest value in the list :",array[max2],"in Month",months[max2])
   fmt.Println("Smallest value in the list:",array[min2],"in Month",months[min2])
   fmt.Println("Average Rainfall:",math.Floor(average*100)/100)


}
