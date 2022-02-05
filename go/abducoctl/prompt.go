package abducoctl

import (
	"github.com/AlecAivazis/survey/v2"
)

func getQuestions() []*survey.Question {
	return []*survey.Question{
		{
			Name: "session",
			Prompt: &survey.Select{
				Message: "Select Session:",
				Options: Names(),
			},
			Validate: survey.Required,
		},
	}
}

func Prompt() {
	answers := struct {
		Session string
	}{}
	survey.Ask(getQuestions(), &answers)
	if Exists(answers.Session) {
		Connect(ctx, answers.Session)
	}
}
