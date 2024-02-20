package main

import (
	"bufio"
	"context"
	"github.com/Manav-Aggarwal/herd-mentality/server/pb/hm"
	"io"
	"time"
)

type server struct {
	questions []string

	ticker *time.Ticker
}

func (s *server) start(ctx context.Context) {
	s.ticker = time.NewTicker(60 * time.Second)
	go s.timerRoutine(ctx)
}

func (s *server) CurrentQuestion(ctx context.Context, request *hm.QuestionRequest) (*hm.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) SubmitAnswer(ctx context.Context, answer *hm.Answer) (*hm.Confirmation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetResults(ctx context.Context, request *hm.ResultsRequest) (*hm.Results, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) loadQuestions(questions io.Reader) {
	scanner := bufio.NewScanner(questions)

	for scanner.Scan() {
		s.questions = append(s.questions, scanner.Text())
	}
}

func (s *server) submitQuestion() {
	// TODO
}

func (s *server) timerRoutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		case <-s.ticker.C:
			s.submitQuestion()
		default:

		}
	}
}
