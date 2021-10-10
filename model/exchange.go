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

func(e *Exchange)GetAll()(list []Exchange,err error){
	err = DB.Find(&list).Error
	return list, err
}

func(e *Exchange)Get()(list []Exchange, err error){
	err=DB.Debug().Where("exchange_name In(?)", e.ExchangeName).Find(&list).Error
	return list, err
}

