package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Ev     string `json:"ev,omitempty"`
	Et     string `json:"et,omitempty"`
	Id     string `json:"id,omitempty"`
	Uid    string `json:"uid,omitempty"`
	Mid    string `json:"mid,omitempty"`
	T      string `json:"t,omitempty"`
	P      string `json:"p,omitempty"`
	L      string `json:"l,omitempty"`
	Sc     string `json:"sc,omitempty"`
	Atrk1  string `json:"atrk1,omitempty"`
	Atrv1  string `json:"atrv1,omitempty"`
	Atrt1  string `json:"atrt1,omitempty"`
	Atrk2  string `json:"atrk2,omitempty"`
	Atrv2  string `json:"atrv2,omitempty"`
	Atrt2  string `json:"atrt2,omitempty"`
	Atrk3  string `json:"atrk3,omitempty"`
	Atrv3  string `json:"atrv3,omitempty"`
	Atrt3  string `json:"atrt3,omitempty"`
	Atrk4  string `json:"atrk4,omitempty"`
	Atrv4  string `json:"atrv4,omitempty"`
	Atrt4  string `json:"atrt4,omitempty"`
	Atrk5  string `json:"atrk5,omitempty"`
	Atrv5  string `json:"atrv5,omitempty"`
	Atrt5  string `json:"atrt5,omitempty"`
	Uatrk1 string `json:"uatrk1,omitempty"`
	Uatrv1 string `json:"uatrv1,omitempty"`
	Uatrt1 string `json:"uatrt1,omitempty"`
	Uatrk2 string `json:"uatrk2,omitempty"`
	Uatrv2 string `json:"uatrv2,omitempty"`
	Uatrt2 string `json:"uatrt2,omitempty"`
	Uatrk3 string `json:"uatrk3,omitempty"`
	Uatrv3 string `json:"uatrv3,omitempty"`
	Uatrt3 string `json:"uatrt3,omitempty"`
	Uatrk4 string `json:"uatrk4,omitempty"`
	Uatrv4 string `json:"uatrv4,omitempty"`
	Uatrt4 string `json:"uatrt4,omitempty"`
	Uatrk5 string `json:"uatrk5,omitempty"`
	Uatrv5 string `json:"uatrv5,omitempty"`
	Uatrt5 string `json:"uatrt5,omitempty"`
	Uatrk6 string `json:"uatrk6,omitempty"`
	Uatrv6 string `json:"uatrv6,omitempty"`
	Uatrt6 string `json:"uatrt6,omitempty"`
}

type Output struct {
	Event           string                 `json:"event"`
	EventType       string                 `json:"event_type"`
	AppID           string                 `json:"app_id"`
	UserID          string                 `json:"user_id"`
	MessageID       string                 `json:"message_id"`
	PageTitle       string                 `json:"page_title"`
	PageURL         string                 `json:"page_url"`
	BrowserLanguage string                 `json:"browser_language"`
	ScreenSize      string                 `json:"screen_size"`
	Attributes      map[string]interface{} `json:"attributes"`
	Traits          map[string]interface{} `json:"traits"`
}

func main() {
	router := gin.Default()

	inputChannel := make(chan Input)

	router.POST("/submit", func(c *gin.Context) {
		var input Input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		inputChannel <- input

		c.JSON(http.StatusOK, gin.H{"message": "Request has been submitted and is being processed"})
	})

	go worker(inputChannel)

	// Run the server on port 8080
	router.Run(":8080")
}

func worker(inputChannel <-chan Input) {
	for {
		select {
		case input := <-inputChannel:
			// Process the input data
			output := Output{
				Event:           input.Ev,
				EventType:       input.Et,
				AppID:           input.Id,
				UserID:          input.Uid,
				MessageID:       input.Mid,
				PageTitle:       input.T,
				PageURL:         input.P,
				BrowserLanguage: input.L,
				ScreenSize:      input.Sc,
				Attributes: map[string]interface{}{
					input.Atrk1: map[string]interface{}{
						"value": input.Atrv1,
						"type":  input.Atrt1,
					},
					input.Atrk2: map[string]interface{}{
						"value": input.Atrv2,
						"type":  input.Atrt2,
					},
					input.Atrk3: map[string]interface{}{
						"value": input.Atrv3,
						"type":  input.Atrt3,
					},
					input.Atrk4: map[string]interface{}{
						"value": input.Atrv4,
						"type":  input.Atrt4,
					},
				},
				Traits: map[string]interface{}{
					input.Uatrk1: map[string]interface{}{
						"value": input.Uatrv1,
						"type":  input.Uatrt1,
					},
					input.Uatrk2: map[string]interface{}{
						"value": input.Uatrv2,
						"type":  input.Uatrt2,
					},
					input.Uatrk3: map[string]interface{}{
						"value": input.Uatrv3,
						"type":  input.Uatrt3,
					},
					input.Uatrk4: map[string]interface{}{
						"value": input.Uatrv4,
						"type":  input.Uatrt4,
					},
					input.Uatrk5: map[string]interface{}{
						"value": input.Uatrv5,
						"type":  input.Uatrt5,
					},
					input.Uatrk6: map[string]interface{}{
						"value": input.Uatrv6,
						"type":  input.Uatrt6,
					},
				},
			}

			// Convert the output to JSON
			outputJSON, err := json.Marshal(output)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println(string(outputJSON))
		}
	}
}
