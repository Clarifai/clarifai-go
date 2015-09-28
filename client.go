package clarifai

type Client interface {
	clientID() string
	clientSecret() string
	accessToken() string
	apiHost() string
	apiPort() string
}

type ClarifaiClient struct {
	clientID     string
	clientSecret string
	accessToken  string
	apiHost      string
	apiPort      string
}

func initClient(clientID, clientSecret, accessToken string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, accessToken, "api.clarifai.com", "443"}
}
