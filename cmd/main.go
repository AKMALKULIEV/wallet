package main

import (


	"github.com/AKMALKULIEV/wallet/pkg/wallet")



   func main() {
      svc := &wallet.Service{}
   
      svc.RegisterAccount("1234")
      svc.Deposit(5, 100)
   
      svc.RegisterAccount("5678")
      svc.Deposit(425, 100)
   
      svc.RegisterAccount("91011")
      svc.Deposit(265, 4242464200)
   
      svc.ExportToFile("../data/export.txt")
      svc.ImportFromFile("../data/export.txt")
   }
