package main

import "fmt"

type strMap = map[string]string  // type alias, чтобы заменить сложные типы алиасом

func addToSlice(a []string) {
	a[0] = "2"
}

func addToArray(b [1]string) {
	b[0] = "2"
}

func main() {
	a := []string{"1"}  // срез
	addToSlice(a)
	fmt.Println(a)  // [2], содержимое изменится, потому что срезы в Go являются ссылочным типом данных.

	b := [1]string{"1"}  // массив
	addToArray(b)
	fmt.Println(b)  // [1], содержимое не изменится, потому что массивы в Go относятся к агрегирующему типу данных. 
	// Создаётся виртуальная копия массива, но исходный массив не меняется. К этой копии можно было бы добраться через c := addToArray(b).
	// К ссылочным типам относятся срезы, мапы, функции и каналы. 
	// К агрегирующим типам относятся массивы и структуры.

	// Поэтому в приложении для закладок ниже в функциях можно убрать возвраты, т.к. они работают с мапами, с ссылочным типом данных, в которых изменения не нужно сохранять
	// путём переприсваивания. Все операции будут производиться над одним и тем же объектом в памяти, лежащим в переменной bookmarks.
	// Нюанс: т.к. при передаче функции переменной в качестве параметра Go всегда создает копию значения переменной, а срез содержит указатель на массив, длину и ёмкость,
	// то при передаче среза в функцию внутри функции копируется сам срез как структура данных, ссылка на массив остаётся той же. Но если внутри функции изменить длину или
	// ёмкость среза, то эти изменения будут действительны лишь для локальной копии среза внутри функции, внешний срез останется прежним.

	// Приложение для закладок:
	bookmarks := strMap{}  // меняем сложный тип на алиас (и везде ниже)
	Menu: for {  // Menu - пример метки.
		fmt.Println("Приложение для ведения закладок")
		choice := getMenu()
		Switch: switch choice {
		case 1: printBookmarks(bookmarks)
		case 2: bookmarks = addBookmark(bookmarks)  // можно переписать: addBookmark(bookmarks)
		case 3: bookmarks = deleteBookmark(bookmarks)  // можно переписать: deleteBookmark(bookmarks)
		// case 4: break  // в данном случае компилятор не понимает, к чему именно относится break - к конструкции switch-case или к внешнему циклу.
		// Чтобы такой коллизии не происходило, можно использовать метки (labels). Например, как-то назвать цикл, чтобы различать его.
		case 4: break Menu  // с применённой меткой программе понятно, что прервать требуется именованный внешний цикл.
		case 5: break Switch  // а это было бы прерывание работы именованной конструкции switch-case.
		}
	}
}

func getMenu() int {
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Просмотреть закладки;")
	fmt.Println("2. Добавить закладку;")
	fmt.Println("3. Удалить закладку;")
	fmt.Println("4. Выход")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func printBookmarks(bookmarks strMap) {
	if len(bookmarks) == 0 {
		println("Нет добавленных закладок")
		return
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)
	}
}

func addBookmark(bookmarks strMap) strMap {
	var newBookmarkKey, newBookmarkValue string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите значение: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks  // можно убрать возврат в силу того, что мапы - ссылочный тип данных, изменения происходят в одном и том же объекте.
}

