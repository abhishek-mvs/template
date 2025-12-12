package orderManager

import (
	"template/internal/app/dto"
	"template/internal/app/service/orderService"

	"github.com/google/uuid"
)

type OrderManager struct {
	orderService *orderService.OrderService
}

func NewOrderManager(orderService *orderService.OrderService) *OrderManager {
	return &OrderManager{
		orderService: orderService,
	}
}

func (o *OrderManager) CreateOrder(userId string, restaurantId uuid.UUID, dishId uuid.UUID) (dto.CreateOrderResponse, error) {
	return o.orderService.CreateOrder(userId, restaurantId, dishId)
}

func (o *OrderManager) GetUserOrders(userId string) (dto.UserOrdersResponse, error) {
	return o.orderService.GetUserOrders(userId)
}

