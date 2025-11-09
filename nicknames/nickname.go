package nicknames

import (
	"fmt"
	"math/rand"
)

func GenerateNicknames() map[string]bool {
	var firstPart []string = []string{
		"Чёрный", "Белый", "Красный", "Синий", "Золотой",
		"Дикий", "Спокойный", "Быстрый", "Сильный", "Хитрый",
	}

	var secondPart []string = []string{
		"Волк", "Орёл", "Медведь", "Лис", "Тигр",
		"Ворон", "Сокол", "Барс", "Гриф", "Кедр",
	}
	result := make(map[string]bool)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			nick := firstPart[i] + " " + secondPart[j]
			result[nick] = true
		}

	}
	return result
}

func GetRandomNickname(nickNames map[string]bool) (string, error) {
	var freeNicks []string
	for nick, isFree := range nickNames {
		if isFree == true {
			freeNicks = append(freeNicks, nick)
		}
	}
	if len(freeNicks) == 0 {
		fmt.Println("ошибка нет доступных ников")
		return "", fmt.Errorf("ошибка нет доступных ников")
	}
	randomNumber := rand.Intn(len(freeNicks))
	nickName := freeNicks[randomNumber]
	nickNames[nickName] = false
	return nickName, nil
	// Находим свободный ник
	// Отмечаем как занятый
	// Возвращаем ник (и может быть ошибку)
}

func ReleaseNickname() error {

}
