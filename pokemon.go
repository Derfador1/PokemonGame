package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type character struct {
	age int
	name string
	hometown string
}

func main() {
	//var age string
	//var new_age string
	player := character{}

	fmt.Println("Welcome to the world of Pokemon\n")
	fmt.Println("Version Alpha 1.0")
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Please tell me your name:")
	player.name, _ = reader.ReadString('\n')
	player.name = strings.TrimSuffix(player.name, "\n")
	
	//player.name = "Bob"
	player.hometown = "Pallet"
	player.age = 2

	fmt.Printf("%v\n", player)

	
	//line := "hello\n"
	//line = strings.TrimSuffix(line, "\n") // remove newline
	//fmt.Printf("'%s'", line)
	
	
	
	
	
	//fmt.Println("How old are you:")
	//age, _ = reader.ReadString('\n')
	//fmt.Println(age)	
	//new_age = strings.TrimSuffix(age, "\n")
	//fmt.Println(new_age)
	//player.age, _ = strconv.Atoi(age)
	//fmt.Println(player)
	
	
	//fmt.Println("Oh I almost forgot, where are you from:")
	//player.hometown, _ = reader.ReadString('\n')
	//player.hometown = strings.Trim(player.hometown, "\n")
	//fmt.Println("\n")
	//fmt.Println(player)

	
}