package ma

import (
	"crypto/tls"
	"errors"
	"time"

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

// ErrResponse error format
//
//	{
//	  "Errors": [
//	      {
//	          "ErrorTicket": "D13BAC04-1234-5678-804B-F5140B37646E",
//	          "ErrorMsg": "No such booking id",
//	          "ErrorId": "30"
//	      }
//	  ]
//	}
type ErrResponse struct {
	ErrorTicket string `json:"ErrorTicket"`
	ErrorMsg    string `json:"ErrorMsg"`
	ErrorID     int    `json:"ErrorId"`
}

// ARIUpdateResponse ARI Update Response
type ARIUpdateResponse struct {
	Errors   []ErrResponse      `json:"Errors"`
	Messages []ARIUpdateMessage `json:"Messages"`
	Success  bool               `json:"Success"`
	Channels struct{}           `json:"Channels"`
	UpdateID interface{}        `json:"UpdateId"`
	Options  struct {
		FIUA int `json:"FIUA"`
	} `json:"Options"`
}

// ARIRequestOptions ....
// If QueryForStatus is true then the ARIUpdateResponse will contain the additional parameter UpdateId,
// which is needed for ARIUpdateStatus. It is extremely likely that blocking updates will be phased out
// in favour of QueryForStatus so we highly recommend that you send only non-blocking ARIUpdate requests
// and poll myallocator using ARIUpdateStatus.
type ARIRequestOptions struct {
	QueryForStatus     bool `json:"QueryForStatus"`
	IgnoreInvalidRooms bool `json:"IgnoreInvalidRooms"`
}

// ARIRequest ARIUpdate Rule
// https://myallocator.github.io/apidocs/#api-3_API_Methods-ARIUpdate
type ARIRequest struct {
	AuthUserToken      string            `json:"Auth/UserToken"`
	AuthPropertyID     string            `json:"Auth/PropertyId"`
	AuthVendorID       string            `json:"Auth/VendorId"`
	AuthVendorPassword string            `json:"Auth/VendorPassword"`
	Channels           []string          `json:"Channels"`
	Allocations        []Allocation      `json:"Allocations"`
	Options            ARIRequestOptions `json:"Options"`
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

// BookingCancelReq allows you to cancel a booking both in myallocator and the channel.
// Only very few channels support this, check the 'api_cancellation' flag on the ChannelList
// call to see which channels support this.
type BookingCancelReq struct {
	AuthUserToken      string `json:"Auth/UserToken"`
	AuthVendorID       string `json:"Auth/VendorId"`
	AuthVendorPassword string `json:"Auth/VendorPassword"`
	AuthPropertyID     string `json:"Auth/PropertyId"`
	MyAllocatorID      string `json:"MyAllocatorId"`
	CancellationReason string `json:"CancellationReason"`
}

// BookingCancelResp for booking cacnel response.
type BookingCancelResp struct {
	StartDate string `json:"start_date"`
	Success   bool   `json:"Success"`
	QueryType string `json:"query_type"`
	Bookings  []struct {
		MyallocatorID           string `json:"MyallocatorId"`
		OrderSource             string `json:"OrderSource"`
		MyallocatorCreationDate string `json:"MyallocatorCreationDate"`
		Version                 int    `json:"Version"`
		IsModification          bool   `json:"IsModification"`
		MyallocatorCreationTime string `json:"MyallocatorCreationTime"`
		TotalCurrency           string `json:"TotalCurrency"`
		OrderTime               string `json:"OrderTime"`
		Channel                 string `json:"Channel"`
		IsCancellation          bool   `json:"IsCancellation"`
		Customers               []struct {
			CustomerIP          string `json:"CustomerIP"`
			CustomerNote        string `json:"CustomerNote"`
			CustomerLName       string `json:"CustomerLName"`
			CustomerPhone       string `json:"CustomerPhone"`
			CustomerNationality string `json:"CustomerNationality"`
			CustomerFName       string `json:"CustomerFName"`
			CustomerEmail       string `json:"CustomerEmail"`
		} `json:"Customers"`
		Rooms []struct {
			Units                  string        `json:"Units"`
			EndDate                string        `json:"EndDate"`
			RoomTypeIds            []int         `json:"RoomTypeIds"`
			MyallocatorRateplanIds []interface{} `json:"MyallocatorRateplanIds"`
			RateID                 string        `json:"RateId"`
			StartDate              string        `json:"StartDate"`
			Price                  int           `json:"Price"`
			RateDesc               string        `json:"RateDesc"`
			ChannelRoomType        string        `json:"ChannelRoomType"`
			Currency               string        `json:"Currency"`
			DayRates               []struct {
				Date        string      `json:"Date"`
				Currency    string      `json:"Currency"`
				Rate        int         `json:"Rate"`
				Description interface{} `json:"Description"`
			} `json:"DayRates"`
		} `json:"Rooms"`
		MarkedAsRead                bool   `json:"MarkedAsRead"`
		MyallocatorModificationDate string `json:"MyallocatorModificationDate"`
		OrderDate                   string `json:"OrderDate"`
		CancellationReason          string `json:"CancellationReason"`
		StartDate                   string `json:"StartDate"`
		Acknowledged                bool   `json:"Acknowledged"`
		TotalPrice                  string `json:"TotalPrice"`
		EndDate                     string `json:"EndDate"`
		PropertyID                  int    `json:"PropertyId"`
		MyallocatorModificationTime string `json:"MyallocatorModificationTime"`
		OrderID                     string `json:"OrderId"`
	} `json:"Bookings"`
	EndDate string `json:"end_date"`
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

// VendorReq just for vendor id and password
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

// BookingListReq allows you to query for bookings made to a specific property by booking date,
// modification date or arrival date.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-BookingList
type BookingListReq struct {
	AuthUserToken         string `json:"Auth/UserToken"`
	AuthVendorID          string `json:"Auth/VendorId"`
	AuthVendorPassword    string `json:"Auth/VendorPassword"`
	AuthPropertyID        string `json:"Auth/PropertyId"`
	ModificationStartDate string `json:"ModificationStartDate"`
	ModificationEndDate   string `json:"ModificationEndDate"`
}

// BookingListRes for BookingList response
type BookingListRes struct {
	QueryType string `json:"query_type"`
	EndDate   string `json:"end_date"`
	Success   bool   `json:"Success"`
	Bookings  []struct {
		TotalCurrency string `json:"TotalCurrency"`
		OrderID       string `json:"OrderId"`
		Rooms         []struct {
			RoomTypeIds []int `json:"RoomTypeIds"`
			Adults      int   `json:"Adults"`
			DayRates    []struct {
				RateID      string `json:"RateId"`
				Description string `json:"Description"`
				Date        string `json:"Date"`
				Rate        string `json:"Rate"`
			} `json:"DayRates"`
			RateID                 string        `json:"RateId"`
			Policy                 string        `json:"Policy"`
			RoomDesc               string        `json:"RoomDesc"`
			ChannelRoomType        string        `json:"ChannelRoomType"`
			RateDesc               string        `json:"RateDesc"`
			MyallocatorRateplanIds []interface{} `json:"MyallocatorRateplanIds"`
			StartDate              string        `json:"StartDate"`
			Occupancy              int           `json:"Occupancy"`
			EndDate                string        `json:"EndDate"`
			Children               int           `json:"Children"`
			Units                  string        `json:"Units"`
			Breakfast              bool          `json:"Breakfast"`
		} `json:"Rooms"`
		PropertyID                  int    `json:"PropertyId"`
		StartDate                   string `json:"StartDate"`
		EndDate                     string `json:"EndDate"`
		MyallocatorCreationDate     string `json:"MyallocatorCreationDate"`
		PaymentCollect              string `json:"PaymentCollect,omitempty"`
		IsModification              bool   `json:"IsModification"`
		MarkedAsRead                bool   `json:"MarkedAsRead"`
		MyallocatorModificationTime string `json:"MyallocatorModificationTime"`
		OrderDate                   string `json:"OrderDate"`
		OrderTime                   string `json:"OrderTime"`
		IsCancellation              bool   `json:"IsCancellation"`
		Acknowledged                bool   `json:"Acknowledged"`
		Version                     int    `json:"Version"`
		Channel                     string `json:"Channel"`
		MyallocatorCreationTime     string `json:"MyallocatorCreationTime"`
		MyallocatorID               string `json:"MyallocatorId"`
		MyallocatorModificationDate string `json:"MyallocatorModificationDate"`
		TotalPrice                  string `json:"TotalPrice"`
		Customers                   []struct {
			CustomerLName       string `json:"CustomerLName"`
			CustomerPhoneMobile string `json:"CustomerPhoneMobile"`
			CustomerNote        string `json:"CustomerNote"`
			CustomerArrivalTime string `json:"CustomerArrivalTime"`
			CustomerEmail       string `json:"CustomerEmail"`
			CustomerFName       string `json:"CustomerFName"`
		} `json:"Customers"`
		OrderSource string `json:"OrderSource,omitempty"`
	} `json:"Bookings"`
	StartDate string `json:"start_date"`
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

// BookingList  query for bookings made to a specific property
// by booking date, modification date or arrival date.
func (m *MyAllocator) BookingList(req *BookingListReq) (*BookingListRes, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(BookingListRes)

	if err := SendRequest(BookingListURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// BookingCancel allows you to cancel a booking both in myallocator
// and the channel. Only very few channels support this.
func (m *MyAllocator) BookingCancel(req *BookingCancelReq) (*BookingCancelResp, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(BookingCancelResp)

	if err := SendRequest(BookingCancelURL, req, res, m.Debug); err != nil {
		return nil, err
	}

	return res, nil
}

// SetDebug to set debug mode
func (m *MyAllocator) SetDebug(enable bool) *MyAllocator {
	m.Debug = enable

	return m
}

// SendRequest for myallocator api
func SendRequest(url string, body interface{}, res interface{}, debug bool) error {
	req := httplib.Post(url).Debug(debug)
	//nolint:gosec
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(20*time.Second, 20*time.Second)
	if _, err := req.JSONBody(body); err != nil {
		return err
	}
	return req.ToJSON(res)
}
