package shows

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func fetchEpisodateDataByShowId(episodateShowId uint) (error, *TvShow) {
	client := http.Client{
		Timeout: time.Duration(30*time.Second),
	}

	var err error
	var req *http.Request
	var resp *http.Response

	url := "https://www.episodate.com/api/show-details"
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err, nil
	}

	urlQueries := req.URL.Query()
	urlQueries.Set("q", fmt.Sprintf("%d", episodateShowId))
	req.URL.RawQuery = urlQueries.Encode()

	resp, err = client.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	respTemplate := map[string]TvShow{}
	err = json.NewDecoder(resp.Body).Decode(&respTemplate)
	if err != nil {
		return err, nil
	}


	show, exists := respTemplate["tvShow"]
	if !exists {
		return errors.New("response was not successful, no tvShow data found"), nil
	}
	return nil, &show
}
