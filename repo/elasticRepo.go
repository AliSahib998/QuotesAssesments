package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AliSahib998/QuotesAssesments/config"
	"github.com/AliSahib998/QuotesAssesments/model"
	"io/ioutil"
	"strings"
)

func SaveDocument(index string, data *bytes.Buffer, id string) error {
	var esClient = config.ElasticDb
	_, err := esClient.Index(index, data, esClient.Index.WithDocumentID(id))
	return err
}

func UpdateDocument(index string, data *bytes.Buffer, id string) error {
	var esClient = config.ElasticDb
	_, err := esClient.Index(index, data, esClient.Index.WithDocumentID(id), esClient.Index.WithRefresh("true"))
	return err
}

func SearchUser(query *model.SearchQuery, index string) ([]model.UserRegistration, error) {
	var esClient = config.ElasticDb

	resp, err := esClient.Search(
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(strings.NewReader(buildSearchQuery(query))),
	)

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return extractHitsForUser(string(body))
	}

	return nil, err
}

func SearchQuote(query *model.SearchQuery, index string) ([]model.QuoteDocument, error) {
	var esClient = config.ElasticDb

	resp, err := esClient.Search(
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(strings.NewReader(buildSearchQuery(query))),
	)

	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return extractHitsForQuote(string(body))
	}

	return nil, err
}

func extractHitsForUser(resp string) ([]model.UserRegistration, error) {
	type Response struct {
		Hits struct {
			Hits []struct {
				Source model.UserRegistration `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	var response Response
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return nil, err
	}

	hits := response.Hits.Hits

	var users []model.UserRegistration
	for _, hit := range hits {
		users = append(users, hit.Source)
	}

	return users, nil
}

func extractHitsForQuote(resp string) ([]model.QuoteDocument, error) {
	type Response struct {
		Hits struct {
			Hits []struct {
				Source model.QuoteDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	var response Response
	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return nil, err
	}

	hits := response.Hits.Hits

	var quotes []model.QuoteDocument
	for _, hit := range hits {
		quotes = append(quotes, hit.Source)
	}

	return quotes, nil
}

func buildSearchQuery(query *model.SearchQuery) string {
	var parts []string

	if len(query.QueryString) > 0 && len(query.SearchField) > 0 && !query.IsFullSearch {
		matchQuery := fmt.Sprintf(`"query": {
            "match": {
                "%s": "%s"
            }
        }`, query.SearchField, query.QueryString)
		parts = append(parts, matchQuery)
	} else if len(query.QueryString) > 0 && len(query.SearchField) > 0 && query.IsFullSearch {
		matchQuery := fmt.Sprintf(`"query": {
        "bool": {
            "should": [
                {
                    "wildcard": {
                        "%s": "*%s*"
                    }
                }
            ]
        }
    }`, query.SearchField, query.QueryString)
		parts = append(parts, matchQuery)
	} else {
		parts = append(parts, `"query": { "match_all": {} }`)
	}

	if len(query.SortOrder) > 0 && len(query.SortField) > 0 {
		sortCriteria := fmt.Sprintf(`"sort": [
                { "%s": { "order": "%s" } }
            ]`, query.SortField, query.SortOrder)
		parts = append(parts, sortCriteria)
	}

	return "{" + strings.Join(parts, ",") + "}"
}
