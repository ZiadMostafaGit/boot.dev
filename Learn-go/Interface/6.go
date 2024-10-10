package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}
type Code struct {
	message string
}

func (pl PlainText) Format() string {

	return pl.message

}

func (bo Bold) Format() string {
	return "**" + bo.message + "**"
}

func (co Code) Format() string {
	return "`" + co.message + "`"
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
