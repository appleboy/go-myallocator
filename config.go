package ma

const (
	// ARIUpdateURL for an allocation sets the number of rooms or beds API
	ARIUpdateURL = "https://api.myallocator.com/pms/v201804/json/ARIUpdate"
	// AssociateUserToPMSURL for Creates a permanent link between an existing myallocator user and a PMS Vendor
	AssociateUserToPMSURL = "https://api.myallocator.com/pms/v201804/json/AssociateUserToPMS"
	// VendorSetURL for if you're using booking callbacks to set the callback URL and password.
	VendorSetURL = "https://api.myallocator.com/pms/v201804/json/VendorSet"
	// RoomAvailabilityListURL is RoomAvailabilityList API URL
	RoomAvailabilityListURL = "https://api.myallocator.com/pms/v201804/json/RoomAvailabilityList"
	// ChannelListURL is ChannelList API URL
	ChannelListURL = "https://api.myallocator.com/pms/v201804/json/ChannelList"
	// PropertyChannelListURL is PropertyChannelList API URL
	PropertyChannelListURL = "https://api.myallocator.com/pms/v201908/json/PropertyChannelList"
	// BookingListURL is BookingList API URL
	BookingListURL = "https://api.myallocator.com/pms/v201804/json/BookingList"
	// BookingCancelURL allows you to cancel a booking both in myallocator and the channel.
	BookingCancelURL = "https://api.myallocator.com/pms/v201804/json/BookingCancel"
)
