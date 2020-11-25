package combination

import (
	"errors"
	"strings"
)

// параметры
type setting struct {
	long int
	com  []string
}

type str struct {
	s string
}

// допустимые символы
var pull = setting{
	20,
	[]string{

		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",

		"a", "b", "c", "d",
		"e", "f", "g", "h",
		"i", "j", "k", "l",
		"m", "n", "o", "p",
		"q", "r", "s", "t",
		"u", "v", "w", "x",
		"y", "z",

		"-", "_", "^", "%",
	},
}

// строка в массив
func strToMass(s string) []string {

	var m []string

	for _, r := range s {
		m = append(m, string(r))
	}
	return m

}

//убираем повторяющиеся символы из массива (Внимание!!! итерация по картам происходит не по порядку, поэтому порядок в массве меняется это не допустимо надо переделать без карт)
func clearMass() {

	var clear []string

	chek := func(s string) bool {
		for _, r := range clear {
			if s == r {
				return true
			}
		}

		return false
	}

	for _, v := range pull.com {

		if !chek(v) {
			clear = append(clear, v)
		}

	}

	// обновляем массив
	pull.com = clear

}

// задать свой надор символов и динну
func Set(l int, m interface{}) {

	switch f := m.(type) {
	case string:
		//преобразуем в массив
		pull = setting{l, strToMass(f)}
	case []string:
		pull = setting{l, f}
	default:
		panic("Задан не подходящий типа набора символов")
	}

	// чистим массив от повторений
	clearMass()

}

//проверка строки на количество символов
//и на наличие неизвестных символов
func (st *str) check() error {

	//----------------------------------------
	//1. проверяем на количество
	if len(st.s) > pull.long {
		return errors.New("Ошибка: Привышенно допустимое количество символов в строке ")
	}

	//-----------------------------------------
	//2. проверяем наличие неизвестных символов
	var r string      //символ заданной строки
	var fail []string // массив неизвестных символов

	// передираем символы в заданной строке
OSN:
	for _, y := range st.s {

		//rune to string
		r = string(y)

		// если нашли сивол из заданной строки в перечне допустимых
		for _, x := range pull.com {
			if r == x {
				continue OSN
			}
		}

		// если символ не найденн
		// добовляем в список неизвестных символов
		fail = append(fail, r)
	}

	// проверим были ли найденны неизвестные символы
	if len(fail) > 0 {
		return errors.New("Ошибка: присутствуют неизвестные символы: " + strings.Join(fail, ", "))
	}

	return nil

}

// возвращаем следующий символ
func lastS(s string) (string, error) {

	for i, r := range pull.com {
		if r == s {
			if i < len(pull.com)-1 {
				return pull.com[i+1], nil
			} else {
				return "-1", errors.New("последний символ")
			}
		}
	}

	return "", errors.New("Ошибка: последний символ заданной строки, не обноружен в допустимых ( " + s + " )")

}

// сброс символов на первые
// r - массив рун заданной строки
// i - порядковый номер элемента массива рун (r), с которого надо начать сброс на первый символ
func reset(r []rune, i int) []rune {

	y := []rune(pull.com[0])[0]

	for t := i; t < len(r); t++ {
		r[t] = y
	}

	return r

}

// добавляем одну позицию к текушей строуке, так как закончились возможные комбинации
// i - длинна прежней строки
func (st *str) add() string {

	var a string

	// формируем новую строку с добавлением еще одной позиции
	for y := 0; y <= len(st.s); y++ {
		a += pull.com[0]
	}

	return a
}

// получаем следующую комбинацию
func (st *str) next() (string, error) {

	na := []rune(st.s)

	for i := len(na) - 1; i >= 0; i-- {
		//fmt.Println(string(na[i]))

		l, err := lastS(string(na[i]))

		// если есть следующий символ
		if err == nil {
			na[i] = []rune(l)[0]
			return string(reset(na, i+1)), nil
		}

		//если это последний символ
		if err != nil {
			if l == "-1" {
				if i == 0 {
					// это значит что первый символ в заданной строке последний и его не возможно поменять на следующий
					// значит больше нет возможных комбинаций
					// значит надо добавить еще одну позицию если это позволяют условия
					if len(na) == pull.long {
						return "", errors.New("Болше комбинаций не существует ")
					}
					return st.add(), nil
				}
			} else {
				return "", errors.New("Ошибка такого символа нет ")
			}
		}
	}

	return "", errors.New("Что то побло не так ")
}

// основная функция выводит следующую комбинацию относительно переданной
func Next(s string) (string, error) {

	sroc := &str{s}

	// преверим переданную строку
	if err := sroc.check(); err != nil {
		return s, err
	}

	// получим новую комбинацию
	if n, err := sroc.next(); err != nil {
		return s, err
	} else {
		return n, nil
	}

	return s, errors.New("Ошибка: основная функция Next завершилась ничем ")

}

// получить первый символ из набра
func GetFirst() string {
	return pull.com[0]
}
