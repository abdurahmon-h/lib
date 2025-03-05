package main

import "fmt"

func movieData() map[string]string {
	categories := make(map[string]string)

	categories["1"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["2"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["3"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["4"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["5"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["6"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["7"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"
	categories["8"] = "https://hd.kinopoisk.ru/ru-tj/selection/138"

	return categories
}

func findMovieByCat(categories map[string]string, movie string) string {
	if link, exists := categories[movie]; exists {
		return link
	}
	return ""
}

func main() {
	categories := movieData()
	var movieInput string
	fmt.Println("1. Приключения\n2. Детективы\n3. Комедии\n4. Мультфильмы\n5. Фэнтези\n6. Боевики\n7. Романтические комедии\n8. Фантастика")
	fmt.Print("Выберите категорию: ")
	fmt.Scan(&movieInput)

	movieLink := findMovieByCat(categories, movieInput)
	if movieLink != "" {
		fmt.Println("Ссылка на фильм:", movieLink)
	} else {
		fmt.Println("Фильм не найден")
	}
}
