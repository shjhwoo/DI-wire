package notifications

import (
	"fmt"
	"os/user"
)

func InformOrderShipped(receiver user.User, orderID string, sendSMS func(user.User, string) error) bool {
	message := fmt.Sprintf("Your order #%s is shipped!", orderID)
	//function closure argument
	err := sendSMS(receiver, message)

	if err != nil {
		return false
	}

	return true
}
