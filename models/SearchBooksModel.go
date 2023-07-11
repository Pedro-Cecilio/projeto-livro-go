package models

type SearchBooksModel struct {
	TotalItems int    `json:"totalItems"`
	Items      []Book `json:"items"`
}

type Book struct {
	Id         string     `json:"id"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title         string     `json:"title"`
	Authors       []string   `json:"authors"`
	Publisher     string     `json:"publisher"`
	PublishedDate string     `json:"publishedDate"`
	Description   string     `json:"description"`
	PageCount     int        `json:"pageCount"`
	Categories    []string   `json:"categories"`
	ImageLinks    ImageLinks `json:"imageLinks"`
	Language      string     `json:"language"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

