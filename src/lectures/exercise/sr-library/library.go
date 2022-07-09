//--Summary:
//  Create a program to manage lending of library books.
//

//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	yes = true
	no  = false
)

//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
type Book struct {
	name         string
	checkedOut   CheckedOut
	checkOutTime time.Time
	checkInTime  time.Time
	borrowedBy   Member
}

type Member string
type CheckedOut bool

func printLibraryState(books []Book, members []Member) {
	for i := 0; i < len(books); i++ {
		book := books[i]
		if book.checkedOut {
			fmt.Println(book.name, "was checked out", book.checkOutTime, "by", book.borrowedBy)
		} else {
			fmt.Println(book.name, "is not checked out. Was checked in", book.checkInTime)
		}
	}
	fmt.Println("These are the current members in the library")
	for _, member := range members {
		fmt.Println(member)
	}
}

func checkOut(book *Book, borrowedBy Member) (CheckedOut, error) {
	if book.checkedOut {
		return no, errors.New(book.name + " already checkout out to " + string(book.borrowedBy))
	}
	book.checkedOut = yes
	book.borrowedBy = borrowedBy
	book.checkOutTime = time.Now()
	return yes, nil
}

func checkIn(book *Book) (CheckedOut, error) {
	if !book.checkedOut {
		return yes, errors.New(book.name + " is not checked out and thus can not be checked in")
	}
	book.checkedOut = no
	book.checkInTime = time.Now()
	book.borrowedBy = ""
	return no, nil
}

func main() {
	//  - Add at least 4 books and at least 3 members to the library
	members := []Member{"Marius", "Kjell", "Håkon"}
	books := make([]Book, 4)
	books[0] = Book{name: "Lord of the flies"}
	books[1] = Book{name: "Lord of the rings"}
	books[2] = Book{name: "Lord of the flings"}
	books[3] = Book{name: "Lord of the lies"}
	printLibraryState(books, members)
	_, err := checkOut(&books[3], members[2])
	if err != nil {
		panic(err)
	} else {
		printLibraryState(books, members)
	}
	_, err = checkIn(&books[3])
	if err != nil {
		panic(err)
	} else {
		printLibraryState(books, members)
	}
	//  - Check out a book
	//  - Check in a book
	//  - Print out initial library information, and after each change
	//* There must only ever be one copy of the library in memory at any time
}
