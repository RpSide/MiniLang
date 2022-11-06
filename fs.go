package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadToString(file string) string {

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadByWord(file string) []string {

	var fe []string
	f, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		fe = append(fe, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return fe
}

func ReadByLine(file string) []string {

	var fe []string
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fe = append(fe, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fe
}

func WriteFile(file string, data string) {

	f, err := os.Create(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(data)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func CreateDir(DirName string) {

	os.Mkdir(DirName, os.ModePerm)

}
