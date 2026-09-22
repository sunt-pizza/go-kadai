package main

import "fmt"

type ToDo struct {
	ID int
	Title string
	Completed bool
}

func Completed(c *ToDo) {
	fmt.Printf("ID:%dのToDoを完了に更新します\n", c.ID)
	c.Completed = true
}

func printTodos(p ToDo) {
	if (p.Completed == true) {
		fmt.Printf("[完 了]（ID：%d）%s\n", p.ID, p.Title)
	} else {
		fmt.Printf("[未完了]（ID：%d）%s\n", p.ID, p.Title)
	}
}

func main() {
	item1 := ToDo{
		ID: 1,
		Title: "学習計画",
		Completed: false, 
	}

	item2 := ToDo{
		ID: 2,
		Title: "環境構築",
		Completed: false, 
	}

	item3 := ToDo{
		ID: 3,
		Title: "基礎文法",
		Completed: false, 
	}

	Todos := []ToDo{
		item1, item2, item3,
	}

	fmt.Println("--- ToDoリスト（初期状態） ---")
	for _, t := range Todos {
		printTodos(t)
	}
	fmt.Println()

	Completed(&Todos[0])
	Completed(&Todos[1])
	fmt.Println()

	fmt.Println("--- ToDoリスト（最終状態） ---")
	for _, t := range Todos {
		printTodos(t)
	}
}