package internal

var cover = []Cover{
	{
		Url:  "http://localhost:8090/shop/demo/product_2.png",
		Type: "photo",
	},
	{
		Url:  "http://localhost:8090/shop/demo/product_9.png",
		Type: "photo",
	},
	{
		Url:  "http://localhost:8090/shop/demo/product_7.png",
		Type: "photo",
	},
}

var products = []Product{
	{
		ID:          "2",
		Cover:       cover,
		Title:       "Золоте кольє",
		Description: "Романтичний та тендітний акцент твоїх образів.",
		Price:       1699,
	},
}
