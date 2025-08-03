package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
	// Тип тренировки: Бег
	// Длительность: 0.75 ч.
	// Дистанция: 10.00 км.
	// Скорость: 13.34 км/ч
	// Сожгли калорий: 18621.75
	returnTrainingInfoString = `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	// Разделить строку на слайс строк.
	slicedString := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.
	if len(slicedString) != 3 {
		err := fmt.Errorf("String %s has not right format, slice %v not 3", data, slicedString)
		return 0, "", 0, err
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. Обработать возможные ошибки.
	// При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	firstReturn, err := strconv.Atoi(slicedString[0]) // Atoi returns an int and an error
	if err != nil {
		fmt.Println("Error converting 1st string to int:", err)
		return 0, "", 0, err
	}
	// Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.
	if firstReturn <= 0 {
		err := fmt.Errorf("Steps count must be greater than 0, got: %d", firstReturn)
		return 0, "", 0, err
	}
	// Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	thirdReturn, err := time.ParseDuration(slicedString[2])
	if err != nil {
		fmt.Println("Error converting 3rd string to time:", err)
		return 0, "", 0, err
	}
	// Проверить: продолжительность должна быть больше 0. Если это не так, вернуть нули и ошибку.
	if thirdReturn <= 0 {
		err := fmt.Errorf("Duration must be greater than 0, got: %v", thirdReturn)
		return 0, "", 0, err
	}

	// Если всё прошло без ошибок, верните количество шагов, вид активности, продолжительность и nil (для ошибки).
	return firstReturn, slicedString[1], thirdReturn, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	// 	рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	// Соответствующая константа уже определена в пакете.
	stepLength := height * stepLengthCoefficient
	// умножьте пройденное количество шагов на длину шага.
	// разделите полученное значение на число метров в километре (mInKm, константа определена в пакете).
	// Обратите внимание, что целочисленную переменную steps необходимо будет привести к другому числовому типу.
	returnVal := (float64(steps) * stepLength) / float64(mInKm)
	return returnVal
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	// Вычислить дистанцию с помощью distance().
	distanceVariable := distance(steps, height)
	// Вычислить и вернуть среднюю скорость. Для этого разделите дистанцию на продолжительность в часах.
	// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	// https://cs.opensource.google/go/go/+/go1.24.5:src/time/time.go;l=1100
	hoursDuration := duration.Hours() // Get the duration in hours as a float64
	averageSpeed := distanceVariable / hoursDuration
	return averageSpeed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Алгоритм реализации функции:
	// Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	// реализовал через switch, для тренировки, а все до этого через if
	// про параметры не совсем понял, спросил у chatgpt, то что он предложил показалось ок
	switch {
	case steps <= 0:
		return 0, errors.New("steps must be greater than 0")
	case weight <= 0 || weight > 500:
		return 0, errors.New("weight must be between 0 and 500 kg")
	case height <= 0.5 || height > 2.5:
		return 0, errors.New("height must be between 0.5 and 2.5 meters")
	case duration <= 0:
		return 0, errors.New("duration must be greater than 0")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	averageSpeed := meanSpeed(steps, height, duration)
	// Рассчитать и вернуть количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	caloriesCount := (weight * averageSpeed * duration.Minutes()) / 60.0
	return caloriesCount, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// 	Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	switch {
	case steps <= 0:
		// Функция errors.New из пакета errors: Создаёт новый объект ошибки (error) с указанным сообщением.
		return 0, errors.New("steps must be greater than 0")
	case weight <= 0 || weight > 500:
		return 0, errors.New("weight must be between 0 and 500 kg")
	case height <= 0.5 || height > 2.5:
		return 0, errors.New("height must be between 0.5 and 2.5 meters")
	case duration <= 0:
		return 0, errors.New("duration must be greater than 0")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	averageSpeed := meanSpeed(steps, height, duration)
	// Рассчитать количество калорий. Для этого:
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient.
	// Соответствующая константа объявлена в пакете. Вернуть полученное значение.
	caloriesCount := ((weight * averageSpeed * duration.Minutes()) / 60.0) * walkingCaloriesCoefficient
	return caloriesCount, nil

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	// Получить значения из строки данных с помощью функции parseTraining(),
	// обработать возможные ошибки и вывести их в лог с помощью log.Println(err).
	steps, traininType, trainTiming, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Проверить, какой вид тренировки был передан в строке, которую парсили (лучше использовать switch).
	// Для каждого из видов тренировки рассчитать дистанцию, среднюю скорость и калории.
	distancia := distance(steps, height)
	speed := meanSpeed(steps, height, trainTiming)
	// Для каждого вида тренировки сформировать и вернуть строку, образец которой был представлен выше.
	// Если был передан неизвестный тип тренировки, вернуть ошибку с текстом неизвестный тип тренировки.
	// var caloriesBurned float64
	switch {
	case traininType == "Ходьба":
		caloriesBurned, err := WalkingSpentCalories(steps, weight, height, trainTiming)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(returnTrainingInfoString, traininType, trainTiming.Hours(), distancia, speed, caloriesBurned), nil
	case traininType == "Бег":
		caloriesBurned, err := RunningSpentCalories(steps, weight, height, trainTiming)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(returnTrainingInfoString, traininType, trainTiming.Hours(), distancia, speed, caloriesBurned), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// Пример возвращаемой строки:
	// Тип тренировки: Бег
	// Длительность: 0.75 ч.
	// Дистанция: 10.00 км.
	// Скорость: 13.34 км/ч
	// Сожгли калорий: 18621.75
	// return fmt.Sprintf(returnTrainingInfoString, traininType, trainTiming.Hours(), distancia, speed, caloriesBurned), nil

}
