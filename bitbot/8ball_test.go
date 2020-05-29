package bitbot

import (
	"math/rand"
	"testing"
)

func TestMake8BallAnswer(t *testing.T) {
	rand.Seed(1)

	var testAnswers string

	answersMap := make(map[string]int)
	for _, ans := range magic8responses {
		answersMap[ans] = 0
	}

	// We get 200 random answers, and test that every one of them is awaited
	for i := 1; i <= 200; i++ {
		testAnswers = make8BallAnswer()

		// It's one of the allowed answers
		if _, ok := answersMap[testAnswers]; !ok {
			t.Errorf("Wrong answer for 8ball : %s", testAnswers)
		}
	}
}
