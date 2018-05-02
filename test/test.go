package test

import (
	"github.com/DisplaySweet/portal-go/src"
)

//A space for testing the logic and the functionality of the portal pkg

func checkAccounts() {
	//Initialise an account with any data
	account := portal.Account{
		AccountName: "Mr. Account",
		Email:       "mraccount@email.domain",
	}

	//Check and confirm the functionality available to an account
	accContacts, err := account.GetOwnedContacts()
	if err != nil {
		println(err)
	}
	println(accContacts)

	accDeposits, err := account.GetOwnedDeposits()
	if err != nil {
		println(err)
	}
	println(accDeposits)

	err = account.Update()
	if err != nil {
		println(err)
	}

	err = account.Delete()
	if err != nil {
		println(err)
	}

}

func checkCompanies() {
	//Initialise a company object with any data
	company := portal.Company{
		Name:   "A Real Company",
		Active: true,
	}

	//Init an agent for use in a method below
	agent := &portal.Agent{}
	agents := []*portal.Agent{
		agent,
	}

	//Check and confirm the functionality available to the company object
	accounts, contacts, err := company.GetAccountsContacts()
	if err != nil {
		println(err)
	}
	println(accounts)
	println(contacts)

	err = company.AddUsers(agents)
	if err != nil {
		println(err)
	}

	err = company.Delete()
	if err != nil {
		println(err)
	}

	err = company.Update()
	if err != nil {
		println(err)
	}
}

func checkContacts() {
	contact := portal.Contact{
		Firstname: "steve",
		Lastname:  "jobs",
	}

	err := contact.SendUpdate()
	if err != nil {
		println(err)
	}
}

func checkEvents() {
	event := portal.Event{
		Name:      "Smellrose Event A",
		Terminals: 8,
	}

	err := event.Update()
	if err != nil {
		println(err)
	}

	err = event.Delete()
	if err != nil {
		println(err)
	}
}

//Listing is missing a lot of functionality
func checkListings() {
	listing := portal.Listing{
		Name:         "A304",
		Availability: "Under Offer",
		Bedrooms:     "1",
		Bathrooms:    "1",
	}

	err := listing.Update()
	if err != nil {
		println(err)
	}

	err = listing.Delete()
	if err != nil {
		println(err)
	}

}

func checkOffer() {
	offer := portal.Offer{
		Status:        "Active",
		Price:         5087.87,
		FundsReceived: 5087.87,
	}

	completedOffer, err := offer.Complete()
	if err != nil {
		println(err)
	}
	println(completedOffer)

	cancelledOffer, err := offer.Cancel()
	if err != nil {
		println(err)
	}
	println(cancelledOffer)

	updatedOffer, err := offer.Update()
	if err != nil {
		println(err)
	}
	println(updatedOffer)
}

func checkProspects() {
	prospect := portal.Prospect{}

	agent := &portal.Agent{}
	schedule := &portal.Schedule{}

	err := prospect.UpdateAgent(agent)
	if err != nil {
		println(err)
	}

	err = prospect.Update()
	if err != nil {
		println(err)
	}

	err = prospect.UpdateSchedule(schedule)
	if err != nil {
		println(err)
	}
}
