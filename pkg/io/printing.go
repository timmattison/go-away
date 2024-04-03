package io

import (
	"github.com/timmattison/go-away/pkg/environment"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetPrinter() (*message.Printer, language.Tag) {
	locale := environment.GetLocale()

	tag := language.Make(locale)

	return message.NewPrinter(tag), tag
}
