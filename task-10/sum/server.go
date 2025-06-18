package task10

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrInvalidValue = fmt.Errorf("value must be positive")
)

type Request struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Response struct {
	Result int `json:"result"`
	Error string `json:"error"`
}

func MultiplyInteger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		err := render.Decode(r, &req)
		if err != nil {
			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, Response{
				Result: 0,
				Error: "failed to decode request",
			})

			return
		}

		if req.X < 0 || req.Y < 0 {
			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, Response{
				Result: 0,
				Error: "value must be positive",
			})

			return
		}

		res := req.X + req.Y

		render.JSON(w, r, Response{
			Result: res,
		})
	}
}