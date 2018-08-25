package waldorf

import (
	"fmt"
)

type observation struct {
	complaints []string
}

func Observe() observation {
	return observation{
		complaints: []string{},
	}
}

func (r *observation) IsTrue(statement bool, msg string, formatting ...interface{}) (validated bool) {
	if !statement {
		r.complaints = append(r.complaints, generateMSG(msg, formatting...))
	}
	return statement
}

func (r *observation) IsFalse(statement bool, msg string, formatting ...interface{}) (validated bool) {
	if statement {
		r.complaints = append(r.complaints, generateMSG(msg, formatting...))
	}
	return !statement
}

func (r *observation) Ridicule() []string {
	if len(r.complaints) == 0 {
		return nil
	}
	return r.complaints
}

func generateMSG(msg string, formatting ...interface{}) string {
	if len(formatting) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, formatting)
}
