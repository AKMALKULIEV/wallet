package wallet

import (
	"errors"
	"github.com/AKMALKULIEV/wallet/pkg/types"
)


type Service struct{
	nextAccountID int64
	accounts []*types.Account
	payments []*types.Payment 
}

var ErrAccountNotFound = errors.New("Нет такого аакаунта")
func (s *Service)RegisterAccount(phone types.Phone)  {
	for _, account := range s.accounts {
		if account.Phone == phone{
			return
		}
	}
	s.nextAccountID++
	s.accounts =append(s.accounts, &types.Account{
		ID: s.nextAccountID,
		Phone : phone,
		Balance: 0,
	})
}

func (s *Service) Deposit(accountID int64, amount types.Money)  {
	if amount <= 0{
		return
	}
	var account *types.Account

	for _, acc := range s.accounts {
		if acc.ID== accountID{
			account = acc
			break 
            	
		}
	}
	if account == nil{
		return
	}	

	account.Balance += amount 
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, accounts := range s.accounts {
		if accounts.ID == accountID {
			account = accounts
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}