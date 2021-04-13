package main

import (
	"fmt"

	"github.com/AKMALKULIEV/wallet/pkg/wallet")



func main() {
   svc := &wallet.Service{}
   account, err := svc.RegisterAccount("65464615")
   if err != nil{
      fmt.Println(err) 
      return
   }
   fmt.Println(account)
   
   // account, err = svc.RegisterAccount("65464615")
   // if err != nil{
   //    fmt.Println(err) 
   //    return
   // }
   // fmt.Println(account)
   
   err = svc.Deposit(account.ID, 10604)
   if err != nil{
      fmt.Println(err) 
      return
   }
   fmt.Println(account.Balance)
}
