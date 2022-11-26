package mubi

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserLists struct {
	Lists []List `json:"lists"`
}

type List struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	FilmIds []int  `json:"film_ids"`
}

type Film struct {
	id            int
	title         string
	originalTitle string
	year          int
}
