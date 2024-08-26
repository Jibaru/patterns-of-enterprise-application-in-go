package domain

type Artist struct {
	Name string
}

type Album struct {
	Title  string
	Artist Artist
}
