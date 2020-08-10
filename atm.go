// Support multiple transactions in same run i.e. without restarting the program user can withdraw multiple times
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var amount int = 4563

func withdraw() { //defining a function to withdraw money
	var WdwAmount int
	fmt.Printf("\n\tChecking Balance ....\n\n\tBalance Amount : %d\n", amount) //Check for balance before allowing to withdraw
	fmt.Println("\n\tEnter the amount to withdraw")                            //printing a message to enter the amount to withdraw using println function
	fmt.Scanln(&WdwAmount)                                                     //takes the user input and scans the value
	if WdwAmount%100 != 0 {                                                    //checks if the amount entered by the user is not a multiple of 100 and not less than or equal to zero
		fmt.Println("\n\tPlease enter the amount in multiples of 100") //if the condition is true ,this message is printed using println function
	} else if WdwAmount == 0 { //if withdraw amount is equal to zero
		fmt.Println("Please enter a valid amount") //print this message
	} else if WdwAmount < 0 { //if withdraw amount less than zero
		fmt.Println("Please enter a positive number") //print this message
	} else if WdwAmount >= 2000 { //Maximum allowed per transaction is 2000
		fmt.Println("\nSorry ,the transaction limit is Rs 2000.") //else print the message using println function
	} else {
		//code to output the notes and number of notes based on the withdrawal amount
		temp := WdwAmount             //storing the value in WdwAmount to the new variable temp
		n := []int{500, 200, 100}     //creating a slice of int and storing the notes as integers in variable n using short declaration operator
		counter := []int{0, 0, 0}     //creating a slice of int and setting zero as values,counter represents the no of notes for that value
		for j := 0; j < len(n); j++ { //creating a for loop with  condition: length less than n
			if temp >= n[j] { //if the value in temp is greater than or equal to each item in the slice
				counter[j] = temp / n[j]      //if the condition is true,the coressponding item in counter is updated with the  value(temp divided by item in slice)
				temp = temp - counter[j]*n[j] //the temp value  is updated .
			}
		}
		var s []string
		for j := 0; j < len(n); j++ { //creating a for loop with  condition: length less than n
			if counter[j] != 0 { //checks if each item of in slice counter is !=0
				t := strconv.Itoa(n[j]) + "*" + strconv.Itoa(counter[j]) //coverting the int values to string and storing in new variable t
				s = append(s, t)                                         //appending t to s
			}
		}
		fmt.Println("\n\t", strings.Join(s, "+")) //joins the string where s is the string from which we can concatenate elements and "+" is the separator  which is placed between the elements in the final string.
		if WdwAmount <= amount {                  //if condition to check the withdrawal amount is less than or equal to amount
			amount = amount - WdwAmount                          //deduct the withdrawal amount from the original amount of that user
			fmt.Println("\n\tyour transaction is processing...") //prints a meassage using println function
			fmt.Println("\n\tplease collect your cash")          //prints a meassage using println function
			fmt.Println("\n\tBalance amount:", amount)           //print the balance amount after withdrawal
		} else { //low balance  messages.
			fmt.Println("Sorry,Insufficient Balance!") //if the condition is false prints the message
		}
	}
}
func main() {
	var choice int
	//Max 5 transactions per day
	exit := false            //assign a boolean value false to a new variable x
	for i := 0; i < 5; i++ { //defining a for loop
		withdraw()                                                                                            //call withdraw function
		fmt.Println("\nDo you want to continue?,please enter your choice press 1 to continue or 2 to cancel") //printing the message asking to continue withdrawal
		fmt.Scanln(&choice)                                                                                   //scan the value of type int
		if i == 4 {                                                                                           //if condition to check if i ==4
			fmt.Println("Sorry you have reached the maximum limit") //prints this message
			break                                                   //break the loop
		}
		switch choice { //switch case to select the choice
		case 1: //if choice is 1
			break
		case 2: //if choice is 2
			fmt.Println("Thank you for using ATM") //print the message
			exit = true                            //reassign value in exit variable in case2
		default: //In default case
			fmt.Println("Please enter your choice correctly") // print this  message
			exit = true                                       //reassign value in exit variable in default case
		}
		if exit { //if exit is true
			break //break the loop
		}
	}

}
