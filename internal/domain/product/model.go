package internal

type Product struct {
	ID          string  `json:"uid"`
	Cover       []Cover `json:"cover"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
}

type Cover struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}
