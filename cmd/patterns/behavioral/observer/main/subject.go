package main

// Unused
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}
