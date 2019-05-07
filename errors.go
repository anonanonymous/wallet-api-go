package walletapi

// ERRORS - wallet-api errors
var ERRORS = map[int]string{
	400: "A parse error occurred, or an error occurred processing your request.",
	401: "API key is missing or invalid.",
	403: "This operation requires a wallet to be open, and one has not been opened.",
	404: "The transaction hash was not found.",
	500: "An exception was thrown whilst processing the request.",
}
