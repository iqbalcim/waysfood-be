package handlers

import (
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var cartModels models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&cartModels); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var cart models.Transaction
	cartExist, err := h.CartRepository.GetChartByUser(userId, productId)
	// fmt.Println(cartExist)

	if err == nil {
		cartModels.UsersID = userId
		cartModels.ProductID = productId
		cartModels.Qty = cartExist.Qty + 1
		cartModels.Price = cartModels.Price + cartExist.Price
		// fmt.Println(cartModels)
		// update cart
		cart, err = h.CartRepository.UpdateCartQty(cartModels, userId, productId)
		fmt.Println(cart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

	} else {
		cartModels.UsersID = userId
		cartModels.ProductID = productId
		cartModels.Qty = 1

		// fmt.Println(cartModels)

		cart, err = h.CartRepository.AddToCart(cartModels)
		fmt.Println(cart)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCart) GetChartByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	cart, err := h.CartRepository.GetChartByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, c := range cart {
		cart[i].Products.Image = os.Getenv("PATH_FILE") + c.Products.Image

	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCart) DeleteChartByQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	cartItem, _ := h.CartRepository.GetChartByUser(userId, productId)

	var cartUpdate models.Transaction

	var cart models.Transaction
	var updateErr error

	cartUpdate.Qty = cartItem.Qty - 1
	cartUpdate.Price = cartItem.Price - (cartItem.Price / cartItem.Qty)
	cart, updateErr = h.CartRepository.UpdateCartQty(cartUpdate, userId, productId)

	if updateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: updateErr.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteChartByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	productId, _ := strconv.Atoi(mux.Vars(r)["productID"])

	cartItem, _ := h.CartRepository.GetChartByUser(userId, productId)

	cartDelete, err := h.CartRepository.DeleteCartByID(cartItem, cartItem.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: cartDelete}
	json.NewEncoder(w).Encode(response)
}
