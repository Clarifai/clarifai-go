package clarifai

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
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
		MaxVideoDuration  int     `json:"max_video_duration"`
		MaxVideoBatchSize int     `json:"max_video_batch_size"`
		MinVideoSize      int     `json:"min_video_size"`
		MinImageSize      int     `json:"min_image_size"`
		MaxBatchSize      int     `json:"max_batch_size"`
		APIVersion        float32 `json:"api_version"`
	}
}

// TagResp represents the expected JSON response from /tag/
type TagResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Meta          struct {
		Tag struct {
			Timestamp json.Number `json:"timestamp"`
			Model     string      `json:"model"`
			Confid    string      `json:"config"`
		}
	}
	Results []TagResult
}

// TagResult represents the expected data for a single tag result
type TagResult struct {
	DocID         uint64 `json:"docid"`
	URL           string `json:"url"`
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	LocalID       string `json:"local_id"`
	Result        struct {
		Tag struct {
			Classes []string  `json:"classes"`
			CatIDs  []string  `json:"catids"`
			Probs   []float32 `json:"probs"`
		}
	}
	DocIDString string `json:"docid_str"`
}

// FeedbackForm is used to send feedback back to Clarifai
type FeedbackForm struct {
	DocIDs           []string
	URLs             []string
	AddTags          []string
	RemoveTags       []string
	DissimilarDocIDs []string
	SimilarDocIDs    []string
	SearchClick      []string
}

// FeedbackResp is the expected response from /feedback/
type FeedbackResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
}

// Info will return the current status info for the given client
func (client *Client) Info() (*InfoResp, error) {
	res, err := client.commonHTTPRequest(nil, "info", "GET", false)

	if err != nil {
		return nil, err
	}

	info := new(InfoResp)
	err = json.Unmarshal(res, info)

	return info, err
}

// Tag allows the client to request tag data on a single, or multiple photos
func (client *Client) Tag(urls, localIDs []string) (*TagResp, error) {
	if urls == nil {
		return nil, errors.New("Requires at least one url")
	}

	form := url.Values{}
	for _, url := range urls {
		form.Add("url", url)
	}
	if localIDs != nil {
		for _, localID := range localIDs {
			form.Add("local_id", localID)
		}
	}

	res, err := client.commonHTTPRequest(form, "tag", "POST", false)

	if err != nil {
		return nil, err
	}

	tagres := new(TagResp)
	err = json.Unmarshal(res, tagres)

	return tagres, err
}

// Feedback allows the user to provide contextual feedback to Clarifai in order to improve their results
func (client *Client) Feedback(params FeedbackForm) (*FeedbackResp, error) {
	if params.DocIDs == nil && params.URLs == nil {
		return nil, errors.New("Requires at least one docid or url")
	}

	if params.DocIDs != nil && params.URLs != nil {
		return nil, errors.New("Request must provide exactly one of the following fields: {'DocIDs', 'URLs'}")
	}

	form := url.Values{}

	if params.DocIDs != nil {
		form.Add("docids", strings.Join(params.DocIDs, ","))
	} else {
		form.Add("url", strings.Join(params.URLs, ","))
	}

	if params.AddTags != nil {
		form.Add("add_tags", strings.Join(params.AddTags, ","))
	}

	if params.RemoveTags != nil {
		form.Add("remove_tags", strings.Join(params.RemoveTags, ","))
	}

	if params.DissimilarDocIDs != nil {
		form.Add("dissimilar_docids", strings.Join(params.DissimilarDocIDs, ","))
	}

	if params.SimilarDocIDs != nil {
		form.Add("similar_docids", strings.Join(params.SimilarDocIDs, ","))
	}

	if params.SearchClick != nil {
		form.Add("search_click", strings.Join(params.SearchClick, ","))
	}

	res, err := client.commonHTTPRequest(form, "feedback", "POST", false)

	feedbackres := new(FeedbackResp)
	err = json.Unmarshal(res, feedbackres)

	return feedbackres, err

}
