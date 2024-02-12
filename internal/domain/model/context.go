package model

import "context"

type Context interface {
	context.Context
	GetUser() User
	GetToken() string
}

type AppContext struct {
	context.Context
}

func (ctx AppContext) GetUser() User {
	user := ctx.Value(ContextTagUser).(User)
	return user
}

func (ctx AppContext) GetToken() string {
	token := ctx.Value(ContextTagToken).(string)
	return token
}
