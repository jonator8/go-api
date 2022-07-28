package internal

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger     *logrus.Logger
	Commands   *CommandsBootstrap
	Validator  *validator.Validate
	Translator *ut.Translator
}

func NewApp(
	logger *logrus.Logger,
	cmd *CommandsBootstrap,
) *App {
	english := en.New()
	universal := ut.New(english, english)
	trans, _ := universal.GetTranslator("en")
	val := validator.New()

	customTranslations(val, trans)

	return &App{
		Logger:     logger,
		Commands:   cmd,
		Validator:  val,
		Translator: &trans,
	}
}
