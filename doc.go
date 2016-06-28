//Package gocensys provides structs and functions for accessing version 1
//of the Censys API.
//
//Successful API queries return native Go structs that can be used immediately,
//with no need for type assertions.
//
//Authentication
//
//If you already have the uid and secret for your user, creating the client is simple:
//
//  api := gocensys.NewCensysAPI("your-uid", "your-secret", nil)
//
//
//Queries
//
//Executing queries on an authenticated CensysAPI struct is simple.
//
//  series, _ := api.GetSeries()
//
//Endpoints
//
//gocensys implements most of the endpoints defined in the Censys API documentation: https://www.censys.io/api/v1.
//
//
//More detailed information about the behavior of each particular endpoint can be found at the official Censys API documentation.
package gocensys
