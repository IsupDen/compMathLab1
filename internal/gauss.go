package gauss

import (
	"errors"
	"fmt"
	"math"
)

func PrintMatrix(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%.2f\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func PrintRoots(roots []float64) {
	fmt.Println("\nКорни СЛАУ:")
	for i := 0; i < len(roots); i++ {
		fmt.Printf("x%v = %.2f\n", i+1, roots[i])
	}
}

type TriangleMatrix struct {
	matrix [][]float64
}

type OriginalMatrix struct {
	matrix [][]float64
}

func NewOriginalMatrix(m [][]float64) *OriginalMatrix {
	return &OriginalMatrix{m}
}

func (t *TriangleMatrix) Determinant() float64 {
	det := 1.0
	for i := 0; i < len(t.matrix); i++ {
		det *= t.matrix[i][i]
	}
	return det
}

func (t *TriangleMatrix) Mistake(roots []float64) {
	fmt.Println("\nВектор невязки:")
	size := len(t.matrix)
	mistake := make([]float64, size)
	for i := 0; i < size; i++ {
		sum := 0.0
		for j := 0; j < size; j++ {
			sum += roots[j] * t.matrix[i][j]
		}
		mistake[i] = t.matrix[i][size] - sum
		fmt.Println(mistake[i])
	}
}

func (o *OriginalMatrix) Triangle() (*TriangleMatrix, error) {
	size := len(o.matrix)
	triangle := make([][]float64, size)
	copy(triangle, o.matrix)
	for i := 0; i < size; i++ {
		fmt.Println("\n-----------------------------")
		fmt.Printf("\t%v-й шаг\n", i+1)
		fmt.Println("-----------------------------")
		maxRow := 0
		max := 0.0
		for j := i; j < size; j++ {
			if max < math.Abs(triangle[j][i]) {
				max = math.Abs(triangle[j][i])
				maxRow = j
			}
		}

		fmt.Printf("Максимальный элемент %v-го столбца - %.2f\n", i+1, max)

		if max == 0 {
			return &TriangleMatrix{}, errors.New("нет решений")
		}

		if maxRow != i {
			fmt.Printf("Меняем местами %v-ю и %v-ю строки\n", i+1, maxRow+1)
			triangle[i], triangle[maxRow] = triangle[maxRow], triangle[i]

			fmt.Println("Матрица после перестоновки:")
			PrintMatrix(triangle)
		}

		for j := i + 1; j < size; j++ {
			k := triangle[i][i] / triangle[j][i]
			for l := i; l <= size; l++ {
				triangle[j][l] = triangle[j][l]*k - triangle[i][l]
			}
		}
		fmt.Printf("\nМатрица после %v-го шага:\n", i+1)
		PrintMatrix(triangle)
	}
	fmt.Println("-----------------------------")
	return &TriangleMatrix{triangle}, nil
}

func (t *TriangleMatrix) Roots() []float64 {
	size := len(t.matrix)
	roots := make([]float64, size)
	for i := size - 1; i >= 0; i-- {
		root := t.matrix[i][size]
		for j := i + 1; j < size; j++ {
			root -= roots[j] * t.matrix[i][j]
		}
		roots[i] = root / t.matrix[i][i]
	}
	return roots
}
