# Go Fiber MongoDB App

## Overview
A simple REST API built with Go, Fiber, and MongoDB to manage users and phone numbers.

## Prerequisites
- Go 1.24.4+
- MongoDB
- Fiber framework
- MongoDB Go driver

## Setup
1. Install dependencies:
   ```bash
   go mod tidy
   
### Notes:
- The module name is updated to `go-fiber-app` to match your `go.mod`.
- Fiber (`github.com/gofiber/fiber/v2`) replaces Gin, and the handler and routes are adjusted accordingly.
- The `go.mod` includes all specified dependencies, and `go mod tidy` will resolve them.
- The project uses the `fiber_db` database as specified in the `.env` file.
- Ensure MongoDB is running locally at `mongodb://localhost:27017`.
- Let me know if you need additional endpoints or features!