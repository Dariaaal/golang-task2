package internal

var products = []Product{
	{
		ID: "2",
		Cover: []Cover{
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
		},
		Title:       "Золоте кольє",
		Description: "Романтичний та тендітний акцент твоїх образів.",
		Price:       1699,
	},
}
