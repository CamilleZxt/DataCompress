package rwfile

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(file,str string){
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = fmt.Fprint(f, str)
	//_, err = fmt.Fprintln(f, str)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFile2(file,str string){
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	l, err := f.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l, " bytes written successfully!")
}

func WriteFile3(file string,byt [] byte){
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//d := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	l, err := f.Write(byt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l, " bytes written successfully!")
}

