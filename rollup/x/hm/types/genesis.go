package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		QuestionList: []Question{},
		AnswerList:   []Answer{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in question
	questionIndexMap := make(map[string]struct{})

	for _, elem := range gs.QuestionList {
		index := string(QuestionKey(elem.Index))
		if _, ok := questionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for question")
		}
		questionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in answer
	answerIndexMap := make(map[string]struct{})

	for _, elem := range gs.AnswerList {
		index := string(AnswerKey(elem.Index))
		if _, ok := answerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for answer")
		}
		answerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
