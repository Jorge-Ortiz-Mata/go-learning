package main

import (
	"example/note/note"
	"example/note/todo"
	"fmt"
)

type displayer interface {
	Display()
}

type saver interface {
	Save() string
}

// Option 01
// type outputtable interface {
// 	Display()
// 	Save() string
// }

// Option 02 - I like this
type outputtable interface {
	displayer
	saver
}

func main() {
	noteOne := note.New("Buy eggs", "Go to the store after 1 pm")
	todoOne := todo.New("This the first text of my first todo")
	Display(noteOne)
	Display(todoOne)
	Save(noteOne)
	Save(todoOne)
	DisplayAndSave(noteOne)
	DisplayAndSave(todoOne)
}

func DisplayAndSave(data outputtable) {
	fmt.Println("Display and Save")
	contentToSave := data.Save()
	fmt.Println(contentToSave)
	data.Display()
}

func Save(data saver) {
	contentToSave := data.Save()
	fmt.Println(contentToSave)
}

func Display(data displayer) {
	data.Display()
}
