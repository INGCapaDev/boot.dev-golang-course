package main

func getExpenseReport(e expense) (string, float64) {
	// email, isEmail := e.(email)
	// sms, isSMS := e.(sms)

	// if isEmail {
	// 	return email.toAddress, email.cost()
	// }
	// if isSMS {
	// 	return sms.toPhoneNumber, sms.cost()
	// }

	// return "", e.cost()

	switch expense := e.(type) {
	case email:
		return expense.toAddress, expense.cost()
	case sms:
		return expense.toPhoneNumber, expense.cost()
	default:
		return "", expense.cost()
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
