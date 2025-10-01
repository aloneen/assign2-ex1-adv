# Exercise 1: Advanced PostgreSQL Operations with database/sql
### Objective: Connect to PostgreSQL, perform advanced operations, and handle transactions and error management.
##### 1.	Setup PostgreSQL Connection:
      ○	Create a Go program that connects to your PostgreSQL database using the pq driver.
      ○	Implement connection pooling with sql.DB.
##### 2.	Create a Table with Constraints:
      ○	Write a function to create a users table with the following constraints:
      ■	id as a primary key, auto-incremented.
      ■	name as a unique, non-null field.
      ■	age as a non-null integer field.
##### 3.	Insert Data with Transactions:
      ○	Write a function to insert multiple users into the users table within a transaction.
      ○	Implement error handling to roll back the transaction if any error occurs during insertion.
##### 4.	Query Data with Filtering and Pagination:
      ○	Write a function to query and print users with optional filters for age and pagination support.
      ○	Implement pagination to return a specific number of results per page.
##### 5.	Update and Delete Data:
      ○	Write functions to update a user’s details and delete a user by their ID, including error handling.
