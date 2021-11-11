package entity

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("value not found")
	ErrExpired = errors.New("exp date have been passed")
)
type Key string

type Value struct {
	Val interface{}
	Exp time.Time
}

func (v Value) IsValid()bool{
	return v.Exp.Before(time.Now())
}

// GetValue returns the value or nil and the error indicates whether expiry date is passed.
func (v Value)GetValue() (interface{}, error){
	if v.IsValid(){
		return v.Val, nil
	}
	return nil, ErrExpired
}

func (v *Value) SetValue(val interface{}){
	v.Val = val
}

func (v *Value) SetExp (t time.Time){
	v.Exp = t
}