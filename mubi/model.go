package mubi

type Metadata struct {
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	TotalCount  int `json:"total_count"`
}

type MetaCursor struct {
	TotalCount int  `json:"total_count"`
	NextCursor *int `json:"next_cursor"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserListsResponse struct {
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

type UserRatings struct {
	Ratings []RatingResponse `json:"ratings"`
	Meta    MetaCursor       `json:"meta"`
}

type RatingResponse struct {
	FilmId int `json:"film_id"`
	Stars  int `json:"overall"`
}

// RatingsLookup Map from `FilmId` to `Stars`
type RatingsLookup map[int]int

type DiaryEntry struct {
	Film        Film
	WatchedDate string
	Rating      *int
}
