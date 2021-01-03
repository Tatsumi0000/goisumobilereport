package sugoimobilereport

// https://developer.apple.com/forums/thread/24203
// type ProductTypeIdentifier string
type ProductTypeIdentifier string

const (
	// FreeOrPaidiPhoneAndiPod 1 Free or paid app iPhone and iPod touch (iOS)
	FreeOrPaidiPhoneAndiPod = ProductTypeIdentifier("1")
	// UpdateiPhoneAndiPod 7 Update iPhone and iPod touch (iOS)
	UpdateiPhoneAndiPod = ProductTypeIdentifier("7")
	// AppBundle 1-B App Bundle
	AppBundle = ProductTypeIdentifier("1-B")
	// PaidAppiPhoneAndiPod 1E Paid app Custom iPhone and iPod touch (iOS)
	PaidAppiPhoneAndiPod = ProductTypeIdentifier("1E")
	// PaidAppCustomiPad 1EP Paid app Custom iPad (iOS)
	PaidAppCustomiPad = ProductTypeIdentifier("1EP")
	// PaidAppCustomUniversal 1EU Paid app Custom universal (iOS)
	PaidAppCustomUniversal = ProductTypeIdentifier("1EU")
	// FreeOrPaidAppUniversal 1F Free or paid app Universal (iOS)
	FreeOrPaidAppUniversal = ProductTypeIdentifier("1F")
	// FreeOrPaidAppiPad 1T Free or paid app iPad (iOS)
	FreeOrPaidAppiPad = ProductTypeIdentifier("1T")
	// UpdateUniversal 7F Update Universal (iOS)
	UpdateUniversal = ProductTypeIdentifier("7F")
	// UpdateiPad 7TUpdate iPad (iOS)
	UpdateiPad = ProductTypeIdentifier("7T")
	// FreeOrPaidAppMacApp F1 Free or paid app Mac app
	FreeOrPaidAppMacApp = ProductTypeIdentifier("F1")
	// UpdateMacApp F7 Update Mac app
	UpdateMacApp = ProductTypeIdentifier("F7")
	// InAppPurchaseMacApp FI1 In-App Purchase Mac app
	InAppPurchaseMacApp = ProductTypeIdentifier("FI1")
	// InAppPurchasePurchaseiOS IA1 In-App Purchase Purchase (iOS)
	InAppPurchasePurchaseiOS = ProductTypeIdentifier("IA1")
	// InAppPurchasePurchaseMac IA1-M In-App Purchase Purchase (Mac)
	InAppPurchasePurchaseMac = ProductTypeIdentifier("IA1-M")
	// InAppPurchaseSubscriptioniOS IA9 In-App Purchase Subscription (iOS)
	InAppPurchaseSubscriptioniOS = ProductTypeIdentifier("IA9")
	// InAppPurchaseSubscriptionMac IA9-M In-App Purchase Subscription (Mac)
	InAppPurchaseSubscriptionMac = ProductTypeIdentifier("IA9-M")
	// InAppPurchaseFreeSubscriptioniOS IAC In-App Purchase Free subscription (iOS)
	InAppPurchaseFreeSubscriptioniOS = ProductTypeIdentifier("IAC")
	// InAppPurchaseFreeSubscriptionMac IAC-M In-App Purchase Free subscription (Mac)
	InAppPurchaseFreeSubscriptionMac = ProductTypeIdentifier("IAC-M")
	// InAppPurchaseAutoRenewableSubscriptioniOS IAY In-App Purchase Auto-renewable subscription (iOS)
	InAppPurchaseAutoRenewableSubscriptioniOS = ProductTypeIdentifier("IAY")
	// InAppPurchaseAutoRenewableSubscriptionMac IAY-M In-App Purchase Auto-renewable subscription (Mac)
	InAppPurchaseAutoRenewableSubscriptionMac = ProductTypeIdentifier("IAY-M")
	// RedownloadOfiPhoneOnlyOriOSAndtvOSApp 3 excluding iPad-only
	RedownloadOfiPhoneOnlyOriOSAndtvOSApp = ProductTypeIdentifier("3")
	// RedownloadOfUniversalApp 3F excluding tvOS
	RedownloadOfUniversalApp = ProductTypeIdentifier("3F")
	// RedownloadOfiPadOnlyApp 3T
	RedownloadOfiPadOnlyApp = ProductTypeIdentifier("3T")
	// RedownloadOfMacApp F3
	RedownloadOfMacApp = ProductTypeIdentifier("F3")
)
