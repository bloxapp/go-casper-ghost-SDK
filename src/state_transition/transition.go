package state_transition

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/bloxapp/go-casper-ghost-SDK/src/core"
	"github.com/bloxapp/go-casper-ghost-SDK/src/shared"
)

type IStateTransition interface {
	// ExecuteStateTransition runs ExecuteNoVerify and verifies post transition state root
	//
	// Spec pseudocode definition:
	//  def state_transition(state: BeaconState, signed_block: SignedBeaconBlock, validate_result: bool=True) -> BeaconState:
	//    block = signed_block.message
	//    # Process slots (including those with no blocks) since block
	//    process_slots(state, block.slot)
	//    # Verify signature
	//    if validate_result:
	//        assert verify_block_signature(state, signed_block)
	//    # Process block
	//    process_block(state, block)
	//    if validate_result:
	//        assert block.state_root == hash_tree_root(state)
	//    # Return post-state
	//    return state
	ExecuteStateTransition(state *core.State, signedBlock *core.SignedBlock, validateResult bool) (newState *core.State, err error)

	// ProcessBlock creates a new, modified beacon state by applying block operation
	// transformations as defined in the Ethereum Serenity specification, including processing proposer slashings,
	// processing block attestations, and more.
	//
	// Spec pseudocode definition:
	//
	//  def process_block(state: BeaconState, block: BeaconBlock) -> None:
	//    process_block_header(state, block)
	//    process_randao(state, block.block)
	//    process_eth1_data(state, block.block)
	//    process_operations(state, block.block)
	ProcessBlock(state *core.State, newBlockBody *core.SignedBlock) error
	// ProcessSlots process through skip slots and apply epoch transition when it's needed
	//
	// Spec pseudocode definition:
	//  def process_slots(state: BeaconState, slot: Slot) -> None:
	//    assert state.slot <= slot
	//    while state.slot < slot:
	//        process_slot(state)
	//        # Process epoch on the first slot of the next epoch
	//        if (state.slot + 1) % SLOTS_PER_EPOCH == 0:
	//            process_epoch(state)
	//        state.slot += 1
	//    ]
	ProcessSlots(state *core.State, slot uint64) error
}

type StateTransition struct {}
func NewStateTransition() *StateTransition { return &StateTransition{} }

func (st *StateTransition)ExecuteStateTransition(state *core.State, signedBlock *core.SignedBlock, validateResult bool) (newState *core.State, err error) {
	newState = shared.CopyState(state)

	if err := st.ProcessSlots(newState, signedBlock.Block.Slot); err != nil {
		return nil, fmt.Errorf("ExecuteStateTransition: %s", err.Error())
	}

	if validateResult {
		err := shared.VerifyBlockSig(newState, signedBlock)
		if err != nil {
			return nil, fmt.Errorf("ExecuteStateTransition: %s", err.Error())
		}
	}

	if err := st.ProcessBlock(newState, signedBlock.Block); err != nil {
		return nil, fmt.Errorf("ExecuteStateTransition: %s", err.Error())
	}

	if validateResult {
		postStateRoot, err := newState.HashTreeRoot()
		if err != nil {
			return nil, fmt.Errorf("ExecuteStateTransition: %s", err.Error())
		}
		if !bytes.Equal(signedBlock.Block.StateRoot, postStateRoot[:]) {
			return nil, fmt.Errorf("ExecuteStateTransition: new block state root is wrong, expected %s", hex.EncodeToString(postStateRoot[:]))
		}
	}

	return newState, nil
}
