package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerSecond)
	shirtItem.register(observerFirst)
	shirtItem.updateAvailability()
	shirtItem.updateAvailability()
	shirtItem.deregister(observerFirst)
	shirtItem.deregister(observerSecond)
	shirtItem.updateAvailability()
}
