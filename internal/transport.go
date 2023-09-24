package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type ProductTransport struct {
	log *zap.Logger
}

func newProductTransport(log *zap.Logger) ProductTransport {
	return ProductTransport{
		log: log,
	}
}

func (t *ProductTransport) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "productId")
		t.log.Debug("product id", zap.String("id", productID))
		ctx := context.WithValue(r.Context(), "id", productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	foundProducts, err := json.Marshal(products)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
	w.Write(foundProducts)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
	var newProduct *Product
	json.Unmarshal(reqBody, &newProduct)
	products = append(products, newProduct)
	json.NewEncoder(w).Encode(newProduct)
}

func (t *ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {

	rowProductID := r.Context().Value("id")

	if rowProductID == nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	productId := rowProductID.(string)

	t.log.Debug("get product", zap.String("id", productId))

	for _, product := range products {
		if product.ID == productId {
			responseData, err := json.Marshal(product)
			if err != nil {
				http.Error(w, http.StatusText(400), 400)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseData)
			return
		}
	}

	http.Error(w, http.StatusText(400), 400)
}

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
//     var product Product
//     id := chi.URLParam(r, "id")
//     json.NewDecoder(r.Body).Decode(&post)

//     query, err := db.Prepare("Update posts set title=?, content=? where id=?")
//     catch(err)
//     _, er := query.Exec(post.Title, post.Content, id)
//     catch(er)

//     defer query.Close()

//     respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
//     id := chi.URLParam(r, "id")

//     query, err := db.Prepare("delete from posts where id=?")
//     catch(err)
//     _, er := query.Exec(id)
//     catch(er)
//     query.Close()

//     respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
// }
