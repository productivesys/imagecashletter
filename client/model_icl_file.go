/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IclFile struct for IclFile
type IclFile struct {
	// File ID
	ID          string         `json:"ID,omitempty"`
	FileHeader  IclFileHeader  `json:"fileHeader,omitempty"`
	CashLetters []CashLetter   `json:"cashLetters,omitempty"`
	Bundles     []Bundle       `json:"bundles,omitempty"`
	FileControl IclFileControl `json:"fileControl,omitempty"`
}