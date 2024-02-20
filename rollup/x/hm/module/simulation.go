package hm

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"hm/testutil/sample"
	hmsimulation "hm/x/hm/simulation"
	"hm/x/hm/types"
)

// avoid unused import issue
var (
	_ = hmsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitQuestion = "op_weight_msg_submit_question"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitQuestion int = 100

	opWeightMsgSubmitAnswer = "op_weight_msg_submit_answer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitAnswer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	hmGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&hmGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitQuestion int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitQuestion, &weightMsgSubmitQuestion, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitQuestion = defaultWeightMsgSubmitQuestion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitQuestion,
		hmsimulation.SimulateMsgSubmitQuestion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitAnswer int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitAnswer, &weightMsgSubmitAnswer, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitAnswer = defaultWeightMsgSubmitAnswer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitAnswer,
		hmsimulation.SimulateMsgSubmitAnswer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitQuestion,
			defaultWeightMsgSubmitQuestion,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hmsimulation.SimulateMsgSubmitQuestion(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitAnswer,
			defaultWeightMsgSubmitAnswer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hmsimulation.SimulateMsgSubmitAnswer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
