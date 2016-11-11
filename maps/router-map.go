package main

import "fmt"

type Route struct {
	Name    string
	Method  string
	Pattern string
}

var routes []Route

func init() {
	// Add user routes.
	userRoutes := []Route{
		Route{"User", "GET", "/user/{id}"},
		Route{"Users", "GET", "/users"},
		Route{"User Form", "GET", "/user/add"},
		Route{"Add User", "POST", "/user/add"},
	}
	routes = append(routes, userRoutes...)
}

func init() {
	// Add page routes
	pageRoutes := []Route{
		Route{"Home", "GET", "/"},
		Route{"About", "GET", "/about"},
	}
	routes = append(routes, pageRoutes...)
}

func main() {
	fmt.Printf("%#v\n", routes)
}
