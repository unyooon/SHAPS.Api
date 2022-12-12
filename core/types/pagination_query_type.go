package types

type Query struct {
	Page int `form:"page"`
	Size int `form:"size"`
}