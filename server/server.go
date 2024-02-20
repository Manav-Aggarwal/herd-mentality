package main

import (
	"context"
	"github.com/Manav-Aggarwal/herd-mentality/server/pb/hm"
	"io"
)

type server struct {
	questions []string
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
}
