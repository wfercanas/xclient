package dto

type Contribution struct {
	ContributionName string  `json:"contributionName"`
	ContributionId   string  `json:"contributionId"`
	Balance          float64 `json:"balance"`
}
