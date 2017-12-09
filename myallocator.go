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
		Hola struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
		} `json:"hola"`
		Enbc struct {
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"enbc"`
		Net struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"net"`
		Hcu struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Validate   int    `json:"validate"`
				Type       string `json:"type"`
				Hidden     int    `json:"hidden"`
				Activation int    `json:"activation"`
				Label      string `json:"label"`
			} `json:"PropertyInput"`
			Hidden int    `json:"hidden"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"hcu"`
		Ct struct {
			PropertyInput []struct {
				Type     string `json:"type"`
				Label    string `json:"label"`
				Length   int    `json:"length"`
				Stored   int    `json:"stored"`
				Altlabel string `json:"altlabel"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"ct"`
		Kch struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
		} `json:"kch"`
		Htrv struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Rs   int    `json:"rs"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Validate   int    `json:"validate,omitempty"`
				Label      string `json:"label"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"htrv"`
		Vjo struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"vjo"`
		Luna struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"luna"`
		Lvur struct {
			Hidden        int `json:"hidden"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"lvur"`
		Gr struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"gr"`
		Outp struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Hidden        int `json:"hidden"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"outp"`
		Ctr struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Label      string `json:"label"`
				Type       string `json:"type"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Length     int    `json:"length"`
				Stored     int    `json:"stored"`
				Hint       string `json:"hint,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"ctr"`
		Ysh struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Label    string `json:"label,omitempty"`
				Validate int    `json:"validate,omitempty"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"ysh"`
		Saas struct {
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"saas"`
		Tavr struct {
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Hidden int    `json:"hidden"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"tavr"`
		Aveo struct {
			RoomInput []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Rs   int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Type     string `json:"type,omitempty"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"aveo"`
		Hom struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"hom"`
		Bnw struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Type            string `json:"type"`
				AdditionalSetup int    `json:"additional_setup"`
				Label           string `json:"label"`
				Stored          int    `json:"stored"`
				Name            string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string        `json:"api_cancellation"`
			RoomInput       []interface{} `json:"RoomInput"`
		} `json:"bnw"`
		Ini struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Type     string `json:"type"`
				Validate int    `json:"validate"`
				Label    string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"ini"`
		Eba struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Hidden int    `json:"hidden"`
				Type   string `json:"type"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"eba"`
		Hotu struct {
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type       string `json:"type"`
				Label      string `json:"label"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Length     int    `json:"length"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"hotu"`
		Bdrm struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"bdrm"`
		Tmd struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"tmd"`
		Hrs struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Type     string `json:"type"`
				Validate int    `json:"validate"`
				Label    string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length,omitempty"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"hrs"`
		Inst struct {
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
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"inst"`
		Btob struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
		} `json:"btob"`
		Hd struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"hd"`
		Loc struct {
			APICancellation string        `json:"api_cancellation"`
			PropertyInput   []interface{} `json:"PropertyInput"`
			Hidden          string        `json:"hidden"`
			RoomInput       []interface{} `json:"RoomInput"`
			IsLive          string        `json:"isLive"`
			Name            string        `json:"name"`
			IsLockedLogin   string        `json:"isLockedLogin"`
			SignupLink      string        `json:"signupLink"`
			IsCalendarBased string        `json:"isCalendarBased"`
			IsBeta          string        `json:"isBeta"`
		} `json:"loc"`
		Liko struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"liko"`
		Ho struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"ho"`
		Asyo struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"asyo"`
		Ya struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"ya"`
		Pre struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"pre"`
		Ago struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length,omitempty"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Validate   int    `json:"validate"`
				Type       string `json:"type"`
				Activation int    `json:"activation"`
				Label      string `json:"label"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"ago"`
		Trpz struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"trpz"`
		Lota struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"lota"`
		Wim struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
				Hidden int    `json:"hidden"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"wim"`
		Yoho struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"yoho"`
		Toqt struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"toqt"`
		Gta struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type       string `json:"type"`
				Label      string `json:"label,omitempty"`
				Activation int    `json:"activation,omitempty"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Hidden     int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"gta"`
		Chot struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"chot"`
		Eb struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
				Length   int    `json:"length,omitempty"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
		} `json:"eb"`
		Whoo struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type       string `json:"type"`
				Hidden     int    `json:"hidden"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid"`
				Stored     int    `json:"stored"`
				Length     int    `json:"length"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"whoo"`
		Ote struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				Label      string `json:"label"`
				Activation int    `json:"activation"`
				Validate   int    `json:"validate"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"ote"`
		Hopa struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"hopa"`
		Cb struct {
			IsLive        string `json:"isLive"`
			Name          string `json:"name"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Length   int    `json:"length"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"cb"`
		Bool struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"bool"`
		Nom struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"nom"`
		Orb struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Label    string `json:"label,omitempty"`
				Validate int    `json:"validate,omitempty"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"orb"`
		Myb struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"myb"`
		Bbnl struct {
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label    string `json:"label"`
				Validate int    `json:"validate"`
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"bbnl"`
		Ktel struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored     int    `json:"stored"`
				Length     int    `json:"length"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Name       string `json:"name"`
				Type       string `json:"type"`
				Label      string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"ktel"`
		Coco struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"coco"`
		Bp struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Stored int    `json:"stored,omitempty"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Validate   int    `json:"validate,omitempty"`
				Type       string `json:"type"`
				Label      string `json:"label"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Activation int    `json:"activation,omitempty"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"bp"`
		Bon struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
		} `json:"bon"`
		Flh struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"flh"`
		Bkb struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"bkb"`
		Bh struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"bh"`
		Kqbe struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"kqbe"`
		Tlok struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored     int    `json:"stored"`
				ResolvePid int    `json:"resolve_pid"`
				Name       string `json:"name"`
				Type       string `json:"type"`
				Label      string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type string `json:"type"`
				Rs   int    `json:"rs"`
				Name string `json:"name"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"tlok"`
		Art struct {
			Name          string `json:"name"`
			IsLive        string `json:"isLive"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"art"`
		Rsrv struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"rsrv"`
		Sbkr struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"sbkr"`
		Bnb struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Validate int    `json:"validate"`
				Label    string `json:"label"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
		} `json:"bnb"`
		Sawa struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"sawa"`
		Ohma struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"ohma"`
		Xmlt struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"xmlt"`
		Hotq struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"hotq"`
		Bole struct {
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"bole"`
		Odi struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label      string `json:"label"`
				Validate   int    `json:"validate"`
				Type       string `json:"type"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid"`
				Stored     int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"odi"`
		Han struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"han"`
		Hb struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Length   int    `json:"length"`
				Validate int    `json:"validate,omitempty"`
				Type     string `json:"type"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"hb"`
		Iwb struct {
			Name          string `json:"name"`
			IsLive        string `json:"isLive"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Label    string `json:"label"`
				Validate int    `json:"validate,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"iwb"`
		Hi struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Type       string `json:"type"`
				Validate   int    `json:"validate,omitempty"`
				Label      string `json:"label"`
				Activation int    `json:"activation"`
				Length     int    `json:"length"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"hi"`
		Sole struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"sole"`
		Mmt struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Validate   int    `json:"validate,omitempty"`
				Label      string `json:"label"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length,omitempty"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"mmt"`
		Ht struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Validate int    `json:"validate,omitempty"`
				Type     string `json:"type"`
				Label    string `json:"label"`
				Length   int    `json:"length"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"ht"`
		Enz struct {
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"enz"`
		Adve struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"adve"`
		Fk struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Hidden int    `json:"hidden"`
				Type   string `json:"type"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"fk"`
		Advt struct {
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"advt"`
		Wix struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"wix"`
		Vrbo struct {
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Hidden int    `json:"hidden"`
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"vrbo"`
		Zldv struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"zldv"`
		Vcay struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"vcay"`
		Adv struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Name     string `json:"name"`
				Length   int    `json:"length"`
				Stored   int    `json:"stored"`
				Validate int    `json:"validate,omitempty"`
				Type     string `json:"type"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"adv"`
		Etb struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Type       string `json:"type"`
				Validate   int    `json:"validate"`
				Label      string `json:"label"`
				Activation int    `json:"activation,omitempty"`
				Stored     int    `json:"stored"`
				Length     int    `json:"length"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length,omitempty"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"etb"`
		Wpk struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"wpk"`
		Pzg struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"pzg"`
		Hw2 struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Activation int    `json:"activation"`
				Label      string `json:"label"`
				Altlabel   string `json:"altlabel"`
				Hidden     int    `json:"hidden"`
				Validate   int    `json:"validate"`
				Type       string `json:"type"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid"`
				Stored     int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"hw2"`
		Clt struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Activation int    `json:"activation"`
				Label      string `json:"label"`
				Type       string `json:"type"`
				Validate   int    `json:"validate,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Rs   int    `json:"rs"`
			} `json:"RoomInput"`
		} `json:"clt"`
		Prsc struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"prsc"`
		Skd struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"skd"`
		Hw struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Length   int    `json:"length"`
				Label    string `json:"label,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"hw"`
		Max struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				Activation int    `json:"activation,omitempty"`
				Label      string `json:"label"`
				Validate   int    `json:"validate,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"max"`
		Trav struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"trav"`
		Mrbb struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"mrbb"`
		Lmg struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate int    `json:"validate"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"lmg"`
		Oa struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type   string `json:"type,omitempty"`
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"oa"`
		Goi struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"goi"`
		Air struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Type   string `json:"type"`
				Hidden int    `json:"hidden"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"air"`
		Sur struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"sur"`
		Hou struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Type   string `json:"type"`
				Hidden int    `json:"hidden"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"hou"`
		Hde struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Type   string `json:"type"`
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"hde"`
		Buen struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"buen"`
		Bbit struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Length int    `json:"length"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Type   string `json:"type"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"bbit"`
		Vive struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"vive"`
		Deco struct {
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Label  string `json:"label"`
				Type   string `json:"type"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"deco"`
		Rtel struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Length     int    `json:"length"`
				Stored     int    `json:"stored"`
				Label      string `json:"label"`
				Type       string `json:"type"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
		} `json:"rtel"`
		Hort struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type       string `json:"type"`
				Hidden     int    `json:"hidden,omitempty"`
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				Length     int    `json:"length,omitempty"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
		} `json:"hort"`
		Mcs struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"mcs"`
		Ptch struct {
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Hidden int    `json:"hidden"`
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs,omitempty"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"ptch"`
		Akt struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"akt"`
		Esc struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate   int    `json:"validate"`
				Type       string `json:"type"`
				Label      string `json:"label"`
				Activation int    `json:"activation"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"esc"`
		Bb struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label    string `json:"label,omitempty"`
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Length   int    `json:"length,omitempty"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
		} `json:"bb"`
		Bv struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored   int    `json:"stored"`
				Length   int    `json:"length"`
				Name     string `json:"name"`
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"bv"`
		Boo struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Length         int    `json:"length"`
				ActivationHelp int    `json:"activation_help"`
				Label          string `json:"label"`
				Activation     int    `json:"activation"`
				ResolvePid     int    `json:"resolve_pid"`
				Stored         int    `json:"stored"`
				Name           string `json:"name"`
				Type           string `json:"type"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"boo"`
		Hipc struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Type   string `json:"type"`
				Hidden int    `json:"hidden"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"hipc"`
		G2I struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"g2i"`
		Adsh struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"adsh"`
		Buun struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"buun"`
		Mal struct {
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"mal"`
		Vaca struct {
			IsLive        string `json:"isLive"`
			Name          string `json:"name"`
			Hidden        int    `json:"hidden"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"vaca"`
		H2H struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"h2h"`
		Trab struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"trab"`
		Wanu struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"wanu"`
		Exp struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length,omitempty"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label      string `json:"label"`
				Validate   int    `json:"validate,omitempty"`
				Type       string `json:"type"`
				Name       string `json:"name"`
				Length     int    `json:"length,omitempty"`
				Stored     int    `json:"stored"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"exp"`
		Hoga struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				Length     int    `json:"length"`
				ResolvePid int    `json:"resolve_pid"`
				Hidden     int    `json:"hidden"`
				Type       string `json:"type"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"hoga"`
		Tpvl struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
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
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"tpvl"`
		Hb2 struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate int    `json:"validate,omitempty"`
				Type     string `json:"type"`
				Label    string `json:"label"`
				Length   int    `json:"length,omitempty"`
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"hb2"`
		Cs struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"cs"`
		Ibc struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"ibc"`
		Evry struct {
			RoomInput []struct {
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
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"evry"`
		Hand struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"hand"`
		Ddhm struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"ddhm"`
		Rfc struct {
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
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"rfc"`
		Ta struct {
			IsBeta          string        `json:"isBeta"`
			IsCalendarBased string        `json:"isCalendarBased"`
			SignupLink      string        `json:"signupLink"`
			IsLockedLogin   string        `json:"isLockedLogin"`
			Name            string        `json:"name"`
			IsLive          string        `json:"isLive"`
			RoomInput       []interface{} `json:"RoomInput"`
			PropertyInput   []struct {
				Stored          int    `json:"stored"`
				Name            string `json:"name"`
				Type            string `json:"type"`
				Label           string `json:"label"`
				Activation      int    `json:"activation,omitempty"`
				AdditionalSetup int    `json:"additional_setup,omitempty"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
		} `json:"ta"`
		Hmi struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"hmi"`
		Bnb2 struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"bnb2"`
		Mega struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name   string `json:"name"`
			IsLive string `json:"isLive"`
		} `json:"mega"`
		Apal struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Rs   int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label    string `json:"label,omitempty"`
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Name     string `json:"name"`
				Stored   int    `json:"stored"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"apal"`
		Hhb struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"hhb"`
		Lmin struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Length     int    `json:"length"`
				Activation int    `json:"activation"`
				Label      string `json:"label"`
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				ResolvePid int    `json:"resolve_pid"`
				Type       string `json:"type"`
				Validate   int    `json:"validate"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"lmin"`
		Tri struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			Hidden          int    `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"tri"`
		Kni struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"kni"`
		Rep struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate   int    `json:"validate,omitempty"`
				Label      string `json:"label"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
		} `json:"rep"`
		Tom struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label      string `json:"label"`
				Activation int    `json:"activation,omitempty"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Validate   int    `json:"validate,omitempty"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
		} `json:"tom"`
		Rrep struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    int `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"rrep"`
		All struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"all"`
		Stya struct {
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"stya"`
		Vhh struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"vhh"`
		Hr struct {
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"hr"`
		Hc struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Validate   int    `json:"validate,omitempty"`
				Type       string `json:"type"`
				Label      string `json:"label"`
				Activation int    `json:"activation"`
				Stored     int    `json:"stored"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
				Name       string `json:"name"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"hc"`
		Cwd struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"cwd"`
		Snrg struct {
			IsLive    string `json:"isLive"`
			Name      string `json:"name"`
			RoomInput []struct {
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
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"snrg"`
		R2N struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"r2n"`
		Iper struct {
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			IsLive string `json:"isLive"`
			Name   string `json:"name"`
		} `json:"iper"`
		Bud struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          string `json:"hidden"`
			PropertyInput   []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
		} `json:"bud"`
		Hcu2 struct {
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"hcu2"`
		Bd struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Stored   int    `json:"stored"`
				Length   int    `json:"length"`
				Name     string `json:"name"`
				Validate int    `json:"validate,omitempty"`
				Type     string `json:"type"`
				Label    string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"bd"`
		Rec struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Label      string `json:"label"`
				Activation int    `json:"activation"`
				Type       string `json:"type,omitempty"`
				Validate   int    `json:"validate,omitempty"`
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length,omitempty"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"rec"`
		NineFl struct {
			IsLive        string `json:"isLive"`
			Name          string `json:"name"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Name   string `json:"name"`
				Stored int    `json:"stored"`
				Length int    `json:"length"`
				Type   string `json:"type"`
				Hidden int    `json:"hidden"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"9fl"`
		Pand struct {
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          string `json:"hidden"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
		} `json:"pand"`
		Lr struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Stored   int    `json:"stored"`
				Name     string `json:"name"`
				Type     string `json:"type,omitempty"`
				Validate int    `json:"validate,omitempty"`
				Label    string `json:"label"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"lr"`
		Gom struct {
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Name     string `json:"name"`
				Length   int    `json:"length"`
				Stored   int    `json:"stored"`
				Label    string `json:"label"`
				Type     string `json:"type"`
				Validate int    `json:"validate,omitempty"`
				Hidden   int    `json:"hidden,omitempty"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"gom"`
		Ost struct {
			Name          string `json:"name"`
			IsLive        string `json:"isLive"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				ResolvePid int    `json:"resolve_pid"`
				Stored     int    `json:"stored"`
				Name       string `json:"name"`
				Validate   int    `json:"validate"`
				Activation int    `json:"activation"`
				Label      string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length,omitempty"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"ost"`
		Lsbv struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			Hidden          int    `json:"hidden"`
			PropertyInput   []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
		} `json:"lsbv"`
		Tob struct {
			RoomInput []struct {
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Type   string `json:"type"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
		} `json:"tob"`
		Fam struct {
			Name          string `json:"name"`
			IsLive        string `json:"isLive"`
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Stored int    `json:"stored"`
				Name   string `json:"name"`
				Label  string `json:"label"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
		} `json:"fam"`
		Grup struct {
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			Hidden          int    `json:"hidden"`
		} `json:"grup"`
		Loop struct {
			Name      string `json:"name"`
			IsLive    string `json:"isLive"`
			RoomInput []struct {
				Type   string `json:"type"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
				Length int    `json:"length"`
			} `json:"RoomInput"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Stored int    `json:"stored"`
				Name   string `json:"name"`
			} `json:"PropertyInput"`
			Hidden          string `json:"hidden"`
			APICancellation string `json:"api_cancellation"`
			IsBeta          string `json:"isBeta"`
			IsCalendarBased string `json:"isCalendarBased"`
			SignupLink      string `json:"signupLink"`
			IsLockedLogin   string `json:"isLockedLogin"`
		} `json:"loop"`
		Bal struct {
			Hidden        string `json:"hidden"`
			PropertyInput []struct {
				Label  string `json:"label"`
				Name   string `json:"name"`
				Stored int    `json:"stored"`
			} `json:"PropertyInput"`
			APICancellation string `json:"api_cancellation"`
			RoomInput       []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Name   string `json:"name"`
				Rs     int    `json:"rs"`
			} `json:"RoomInput"`
			Name            string `json:"name"`
			IsLive          string `json:"isLive"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
		} `json:"bal"`
		Hbe struct {
			IsCalendarBased string `json:"isCalendarBased"`
			IsBeta          string `json:"isBeta"`
			IsLockedLogin   string `json:"isLockedLogin"`
			SignupLink      string `json:"signupLink"`
			IsLive          string `json:"isLive"`
			Name            string `json:"name"`
			APICancellation string `json:"api_cancellation"`
			PropertyInput   []struct {
				Label      string `json:"label,omitempty"`
				Activation int    `json:"activation,omitempty"`
				Altlabel   string `json:"altlabel,omitempty"`
				Length     int    `json:"length,omitempty"`
				Type       string `json:"type"`
				Validate   int    `json:"validate,omitempty"`
				Name       string `json:"name"`
				Stored     int    `json:"stored"`
				Hidden     int    `json:"hidden,omitempty"`
				ResolvePid int    `json:"resolve_pid,omitempty"`
			} `json:"PropertyInput"`
			Hidden    string `json:"hidden"`
			RoomInput []struct {
				Type   string `json:"type"`
				Length int    `json:"length"`
				Rs     int    `json:"rs"`
				Name   string `json:"name"`
			} `json:"RoomInput"`
		} `json:"hbe"`
	} `json:"Channels"`
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

// ChannelList dumps information about all known channels, their status, and supported fields.
// https://myallocator.github.io/apidocs/#api-3_API_Methods-ChannelList
func (m *MyAllocator) ChannelList(req *VendorReq) (*ChannelListRes, error) {
	req.AuthVendorID = m.AuthVendorID
	req.AuthVendorPassword = m.AuthVendorPassword

	res := new(ChannelListRes)

	if err := SendRequest(ChannelListURL, req, res, m.Debug); err != nil {
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
