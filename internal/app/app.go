package app

import (
	"bufio"
	"fmt"
	gauss "lab1/internal"
	"os"
	"strconv"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	fmt.Println(`Выберете способ ввода матрицы:
(1) Ввод с клавиатуры;
(2) Ввод из файла;
(0) Выход из программы;`)
	var inputType int

	for {
		if _, err := fmt.Scanln(&inputType); err != nil {
			fmt.Println("Ошибка при вводе!")
			return
		}
		switch inputType {
		case 1:
			inputFromKeyboard()
			return
		case 2:
			inputFromFile()
			return
		case 0:
			return
		}
	}
}

func inputFromKeyboard() {
	fmt.Println("\nВведите размер матрицы:")
	var size int
	if _, err := fmt.Scanln(&size); err != nil {
		fmt.Println("Ошибка при вводе!")
		return
	}
	matrix := make([][]float64, size)
	fmt.Println("\nВведите элементы матрицы:")
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size+1)
		for j := 0; j <= size; j++ {
			if _, err := fmt.Scan(&matrix[i][j]); err != nil {
				fmt.Println("Ошибка при вводе!")
				return
			}
		}
	}
	fmt.Println(matrix)
	calculate(matrix)
}

func inputFromFile() {
	fmt.Println("\nВведите расположение файла:")
	var path string
	if _, err := fmt.Scanln(&path); err != nil {
		fmt.Println("Ошибка при вводе!")
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size+1)
		for j := 0; j <= size; j++ {
			scanner.Scan()
			matrix[i][j], _ = strconv.ParseFloat(scanner.Text(), 64)
		}
	}
	fmt.Println()
	gauss.PrintMatrix(matrix)
	calculate(matrix)

}

func calculate(matrix [][]float64) {
	o := gauss.NewOriginalMatrix(matrix)
	t, err := o.Triangle()
	if err != nil {
		fmt.Println(err)
		return
	}
	if det := t.Determinant(); det != 0 {
		roots := t.Roots()
		gauss.PrintRoots(roots)
		t.Mistake(roots)
	} else {
		fmt.Println("СЛАУ имеет бесконечное множество решений")
	}

}
