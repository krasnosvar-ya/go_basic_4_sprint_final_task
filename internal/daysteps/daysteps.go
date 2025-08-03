package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
	// https://forum.golangbridge.org/t/how-to-interpolate-multiline-string/22470
	returnDayActionInfoString = `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slicedString := strings.Split(data, ",")
	// for _, v := range slicedString {

	// Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	if len(slicedString) != 2 {
		err := fmt.Errorf("String %s has not right format, slice %v not 2", data, slicedString)
		return 0, 0, err
	}

	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	firstReturn, err := strconv.Atoi(slicedString[0]) // Atoi returns an int and an error
	if err != nil {
		// test expects "log" not "fmt"- should mention in task(((
		log.Println("Error converting 1st string to int:", err)
		return 0, 0, err
	}
	// Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.
	if firstReturn <= 0 {
		err := fmt.Errorf("Steps count is %d:", firstReturn)
		return 0, 0, err
	}
	// Преобразовать второй элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	// https://www.geeksforgeeks.org/go-language/time-parseduration-function-in-golang-with-examples/
	secondReturn, err := time.ParseDuration(slicedString[1])
	if err != nil {
		log.Println("Error converting 2nd string to time:", err)
		return 0, 0, err
	}
	// Проверить: продолжительность должна быть больше 0. Если это не так, вернуть нули и ошибку.
	if secondReturn <= 0 {
		err := fmt.Errorf("Duration must be greater than 0, got: %v", secondReturn)
		return 0, 0, err
	}
	return firstReturn, secondReturn, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	// Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage().
	// В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	steps, timeFromParsePackage, err := parsePackage(data)
	if err != nil {
		log.Println("Error fetching data from parsePackage:", err)
		// https://go.dev/doc/tutorial/handle-errors
		return ""
	}

	// Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага.
	// Константа stepLength (длина шага) уже определена в коде.
	// Перевести дистанцию в километры, разделив её на число метров в километре (константа mInKm, определена в пакете).
	distance := (float64(steps) * stepLength) / float64(mInKm)
	// Вычислить количество калорий, потраченных на прогулке. Функция для вычисления калорий WalkingSpentCalories()
	// будет определена в пакете spentcalories, которую вы тоже реализуете.

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, timeFromParsePackage)
	if err != nil {
		log.Println("Error calculating calories:", err)
		return ""
	}
	return fmt.Sprintf(returnDayActionInfoString, steps, distance, calories)
}
