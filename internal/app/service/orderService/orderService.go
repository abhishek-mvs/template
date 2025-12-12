package orderService

import (
	"errors"
	"time"
	"template/internal/app/dto"
	"template/internal/app/service/dishService"
	"template/internal/app/service/restaurantService"
	"template/internal/pkg/db"

	"github.com/google/uuid"
)

type OrderService struct {
	DB                *db.DB
	DishService       *dishService.DishService
	RestaurantService *restaurantService.RestaurantService
}

func NewOrderService(db *db.DB, dishService *dishService.DishService, restaurantService *restaurantService.RestaurantService) *OrderService {
	return &OrderService{
		DB:                db,
		DishService:       dishService,
		RestaurantService: restaurantService,
	}
}

func (o *OrderService) CreateOrder(userId string, restaurantId uuid.UUID, dishId uuid.UUID) (dto.CreateOrderResponse, error) {
	// Validate restaurant exists
	restaurant, err := o.RestaurantService.GetRestaurantById(restaurantId)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}
	if restaurant == nil {
		return dto.CreateOrderResponse{}, errors.New("restaurant not found")
	}

	// Validate dish exists and belongs to restaurant
	dish, err := o.DishService.GetDishById(dishId)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}
	if dish == nil {
		return dto.CreateOrderResponse{}, errors.New("dish not found")
	}
	if dish.RestaurantId != restaurantId {
		return dto.CreateOrderResponse{}, errors.New("dish does not belong to the restaurant")
	}
	if !dish.IsAvailable {
		return dto.CreateOrderResponse{}, errors.New("dish is not available")
	}

	// Create order
	order := &dto.Order{
		Id:           uuid.New(),
		UserId:       userId,
		RestaurantId: restaurantId,
		DishId:       dishId,
		DishName:     dish.Name,
		Price:        dish.Price,
		Status:       "placed",
		CreatedAt:    time.Now().Format(time.RFC3339),
	}

	id, err := o.DB.Save("orders", order)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}
	order.Id = id

	return dto.CreateOrderResponse{
		Order: *order,
	}, nil
}

func (o *OrderService) GetUserOrders(userId string) (dto.UserOrdersResponse, error) {
	allOrders := o.DB.GetAll("orders")
	
	orders := make([]dto.Order, 0)
	for _, orderVal := range allOrders {
		order, ok := orderVal.(*dto.Order)
		if !ok {
			continue
		}
		if order.UserId == userId {
			orders = append(orders, *order)
		}
	}

	return dto.UserOrdersResponse{
		Orders: orders,
	}, nil
}

