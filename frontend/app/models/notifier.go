package models

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
)

type Notification string

func NewNotification(message string) (Notification, error) {
	return Notification(message), nil
}

func (n *Notifier) Notify(not Notification) error {
	fmt.Println(string(not))
	return nil
}

type Notifier struct {
}

func (n *Notifier) Validate(v *revel.Validation) {
}

func newNotifier() *Notifier {
	return new(Notifier)
}

func (n *Notifier) Marshal() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Notifier) UnMarshal(m []byte) error {
	return json.Unmarshal(m, n)
}
