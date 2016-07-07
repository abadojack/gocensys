## gocensys

[![Build Status](https://travis-ci.org/abadojack/gocensys.svg?branch=master)](https://travis-ci.org/abadojack/gocensys)  [![GoDoc](https://godoc.org/github.com/abadojack/gocensys?status.png)](http://godoc.org/github.com/abadojack/gocensys)

gocensys is a simple Go package for accessing the [Censys API](https://www.censys.io/api).

Successful API queries return native Go structs that can be used immediately,
with no need for type assertions.

gocensys implements endpoints defined in the documentation: https://www.censys.io/api/v1.
More detailed information about the behavior of each particular endpoint can be found at the official documentation.


## Milestone - based on endpoints implemented.

[data](https://censys.io/api/v1/docs/data) - Done

[view](https://censys.io/api/v1/docs/view) - Done

[report](https://censys.io/api/v1/docs/report) - Done

[query](https://censys.io/api/v1/docs/query) - Doing

[export](https://censys.io/api/v1/docs/export) - Done

[search](https://censys.io/api/v1/docs/search) - Done



## Installation

	$ go get -u github.com/abadojack/gocensys

## Usage

```Go
	import "github.com/abadojack/gocensys"
```

## Authentication

If you already have the uid and secret for your user, creating the client is simple:

```Go
  api := gocensys.NewCensysAPI("your-uid", "your-secret", nil)
```

## Queries

Executing queries on an authenticated CensysAPI struct is simple.
```Go
  series, _ := api.GetSeries()
```

## Licence
gocensys is free software licenced under the GNU LGPL licence. Details are provided in the LICENCE file.
