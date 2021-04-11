package main

import "github.com/AKMALKULIEV/wallet/pkg/wallet"


func main() {
   svc := &wallet.Service{}
   svc.RegisterAccount("65464615")
   svc.Deposit(1,10 )
   svc.RegisterAccount("4667765476")

}
