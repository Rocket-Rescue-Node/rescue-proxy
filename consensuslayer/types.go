package consensuslayer

type PrepareBeaconProposerRequest []struct {
	ValidatorIndex string `json:"validator_index"`
	FeeRecipient   string `json:"fee_recipient"`
}

type RegisterValidatorMessage struct {
	FeeRecipient string `json:"fee_recipient"`

	// Omitting gas_limit and timestamp

	Pubkey string `json:"pubkey"`
}

type RegisterValidatorRequest []struct {
	Message RegisterValidatorMessage `json:"message"`

	// Omitting signature. The BN will validate it for us.
}
