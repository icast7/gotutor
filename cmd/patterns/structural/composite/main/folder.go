package main

import "fmt"

type folder struct {
	components []component
	name       string
}

func (f *folder) getName() string {
	return f.name
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.getName())
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
