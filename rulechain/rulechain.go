package rulechain

import (
	"context"
	"time"
)

//DataSource datasource
type DataSource struct {
	Name         string
	Protocol     string
	IsProvider   bool
	ServicePort  string
	IsTLSEnabled bool
	ConnectURL   string
	CertFile     []byte
	KeyFile      []byte
}

//RuleChain rulechain
type RuleChain struct {
	Name         string
	ID           string
	Description  string
	DebugMode    bool
	UserID       string
	Type         string
	Domain       string
	Status       string
	Payload      []byte
	Root         bool
	Channel      string
	SubTopic     string
	CreateAt     time.Time
	LastUpdateAt time.Time
	DataSource   DataSource
}

// Validate returns an error if representtation is invalid
func (r RuleChain) Validate() error {
	if r.ID == "" || r.UserID == "" {
		return ErrMalformedEntity
	}
	return nil
}

//RuleChainRepository specifies realm persistence API
type RuleChainRepository interface {
	//Save save the rulechain
	Save(context.Context, RuleChain) error

	//Update the rulechain
	Update(context.Context, RuleChain) error

	//Retrieve return rulechain by userid and rulechain id
	Retrieve(context.Context, string, string) (RuleChain, error)

	//Revoke remove rulechain by userid and rulechain id
	Revoke(context.Context, string, string) error

	//List return all rulechains
	List(context.Context, string) ([]RuleChain, error)
}

// RuleChainCache contains thing caching interface.
type RuleChainCache interface {
	// Save stores pair thing key, thing id.
	Save(context.Context, string, string) error

	// ID returns thing ID for given key.
	ID(context.Context, string) (string, error)

	// Removes thing from cache.
	Remove(context.Context, string) error
}
