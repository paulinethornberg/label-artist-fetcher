package model

type SongCollection struct {
	Songs []Song `json:"song"`
}

type Song struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Composer    string `json:"composer"`
	RecordLabel string `json:"recordlabel"`
}
