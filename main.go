package main

import (
	"github.com/aloneen/assign2-ex1-adv/controllers"
	"github.com/aloneen/assign2-ex1-adv/initializers"
	"github.com/aloneen/assign2-ex1-adv/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	// Create table
	controllers.CreateUsersTable()

	// Insert users with transaction
	controllers.InsertUsers([]models.User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	})

	// Query with filter + pagination
	controllers.QueryUsers(25, 1, 2)
	controllers.QueryUsers(25, 2, 2)

	// Update
	controllers.UpdateUser(1, "AliceUpdated", 26)

	// Delete
	controllers.DeleteUser(2)
}
