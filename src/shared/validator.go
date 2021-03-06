package shared

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/bloxapp/go-casper-ghost-SDK/src/core"
	"github.com/bloxapp/go-casper-ghost-SDK/src/shared/params"
	"github.com/prysmaticlabs/prysm/shared/hashutil"
	"github.com/prysmaticlabs/prysm/shared/mathutil"
	"github.com/wealdtech/go-bytesutil"
)

/**
	def is_active_validator(validator: Validator, epoch: Epoch) -> bool:
		"""
		Check if ``validator`` is active.
		"""
		return validator.activation_epoch <= epoch < validator.exit_epoch
 */
func IsActiveValidator(bp *core.Validator, epoch uint64) bool {
	return bp.ActivationEpoch <= epoch && epoch < bp.ExitEpoch
}

/**
	def is_eligible_for_activation_queue(validator: Validator) -> bool:
		"""
		Check if ``validator`` is eligible to be placed into the activation queue.
		"""
		return (
			validator.activation_eligibility_epoch == FAR_FUTURE_EPOCH
			and validator.effective_balance == MAX_EFFECTIVE_BALANCE
		)
*/
func IsEligibleForActivationQueue(bp *core.Validator) bool {
	return bp.ActivationEligibilityEpoch == params.ChainConfig.FarFutureEpoch && bp.EffectiveBalance == params.ChainConfig.MaxEffectiveBalance
}

/**
	def is_eligible_for_activation(state: BeaconState, validator: Validator) -> bool:
		"""
		Check if ``validator`` is eligible for activation.
		"""
		return (
			# Placement in queue is finalized
			validator.activation_eligibility_epoch <= state.finalized_checkpoint.epoch
			# Has not yet been activated
			and validator.activation_epoch == FAR_FUTURE_EPOCH
		)
 */
func IsEligibleForActivation(state *core.State, bp *core.Validator) bool {
	return bp.ActivationEligibilityEpoch <= state.FinalizedCheckpoint.Epoch && // Placement in queue is finalized
					bp.ActivationEpoch == params.ChainConfig.FarFutureEpoch // Has not yet been activated
}

/**
	def is_slashable_validator(validator: Validator, epoch: Epoch) -> bool:
		"""
		Check if ``validator`` is slashable.
		"""
		return (not validator.slashed) and (validator.activation_epoch <= epoch < validator.withdrawable_epoch)
 */
func IsSlashableValidator(bp *core.Validator, epoch uint64) bool {
	return !bp.Slashed && (bp.ActivationEpoch <= epoch && epoch < bp.WithdrawableEpoch)
}

/**
def compute_proposer_index(state: BeaconState, indices: Sequence[ValidatorIndex], seed: Bytes32) -> ValidatorIndex:
    """
    Return from ``indices`` a random index sampled by effective balance.
    """
    assert len(indices) > 0
    MAX_RANDOM_BYTE = 2**8 - 1
    i = uint64(0)
    total = uint64(len(indices))
    while True:
        candidate_index = indices[compute_shuffled_index(i % total, total, seed)]
        random_byte = hash(seed + uint_to_bytes(uint64(i // 32)))[i % 32]
        effective_balance = state.validators[candidate_index].effective_balance
        if effective_balance * MAX_RANDOM_BYTE >= MAX_EFFECTIVE_BALANCE * random_byte:
            return candidate_index
        i += 1
 */
func ComputeProposerIndex(state *core.State, indices []uint64, seed []byte) (uint64, error) {
	if len(indices) == 0 {
		return 0, fmt.Errorf("couldn't compute proposer, indices list empty")
	}
	maxRandomByte := uint64(1<<8-1)
	i := uint64(0)
	total := uint64(len(indices))
	for {
		idx, err := computeShuffledIndex(i % total, total, bytesutil.ToBytes32(seed), true, params.ChainConfig.ShuffleRoundCount)
		if err != nil {
			return 0, err
		}

		candidateIndex := indices[idx]
		b := append(seed[:], bytesutil.Bytes8(i / 32)...)
		randomByte := hashutil.Hash(b)[i%32]

		bp := GetValidator(state, candidateIndex)
		if bp == nil {
			return 0, fmt.Errorf("could not find shuffled BP index %d", candidateIndex)
		}
		effectiveBalance := bp.EffectiveBalance

		if effectiveBalance * maxRandomByte >= params.ChainConfig.MaxEffectiveBalance * uint64(randomByte) {
			return candidateIndex, nil
		}
		i++
	}
}

