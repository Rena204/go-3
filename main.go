package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func main() {
	var items []Item

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Добро пожаловать в коллекцию URL!")
	fmt.Println("Нажмите 'A' для добавления новой ссылки")
	fmt.Println("Нажмите 'D' для удаления ссылки")
	fmt.Println("Нажмите 'L' для вывода списка ссылок")
	fmt.Println("Нажмите 'Q' для выхода")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch strings.ToUpper(string(char)) {
		case "A":
			item := addItem()
			items = append(items, item)
			fmt.Println("Ссылка успешно добавлена!")
		case "D":
			deleteItem(&items)
		case "L":
			printItems(items)
		case "Q":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный ввод. Попробуйте еще раз.")
		}
	}
}

func addItem() Item {
	var item Item

	fmt.Print("Введите имя ссылки: ")
	fmt.Scanln(&item.Name)

	item.Date = time.Now()

	fmt.Print("Введите теги для ссылки через запятую: ")
	fmt.Scanln(&item.Tags)

	fmt.Print("Введите URL: ")
	fmt.Scanln(&item.Link)

	return item
}

func deleteItem(items *[]Item) {
	var index int

	fmt.Print("Введите индекс ссылки, которую хотите удалить: ")
	fmt.Scanln(&index)

	if index >= 0 && index < len(*items) {
		*items = append((*items)[:index], (*items)[index+1:]...)
		fmt.Println("Ссылка успешно удалена!")
	} else {
		fmt.Println("Неверный индекс ссылки.")
	}
}

func printItems(items []Item) {
	fmt.Println("Список ссылок:")

	for i, item := range items {
		fmt.Printf("Индекс: %d\n", i)
		fmt.Printf("Имя: %s\n", item.Name)
		fmt.Printf("Дата добавления: %s\n", item.Date.Format(time.RFC3339))
		fmt.Printf("Теги: %s\n", item.Tags)
		fmt.Printf("URL: %s\n", item.Link)
		fmt.Println()
	}
}
