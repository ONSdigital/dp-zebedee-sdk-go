package zebedee

import (
	"encoding/json"
	"fmt"
)

const (
	//Manual publish - publish is trigger by a user action.
	Manual PublishType = iota
	//Scheduled publish - automatically published by Zebedee at the configured date/time.
	Scheduled

	//CollectionDateFMT is the date/time format expected by the CMS.
	CollectionDateFMT = "2006-01-02T15:04:05.000Z"
)

//PublishType enum defining the different types of collection publishes
type PublishType int

//User defines the CMS user structure
type User struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Inactive          bool   `json:"inactive"`
	LastAdmin         string `json:"lastAdmin"`
	TemporaryPassword bool   `json:"temporaryPassword"`
}

// Credentials is the model representing the user login details
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Session is the model of a CMS user session.
type Session struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// Permissions is the model representing user's CMS permissions
type Permissions struct {
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	Editor bool   `json:"editor"`
}

// collectionBase
type collectionBase struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        PublishType `json:"type"`
	PublishDate string      `json:"publishDate"`
	ReleaseURI  string      `json:"releaseUri,omitempty"`
	Teams       []string    `json:"teams"`
}

//CollectionDescription object representation of a CMS collection metadata
type CollectionDescription struct {
	collectionBase
	Encrypted             bool            `json:"isEncrypted"`
	PublishComplete       bool            `json:"publishComplete"`
	ApprovalStatus        string          `json:"approvalStatus"`
	InProgressUris        []string        `json:"inProgressUris"`
	CompleteUris          []string        `json:"completeUris"`
	ReviewedUris          []string        `json:"reviewedUris"`
	DatasetVersions       []string        `json:"datasetVersions"`
	Datasets              []string        `json:"datasets"`
	EventsByUri           interface{}     `json:"eventsByUri"`
	PendingDeletes        []string        `json:"pendingDeletes"`
	PublishResults        []PublishResult `json:"publishResults"`
	TimeseriesImportFiles []string        `json:"timeseriesImportFiles"`
}

type PublishResult struct {
	Message      string                `json:"message"`
	Error        bool                  `json:"error"`
	Transactions PublishingTransaction `json:"transaction"`
}

type PublishingTransaction struct {
	ID        string    `json:"id"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	UriInfos  []URIInfo `json:"uriInfos"`
	Errors    []string  `json:"errors"`
}

type URIInfo struct {
	Action                 string `json:"action"`
	URI                    string `json:"uri"`
	Start                  string `json:"start"`
	End                    string `json:"end"`
	Duration               int    `json:"duration"`
	VerificationStatus     string `json:"verificationStatus"`
	VerificationEnd        string `json:"verificationEnd"`
	VerificationRetryCount int    `json:"verificationRetryCount"`
	VerifyMessage          int    `json:"verifyMessage"`
	SHA                    int    `json:"sha"`
	Size                   int    `json:"size"`
	Error                  string `json:"error"`
}

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

type TeamsList struct {
	Teams []Team `json:"teams"`
}

func (pt PublishType) Name() string {
	switch pt {
	case 0:
		return "manual"
	default:
		return "scheduled"
	}
}

func (pt PublishType) ValueOf(val string) PublishType {
	switch val {
	case Manual.Name():
		return Manual
	default:
		return Scheduled
	}
}

func (pt PublishType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.Name())
}

func (pt *PublishType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch raw {
	case Manual.Name():
		*pt = Manual
	case Scheduled.Name():
		*pt = Scheduled
	default:
		return fmt.Errorf("JSON unmarshing error invalid PublishType value %q", raw)
	}

	return nil
}
