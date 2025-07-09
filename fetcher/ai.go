package fetcher

import (
	"context"
	"fmt"
	"strings"
	"trainer-helper/config"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AI struct {
	AppConfig *config.Config
}

const PROMPT = `
You are a developer you have coppied a text from google word doc and have a json schema to fill using the data, transform the raw string into a json that will fit the json schema.
If RPE appears only in on row use it everywhere.
If multiple sets have the same data, group them and increase the set_count, add Kg to intensity where applicible.
Group the exercises by the exercise name.
Don't just create one big list of work sets. Give no additional comments
If there is "alebo", split the work sets into 2 exercises.
Give me the JSON output only no extra spaces no identation so that I can machine parse it.
If there are mulitple exercise like C1 C2 and C3 Split them so that each exercise is its own exercise
If there is Počet sérií a nubmer but only one itensity and repetitions, copy the work set to the set count
Always provide RPE, if its missing use null
If you split exercises of similiar cateogry like C1 and C2, also split the work sets and intensity and number of sets RPE correcly between the two exercises
Use only these exercise names translated from the slovak language. However if the exercise name is not in the list just use the original name.
%s


Here is the schema:
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Workout List Schema",
  "description": "Schema for a list of workout entries.",
  "type": "array",
  "items": {
    "type": "object",
    "properties": {
      "exercise_name": {
        "type": "string",
        "description": "The name of the exercise."
      },
      "note": {
        "type": "string",
        "description": "A note or description for the workout."
      },
      "work_sets": {
        "type": "array",
        "description": "An array of individual work sets.",
        "items": {
          "type": "object",
          "properties": {
            "reps": {
              "type": "integer",
              "description": "The number of repetitions for the set.",
              "minimum": 1
            },
            "intensity": {
              "type": "string",
              "description": "The intensity of the set, represented as a string (e.g., weight, percentage)."
            },
            "rpe": {
              "type": "integer",
              "description": "The Rate of Perceived Exertion (RPE) for the set.",
              "minimum": 1,
              "maximum": 10
            }
          },
          "required": [
            "reps",
            "intensity"
          ]
        }
      }
    },
    "required": [
      "exercise_name",
      "note",
      "work_sets"
    ]
  }
}



Here is the raw text:
%s
`

func (ai AI) RawStringToJson(exerciseList []string, rawString string) (result string, err error) {
	client := anthropic.NewClient(
		// TODO: Config
		option.WithAPIKey(ai.AppConfig.ClaudeToken),
	)
	completePrompt := fmt.Sprintf(PROMPT, strings.Join(exerciseList, ", "), rawString)

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(completePrompt)),
		},
		Model: anthropic.ModelClaudeSonnet4_20250514,
	})
	if err != nil {
		return "", err
	}

	return message.Content[0].Text, nil
}
