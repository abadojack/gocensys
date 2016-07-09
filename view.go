package gocensys

// ViewDocument used to fetch full document once its ID is known (e.g., IP address or domain).
func (c CensysAPI) ViewDocument(index, id string) (interface{}, error) {
	var data interface{} //TODO : replace the interface with a struct
	err := c.apiGet("/view/"+index+"/"+id, nil, &data)

	return data, err
}
