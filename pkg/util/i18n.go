package util

import (
	"fmt"
	"gateway/locales"

	"github.com/gobuffalo/buffalo"
	i18n "github.com/gobuffalo/mw-i18n/v2"
)

var (
	Translator *i18n.Translator
)

func init() {
	var err error
	Translator, err = i18n.New(locales.FS(), "en-US")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func T(c buffalo.Context, id string, args ...any) string {
	return Translator.Translate(c, id, args...)
}
