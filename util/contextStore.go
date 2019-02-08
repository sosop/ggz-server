package util

import "context"

type ContextKey string

func Get(context context.Context, key ContextKey) interface{} {
	if v := context.Value(key); v != nil {
		return v
	}
	return nil
}
