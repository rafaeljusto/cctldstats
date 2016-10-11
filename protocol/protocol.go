package protocol

// RegisteredDomainsResponse store statistic information about the registered domains of the ccTLD.
type RegisteredDomainsResponse struct {
	// Number quantity of domains registered.
	Number int `json:"number"`
}
