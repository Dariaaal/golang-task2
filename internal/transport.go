package internal

import (
	"context"
	"encoding/json"

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

func ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        productID := chi.URLParam(r, "id")

		ctx := context.WithValue(r.Context(), "id", productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.Context().Value("id").(*Product)

	t.log.Debug("get product", zap.String("id", productID.ID))

	// for _, product := range products {
	// 	if product.ID == productID.ID {
	// 		json.NewEncoder(w).Encode(product)
	// 	}
	// }

	if productID != nil {
		foundProduct, err := json.Marshal(products)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(foundProduct)
	}

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
