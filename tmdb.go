package tmdb

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const ApiBaseUrl = "https://api.themoviedb.org/3"
const ApiFindUrl = ApiBaseUrl + "/find/"
const ApiTvUrl = ApiBaseUrl + "/tv/"

type Provider string

const ProviderIMDB Provider = "imdb_id"

func FindByExternalId(client *http.Client, apiKey string, provider Provider, id string) (ExternalFindResults, error) {
	url := ApiFindUrl + id + "?api_key=" + apiKey + "&language=en-US&external_source=" + string(provider)
	res, err := client.Get(url)
	if err != nil {
		return ExternalFindResults{}, err
	}
	if res.StatusCode != http.StatusOK {
		return ExternalFindResults{}, errors.New("invalid response code")
	}
	var r ExternalFindResults
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return ExternalFindResults{}, err
	}
	return r, nil
}

func ShowById(client *http.Client, apiKey string, id int) (ShowInfo, error) {
	url := ApiTvUrl + strconv.Itoa(id) + "?api_key=" + apiKey + "&language=en-US"
	res, err := client.Get(url)
	if err != nil {
		return ShowInfo{}, err
	}
	var r ShowInfo
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return ShowInfo{}, err
	}
	return r, nil
}

func EpisodesForShowById(client *http.Client, apiKey string, id int, seasonNumber int) ([]EpisodeInfo, error) {
	url := ApiTvUrl + strconv.Itoa(id) + "/season/" + strconv.Itoa(seasonNumber) + "?api_key=" + apiKey + "&language=en-US"
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	var r episodeInfoContainer
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}
	return r.Episodes, nil
}
