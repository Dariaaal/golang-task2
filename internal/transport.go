package main

import "fmt"

var cover = []Cover{
	Cover{
		"url": "http://localhost:8090/shop/demo/product_2.png",
		"type": "photo",
	},
	Cover{
		"url": "http://localhost:8090/shop/demo/product_9.png",
		"type": "photo",
	},
	Cover{
		"url": "http://localhost:8090/shop/demo/product_7.png",
		"type": "photo",
	},
}

var product = Product{
		"uid": "2",
		"cover": cover,
		"title": "Золоте кольє",
		"description": "Романтичний та тендітний акцент твоїх образів.",
		"price": 1699
	  }

