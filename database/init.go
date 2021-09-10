package database

import (
	"GoIndustryProject/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Initialize tables and admin account
func Init() {
	CreateUserTable()
	CreateSessionTable()
	CreateRestaurantTable()
	CreateFoodTable()
	CreateOrderTable()
	CreateOrderItemTable()

	CreateAdminAccount()
}

// Creates initial admin account. If account already exist, error will be printed.
func CreateAdminAccount() {
	bPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	myUser := models.User{
		Username: "admin",
		Password: bPassword,
		First:    "first",
		Last:     "last",
	}
	err := InsertUser(myUser) //previously mapUsers["admin"] = myUser
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Admin Account Created")
	}
}

// Creates "users" table in database
func CreateUserTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"users" +
		" (" +
		"Username VARCHAR(255) PRIMARY KEY, " +
		"Password BLOB, " +
		"First VARCHAR(255), " +
		"Last VARCHAR(255), " +
		"Gender VARCHAR(6), " +
		"Birthday DATE, " +
		"Height SMALLINT UNSIGNED, " +
		"Weight SMALLINT UNSIGNED, " +
		"ActivityLevel SMALLINT UNSIGNED, " +
		"CaloriesPerDay FLOAT UNSIGNED, " +
		"Halal BOOL, " +
		"Vegan BOOL, " +
		"Address VARCHAR(255), " +
		"PostalCode MEDIUMINT UNSIGNED, " +
		"Lat FLOAT, " +
		"Lng FLOAT" +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: users")
	}
}

// Creates "sessions" table in database
func CreateSessionTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"sessions" +
		" (" +
		"UUID VARCHAR(255) PRIMARY KEY, " +
		"Username VARCHAR(255)" +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: sessions")
	}
}

// Creates "restaurants" table in database
func CreateRestaurantTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"restaurants" +
		" (" +
		"ID MEDIUMINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, " +
		"Name VARCHAR(255), " +
		"Description VARCHAR(255), " +
		"Halal BOOL, " +
		"Vegan BOOL, " +
		"Address VARCHAR(255), " +
		"PostalCode MEDIUMINT UNSIGNED, " +
		"Lat FLOAT, " +
		"Lng FLOAT" +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: restaurants")
	}
}

// Creates "foods" table in database
func CreateFoodTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"foods" +
		" (" +
		"ID MEDIUMINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, " +
		"RestaurantID MEDIUMINT UNSIGNED, " +
		"Name VARCHAR(255), " +
		"Price FLOAT UNSIGNED, " +
		"Calories FLOAT UNSIGNED, " +
		"Halal BOOL, " +
		"Vegan BOOL " +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: foods")
	}
}

// Creates "orders" table in database
func CreateOrderTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"orders" +
		" (" +
		"ID MEDIUMINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, " +
		"Username VARCHAR(255), " +
		"Status VARCHAR(255), " + // cart, past
		"Date DATE, " +
		"Address VARCHAR(255), " +
		"PostalCode MEDIUMINT UNSIGNED, " +
		"Lat FLOAT, " +
		"Lng FLOAT" +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: orders")
	}
}

// Creates "order_items" table in database
func CreateOrderItemTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS " +
		"order_items" +
		" (" +
		"ID MEDIUMINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, " +
		"OrderID MEDIUMINT UNSIGNED, " +
		"FoodID MEDIUMINT UNSIGNED, " +
		"Quantity SMALLINT UNSIGNED, " +
		"Subtotal FLOAT UNSIGNED" +
		")")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: order_items")
	}
}