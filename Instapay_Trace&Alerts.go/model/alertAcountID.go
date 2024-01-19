package model

import "time"

type (
	AccountAlertRequest1 struct {
		AccountAlertID string `json:"accountalertid"`
	}

	ResponseInfo1 struct {
		ID                           string    `json:"id"`
		NetworkAlertID               string    `json:"networkAlertID"`
		AccountID                    string    `json:"accountID"`
		NetworkID                    string    `json:"networkID"`
		OwningBankID                 string    ` json:"owningBankID"`
		OwningBankName               string    ` json:"owningBankName"`
		Time                         time.Time `json:"time"`
		Name                         string    `json:"name"`
		MuleScore                    float64   `json:"muleScore"`
		SourceTransactionValue       int       `json:"sourceTransactionValue"`
		EndpointFlag                 bool      ` json:"endpointFlag"`
		NumOutboundRelationships     int       `json:"numOutboundRelationships"`
		NumInboundRelationships      int       `json:"numInboundRelationships"`
		NumScheduledMandates         int       `json:"numScheduledMandates"`
		FirstAppearance              time.Time `json:"firstAppearance"`
		MostRecentAppearance         time.Time `json:"mostRecentAppearance"`
		FirstTransactionTime         time.Time `json:"firstTransactionTime"`
		MostRecentTransactionTime    time.Time `json:"mostRecentTransactionTime"`
		ReceivesSalary               bool      `json:"receivesSalary"`
		DwellTime                    string    `json:"dwellTime"`
		NumNetworks                  int       `json:"numNetworks"`
		NumTracedNetworks            int       `json:"numTracedNetworks"`
		Generation                   int       `json:"generation"`
		TraceType                    string    `json:"traceType"`
		TotalSuspiciousValueInbound  int       `json:"totalSuspiciousValueInbound"`
		TotalSuspiciousValueOutbound int       `json:"totalSuspiciousValueOutbound"`
		TotalValueInbound            int       `json:"totalValueInbound"`
		TotalValueOutbound           int       `json:"totalValueOutbound"`
		Generations                  []int     `json:"generations"`
		MostRecentFeedback           string    `json:"mostRecentFeedback"`
		ParentAlertID                string    `json:"parentAlertID"`
		DecisionDate                 time.Time `json:"decisionDate"`
	}
)
