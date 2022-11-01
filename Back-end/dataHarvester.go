package server



func main (){

	type Artist struct {
		Id           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
		Dates        string   `json:"dates"`
	}
	
	//Gets data from URL and returns the body of data as []byte and an error
	func harvest(url string) []byte, error {
		res, err := http.Get(url) //get contents of API
		if err != nil {
			return []byte{}, err
		}
		defer res.Body.Close()
	
		response, err := ioutil.ReadAll(res.Body) //read body of contents
		if err != nil {
			return []byte{}, err
		}
		return response, err
	}
