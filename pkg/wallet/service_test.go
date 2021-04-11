package wallet

import (
	// "reflect"
	"testing"
	// "github.com/AKMALKULIEV/wallet/pkg/types"
)

func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("654679654646")

	account, err := svc.FindAccountByID(46)
	if err == nil {
		t.Errorf("got > %v want > nil", account)
	}
}
