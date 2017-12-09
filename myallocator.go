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

// VendorSetRequest Use this call if you're using booking callbacks to set the callback URL and password.
// See https://myallocator.github.io/apidocs/#api-3_API_Methods-VendorSet
type VendorSetRequest struct {
	AuthVendorID          string `json:"Auth/VendorId"`
	AuthVendorPassword    string `json:"Auth/VendorPassword"`
	CallbackURL           string `json:"Callback/URL"`
	CallbackPassword      string `json:"Callback/Password"`
	CallbackNotifyBooking bool   `json:"Callback/NotifyBooking"`
}

// VendorSetResponse for vendor set response
type VendorSetResponse struct {
	Success               bool   `json:"Success"`
	CallbackPassword      string `json:"Callback/Password"`
	CallbackNotifyBooking bool   `json:"Callback/NotifyBooking"`
	CallbackURL           string `json:"Callback/URL"`
}

// RoomAvailabilityListReq can be used to query for all data that we hold
// for a specific property and date range.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-RoomAvailabilityList
type RoomAvailabilityListReq struct {
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
	AuthUserToken      string `json:"Auth/UserToken"`
	AuthPropertyID     string `json:"Auth/PropertyId"`
	StartDate          string `json:"StartDate"`
	EndDate            string `json:"EndDate"`
}

// RoomAvailabilityListRes for RoomAvailabilityList response
type RoomAvailabilityListRes struct {
	Success bool `json:"Success"`
	Rooms   []struct {
		RatePlanID int  `json:"RatePlanId"`
		PropertyID int  `json:"PropertyId"`
		IsPrivate  bool `json:"isPrivate"`
		RoomID     int  `json:"RoomId"`
		Dates      []struct {
			Units              int    `json:"Units"`
			Date               string `json:"Date"`
			ClosedForDeparture int    `json:"ClosedForDeparture"`
			ClosedForArrival   int    `json:"ClosedForArrival"`
			Price              string `json:"Price"`
			MinStay            int    `json:"MinStay"`
			MaxStay            int    `json:"MaxStay"`
			Closed             int    `json:"Closed"`
		} `json:"Dates"`
		RoomName string `json:"RoomName"`
	} `json:"Rooms"`
}

// VendorReq jsut for vendor id and password
type VendorReq struct {
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
}

// ChannelListRes for all channel struct
type ChannelListRes struct {
	Success  bool `json:"Success"`
	Channels struct {
		Koh struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"koh"`
	} `json:"Channels"`
}

// PropertyChannelListReq Lists the channel details associated with a property. Formerly (v1) GetProperties
// https://myallocator.github.io/apidocs/#api-3_API_Methods-PropertyChannelList
type PropertyChannelListReq struct {
	AuthUserToken      string `json:"Auth/UserToken"`
	AuthPropertyID     string `json:"Auth/PropertyId"`
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
}

// PropertyChannelListRes for PropertyChannelList response
// missing some field
type PropertyChannelListRes struct {
	Properties []struct {
		Channels struct {
			Ote struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ote"`
		} `json:"Channels"`
		PropertyID int `json:"PropertyId"`
	} `json:"Properties"`
	Success bool `json:"Success"`
}

// AssociatePropertyToPMSReq Creates a permanent link between an existing myallocator property and a PMS Vendor.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-AssociatePropertyToPMS
type AssociatePropertyToPMSReq struct {
	AuthUserID         string `json:"Auth/UserId"`
	AuthUserPassword   string `json:"Auth/UserPassword"`
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
	AuthPropertyID     string `json:"Auth/PropertyId"`
	PMSPropertyID      string `json:"PMSPropertyId,omitempty"`
}

// AssociatePropertyToPMSRes for AssociatePropertyToPMS response
type AssociatePropertyToPMSRes struct {
	Success           bool   `json:"Success"`
	AuthPropertyToken string `json:"Auth/PropertyToken"`
	Warning           string `json:"warning"`
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

// VendorSet Use this call if you're using booking callbacks to set the callback URL and password.
func (m *MyAllocator) VendorSet(req *VendorSetRequest) (*VendorSetResponse, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(VendorSetResponse)

	if err := SendRequest(VendorSetURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// RoomAvailabilityList can be used to query for all data that we hold for a specific property
// and date range. The date range can only be 31 days as the maximum. You can query multiple times
// if you need a longer date range.
func (m *MyAllocator) RoomAvailabilityList(req *RoomAvailabilityListReq) (*RoomAvailabilityListRes, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(RoomAvailabilityListRes)

	if err := SendRequest(RoomAvailabilityListURL, req, res, m.Debug); err != nil {
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

// ChannelList dumps information about all known channels, their status, and supported fields.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-ChannelList
func (m *MyAllocator) ChannelList() (*ChannelListRes, error) {
	req := &VendorReq{
		AuthVendorID:       m.AuthVendorID,
		AuthVendorPassword: m.AuthVendorPassword,
	}

	res := new(ChannelListRes)

	if err := SendRequest(ChannelListURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// PropertyChannelList Lists the channel details associated with a property. Formerly (v1) GetProperties
func (m *MyAllocator) PropertyChannelList(req *PropertyChannelListReq) (*PropertyChannelListRes, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(PropertyChannelListRes)

	if err := SendRequest(PropertyChannelListURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// AssociatePropertyToPMS creates a permanent link between an existing myallocator property and a PMS Vendor.
func (m *MyAllocator) AssociatePropertyToPMS(req *AssociatePropertyToPMSReq) (*AssociatePropertyToPMSRes, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(AssociatePropertyToPMSRes)

	if err := SendRequest(AssociatePropertyToPMSURL, req, res, m.Debug); err != nil {
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