/**
def compute_activation_exit_epoch(epoch: Epoch) -> Epoch:
    """
    Return the epoch during which validator activations and exits initiated in ``epoch`` take effect.
    """
    return Epoch(epoch + 1 + MAX_SEED_LOOKAHEAD)
 */
func ComputeActivationExitEpoch(epoch uint64) uint64 {
	return epoch + 1 + params.ChainConfig.MaxSeedLookahead
}

/**
def get_active_validator_indices(state: BeaconState, epoch: Epoch) -> Sequence[ValidatorIndex]:
    """
    Return the sequence of active validator indices at ``epoch``.
    """
    return [ValidatorIndex(i) for i, v in enumerate(state.validators) if is_active_validator(v, epoch)]
 */
func GetActiveValidators(state *core.State, epoch uint64) []uint64 {
	var activeBps []uint64
	for i, val := range state.Validators {
		if IsActiveValidator(val, epoch) {
			activeBps = append(activeBps, uint64(i))
		}
	}
	return activeBps
}

/**
def get_validator_churn_limit(state: BeaconState) -> uint64:
    """
    Return the validator churn limit for the current epoch.
    """
    active_validator_indices = get_active_validator_indices(state, get_current_epoch(state))
    return max(MIN_PER_EPOCH_CHURN_LIMIT, uint64(len(active_validator_indices)) // CHURN_LIMIT_QUOTIENT)
 */
func GetValidatorChurnLimit(state *core.State) uint64 {
	activeBPs := GetActiveValidators(state, GetCurrentEpoch(state))
	churLimit := uint64(len(activeBPs)) / params.ChainConfig.ChurnLimitQuotient
	if churLimit < params.ChainConfig.MinPerEpochChurnLimit {
		churLimit = params.ChainConfig.MinPerEpochChurnLimit
	}
	return churLimit
}

/**
def get_beacon_proposer_index(state: BeaconState) -> ValidatorIndex:
    """
    Return the beacon proposer index at the current slot.
    """
    epoch = get_current_epoch(state)
    seed = hash(get_seed(state, epoch, DOMAIN_BEACON_PROPOSER) + uint_to_bytes(state.slot))
    indices = get_active_validator_indices(state, epoch)
    return compute_proposer_index(state, indices, seed)
 */
func GetBlockProposerIndex(state *core.State) (uint64, error) {
	epoch := GetCurrentEpoch(state)
	seed := GetSeed(state, epoch, params.ChainConfig.DomainBeaconProposer)
	SeedWithSlot := append(seed[:], bytesutil.Bytes8(state.Slot)...)
	hash := hashutil.Hash(SeedWithSlot)

	validators := GetActiveValidators(state, epoch)
	return ComputeProposerIndex(state, validators, hash[:])
}

/**
def increase_balance(state: BeaconState, index: ValidatorIndex, delta: Gwei) -> None:
    """
    Increase the validator balance at index ``index`` by ``delta``.
    """
    state.balances[index] += delta
 */
func IncreaseBalance(state *core.State, index uint64, delta uint64) {
	state.Balances[index] += delta
}

/**
def decrease_balance(state: BeaconState, index: ValidatorIndex, delta: Gwei) -> None:
    """
    Decrease the validator balance at index ``index`` by ``delta``, with underflow protection.
    """
    state.balances[index] = 0 if delta > state.balances[index] else state.balances[index] - delta
*/
func DecreaseBalance(state *core.State, index uint64, delta uint64) {
	if bp := GetValidator(state, index); bp != nil {
		if delta > state.Balances[index] {
			state.Balances[index] = 0
		} else {
			state.Balances[index]  -= delta
		}
	}
}

/**
def initiate_validator_exit(state: BeaconState, index: ValidatorIndex) -> None:
    """
    Initiate the exit of the validator with index ``index``.
    """
    # Return if validator already initiated exit
    validator = state.validators[index]
    if validator.exit_epoch != FAR_FUTURE_EPOCH:
        return

    # Compute exit queue epoch
    exit_epochs = [v.exit_epoch for v in state.validators if v.exit_epoch != FAR_FUTURE_EPOCH]
    exit_queue_epoch = max(exit_epochs + [compute_activation_exit_epoch(get_current_epoch(state))])
    exit_queue_churn = len([v for v in state.validators if v.exit_epoch == exit_queue_epoch])
    if exit_queue_churn >= get_validator_churn_limit(state):
        exit_queue_epoch += Epoch(1)

    # Set validator exit epoch and withdrawable epoch
    validator.exit_epoch = exit_queue_epoch
    validator.withdrawable_epoch = Epoch(validator.exit_epoch + MIN_VALIDATOR_WITHDRAWABILITY_DELAY)
 */