func deleteBookmark(bookmarks strMap) strMap {
	var bookmarkKeytoDelete string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&bookmarkKeytoDelete)
	delete(bookmarks, bookmarkKeytoDelete)
	return bookmarks  // можно убрать возврат в силу того, что мапы - ссылочный тип данных, изменения происходят в одном и том же объекте.
}

	// Теория:
	/*
	// maps - структутры, содержащие пары "ключ-значение". Похожи на словари из питона, за обоими сокрыты хеш-таблицы.
	// Отличия:
	// - Мапы строго типизированы, как и всё в Go, в них все ключи могут быть только одного типа, как и все значения. При этом ключи и значения могут быть разных типов.
	// - В Python порядок добавления элементов в словарь сохраняется, начиная с версии 3.7; в мапах Go порядок отображения пар случайный.
	// - В Python поиск несуществующего ключа в словаре выводит KeyError с падением программы; в мапах Go программа не падает, а возвращает нулевое значение для указанного типа.
	// - В Python можно с ходу создать пустой словарь и заполнять его. В Go пустая мапа равна nil. Чтобы заполнить её, нужно сперва проинициализировать её в памяти
	//   с помощью либо функции make(), либо при объявлении с помощью фигурных скобок в конце.

	// mNil := map[string]int  // Объявление nil-мапы
	// mEmpty := map[string]float64{}  // Объявление пустой, но инициализированной мапы
	m := map[string]string{  // объявление непустой мапы с попутным добавлением в неё элементов
		"SVO": "Sheremetyevo, Moscow, Russia",
		"VKO": "Vnukovo, Moscow, Russia",
		"DME": "Domodedovo, Moscow, Russia",
		"LED": "Pulkovo, Saint-Petersburg, Russia",
		"UFA": "Ufa, Ufa, Russia",
		"UUD": "Baikal, Ulan-Ude, Russia",
	} 
	fmt.Println(m)  // map[DME:Domodedovo, Moscow, Russia LED:Pulkovo, Saint-Petersburg, Russia SVO:Sheremetyevo, Moscow, Russia UFA:Ufa, Ufa, Russia UUD:Baikal, Ulan-Ude, Russia VKO:Vnukovo, Moscow, Russia]

	// вывод значения по ключу:
	fmt.Println(m["UFA"])  // Ufa, Ufa, Russia

	// вывод значения по несуществующему ключу:
	fmt.Println(m["GGG"])  //   (пустая строка - значение по умолчанию для строкового типа данных)

	// изменение значения по ключу:
	m["UUD"] = "Baikal, Ulaan-Ude, Russia"
	fmt.Println(m["UUD"])  // Baikal, Ulaan-Ude, Russia

	// Добавление элементов (в непустую мапу):
	m["ULN"] = "Buyant-Ukhaa, Ulaanbaatar, Mongolia"
	m["UBN"] = "Chinggis Khaan, Ulaanbaatar, Mongolia"
	m["ZIA"] = "Zhukovskiy, Moscow, Russia"
	fmt.Println(m)  // map[DME:Domodedovo, Moscow, Russia LED:Pulkovo, Saint-Petersburg, Russia SVO:Sheremetyevo, Moscow, Russia UBN:Chinggis Khaan, Ulaanbaatar, Mongolia UFA:Ufa, Ufa, Russia ULN:Buyant-Ukhaa, Ulaanbaatar, Mongolia UUD:Baikal, Ulaan-Ude, Russia VKO:Vnukovo, Moscow, Russia ZIA:Zhukovskiy, Moscow, Russia]
	
	// Попытка добавления ключа неподходящего типа:
	// m[1] = "Error"  // .\main.go:40:4: cannot use 1 (untyped int constant) as string value in map index

	// Удаление элемента с помощью встроенного метода delete():
	delete(m, "ZIA")  // Первым аргументом принимает мапу, вторым - удаляемый ключ.
	fmt.Println(m)  // map[DME:Domodedovo, Moscow, Russia LED:Pulkovo, Saint-Petersburg, Russia SVO:Sheremetyevo, Moscow, Russia UBN:Chinggis Khaan, Ulaanbaatar, Mongolia UFA:Ufa, Ufa, Russia ULN:Buyant-Ukhaa, Ulaanbaatar, Mongolia UUD:Baikal, Ulaan-Ude, Russia VKO:Vnukovo, Moscow, Russia]

	// Если попробовать удалить несуществующий элемент, то программа не упадёт и не выведет ошибку:
	delete(m, "FKU")  // нет ошибки

	// Один из методов проверки наличия ключа:
	_, exist := m["WOW"]
	fmt.Println(exist)  // false; ключа нет в мапе.

	// Итерация по мапе:
	newMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range newMap {
		fmt.Println(key, value)
	}

	// У мап есть проблемы с выделением памяти. У мап нет ёмкости, им можно задать изначальный размер, но окончательный размер определяется только итоговым кол-вом элементов.
	// У мап есть длина (len()). Если заведомо примерно известен размер мапы, то можно с помощью функции make() задать изначальный размер мапе, чтобы сократить время и ресурс 
	// на аллокацию памяти. 

	// m := make(strMap, 3)

	// К слову, инициализация пустой мапы тоже возможна через make():

	// m := make(map[string]int)

	*/