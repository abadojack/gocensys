## gocensys

[![Build Status](https://travis-ci.org/abadojack/gocensys.svg?branch=master)](https://travis-ci.org/abadojack/gocensys)  [![GoDoc](https://godoc.org/github.com/abadojack/gocensys?status.png)](http://godoc.org/github.com/abadojack/gocensys)

gocensys is a simple Go package for accessing the [Censys API](https://www.censys.io/api).

Successful API queries return native Go structs that can be used immediately,
with no need for type assertions.

gocensys implements endpoints defined in the documentation: https://www.censys.io/api.
More detailed information about the behavior of each particular endpoint can be found at the official documentation.


## Installation

	$ go get -u github.com/abadojack/gocensys

## Usage

```Go
	import "github.com/abadojack/gocensys"
```

## Authentication

HTTP basic auth is required using the API ID and secret that are shown under [My Account](https://www.censys.io/account).

```Go
  api := gocensys.NewCensysAPI("your-uid", "your-secret", nil)
```

## Queries

Executing queries on an authenticated CensysAPI struct is simple.
```Go
queryJob, err := api.StartQueryJob("select p110.pop3.ssl_2.certificate.parsed.issuer.province from ipv4.20160708 limit 1000;")
if err != nil {
	log.Fatal(err)
}

fmt.Println(queryJob.JobID)
```

## Licence
gocensys is free software licenced under the GNU LGPL licence. Details are provided in the LICENCE file.
