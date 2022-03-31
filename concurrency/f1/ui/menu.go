package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	SELECT_GRAND_PRIX = "1"
	NEW_CHAMP         = "2"
	CONTINUE_CHAMP    = "3"
	EXIT              = "4"
)

type MenuItem struct {
	Code        string
	Instruction string
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func createMenus() []MenuItem {
	var menus []MenuItem
	menu1 := MenuItem{Code: SELECT_GRAND_PRIX, Instruction: "Select Grand Prix"}
	menu2 := MenuItem{Code: NEW_CHAMP, Instruction: "New Champ"}
	menu3 := MenuItem{Code: CONTINUE_CHAMP, Instruction: "Continue Champ"}
	menu4 := MenuItem{Code: EXIT, Instruction: "Exit"}

	menus = append(menus, menu1)
	menus = append(menus, menu2)
	menus = append(menus, menu3)
	menus = append(menus, menu4)
	return menus
}

func printMenuOptions(menus []MenuItem) {
	for _, menu := range menus {
		fmt.Printf("%s - %s\n", menu.Instruction, menu.Code)
	}
}

func scanInput() string {
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func Menu() {
	menus := createMenus()
	var exit bool = false
	for {
		printMenuOptions(menus)

		input := scanInput()
		fmt.Println("selected: ", input)
		if strings.Compare(input, EXIT) == 0 {
			exit = true
		}

		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()
		if exit {
			fmt.Println("Exiting")
			time.Sleep(2 * time.Second)
			break
		}
	}
	fmt.Println("exit success")
}
