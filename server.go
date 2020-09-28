/*
 * @Author: your name
 * @Date: 2020-09-27 18:28:28
 * @LastEditTime: 2020-09-28 17:55:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gqlgentoto\server.go
 */
package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/testgql/gqlgen-todos/graph"
	"github.com/testgql/gqlgen-todos/graph/generated"
)

const defaultPort = "8686"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()
	//e.SetBinder(&binding.EchoBinder{})

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// }))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	e.Any("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query", echo.WrapHandler(srv))

	e.Start(":" + defaultPort)
}
