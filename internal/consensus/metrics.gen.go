// Code generated by metricsgen. DO NOT EDIT.

package consensus

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "height",
			Help:      "Height of the chain.",
		}, labels).With(labelsAndValues...),
		ValidatorLastSignedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_last_signed_height",
			Help:      "Last height signed by this validator if the node is a validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		Rounds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rounds",
			Help:      "Number of rounds.",
		}, labels).With(labelsAndValues...),
		RoundDuration: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "round_duration",
			Help:      "Histogram of round duration.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		Validators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators",
			Help:      "Number of validators.",
		}, labels).With(labelsAndValues...),
		ValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validators_power",
			Help:      "Total power of all validators.",
		}, labels).With(labelsAndValues...),
		ValidatorPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_power",
			Help:      "Power of a validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		ValidatorMissedBlocks: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validator_missed_blocks",
			Help:      "Amount of blocks missed per validator.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		MissingValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators",
			Help:      "Number of validators who did not sign.",
		}, labels).With(labelsAndValues...),
		MissingValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_validators_power",
			Help:      "Total power of the missing validators.",
		}, labels).With(labelsAndValues...),
		ByzantineValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators",
			Help:      "Number of validators who tried to double sign.",
		}, labels).With(labelsAndValues...),
		ByzantineValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "byzantine_validators_power",
			Help:      "Total power of the byzantine validators.",
		}, labels).With(labelsAndValues...),
		BlockIntervalSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_interval_seconds",
			Help:      "Time between this and the last block.",
		}, labels).With(labelsAndValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "num_txs",
			Help:      "Number of transactions.",
		}, labels).With(labelsAndValues...),
		BlockSizeBytes: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_size_bytes",
			Help:      "Size of the block.",
		}, labels).With(labelsAndValues...),
		TotalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "total_txs",
			Help:      "Total number of transactions.",
		}, labels).With(labelsAndValues...),
		CommittedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "latest_block_height",
			Help:      "The latest block height.",
		}, labels).With(labelsAndValues...),
		BlockSyncing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_syncing",
			Help:      "Whether or not a node is block syncing. 1 if yes, 0 if no.",
		}, labels).With(labelsAndValues...),
		StateSyncing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "state_syncing",
			Help:      "Whether or not a node is state syncing. 1 if yes, 0 if no.",
		}, labels).With(labelsAndValues...),
		BlockParts: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_parts",
			Help:      "Number of block parts transmitted by each peer.",
		}, append(labels, "peer_id")).With(labelsAndValues...),
		StepDuration: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "step_duration",
			Help:      "Histogram of durations for each step in the consensus protocol.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(labels, "step")).With(labelsAndValues...),
		BlockGossipReceiveLatency: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_gossip_receive_latency",
			Help:      "Histogram of time taken to receive a block in seconds, measured between when a new block is first discovered to when the block is completed.",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		BlockGossipPartsReceived: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_gossip_parts_received",
			Help:      "Number of block parts received by the node, separated by whether the part was relevant to the block the node is trying to gather or not.",
		}, append(labels, "matches_current")).With(labelsAndValues...),
		ProposalBlockCreatedOnPropose: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_block_created_on_propose",
			Help:      "Number of proposal blocks created on propose received.",
		}, append(labels, "success")).With(labelsAndValues...),
		ProposalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_txs",
			Help:      "Number of txs in a proposal.",
		}, labels).With(labelsAndValues...),
		ProposalMissingTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_missing_txs",
			Help:      "Number of missing txs when trying to create proposal.",
		}, labels).With(labelsAndValues...),
		MissingTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "missing_txs",
			Help:      "Number of missing txs when a proposal is received.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		QuorumPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "quorum_prevote_delay",
			Help:      "Interval in seconds between the proposal timestamp and the timestamp of the earliest prevote that achieved a quorum.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		FullPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "full_prevote_delay",
			Help:      "Interval in seconds between the proposal timestamp and the timestamp of the latest prevote in a round where all validators voted.",
		}, append(labels, "proposer_address")).With(labelsAndValues...),
		ProposalTimestampDifference: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_timestamp_difference",
			Help:      "Difference between the timestamp in the proposal message and the local time of the validator at the time it received the message.",
			Buckets: []float64{-10, -.5, -.025, 0, .1, .5, 1, 1.5, 2, 10},
		}, append(labels, "is_timely")).With(labelsAndValues...),
		VoteExtensionReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "vote_extension_receive_count",
			Help:      "Number of vote extensions received labeled by application response status.",
		}, append(labels, "status")).With(labelsAndValues...),
		ProposalReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_receive_count",
			Help:      "Total number of proposals received by the node since process start labeled by application response status.",
		}, append(labels, "status")).With(labelsAndValues...),
		ProposalCreateCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "proposal_create_count",
			Help:      "Total number of proposals created by the node since process start.",
		}, labels).With(labelsAndValues...),
		RoundVotingPowerPercent: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "round_voting_power_percent",
			Help:      "A value between 0 and 1.0 representing the percentage of the total voting power per vote type received within a round.",
		}, append(labels, "vote_type")).With(labelsAndValues...),
		LateVotes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "late_votes",
			Help:      "Number of votes received by the node since process start that correspond to earlier heights and rounds than this node is currently in.",
		}, append(labels, "validator_address")).With(labelsAndValues...),
		FinalRound: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "final_round",
			Help:      "Number of rounds going through from this proposer.",
			Buckets: []float64{0,1,2,3,5,10},
		}, append(labels, "proposer_address")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Height:                      discard.NewGauge(),
		ValidatorLastSignedHeight:   discard.NewGauge(),
		Rounds:                      discard.NewGauge(),
		RoundDuration:               discard.NewHistogram(),
		Validators:                  discard.NewGauge(),
		ValidatorsPower:             discard.NewGauge(),
		ValidatorPower:              discard.NewGauge(),
		ValidatorMissedBlocks:       discard.NewGauge(),
		MissingValidators:           discard.NewGauge(),
		MissingValidatorsPower:      discard.NewGauge(),
		ByzantineValidators:         discard.NewGauge(),
		ByzantineValidatorsPower:    discard.NewGauge(),
		BlockIntervalSeconds:        discard.NewHistogram(),
		NumTxs:                      discard.NewGauge(),
		BlockSizeBytes:              discard.NewHistogram(),
		TotalTxs:                    discard.NewGauge(),
		CommittedHeight:             discard.NewGauge(),
		BlockSyncing:                discard.NewGauge(),
		StateSyncing:                discard.NewGauge(),
		BlockParts:                  discard.NewCounter(),
		StepDuration:                discard.NewHistogram(),
		BlockGossipReceiveLatency:   discard.NewHistogram(),
		BlockGossipPartsReceived:    discard.NewCounter(),
		QuorumPrevoteDelay:            discard.NewGauge(),
		FullPrevoteDelay:              discard.NewGauge(),
		ProposalTimestampDifference:   discard.NewHistogram(),
		VoteExtensionReceiveCount:     discard.NewCounter(),
		ProposalReceiveCount:          discard.NewCounter(),
		ProposalCreateCount:           discard.NewCounter(),
		RoundVotingPowerPercent:       discard.NewGauge(),
		LateVotes:                     discard.NewCounter(),
		ProposalBlockCreatedOnPropose: discard.NewCounter(),
		ProposalTxs:                   discard.NewGauge(),
		ProposalMissingTxs:            discard.NewGauge(),
		MissingTxs:                    discard.NewGauge(),
		FinalRound: discard.NewHistogram(),
	}
}
