package structs

// Book is the representation of Book database table
type Book struct {
	Id          int64       `json:"id"`
	Title       string      `json:"title"`
	Isbn        string      `json:"ibsn"`
	About       string      `json:"about"`
	Edition     int16       `json:"edition"`
	PageNumber  int16       `json:"page_number"`
	ReleaseDate interface{} `json:"release_date"`
	AuthorId    int         `json:"author_id"`
}
