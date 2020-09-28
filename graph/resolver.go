/*
 * @Author: your name
 * @Date: 2020-09-27 18:28:28
 * @LastEditTime: 2020-09-27 19:05:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gqlgentoto\graph\resolver.go
 */
package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/testgql/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
