package awardwallet

//Error Respusta de error
type Error struct {
	Error string `json:"error"`
}

//Response Respuesta exitosa
type Response struct {
	Account       Account       `json:"account"`
	ConnectedUser ConnectedUser `json:"connectedUser"`
}

type Account struct {
	AccountID          int         `json:"accountId,omitempty"`
	Code               string      `json:"code,omitempty"`
	DisplayName        string      `json:"displayName,omitempty"`
	Kind               string      `json:"kind,omitempty"`
	Login              string      `json:"login,omitempty"`
	AutoLoginURL       string      `json:"autologinUrl,omitempty"`
	UpdateURL          string      `json:"updateUrl,omitempty"`
	EditURL            string      `json:"editUrl,omitempty"`
	History            []History   `json:"history,omitempty"`
	Balance            string      `json:"balance"`
	BalanceRaw         int         `json:"balanceRaw"`
	Owner              string      `json:"owner,omitempty"`
	ErrorCode          int         `json:"errorCode,omitempty"`
	LastDetectedChange string      `json:"lastDetectedChange,omitempty"`
	ExpirationDate     string      `json:"expirationDate,omitempty"`
	LastRetrieveDate   string      `json:"lastRetrieveDate,omitempty"`
	LastChangeDate     string      `json:"lastChangeDate,omitempty"`
	Properties         []Propertie `json:"properties,omitempty"`
}

type Propertie struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Kind  int    `json:"kind,omitempty"`
	Rank  int    `json:"rank,omitempty"`
}

type History struct {
	Fields []Field `json:"fields,omitempty"`
}

type Field struct {
	Name  string `json:"name,omitempty"`
	Code  string `json:"code,omitempty"`
	Value Value  `json:"value,omitempty"`
}

type ConnectedUser struct {
	UserID                  int    `json:"userId,omitempty"`
	Fullname                string `json:"fullName,omitempty"`
	Status                  string `json:"status,omitempty"`
	Username                string `json:"userName,omitempty"`
	Email                   string `json:"email,omitempty"`
	ForwardingEmail         string `json:"forwardingEmail,omitempty"`
	AccessLevel             string `json:"accessLevel,omitempty"`
	ConnectionType          string `json:"connectionType,omitempty"`
	AccountsAccessLevel     string `json:"accountsAccessLevel,omitempty"`
	AccountsSharedByDefault bool   `json:"accountsSharedByDefault,omitempty"`
	EditConnectionURL       string `json:"editConnectionUrl,omitempty"`
	AccountListURL          string `json:"accountListUrl,omitempty"`
	TimelineURL             string `json:"timelineUrl,omitempty"`
	BookingRequestsURL      string `json:"bookingRequestsUrl,omitempty"`
}

type Value struct {
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}
