package servers

import (
	"stratesms/dto"
	"stratesms/model"
)

type Exchange struct {
}

func NewExchange() *Exchange {
	return &Exchange{}
}

func (e *Exchange) Add(info *dto.Exchange) error {
	exchange := model.Exchange{
		ExchangeName: info.ExchangeName,
		Website:      info.Website,
		Describe:     info.Describe,
	}
	return exchange.Add()
}

func (e *Exchange) GetAll() (exchanges []model.Exchange, err error) {
	err = model.DB.Find(&exchanges).Error
	return exchanges, err
}

func (e *Exchange) Get(info *dto.Exchange) (ex []model.Exchange, err error) {
	err = model.DB.Debug().Where("exchange_name In(?)", &info.ExchangeName).Find(&ex).Error
	return ex, err
}

func (e *Exchange) Put(info *dto.Exchange) error {
	var exchange model.Exchange
	err := model.DB.Where("exchange_name=?", info.ExchangeName).First(&exchange).Error
	if err != nil {
		return err
	}
	model.DB.Debug().Model(&exchange).Update(map[string]interface{}{
		"exchange_name": info.ExchangeName,
		"website":       info.Website,
		"describe":      info.Describe})
	return nil
}
func (e *Exchange) Delete(info string) error {
	var exchange model.Exchange
	err := model.DB.Where("exchange_name=?", info).First(&exchange).Error
	if err != nil {
		return err
	}
	err = model.DB.Debug().Unscoped().Delete(&exchange).Error
	if err != nil {
		return err
	}
	return nil

}
