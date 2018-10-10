package main

import "fmt"

func main() {
	if !candrink(16) { //if not true will evaluate to False
		fmt.Println("You cant drink!!!")
	}

	if candrink(18){ //if not boolean -- so here its true -- equivalent of If
		fmt.Println("Will run")
	}

	fmt.Println(canride(43, 180))
	fmt.Println(canride(13, 99))
}

func candrink(age int) bool{
	if age >= 18{
		return true
	} else {
		return false
	}
}

//also multiple arg definition

func canride(age, heightcm int) string {
	if (age >= 12) && (heightcm > 100){
		return "Can ride!"
	} else {
		return "You cannot ride!"
	}
}

//multiple return

func greet(fname, lname string, age int) (string, string, int){
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname), fmt.Sprint(fname, lname, age)
}


/*
	fmt.Sprint -> like Printf but returns string not stdout
*/

/*
	This if will try to call a method. This method will return something that is not nil
	in the case of an error. We can then decide what to do with any err raised??

	The scope of err ir restricted to the if and the code within

func exampleif() string{
	if err:= file.Chmod(0664); err != nil{
		log.Print(err)
		return err
	}
}
*/