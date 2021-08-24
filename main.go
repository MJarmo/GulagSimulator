package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

var (
	fileURL string = "https://drive.google.com/file/d/1tciMrgszAC-yO_0MMDLc3yA_PA-Sw2Mk/view?usp=sharing"
)

const (
	Prisoner = iota
	Guard
	Doctor
	Special
)
const (
	MAX_AGE int = 60
	MIN_AGE int = 8
	MIN_WEIGHT int = 30
	MAX_WEIGHT int = 130
	MIN_HAPPYNES int = 30
	MAX_HAPPYNESS int = 100

)

type Person struct {
	name string
	lastName string
	age int
	weight int
	role int
	health int 
	happyness int
}


func  createPrisoner()Person {
	var p Person
	p.name = "XX"
	p.lastName = "XD"
	p.age = rand.Intn(MAX_AGE - MIN_AGE) + MIN_AGE
	p.weight = rand.Intn(MAX_WEIGHT - MIN_WEIGHT) + MIN_WEIGHT
	p.role = Prisoner
	p.health = rand.Intn(MAX_HAPPYNESS - MIN_HAPPYNES) + MIN_HAPPYNES
	p.happyness = rand.Intn(MAX_HAPPYNESS - MIN_HAPPYNES) + MIN_HAPPYNES
	return p
}
func main() {
	DownloadFile("kk.txt", fileURL)
	fmt.Println("Welcom in gulag symulator 2021")
	x := createPrisoner()
	fmt.Println(x)
}
func InitGame() {

}

func DownloadFile(filePath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}