/*
 * ImageCashLetter API
 *
 * Moov ImageCashLetter () implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BundleControl struct {
	// BundleControl ID
	Id string `json:"id,omitempty"`
	// BundleItemsCount identifies the total number of items within the bundle.
	BundleItemsCount int32 `json:"bundleItemsCount"`
	// BundleTotalAmount identifies the total amount of item amounts within the bundle.
	BundleTotalAmount int32 `json:"bundleTotalAmount"`
	// MICRValidTotalAmount identifies the total amount of all CheckDetail Records within the bundle which contains 1 in the MICRValidIndicator.
	MicrValidTotalAmount int32 `json:"micrValidTotalAmount,omitempty"`
	// BundleImagesCount identifies the total number of Image ViewDetail Records  within the bundle.
	BundleImagesCount int32 `json:"bundleImagesCount,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
	// CreditTotalIndicator is a code that indicates whether Credits Items are included in the totals. If so they  they will be included in this record’s: Items Within Bundle Count, Bundle Total Amount and, Images Within Bundle Count * ` ` - No Credit Items * `0` - Credit Items are not included in totals * `1` - Credit Items are included in totals 
	CreditTotalIndicator string `json:"creditTotalIndicator,omitempty"`
}
