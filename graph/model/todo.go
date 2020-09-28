/*
 * @Author: your name
 * @Date: 2020-09-27 18:41:27
 * @LastEditTime: 2020-09-27 18:41:36
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \gqlgentoto\graph\model\todo.go
 */
package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
