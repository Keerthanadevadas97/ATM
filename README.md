# ATM
The updated code is atm.go

Points Updated :

There was no naming convention followed 


it is giving '-1500' is not the multiplier of 100 which is not true, its negative number that why the transaction cant be done. 


it is not allowed to withdraw 2000 which is the limit 


your withdrawal function was called outside the loop then inside the loop, it should have called inside the loop only


account was reduced before the money was counted, means you have reduced the money then looking for currency count. 


you have checked multiple places for sufficient funds, which is unnecessary


It was asked to print the currency in specific format and you have printed in different format.


use of switch statement


variable scoping 

there was no error handling in case of user entered string/text not number 

Issue:
In case of invalid data entry , when I run the program output is like this:
go run atm.go

	Checking Balance ....

	Balance Amount : 4563

	Enter the amount to withdraw
ggul
Invalid input, expected integer
❯ gul
zsh: command not found: gul



