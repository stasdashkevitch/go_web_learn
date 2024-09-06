package handlers

import (
	"micro/data"
	"net/http"
	"reflect"
)

// swagger:route GET / products listProducts
// Returns a list of products
// responses:
// 200: productsResponse

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	lp := data.GetProducts()
	// мы должны взратить клиенту все в виде json
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	// }
	//
	// rw.Write(d)

	// у метода Marshal есть проблема - мы алоцируем память
	// проще использовать Encode который будет сразу записывать результат в output stream

	p.l.Println(reflect.TypeOf(lp))
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
