package shared

import (
	"github.com/bloxapp/go-casper-ghost-SDK/src/core"
	"github.com/bloxapp/go-casper-ghost-SDK/src/shared/params"
	"github.com/prysmaticlabs/go-bitfield"
	"github.com/ulule/deepcopier"
)

func CopyState(state *core.State) *core.State {
	if state == nil {
		return nil
	}

	ret := &core.State{}

	ret.CurrentSlot = state.CurrentSlot

	ret.BlockRoots = make([][]byte, len(state.BlockRoots))
	for i, r := range state.BlockRoots {
		ret.BlockRoots[i] = make([]byte, 0)
		deepcopier.Copy(r).To(ret.BlockRoots[i])
	}

	ret.StateRoots = make([][]byte, len(state.StateRoots))
	for i, r := range state.StateRoots {
		ret.StateRoots[i] = make([]byte, 0)
		deepcopier.Copy(r).To(ret.StateRoots[i])
	}

	ret.RandaoMix = make([][]byte, len(state.RandaoMix))
	for i, r := range state.RandaoMix {
		ret.RandaoMix[i] = make([]byte, 0)
		deepcopier.Copy(r).To(ret.RandaoMix[i])
	}

	ret.Validators = make([]*core.Validator, len(state.Validators))
	for i, bp := range state.Validators {
		ret.Validators[i] = &core.Validator{}
		deepcopier.Copy(bp).To(ret.Validators[i])
	}

	ret.PreviousEpochAttestations = make([]*core.PendingAttestation, len(state.PreviousEpochAttestations))
	for i, pe := range state.PreviousEpochAttestations {
		ret.PreviousEpochAttestations[i] = &core.PendingAttestation{}
		deepcopier.Copy(pe).To(ret.PreviousEpochAttestations[i])
	}

	ret.CurrentEpochAttestations = make([]*core.PendingAttestation, len(state.CurrentEpochAttestations))
	for i, pe := range state.CurrentEpochAttestations {
		ret.CurrentEpochAttestations[i] = &core.PendingAttestation{}
		deepcopier.Copy(pe).To(ret.CurrentEpochAttestations[i])
	}

	ret.JustificationBits = make(bitfield.Bitvector4, len(state.JustificationBits))
	deepcopier.Copy(state.JustificationBits).To(ret.JustificationBits)

	if state.PreviousJustifiedCheckpoint != nil {
		ret.PreviousJustifiedCheckpoint = &core.Checkpoint{}
		deepcopier.Copy(state.PreviousJustifiedCheckpoint).To(ret.PreviousJustifiedCheckpoint)
	}

	ret.CurrentJustifiedCheckpoint = &core.Checkpoint{}
	deepcopier.Copy(state.CurrentJustifiedCheckpoint).To(ret.CurrentJustifiedCheckpoint)

	if state.FinalizedCheckpoint != nil {
		ret.FinalizedCheckpoint = &core.Checkpoint{}
		deepcopier.Copy(state.FinalizedCheckpoint).To(ret.FinalizedCheckpoint)
	}

	if state.LatestBlockHeader != nil {
		ret.LatestBlockHeader = &core.BlockHeader{}
		deepcopier.Copy(state.LatestBlockHeader).To(ret.LatestBlockHeader)
	}

	return ret
}

// will return nil if not found or inactive
func GetValidator (state *core.State, id uint64) *core.Validator {
	for _, p := range state.Validators {
		if p.GetId() == id && p.Active {
			return p
		}
	}
	return nil
}

/**
def is_valid_genesis_state(state: BeaconState) -> bool:
    if state.genesis_time < MIN_GENESIS_TIME:
        return False
    if len(get_active_validator_indices(state, GENESIS_EPOCH)) < MIN_GENESIS_ACTIVE_VALIDATOR_COUNT:
        return False
    return True
 */
func IsValidGenesisState(state *core.State) bool {
	if state.GenesisTime < params.ChainConfig.MinGenesisTime {
		return false
	}
	if uint64(len(GetActiveValidators(state, params.ChainConfig.GenesisEpoch))) < params.ChainConfig.MinGenesisActiveBPCount {
		return false
	}
	return true
}

func SumSlashings(state *core.State) uint64 {
	totalSlashing := uint64(0)
	for _, slashing := range state.Slashings {
		totalSlashing += slashing
	}
	return totalSlashing
}