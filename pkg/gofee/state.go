package gofee

import "fmt"

type State interface {
	// This needs to chance, BigInt?
	// The representation is in cents. "12.23$" -> 1223
	// The default account is stored under "default"
	Accounts(string) Account

	// DeepCopy
	DeepCopy() State
}

type Account interface {
	Value() int
	SetValue(int)
	DeepCopy() Account
}

type defaultAccount struct {
	value int
}

func NewDefaultAccount() *defaultAccount {
	return &defaultAccount{
		// Account starts with zero
		value: 0,
	}
}

func (a *defaultAccount) Value() int {
	return a.value
}

func (a *defaultAccount) SetValue(value int) {
	a.value = value
}

func (a *defaultAccount) DeepCopy() Account {
	newA := NewDefaultAccount()
	newA.value = a.value
	return newA
}

type concreeteState struct {
	accounts map[string]Account
}

func NewConcreeteState() *concreeteState {
	return &concreeteState{
		accounts: map[string]Account{
			"default": NewDefaultAccount(),
		},
	}
}

func (cs *concreeteState) Accounts(account string) Account {
	return cs.accounts[account]
}

func (cs *concreeteState) String() string {
	return fmt.Sprintf("{default: %f$}", float64(cs.Accounts("default").Value()/100))
}

func (cs *concreeteState) DeepCopy() State {
	newCs := NewConcreeteState()

	for key, value := range cs.accounts {
		newCs.accounts[key] = value.DeepCopy()
	}

	return newCs
}
