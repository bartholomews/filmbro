package mubi

type Metadata struct {
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	TotalCount  int `json:"total_count"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserLists struct {
	Lists []List   `json:"lists"`
	Meta  Metadata `json:"meta"`
}

type List struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	FilmIds []int  `json:"film_ids"`
}

type Film struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	Year          int    `json:"year"`
}
