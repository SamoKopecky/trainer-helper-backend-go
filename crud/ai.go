package crud

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func main() {
	client := anthropic.NewClient(
		option.WithAPIKey(""),
	)
	text := `
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
Use only these exercise names translated from the slovak language.
+------------------------+
| Squat                  |
| Deadlift               |
| Bench Press            |
| RDL                    |
| Cable Horizontal Row   |
| Hack Squat             |
| Leg Press              |
| Calf Raise             |
| Ring Muscle Up         |
| Pull Up                |
| Machine Hip Abduction  |
| Jefferson Curl         |
| Kettlebell Side Bend   |
| Machine Chest Press    |
| Multipress             |
| Dips                   |
| Machine Shoulder Press |
| Triceps Pushdown       |
| Bent Arm Lateral Raise |
| Bench Crunch           |
| Squat                  |
| Deadlift               |
| Bench Press            |
| RDL                    |
| Cable Horizontal Row   |
| Hack Squat             |
| Leg Press              |
| Calf Raise             |
| Ring Muscle Up         |
| Pull Up                |
| Machine Hip Abduction  |
| Jefferson Curl         |
| Kettlebell Side Bend   |
| Machine Chest Press    |
| Multipress             |
| Dips                   |
| Machine Shoulder Press |
| Triceps Pushdown       |
| Bent Arm Lateral Raise |
| Bench Crunch           |
+------------------------+
However if the exercise name is not in the list just use the original name.


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
Názov cviku
Počet sérií
Počet opakovaní
Intenzita
Poznámka
A0: RMU - row

A1: RMU - full - band assisted - ZAKLON,

A2: RMU negativa
2
4
2
6
2,2,2,1  Zaklon
5, 5
leg assisted
černý!






B1: bar - pullups
1
1
1
1
3
2
5
4
16,25kg
16,25
6,25kg
6,25kg
RPE 8
RPE 8
6,5
6,5


C1: cable - halfkneeling lat pulldown
2





7,6




83 nová






D1: kelso shrug
3
13
32kg
`

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(text)),
		},
		Model: anthropic.ModelClaudeSonnet4_20250514,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", message.Content[0].Text)
}
