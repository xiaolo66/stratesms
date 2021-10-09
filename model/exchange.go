package model

import "github.com/jinzhu/gorm"

type Exchange struct {
	gorm.Model
	ExchangeName string
	Website      string
	Describe     string
}

func (e *Exchange) Add() error {
	err := DB.Debug().Create(e).Error
	if err != nil {
		return err
	}
	return nil
}

