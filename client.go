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
	tokenRetries integer
}

// Initialize a new client object
func initClient(clientID, clientSecret, accessToken string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, accessToken, "api.clarifai.com", "443"}
}

// clienID getter
func (self *ClarifaiClient) getClientID() string {
	return self.clientID
}

// clientSecret getter
func (self *ClarifaiClient) getClientSecret() string {
	return self.clientSecret
}

// accessToken getter
func (self *ClarifaiClient) getAccessToken() string {
	return self.accessToken
}
