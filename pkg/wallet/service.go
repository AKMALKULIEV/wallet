package wallet

import (
	"errors"

	"github.com/AKMALKULIEV/wallet/pkg/types"
	// "github.com/google/uuid"
)

var (
	ErrPhoneRegistered      = errors.New("Уже зарегистрирован")
	ErrAccountNotFound      = errors.New("Нет такого аккаунта")
	ErrAmountMustBePositive = errors.New("Должно быть больше 0")
)

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)

	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}


func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {

	for _, account := range s.accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}