func InitiateValidatorExit(state *core.State, index uint64) {
	validator := GetValidator(state, index)
	if validator == nil {
		return
	}
	if validator.ExitEpoch != params.ChainConfig.FarFutureEpoch {
		return
	}

	// Compute exit queue epoch
	exitEpochs := []uint64{}
	for _, val := range state.Validators {
		if val.ExitEpoch != params.ChainConfig.FarFutureEpoch {
			exitEpochs = append(exitEpochs, val.ExitEpoch)
		}
	}
	exitEpochs = append(exitEpochs, ComputeActivationExitEpoch(GetCurrentEpoch(state)))

	// Obtain the exit queue epoch as the maximum number in the exit epochs array.
	exitQueueEpoch := uint64(0)
	for _, i := range exitEpochs {
		if exitQueueEpoch < i {
			exitQueueEpoch = i
		}
	}

	// We use the exit queue churn to determine if we have passed a churn limit.
	exitQueueChurn := uint64(0)
	for _, val := range state.Validators {
		if val.ExitEpoch == exitQueueEpoch {
			exitQueueChurn ++
		}
	}
	if exitQueueChurn >= GetValidatorChurnLimit(state) {
		exitQueueEpoch ++
	}

	// Set validator exit epoch and withdrawable epoch
	validator.ExitEpoch = exitQueueEpoch
	validator.WithdrawableEpoch = validator.ExitEpoch + params.ChainConfig.MinValidatorWithdrawabilityDelay
}

/**
def slash_validator(state: BeaconState,
                    slashed_index: ValidatorIndex,
                    whistleblower_index: ValidatorIndex=None) -> None:
    """
    Slash the validator with index ``slashed_index``.
    """
    epoch = get_current_epoch(state)
    initiate_validator_exit(state, slashed_index)
    validator = state.validators[slashed_index]
    validator.slashed = True
    validator.withdrawable_epoch = max(validator.withdrawable_epoch, Epoch(epoch + EPOCHS_PER_SLASHINGS_VECTOR))
    state.slashings[epoch % EPOCHS_PER_SLASHINGS_VECTOR] += validator.effective_balance
    decrease_balance(state, slashed_index, validator.effective_balance // MIN_SLASHING_PENALTY_QUOTIENT)

    # Apply proposer and whistleblower rewards
    proposer_index = get_beacon_proposer_index(state)
    if whistleblower_index is None:
        whistleblower_index = proposer_index
    whistleblower_reward = Gwei(validator.effective_balance // WHISTLEBLOWER_REWARD_QUOTIENT)
    proposer_reward = Gwei(whistleblower_reward // PROPOSER_REWARD_QUOTIENT)
    increase_balance(state, proposer_index, proposer_reward)
    increase_balance(state, whistleblower_index, Gwei(whistleblower_reward - proposer_reward))
 */
func SlashValidator(state *core.State, slashedIndex uint64) error {
	epoch := GetCurrentEpoch(state)
	InitiateValidatorExit(state, slashedIndex)
	validator := GetValidator(state, slashedIndex)
	if validator == nil {
		return fmt.Errorf("slash validator: block producer not found")
	}
	validator.Slashed = true
	validator.WithdrawableEpoch = mathutil.Max(validator.WithdrawableEpoch, epoch + params.ChainConfig.EpochsPerSlashingVector)
	state.Slashings[epoch % params.ChainConfig.EpochsPerSlashingVector] += validator.EffectiveBalance
	DecreaseBalance(state, slashedIndex, validator.EffectiveBalance / params.ChainConfig.MinSlashingPenaltyQuotient)

	// Apply proposer and whistleblower rewards
	proposer, err := GetBlockProposerIndex(state)
	if err != nil {
		return err
	}
	whistleblowerIndex := proposer
	whistleblowerReward := validator.EffectiveBalance / params.ChainConfig.WhitstleblowerRewardQuotient
	proposerReward := whistleblowerReward / params.ChainConfig.ProposerRewardQuotient
	IncreaseBalance(state, proposer, proposerReward)
	IncreaseBalance(state, whistleblowerIndex, whistleblowerReward - proposerReward)
	return nil
}


// returns error if not found
func ValidatorIndexByPubkey(state *core.State, pk []byte) (uint64, error) {
	// TODO - ValidatorIndexByPubkey optimize with some kind of map
	for i, bp := range state.Validators {
		if bytes.Equal(pk, bp.PublicKey) {
			return uint64(i), nil
		}
	}
	return 0, fmt.Errorf("validator not found for pk: %s", hex.EncodeToString(pk))
}