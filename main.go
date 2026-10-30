package main

import "fmt"

// ---------- Ejercicio 1 ----------

func swap(a *int, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

// ---------- Ejercicio 2 ----------

func sum(values []int) int {
	total := 0
	for i := 0; i < len(values); i++ {
		total += values[i]
	}
	return total
}

func minimum(values []int) int {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func maximum(values []int) int {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func countNegatives(values []int) int {
	count := 0
	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			count++
		}
	}
	return count
}

// recorre de atrás para adelante sin tocar el slice
func printReverse(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Print(values[i], " ")
	}
	fmt.Println()
}

// ---------- Ejercicio 3 ----------

func copyArray(values []int) []int {
	result := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = values[i]
	}
	return result
}

// ---------- Ejercicio 4 ----------

// inserción, ordena el slice que le pasen
func sortAscending(values []int) {
	for i := 1; i < len(values); i++ {
		current := values[i]
		j := i - 1
		for j >= 0 && values[j] > current {
			values[j+1] = values[j]
			j--
		}
		values[j+1] = current
	}
}

// ---------- Ejercicio 5 ----------

// values debe venir ordenado de menor a mayor
func findPair(values []int, target int) bool {
	left := 0
	right := len(values) - 1
	for left < right {
		s := values[left] + values[right]
		if s == target {
			fmt.Printf("Par: (%d, %d)\n", values[left], values[right])
			return true
		} else if s < target {
			left++
		} else {
			right--
		}
	}
	return false
}

// ---------- Ejercicio 6 ----------

func threeSum(values []int) int {
	if len(values) < 3 {
		return 0
	}
	sorted := copyArray(values)
	sortAscending(sorted)

	count := 0
	n := len(sorted)
	for i := 0; i < n-2; i++ {
		// si el menor ya es positivo no hay forma de llegar a cero
		if sorted[i] > 0 {
			break
		}
		// mismo primer valor que antes: ya se probó
		if i > 0 && sorted[i] == sorted[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			s := sorted[i] + sorted[left] + sorted[right]
			if s < 0 {
				left++
			} else if s > 0 {
				right--
			} else {
				fmt.Printf("(%d, %d, %d)\n", sorted[i], sorted[left], sorted[right])
				count++
				left++
				right--
				// saltar repetidos para no sacar la misma terna otra vez
				for left < right && sorted[left] == sorted[left-1] {
					left++
				}
				for left < right && sorted[right] == sorted[right+1] {
					right--
				}
			}
		}
	}
	return count
}

// ---------- Ejercicio 7 ----------

func buildLeftMax(height []int, leftMax []int) {
	if len(height) == 0 {
		return
	}
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
}

func buildRightMax(height []int, rightMax []int) {
	n := len(height)
	if n == 0 {
		return
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
}

// ---------- Ejercicio 8 ----------

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// con verbose = true imprime el cálculo de cada posición
func trappedWaterDetail(height []int, verbose bool) int64 {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	total := int64(0)
	for i := 0; i < n; i++ {
		level := minInt(leftMax[i], rightMax[i])
		water := level - height[i]
		if verbose {
			fmt.Printf("  pos %d: altura=%d leftMax=%d rightMax=%d nivel=%d agua=%d\n",
				i, height[i], leftMax[i], rightMax[i], level, water)
		}
		total += int64(water)
	}
	return total
}

func trappedWater(height []int) int64 {
	return trappedWaterDetail(height, false)
}

// ---------- Pruebas ----------

func status(ok bool) string {
	if ok {
		return "OK"
	}
	return "FALLO"
}

func testPair(values []int, target int, expected bool) {
	fmt.Println("findPair", values, "objetivo", target)
	got := findPair(values, target)
	if !got {
		fmt.Println("No existe")
	}
	fmt.Println(" ->", got, status(got == expected))
}

func testThreeSum(values []int, expected int) {
	fmt.Println("threeSum", values)
	got := threeSum(values)
	if got == 0 {
		fmt.Println("Ninguna")
	}
	fmt.Println(" -> cantidad:", got, status(got == expected))
	fmt.Println(" -> entrada intacta:", values)
}

func testWater(height []int, expected int64) {
	got := trappedWater(height)
	fmt.Println("agua", height, "->", got, status(got == expected))
}

func main() {
	values := []int{4, -2, 0, 7, -5}
	fmt.Println("Entrada:", values)

	fmt.Println("\n== Ejercicio 1 ==")
	a := 8
	b := -3
	fmt.Println("antes:  a =", a, "b =", b)
	swap(&a, &b)
	fmt.Println("despues: a =", a, "b =", b)

	fmt.Println("\n== Ejercicio 2 ==")
	fmt.Println("Suma:", sum(values))
	fmt.Println("Minimo:", minimum(values))
	fmt.Println("Maximo:", maximum(values))
	fmt.Println("Negativos:", countNegatives(values))
	fmt.Print("Recorrido inverso: ")
	printReverse(values)
	fmt.Println("Entrada despues:", values)

	fmt.Println("\n== Ejercicio 3 ==")
	copia := copyArray(values)
	copia[0] = 99
	fmt.Println("original:", values)
	fmt.Println("copia:   ", copia)

	fmt.Println("\n== Ejercicio 4 ==")
	ordenado := copyArray(values)
	sortAscending(ordenado)
	fmt.Println("ordenado:", ordenado)
	fmt.Println("original:", values)

	fmt.Println("\n== Ejercicio 5 ==")
	testPair([]int{-4, -1, 0, 2, 5}, 1, true)
	testPair([]int{-4, -1, 0, 2, 5}, 8, false)
	testPair([]int{0, 0}, 0, true)
	testPair([]int{3}, 6, false)
	testPair([]int{}, 0, false) // tamaño insuficiente

	fmt.Println("\n== Ejercicio 6 ==")
	testThreeSum([]int{-1, 0, 1, 2, -1, -4}, 2)
	testThreeSum([]int{0, 1, 1}, 0)
	testThreeSum([]int{0, 0, 0}, 1)
	testThreeSum([]int{0, 0, 0, 0}, 1)
	testThreeSum([]int{-2, 0, 0, 2, 2}, 1)
	testThreeSum([]int{1, -1}, 0) // tamaño insuficiente

	fmt.Println("\n== Ejercicio 7 ==")
	height := []int{3, 0, 2, 0, 4}
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)
	fmt.Println("height: ", height)
	fmt.Println("leftMax: ", leftMax)
	fmt.Println("rightMax:", rightMax)

	fmt.Println("\n== Ejercicio 8 ==")
	testWater([]int{5}, 0)
	testWater([]int{0, 1, 2, 3}, 0)
	testWater([]int{3, 2, 1, 0}, 0)
	testWater([]int{2, 2, 2}, 0)
	testWater([]int{3, 0, 3}, 3)
	testWater([]int{3, 0, 2, 0, 4}, 7)
	testWater([]int{4, 2, 0, 3, 2, 5}, 9)
	testWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)

	// caso propio: el nivel lo marca el 4, no el 5
	testWater([]int{5, 1, 3, 1, 4}, 7)

	fmt.Println("\nDetalle por posicion de [4 2 0 3 2 5]:")
	total := trappedWaterDetail([]int{4, 2, 0, 3, 2, 5}, true)
	fmt.Println("  total:", total)
}
