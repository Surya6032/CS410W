package main

import (
    //"fmt"
    //"datafile"
    "bufio"
    "os"
    "fmt"
	"strings"
)
func Menu(){
 fmt.Println("A: Add a new vehicle Registeration")
 fmt.Println("S: Search for Vehicle")
 fmt.Println("Q:Quit the Program")


}
func Reading()(map[int]string){
  m := make(map[int]string)
  f, _ := os.Open("Vehicle.txt")
  scanner := bufio.NewScanner(f)
  var i int
  i=0

  //var value string
  for scanner.Scan() {
      line:=scanner.Text()
      m[i]=line
	  i++
    }



return m

}
func Search(){

 var Lplate int
 fmt.Println("Enter License Plate:")
 fmt.Scanln(&Lplate)
 m := make(map[int]string)
 m=Reading()
 value,ok := m[Lplate]
	if ok {
			result := strings.Split(value, ",")
			fmt.Println("First Name:",result[0])
			fmt.Println("Last Name:",result[1])
			fmt.Println("Phone Number:",result[2])
		}else{
	    fmt.Println("License Plate doesn't exists")
	
	}
 



}
func main() {
    var choice string
    var LPlate int
    var firstName string
    var lastName string
	var PhoneNumber string
	m := make(map[int]string)
	m=Reading()
    Menu()
    fmt.Scanln(&choice)
    if choice=="A"{
      fmt.Println("Enter license plate number:")
      fmt.Scanln(&LPlate)
	  value,ok := m[LPlate]
	
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
		
		}
		
    fmt.Println(Reading())


}

