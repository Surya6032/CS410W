/*

       CSS 410W
       Spring Semester, 2019

      Assignment 8:Go program that will process vehicle registration information. A map must be used to store the information. 
				   The key will be the license plate. The data stored will be a string with comma-separated first name, last name, and phone number. 

       Programmed by: Surya Partap Singh
       Due: Sunday,February 9, 2020

*/


package main

import (
    //"fmt"
    //"datafile"
    "bufio"
    "os"
    "fmt"
	"strings"
)
//**********************************************************************************************************
//Function:Menu
//Purpose:  This function Prints the menu
//Incoming: None
//Outgoing: None
//Return:  None
//********************************************************************************************************
func Menu(){
 fmt.Println("A: Add a new vehicle Registeration")
 fmt.Println("S: Search for Vehicle")
 fmt.Println("P: Print All")
 fmt.Println("Q:Quit the Program")


}
//**********************************************************************************************************
//Function: Reading()(map[int]string)
//Purpose:  This function reads values from file and stores it in a map
//Incoming: None
//Outgoing: None
//Return:  data
//********************************************************************************************************
func Reading()(map[int]string){
  data := make(map[int]string)
  file, _ := os.Open("Vehicle.txt")
  scanner := bufio.NewScanner(file)
  var i int
  i=0

  //var value string
  for scanner.Scan() {
      line:=scanner.Text()
      data[i]=line
	  i++
    }



return data

}
//**********************************************************************************************************
//Function: Search
//Purpose:  This function search the file and checks if the data is there or not
//Incoming: None
//Outgoing: Searched item
//Return:  None
//********************************************************************************************************
func Search(){

 var Lplate int
 fmt.Println("Enter License Plate:")
 fmt.Scanln(&Lplate)
 Data := make(map[int]string)
 Data=Reading()
 value,ok := Data[Lplate]
	if ok {
			result := strings.Split(value, ",")
			fmt.Println("First Name:",result[0])
			fmt.Println("Last Name:",result[1])
			fmt.Println("Phone Number:",result[2])
		}else{
	    fmt.Println("License Plate doesn't exists")
	
	}
 



}
//**********************************************************************************************************
//Function: ShowAll
//Purpose:  This function Prints the data from file
//Incoming: None
//Outgoing: Data of file
//Return:  None
//********************************************************************************************************
func ShowAll(){
    
  file, _ := os.Open("Vehicle.txt")
  scanner := bufio.NewScanner(file)

  //var value string
  for scanner.Scan() {
      line:=scanner.Text()
	  result := strings.Split(line, ",")
	  for i := range result {
        fmt.Print(result[i]," ")
		
    }}




}
//Main function 
func main() {
    var choice string
    var LPlate int
    var firstName string
    var lastName string
	var PhoneNumber string
	Data := make(map[int]string)
	Data=Reading()
    for{
	Menu()
    fmt.Scanln(&choice)
    if choice=="A"{
      fmt.Println("Enter license plate number:")
      fmt.Scanln(&LPlate)
	  value,ok := Data[LPlate]
	
		if ok {
			fmt.Println("License plate already exits")
			fmt.Println(value)
			
		}else{
		fmt.Println("Enter First Name:")
		fmt.Sprintf(firstName)
		fmt.Println("Enter Last Name:")
		fmt.Sprintf(lastName)
	    fmt.Println("Enter Phone Number:")
		fmt.Sprintf(PhoneNumber)
		f, err := os.OpenFile("Vehicle.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
           fmt.Println(err)
           return}
		
		_,err1 := f.WriteString("\n"+firstName+","+lastName+","+PhoneNumber) 
		if err1!=nil{
			fmt.Println(err1)}
			
		
		fmt.Println(f)}
		

		
	    
		}else if choice=="S"{
		
		Search()
		
		}else if choice=="P"{
		
		ShowAll()
		
		
		}else if choice=="Q"{
			break;
		
		}else{
		fmt.Print("Wrong Choice")}}
		
    


}

