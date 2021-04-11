package wallet

import (
	// "reflect"
	"testing"

	// "github.com/AKMALKULIEV/wallet/pkg/types"
)

func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+9920000001")

	account, err := svc.FindAccountByID(4987)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}