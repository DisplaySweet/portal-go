package test

import portal "github.com/DisplaySweet/portal-go/src"

//A space for testing the logic and the functionality of the portal pkg

func checkAccounts() {
	account := portal.Account{
		AccountName: "Mr. Account",
		Email:       "mraccount@email.domain",
	}

	accContacts, err := account.GetOwnedContacts()
	accDeposits, err := account.GetOwnedDeposits()
	statusCode, err := account.Update()
	statuscode, err := account.Delete()
}

func checkCompanies() {
	company := portal.Company{
		Name:   "A Real Company",
		Active: true,
	}

	accounts, contacts, err := company.GetAccountsAndContacts
}

func checkContacts() {
	contact := portal.Contact{
		Firstname: "steve",
		Lastname:  "jobs",
	}

	println(contact.Firstname)

	err := contact.SendUpdate()
}
