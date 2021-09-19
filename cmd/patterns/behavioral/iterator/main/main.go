package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}

	user2 := &user{
		name: "b",
		age:  21,
	}

	userCollection := userCollection{
		users: []*user{user2, user1},
	}

	i := userCollection.createIterator()

	for i.hasNext() {
		user := i.getNext()
		fmt.Printf("User is %+v\n\n", user)
	}
}
