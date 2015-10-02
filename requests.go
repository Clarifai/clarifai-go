package clarifai

import (
	"encoding/json"
	"errors"
	"net/url"
)

// InfoResp represents the expected JSON response from /info/
type InfoResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Results       struct {
		MaxImageSize      int     `json:"max_image_size"`
		DefaultLanguage   string  `json:"default_language"`
		MaxVideoSize      int     `json:"max_video_size"`
		MaxImageBytes     int     `json:"max_image_bytes"`
		DefaultModel      string  `json:"default_model"`
		MaxVideoBytes     int     `json:"max_video_bytes"`
		maxVideoDuration  int     `json:"max_video_duration"`
		MaxVideoBatchSize int     `json:"max_video_batch_size"`
		MinVideoSize      int     `json:"min_video_size"`
		MinImageSize      int     `json:"min_image_size"`
		MaxBatchSize      int     `json:"max_batch_size"`
		APIVersion        float32 `json:"api_version"`
	}
}

type TagResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Meta          struct {
		Tag struct {
			Timestamp float64 `json:"timestamp"`
			Model     string  `json:"model"`
			Confid    string  `json:"config"`
		}
	}
	Results []TagResults
}

type TagResults struct {
	DocID         json.Number `json:"docid"`
	URL           string      `json:"url"`
	StatusCode    string      `json:"status_code"`
	StatusMessage string      `json:"status_msg"`
	LocalID       string      `json:"local_id"`
	Result        struct {
		Tag struct {
			Classes []string  `json:"classes"`
			CatIDs  []string  `json:"catids"`
			Probs   []float32 `json:"probs"`
		}
	}
	DocIDString string `json:"docid_str"`
}

// Info will return the current status info for the given client
func Info(client Client) (*InfoResp, error) {
	res, err := client.commonHTTPRequest(nil, "info", "GET")

	if err != nil {
		return nil, err
	}

	info := new(InfoResp)
	err = json.Unmarshal(res, info)

	return info, err
}

func Tag(client Client, urls, local_ids []string) (*TagResp, error) {
	if urls == nil {
		return nil, errors.New("Requires at least one url")
	}

	form := url.Values{}
	for _, url := range urls {
		form.Add("url", url)
	}
	if local_ids != nil {
		for _, localid := range local_ids {
			form.Add("local_id", localid)
		}
	}

	res, err := client.commonHTTPRequest(form, "tag", "POST")

	if err != nil {
		return nil, err
	}

	tagres := new(TagResp)
	err = json.Unmarshal(res, tagres)

	return tagres, err
}
