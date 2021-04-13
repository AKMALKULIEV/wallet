package wallet

import (
	// "reflect"
	"testing"

	"github.com/AKMALKULIEV/wallet/pkg/types"
)

func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("654679654646")

	account, err := svc.FindAccountByID(4987)
	if err == nil {
		t.Errorf("got > %v want > nil", account)
	}
}
func TestService_Reject_Success(t *testing.T) {
	s := &Service{}
	phone := types.Phone("555555555")
	account, err := s.RegisterAccount(phone)
	if err != nil {
		t.Errorf("Reject():can not register account, error = %v", err)
		return
	}
	err = s.Deposit(account.ID, 10_000_00)
	if err != nil {
		t.Errorf("Reject():can not deposit account, error = %v", err)
		return
	}
	payment, err := s.Pay(account.ID, 10_000_00, "auto")
	if err != nil {
		t.Errorf("Reject():can not create payment, error = %v", err)
		return
	}
	err = s.Reject(payment.ID)
	if err != nil {
		t.Errorf("Reject():can not reject payment, error = %v", err)
		return
	}
}

func TestService_Repeat_success_user(t *testing.T) {

	s := newTestServiceUser()
	s.RegisterAccount("+9922000000")
	account, err := s.FindAccountByID(1)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("получили > %v ожидали > nil", err)
	}

	payment, err := s.Pay(account.ID, 100_00, "auto")
	if err != nil {
		t.Errorf("получили > %v ожидали > nil", err)
	}

	pay, err := s.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("получили > %v ожидали > nil", err)
	}

	pay, err = s.Repeat(pay.ID)
	if err != nil {
		t.Errorf("Repeat(): can not payment for an account(%v), error(%v)", pay.ID, err)
	}
}
