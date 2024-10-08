package cart

import (
	"fmt"

	"github.com/ayinde1993/ecom/types"
)

// this file containe a logiques decitions functions
func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}
	return productIds, nil

}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	// list all the product ID
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInsStock(items, productMap); err != nil {
		return 0, 0, nil
	}
	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity of product in our db
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)

	}

	// create the order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: " 6 rue Delerue",
	})
	if err != nil {
		return 0, 0, err
	}

	// create  order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil

}

func checkIfCartIsInsStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)

		}

	}
	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
