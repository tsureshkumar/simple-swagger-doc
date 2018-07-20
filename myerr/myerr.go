package myerr

import "fmt"

type MyErr struct {
	BaseErr error
	Msg     string
}

func (err MyErr) Error() string {
	if err.BaseErr != nil {
		return fmt.Sprintf("My Error: %s. %s.", err.BaseErr.Error(), err.Msg)
	}
	return "error: unknown"
}

func Wrap(berr error, msg string) MyErr {
	return MyErr{
		BaseErr: berr,
		Msg:     msg,
	}
}

func New(msg string) MyErr {
	return MyErr{
		Msg: msg,
	}
}
