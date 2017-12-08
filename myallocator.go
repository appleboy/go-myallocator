package ma

import (
	"crypto/tls"
	"errors"

	"github.com/astaxie/beego/httplib"
)

// MyAllocator struct
type MyAllocator struct {
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
	Debug              bool
}

// Allocation for ARIUpdateURL API
type Allocation struct {
	RoomID    string `json:"RoomId"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Units     string `json:"Units"`
	MinStay   string `json:"MinStay,omitempty"`
	MaxStay   string `json:"MaxStay,omitempty"`
	Price     string `json:"Price,omitempty"`
}

// ARIUpdateMessage for ARIUpdate Message struct
type ARIUpdateMessage struct {
	EndDate   string `json:"EndDate"`
	ID        int    `json:"id"`
	StartDate string `json:"StartDate"`
	User      string `json:"user"`
	EndDay    int    `json:"end_day"`
	End       int    `json:"end"`
	Type      string `json:"type"`
	Units     string `json:"Units"`
	MaxStay   string `json:"MaxStay"`
	StartDay  int    `json:"start_day"`
	Msg       string `json:"msg"`
	Pid       int    `json:"pid"`
	Source    string `json:"source"`
	Price     string `json:"Price"`
	Start     int    `json:"start"`
	RoomID    string `json:"RoomId"`
	MinStay   string `json:"MinStay"`
}

// ARIUpdateResponse ARI Update Response
type ARIUpdateResponse struct {
	Errors   []interface{}      `json:"Errors"`
	Messages []ARIUpdateMessage `json:"Messages"`
	Success  bool               `json:"Success"`
	Channels struct {
	} `json:"Channels"`
	UpdateID int `json:"UpdateId"`
	Options  struct {
		FIUA int `json:"FIUA"`
	} `json:"Options"`
}

// ARIRequest ARIUpdate Rule
// https://myallocator.github.io/apidocs/#api-3_API_Methods-ARIUpdate
type ARIRequest struct {
	AuthUserToken      string       `json:"Auth/UserToken"`
	AuthPropertyID     string       `json:"Auth/PropertyId"`
	AuthVendorID       string       `json:"Auth/VendorId"`
	AuthVendorPassword string       `json:"Auth/VendorPassword"`
	Channels           []string     `json:"Channels"`
	Allocations        []Allocation `json:"Allocations"`
}

// AssociateUserToPMSRequest creates a permanent link between an existing myallocator user and a PMS Vendor.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-AssociateUserToPMS
type AssociateUserToPMSRequest struct {
	AuthUserID         string `json:"Auth/UserId"`
	AuthUserPassword   string `json:"Auth/UserPassword"`
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
	PMSUserID          string `json:"PMSUserId,omitempty"`
}

// AssociateUserToPMSResponse for AssociateUserToPMS response
type AssociateUserToPMSResponse struct {
	Success       bool   `json:"Success"`
	AuthUserToken string `json:"Auth/UserToken"`
}

// New MyAllocator object
func New(vendorID, vendorPassword string) (*MyAllocator, error) {
	if vendorID == "" || vendorPassword == "" {
		return nil, errors.New("missing myallocator vendor id or password")
	}

	return &MyAllocator{
		AuthVendorID:       vendorID,
		AuthVendorPassword: vendorPassword,
	}, nil
}

// UpdateVendor for update vendor user id and password
func (m *MyAllocator) UpdateVendor(id, password string) *MyAllocator {
	m.AuthVendorID = id
	m.AuthVendorPassword = password

	return m
}

// UpdateUserToken get user token from UpdateUserToken API
func (m *MyAllocator) UpdateUserToken(req *AssociateUserToPMSRequest) (*AssociateUserToPMSResponse, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(AssociateUserToPMSResponse)

	if err := SendRequest(AssociateUserToPMSURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// ARIUpdate for an allocation sets the number of rooms or beds
func (m *MyAllocator) ARIUpdate(req *ARIRequest) (*ARIUpdateResponse, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(ARIUpdateResponse)

	if err := SendRequest(ARIUpdateURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// SendRequest for myallocator api
func SendRequest(url string, body interface{}, res interface{}, debug bool) error {
	req := httplib.Post(url).Debug(debug)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.JSONBody(body)
	return req.ToJSON(res)
}
