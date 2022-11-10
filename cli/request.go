package cli

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewGetRequest[T any](ctx context.Context, url, contentType string) (resp T, err error) {
	cli := DefaultClient()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return resp, err
	}

	if len(contentType) == 0 {
		contentType = "application/json"
	}

	res, err := cli.Do(req)
	if err != nil {
		return resp, err
	}

	res.Close = true

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
