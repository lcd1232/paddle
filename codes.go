package paddle

type ErrorCode int

const (
	ErrorCodeLicenseNotFound            ErrorCode = 100
	ErrorCodeBadMethodCall              ErrorCode = 101
	ErrorCodeBadAPIKey                  ErrorCode = 102
	ErrorCodeTimestampInvalid           ErrorCode = 103
	ErrorCodeLicenseAlreadyUtilized     ErrorCode = 104
	ErrorCodeLicenseIsNotActive         ErrorCode = 105
	ErrorCodeActivationNotFound         ErrorCode = 106
	ErrorCodePermissionError            ErrorCode = 107
	ErrorCodeProductNotFound            ErrorCode = 108
	ErrorCodeCurrencyInvalid            ErrorCode = 109
	ErrorCodePurchaseNotFound           ErrorCode = 110
	ErrorCodeAuthenticationTokenInvalid ErrorCode = 111
	ErrorCodeVerificationTokenInvalid   ErrorCode = 112
	ErrorCodePaddingInvalid             ErrorCode = 113
	ErrorCodeAffiliateInvalid           ErrorCode = 114
	ErrorCodeAffiliateCommissionInvalid ErrorCode = 115
	ErrorCodeArgumentsMissing           ErrorCode = 116
	ErrorCodeExpirationTimeInvalid      ErrorCode = 117
	ErrorCodePriceIsTooLow              ErrorCode = 118
	ErrorCodeSubscriptionNotFound       ErrorCode = 119
	ErrorCodeInternalError              ErrorCode = 120
	ErrorCodePaymentNotFound            ErrorCode = 121
	ErrorCodeDateInvalid                ErrorCode = 122
	ErrorCodeModifierNotFound           ErrorCode = 123
	ErrorCodeModifierAlreadyPaid        ErrorCode = 124
	ErrorCodeNoMainCurrencyPrice        ErrorCode = 125
	ErrorCodeEmailInvalid               ErrorCode = 126
	ErrorCodeCouponTypeInvalid          ErrorCode = 127
	ErrorCodePercentageInvalid          ErrorCode = 128
	ErrorCodeAmountInvalid              ErrorCode = 129
	//130	The allowed uses must be a number.
	//131	The given coupon code is invalid. The code must have at least 5 characters.
	//132	The given coupon code has already been used for the product.
	//133	The given coupon expiration date is invalid. The expected date format is “Y-m-d”.
	//134	The given coupon currency is invalid. The currency must be one of the currencies of your product.
	//135	Unable to find requested coupon
	//136	Allowed uses cannot be less than times used.
	//137	The allowed uses must be a number greater than or equal to 0.
	//138	The expires at value must be either not provided or a future date in the format of Y-m-d.
	//139	The given prices format is not valid. The prices must have the format of [‘currency:amount’, ‘currency:amount’, …].
	//140	The given currency code is unknown to our checkout system.
	//141	Either a product ID or a plan ID should be given, not both.
	//142	The given recurring prices format is not valid. The recurring prices must have the format of [‘currency:amount’, ‘currency:amount’, …].
	//143	Recurring price is too low
	//144	Affiliate split sum must total less than 100%
	//145	Recurring affiliate split must either be not set, or set to an integer equal to or greater than 1.
	//146	The current invoice of this subscription is currently being processed, and cannot be updated at this time
	//147	We were unable to complete the resubscription because we could not charge the customer for the resubscription
	//148	The resubscription requires immediate billing so we cannot complete your request
	//149	The plan interval is invalid
	//150	Initial price is too low
	//151	The subscription cannot be updated at this time. Please try later.
	//152	Plan changes can not be made whilst the customer is in their trial period.
	//153	The trial length must be a positive integer.
	//154	Unable to find requested order
	//155	The given amount is not valid.
	//156	The Order cannot be refunded.
	//157	An unknown coupon error has occurred.
	//158	The coupon currency must match your balance currency
	//159	The parameters combination is incorrect.
	//160	Invalid recurring option.
	//161	The minimum threshold must be numeric and higher than 0.01.
	//162	The group has to be a string with at least 1 character and no more than 50 characters.
	//163	The number of coupons is invalid.
	//164	The can_multiple_in_same_checkout parameter has to be a boolean value.
	//165	The given coupon target is not recognised. The only valid types are product and checkout.
	//166	The description has to be a string.
	//167	You cannot set the amount for Flat and Percentage coupons at the same time.
	//168	The product type must be a subscription plan.
	//170	license_code is not set.
	//171	download_url is not set or invalid.
	//172	The transaction can no longer be refunded.
	//173	The subscription does not allow quantities to be set.
	//174	Cannot move to this plan as it doesn’t support the subscription currency.
	//175	Invalid country code.
	//176	This order already has a license code attached to it
	//177	product_name is not set.
	//180	download_url is invalid.
	//181	charge_name is too long or invalid. Length limit: 50
	//182	The given subscription ID is invalid. IDs must be numeric.
	//183	Charges cannot be made with a negative amount
	//184	Access Denied.
	//185	Subscription billing cycles exceeded. Subscription expired.
	//186	Amount is less than allowed minimum transaction amount.
	//187	Subscription has been deleted.
	//188	Transaction failed.
	//189	Rate limit reached for this type of request. Please try again later.
	//190	No valid payment method found
	//191	You cannot pass an offset without defining a limit
	//192	No other modifications are allowed when pausing/unpausing a subscription
	//193	You can not pause/unpause this subscription
	//194	Changes can not be made whilst the subscription is paused.
	//195	The given cancellation reason is not valid.
	//196	Subscription is in a past_due state and the plan can therefore not be changed.
	//197	The subscription of the customer is not active.
	//198	The selected new plan is invalid.
	//199	No other modifications except pausing and changing the passthrough are allowed on a past due subscription.
)
