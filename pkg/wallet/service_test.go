package wallet

import (
	// "reflect"
	"fmt"
	"os"
	"testing"

	"github.com/AKMALKULIEV/wallet/pkg/types"
)

type testService struct {
	*Service
}

type testAccount struct {
	phone    types.Phone
	balance  types.Money
	payments []struct {
		amount   types.Money
		category types.PaymentCategory
	}
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}
func (s *Service) addAccount(data testAccount) (*types.Account, []*types.Payment, []*types.Favorite, error) {

	// register user
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	//  account top up
	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("can't deposity account, error = %v", err)
	}

	// make a payment to account
	payments := make([]*types.Payment, len(data.payments))
	favorites := make([]*types.Favorite, len(data.payments))

	for i, payment := range data.payments {

		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("can't make paymnet, error = %v", err)
		}

		favorites[i], err = s.FavoritePayment(payments[i].ID, "Favorite payment #i")
		if err != nil {
			return nil, nil, nil, fmt.Errorf("can't make favorite paymnet, error = %v", err)
		}
	}
	return account, payments, favorites, nil
}

var defaultTestAccount = testAccount{
	phone:   "+998204567898",
	balance: 500,
	payments: []struct {
		amount   types.Money
		category types.PaymentCategory
	}{
		{amount: 100, category: "auto"},
	},
}

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
func Transactions(s *testService) {
	s.RegisterAccount("1111")
	s.Deposit(1, 500)
	s.Pay(1, 10, "food")
	s.Pay(1, 10, "phone")
	s.Pay(1, 15, "bank")
	s.Pay(1, 25, "auto")
	s.Pay(1, 30, "restaurant")
	s.Pay(1, 50, "auto")
	s.Pay(1, 60, "bank")
	s.Pay(1, 50, "bank")

	s.RegisterAccount("2222")
	s.Deposit(2, 200)
	s.Pay(2, 40, "phone")

	s.RegisterAccount("3333")
	s.Deposit(3, 300)
	s.Pay(3, 36, "auto")
	s.Pay(3, 12, "food")
	s.Pay(3, 25, "phone")
}

func TestService_SumPayments(t *testing.T) {
	s := newTestService()
	Transactions(s)
	sum := s.SumPayments(0)
	if sum != 363 {
		t.Errorf("TestService_SumPayments(): sum=%v", sum)
	}

}

func BenchmarkSumPayments(b *testing.B) {
	s := newTestService()
	Transactions(s)
	want := types.Money(363)
	for i := 0; i < b.N; i++ {
		result := s.SumPayments(3)
		if result != want {
			b.Fatalf("INVALID: result_we_got %v, result_we_want %v", result, want)
		}
	}
}
func TestService_ExportToFile_success(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.ExportToFile("newdata/hello.txt")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_ExportToFile_notFound(t *testing.T) {
	s := newTestService()
	_, _, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	err = s.ExportToFile("")
	if err == nil {
		t.Error(err)
		return
	}
}

func TestService_ImportFromFile_success(t *testing.T) {
	s := newTestService()
	s.RegisterAccount("1111")
	s.Deposit(1, 500)
	pay, _ := s.Pay(1, 100, "phone")
	s.FavoritePayment(pay.ID, "my_phone")

	err := s.ImportFromFile("newdata/hello.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
func TestService_ImportFromFile_notFound(t *testing.T) {
	s := newTestService()

	err := s.ImportFromFile("")
	if err == nil {
		t.Error(err)
		return
	}
}

func TestService_Export(t *testing.T) {
	s := newTestService()
	_, payments, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	payment := payments[0]
	_, err = s.FavoritePayment(payment.ID, "my favorite payment")
	if err != nil {
		t.Errorf("FavoritePayment(): error: %v", err)
		return
	}

	err = s.Export("data")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_Import_success(t *testing.T) {
	s := newTestService()

	err := s.Import("data")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_Import_notFound1(t *testing.T) {

	s := newTestService()

	err := s.Import("")
	if err != nil {
		t.Error(err)
		return
	}
}
func TestService_Import_notFound2(t *testing.T) {

	s := newTestService()

	err := s.Import("")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_Import_Error(t *testing.T) {

	s := newTestService()
	_, payments, _, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	payment := payments[0]
	_, err = s.FavoritePayment(payment.ID, "my favorite payment")
	if err != nil {
		t.Errorf("FavoritePayment(): error: %v", err)
		return
	}

	err = s.Import("data")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestService_Import_emptyFiles(t *testing.T) {
	s := newTestService()

	file1, _ := os.Create("newdata/accounts.dump")
	defer file1.Close()

	file2, _ := os.Create("newdata/payments.dump")
	defer file2.Close()

	file3, _ := os.Create("newdata/favorites.dump")
	defer file3.Close()

	err := s.Import("newdata")
	if err != nil {
		t.Error(err)
		return
	}
}
