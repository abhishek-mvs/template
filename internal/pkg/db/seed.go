package db

import (
	"template/internal/app/dto"
)

// SeedData initializes the database with sample restaurants and dishes
func SeedData(db *DB) {
	// Create sample restaurants
	restaurant1 := &dto.Restaurant{
		Name:        "Pizza Palace",
		Description: "Authentic Italian pizzas",
		Address:     "123 Main St, City",
		Rating:      4.5,
	}
	id1, _ := db.Save("restaurants", restaurant1)
	restaurant1.Id = id1

	restaurant2 := &dto.Restaurant{
		Name:        "Burger King",
		Description: "Delicious burgers and fries",
		Address:     "456 Oak Ave, City",
		Rating:      4.2,
	}
	id2, _ := db.Save("restaurants", restaurant2)
	restaurant2.Id = id2

	restaurant3 := &dto.Restaurant{
		Name:        "Sushi House",
		Description: "Fresh sushi and Japanese cuisine",
		Address:     "789 Pine Rd, City",
		Rating:      4.8,
	}
	id3, _ := db.Save("restaurants", restaurant3)
	restaurant3.Id = id3

	// Create dishes for Pizza Palace
	dish1 := &dto.Dish{
		RestaurantId: id1,
		Name:         "Margherita Pizza",
		Description:  "Classic tomato, mozzarella, and basil",
		Price:        12.99,
		IsAvailable:  true,
	}
	dishId1, _ := db.Save("dishes", dish1)
	dish1.Id = dishId1

	dish2 := &dto.Dish{
		RestaurantId: id1,
		Name:         "Pepperoni Pizza",
		Description:  "Pepperoni and mozzarella cheese",
		Price:        14.99,
		IsAvailable:  true,
	}
	dishId2, _ := db.Save("dishes", dish2)
	dish2.Id = dishId2

	dish3 := &dto.Dish{
		RestaurantId: id1,
		Name:         "Veggie Supreme",
		Description:  "Loaded with vegetables",
		Price:        13.99,
		IsAvailable:  true,
	}
	dishId3, _ := db.Save("dishes", dish3)
	dish3.Id = dishId3

	// Create dishes for Burger King
	dish4 := &dto.Dish{
		RestaurantId: id2,
		Name:         "Classic Burger",
		Description:  "Beef patty with lettuce, tomato, and special sauce",
		Price:        8.99,
		IsAvailable:  true,
	}
	dishId4, _ := db.Save("dishes", dish4)
	dish4.Id = dishId4

	dish5 := &dto.Dish{
		RestaurantId: id2,
		Name:         "Chicken Burger",
		Description:  "Grilled chicken with mayo",
		Price:        9.99,
		IsAvailable:  true,
	}
	dishId5, _ := db.Save("dishes", dish5)
	dish5.Id = dishId5

	dish6 := &dto.Dish{
		RestaurantId: id2,
		Name:         "French Fries",
		Description:  "Crispy golden fries",
		Price:        3.99,
		IsAvailable:  true,
	}
	dishId6, _ := db.Save("dishes", dish6)
	dish6.Id = dishId6

	// Create dishes for Sushi House
	dish7 := &dto.Dish{
		RestaurantId: id3,
		Name:         "Salmon Sushi Roll",
		Description:  "Fresh salmon with rice and seaweed",
		Price:        15.99,
		IsAvailable:  true,
	}
	dishId7, _ := db.Save("dishes", dish7)
	dish7.Id = dishId7

	dish8 := &dto.Dish{
		RestaurantId: id3,
		Name:         "Tuna Sashimi",
		Description:  "Fresh tuna slices",
		Price:        18.99,
		IsAvailable:  true,
	}
	dishId8, _ := db.Save("dishes", dish8)
	dish8.Id = dishId8

	dish9 := &dto.Dish{
		RestaurantId: id3,
		Name:         "California Roll",
		Description:  "Crab, avocado, and cucumber",
		Price:        12.99,
		IsAvailable:  true,
	}
	dishId9, _ := db.Save("dishes", dish9)
	dish9.Id = dishId9
}

