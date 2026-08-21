

package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	// fmt.Printf("Привет")

	// Переменные перенесены в функцию обработки ввода
	/*
		var userHeight float64 // по умолчанию 0.0, если не будет присвоено значение
		var userWeight float64
	*/
	fmt.Printf(`__ Калькулятор индекса массы тела __
Подсчитывает степень нормальности Вашего веса.
Потребуются Ваши рост и вес.`) // курсивные кавычки `` позволяют выводить строки в несколько рядов. Но требуется убрать отступы в начале строк, символ переноса \n не работает в этом случае.

	// ввод пользователя выведен в соответствующую функцию обработки
	/*
		fmt.Print("\nВведите свой рост (в сантиметрах): ")
		fmt.Scan(&userHeight) // метод для ввода и сохранения в переменную. Принимает указатель на переменную, а не саму переменную.
		fmt.Print("Введите свой вес (в килограммах): ")
		fmt.Scan(&userWeight)
	*/

	userWeight, userHeight := getUserInput()

	// Код расчёта IMT вынесен в отдельную функцию
	/*
		IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	*/

	IMT := calculateIMT(userWeight, userHeight)

	// fmt.Println("Ваш индекс массы тела:", IMT) // аналогично: fmt.Printf("Ваш индекс массы тела: %.2f", IMT)
	// fmt.Sprint, Sprintf, Sprintln позволяют сохранять вывод в переменную

	// Код вывода результата выведен в отдельную функцию обучения ради:
	/*
		result := fmt.Sprintf("Ваш индекс массы тела: %.2f", IMT)
		fmt.Print(result)
	*/

	outputResult(IMT)
}

func calculateIMT(userWeight, userHeight float64) float64 { // через пробел после области аргументов указывается тип возвращаемого функцией значения
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT
}

// Есть альтернативный вариант возврата значения - с именованным возвратом:
/*
func calculateIMT(userWeight, userHeight float64) (IMT float64) {  // Во вторых скобках для возврата указывается не просто тип, а и переменная, в которую будет сохранено возвращаемое значение.
	IMT = userWeight / math.Pow(userHeight/100, IMTPower)  // Тогда эта переменная считается созданной, и знак определения с присваиванием меняется на просто присваивание.
	return  // возвращаемую переменную можно пропустить, т.к. она уже указана
}
*/

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.2f", imt)
	fmt.Print(result)
}

func getUserInput() (float64, float64) { // синтаксис функции с выводом нескольких значений: первые скобки - для аргументов, вторые - для выводимых значений
	var userWeight float64
	var userHeight float64
	fmt.Print("\nВведите свой рост (в сантиметрах): ")
	fmt.Scan(&userHeight) // метод для ввода и сохранения в переменную. Принимает указатель на переменную, а не саму переменную.
	fmt.Print("Введите свой вес (в килограммах): ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight // вывод нескольких значений - через запятую
}
