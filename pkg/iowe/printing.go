package iowe

import (
	"github.com/timmattison/go-away/pkg/save-the-planet"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetPrinter() (*message.Printer, language.Tag) {
	locale := save_the_planet.GetLocale()

	tag := language.Make(locale)

	return message.NewPrinter(tag), tag
}
