// Support multiple transactions in same run i.e. without restarting the program user can withdraw multiple times
package main

import "fmt"

var (
	choice    int
	WdwAmount int
	amount    int = 4563 //Start a user with 4563 rupees balance
)

func withdraw() { //defining a function to withdraw money
	fmt.Printf("\n\tChecking Balance ....\n\n\tBalance Amount : %d\n", amount) //Check for balance before allowing to withdraw
	fmt.Println("\n\tEnter the amount to withdraw")                            //printing a message to enter the amount to withdraw using println function
	fmt.Scanln(&WdwAmount)                                                     //takes the user input and scans the value
	//user can withdraw in multiple of 100
	if WdwAmount%100 != 0 || WdwAmount <= 0 { //checks if the amount entered by the user is not a multiple of 100 and not less than or equalto zero
		fmt.Println("\n\tPlease enter the amount in multiples of 100") //if the condition is true ,this message is printed using println function

	} else if WdwAmount >= 2000 { //Maximum allowed per transaction is 2000
		if WdwAmount >= amount { //if the withdrawal amount is greater than amount
			fmt.Println("\ninsufficient balance") //if the condition is true print the message
		} else {
			fmt.Println("\nSorry ,the transaction limit for each transaction is Rs 2000.") //else print the message using println function

		}
	} else {

		if WdwAmount <= amount { //if condition to check the withdrawal amount is less than or equal to amount
			amount = amount - WdwAmount                          //deduct the withdrawal amount from the original amount of that user
			fmt.Println("\n\tyour transaction is processing...") //prints a meassage using println function
			fmt.Println("\n\tplease collect your cash")          //prints a meassage using println function
			//code to output the notes and number of notes based on the withdrawal amount
			n := []int{2000, 500, 200, 100} //creating a slice of int and storing the notes as integers in variable n using short declaration operator
			counter := []int{0, 0, 0, 0}    //creating a slice of int and setting zero as values,counter represents the no of notes for that value
			for j := 0; j < len(n); j++ {   //creating a for loop with  condition: length less than n
				if WdwAmount >= n[j] { //if the withdrawing amount is greater than or equal to each item in the slice
					counter[j] = WdwAmount / n[j]           //if the condition is true,the coressponding item in counter is updated with the  value(withdrawal amount divided by item in slice)
					WdwAmount = WdwAmount - counter[j]*n[j] //the withdrawing amount is updated .
				}
			}
			for j := 0; j < len(n); j++ { //creating a for loop with  condition: length less than n
				if counter[j] != 0 { //checks if each item of in slice counter is !=0
					fmt.Println("\n\t", n[j], " * ", counter[j]) //if condition is true,print the value in n  and the counter for each j which is !=0
				}

			}
			fmt.Println("\n\tBalance amount:", amount) //print the balance amount after withdrawal
		} else { //low balance  messages.
			fmt.Println("Insufficient Balance!") //if the condition is false prints the message
		}
	}

}

func main() {
	withdraw() //calling the function withdraw()
	//Max 5 transactions per day
	for i := 0; i < 4; i++ { //defining a for loop
		fmt.Println("\nDo you want to continue?Press 1 to continue or Press 2 to cancel\n") //printing the message asking to continue withdrawal
		fmt.Scanln(&choice)                                                                 //scan the value of type int

		if choice == 1 { // if the value is 1
			withdraw() //if the condition is true call the function
		} else if choice == 2 { //if the value is 2
			fmt.Println("Thank you for using ATM") //print the message
			break                                  //break the loop
		} else {
			fmt.Println("Please enter your choice correctly") //else print this  message
		}

	}
}
