package main

import (
	"fmt"
	"github.com/Ddarli/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.SaveFile("text.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}
	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it is restored:", restoredFile)
}
