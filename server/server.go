package main

import (
	"bufio"
	"context"
	"github.com/Manav-Aggarwal/herd-mentality/server/pb"
	"github.com/Manav-Aggarwal/herd-mentality/server/pb/hm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type server struct {
	questions []string

	ticker *time.Ticker

	endpoint string
	client   pb.MsgClient
}

func (s *server) CurrentQuestion(ctx context.Context, request *hm.QuestionRequest) (*hm.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) SubmitAnswer(ctx context.Context, answer *hm.Answer) (*hm.Confirmation, error) {
	log.Printf("submitting answer: %+v\n", answer)
	resp, err := s.client.SubmitAnswer(context.Background(), &pb.MsgSubmitAnswer{
		Creator: "root",
		Qid:     strconv.FormatUint(answer.QuestionId, 10),
		Player:  answer.UserName,
		Answer:  answer.Answer,
	})

	log.Println(resp, err)
	return nil, err
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

func (s *server) start(ctx context.Context) {
	conn, err := grpc.Dial(s.endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	s.client = pb.NewMsgClient(conn)

	s.ticker = time.NewTicker(60 * time.Second)
	s.submitQuestion()
	go s.timerRoutine(ctx)
}

func (s *server) submitQuestion() {
	id, question := s.randomQuestion()
	log.Println("submitting question", id, question)
	resp, err := s.client.SubmitQuestion(context.Background(), &pb.MsgSubmitQuestion{
		Creator:  "root",
		Qid:      strconv.Itoa(id),
		Question: question,
	})

	log.Println(resp, err)
}

func (s *server) randomQuestion() (int, string) {
	id := rand.Int() % len(s.questions)
	return id, s.questions[id]
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
