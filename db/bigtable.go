package db

import (
	"eth2-exporter/types"
	"fmt"

	gcp_bigtable "cloud.google.com/go/bigtable"
)

var BigtableClient *Bigtable

const (
	DEFAULT_FAMILY            = "f"
	VALIDATOR_BALANCES_FAMILY = "vb"
	ATTESTATIONS_FAMILY       = "at"
	PROPOSALS_FAMILY          = "pr"
	SYNC_COMMITTEES_FAMILY    = "sc"

	max_block_number = 1000000000
	max_epoch        = 1000000000
)

type Bigtable struct {
	client           *gcp_bigtable.Client
	tableData        *gcp_bigtable.Table
	tableBlocks      *gcp_bigtable.Table
	tableBeaconchain *gcp_bigtable.Table
	chainId          string
}

func InitBigtable(project, instance, chainId string) (*Bigtable, error) {
	return &Bigtable{}, nil
}

func (bigtable *Bigtable) Close() {
}

func (bigtable *Bigtable) SaveValidatorBalances(epoch uint64, validators []*types.Validator) error {
	return nil
}

func (bigtable *Bigtable) SaveAttestationAssignments(epoch uint64, assignments map[string]uint64) error {
	return nil
}

func (bigtable *Bigtable) SaveProposalAssignments(epoch uint64, assignments map[uint64]uint64) error {
	return nil
}

func (bigtable *Bigtable) SaveSyncCommitteesAssignments(startSlot, endSlot uint64, validators []uint64) error {
	return nil
}

func (bigtable *Bigtable) SaveAttestations(blocks map[uint64]map[string]*types.Block) error {
	return nil
}

func (bigtable *Bigtable) SaveProposals(blocks map[uint64]map[string]*types.Block) error {
	return nil
}

func (bigtable *Bigtable) SaveSyncComitteeDuties(blocks map[uint64]map[string]*types.Block) error {
	return nil
}

func (bigtable *Bigtable) GetValidatorBalanceHistory(validators []uint64, startEpoch uint64, limit int64) (map[uint64][]*types.ValidatorBalance, error) {
	res := make(map[uint64][]*types.ValidatorBalance, len(validators))
	return res, nil
}

func (bigtable *Bigtable) GetValidatorAttestationHistory(validators []uint64, startEpoch uint64, limit int64) (map[uint64][]*types.ValidatorAttestation, error) {
	res := make(map[uint64][]*types.ValidatorAttestation, len(validators))
	return res, nil
}

func (bigtable *Bigtable) GetValidatorSyncDutiesHistory(validators []uint64, startEpoch uint64, limit int64) (map[uint64][]*types.ValidatorSyncParticipation, error) {
	res := make(map[uint64][]*types.ValidatorSyncParticipation, len(validators))
	return res, nil
}

func (bigtable *Bigtable) GetValidatorMissedAttestationsCount(validators []uint64, startEpoch uint64, limit int64) (map[uint64]*types.ValidatorMissedAttestationsStatistic, error) {
	res := make(map[uint64]*types.ValidatorMissedAttestationsStatistic)
	return res, nil
}

func (bigtable *Bigtable) GetValidatorSyncDutiesStatistics(validators []uint64, startEpoch uint64, limit int64) (map[uint64]*types.ValidatorSyncDutiesStatistic, error) {
	res := make(map[uint64]*types.ValidatorSyncDutiesStatistic)
	return res, nil
}

func (bigtable *Bigtable) GetValidatorEffectiveness(validators []uint64, epoch uint64) ([]*types.ValidatorEffectiveness, error) {
	res := make([]*types.ValidatorEffectiveness, 0, len(validators))
	return res, nil
}

func (bigtable *Bigtable) GetValidatorBalanceStatistics(startEpoch, endEpoch uint64) (map[uint64]*types.ValidatorBalanceStatistic, error) {
	res := make(map[uint64]*types.ValidatorBalanceStatistic)
	return res, nil
}

func (bigtable *Bigtable) GetValidatorProposalHistory(validators []uint64, startEpoch uint64, limit int64) (map[uint64][]*types.ValidatorProposal, error) {
	res := make(map[uint64][]*types.ValidatorProposal, len(validators))
	return res, nil
}

func reversedPaddedEpoch(epoch uint64) string {
	return fmt.Sprintf("%09d", max_block_number-epoch)
}

func reversedPaddedSlot(slot uint64) string {
	return fmt.Sprintf("%09d", max_block_number-slot)
}
