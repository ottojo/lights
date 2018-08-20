package lights

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ESPURL = "http://ledesp.lan"

func SetNeoPixel(i int, c Color) error {
	request, err := http.NewRequest("GET", ESPURL, nil)
	if err != nil {
		return err
	}

	q := request.URL.Query()
	q.Add("pixel", strconv.Itoa(i))
	q.Add("r", fmt.Sprintf("%.0f", c.R*255))
	q.Add("g", fmt.Sprintf("%.0f", c.G*255))
	q.Add("b", fmt.Sprintf("%.0f", c.B*255))
	request.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}
