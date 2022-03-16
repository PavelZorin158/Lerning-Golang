/* данные, которые вам потребуется декодировать, содержатся в файле
oked.json. Вам необходимо в качестве ответа записать сумму полей
global_id всех элементов, закодированных в файле.
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// в каждом элементе в файле oked.json массива(среза)
	// есть ключ "global_id" со значением int
	type poked struct {
		// чтобы json.Unmarshal декодировал элемент нашей структуры
		// обязательно должен начинаться с большой буквы.
		// но ему можно присвоить анотацию с маленькой буквой, как в
		// файле
		Global_id int `json:"global_id"`
	}

	// файл состоит из массива структур
	type okd []poked

	var mm okd
	var sum int

	// читаем файл в data
	data, err := os.ReadFile("oked.json")
	if err != nil {
		fmt.Println("ошибка чтения файла \n", err)
		return
	}

	// декодируем в срез наших структур
	err = json.Unmarshal(data, &mm)
	if err != nil {
		fmt.Println("чето пошло не так \n", err)
		return
	}
	fmt.Println(len(mm), " элементов")
	for i := 0; i < len(mm); i++ {
		sum += mm[i].Global_id
	}
	fmt.Println("сумма всех global_id =", sum)
}
