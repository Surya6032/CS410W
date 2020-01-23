/*

       CSS 410W
       Spring Semester, 2019

       Assignment 4.2:Program that asks for the customer's name and 
					  how many copies they need. Your program should output the 
	                  customer's name and the total cost.

       Programmed by: Surya Partap Singh
       Due: Wednesday, 2020

*/

package main

import (
	
        "fmt"
        "bufio"
        "os"
        "log"
       
       
        
        
)

func main() {
	var CustomerName string
	var NumberofCopies int
	var Price int
	var NewPrice int
	var Remaining int
	var Dollars float32

	fmt.Println("Enter your Name:")
	reader := bufio.NewReader(os.Stdin)
	CustomerName,err := reader.ReadString('\n')
	if err != nil {
					log.Fatal(err)
			}
	fmt.Println("Enter number of copies to print:")
	fmt.Scanf("%d",&NumberofCopies)
	if NumberofCopies <= 100{
	  Price = NumberofCopies*5
	}else{
	  Remaining=NumberofCopies-100
	  NewPrice=Remaining*3
	  Price=100*5+NewPrice
	}
	fmt.Print("Customer Name:",CustomerName)
	if Price>100{
	 Dollars=float32(Price)/100
	fmt.Print("Amount to be paid:$",Dollars)
	}else{
	Dollars=float32(Price)/100
	fmt.Print("Amount to be paid:$",Dollars)
	}



    


}
