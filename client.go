package clarifai

type Client interface {
	clientID() string
	clientSecret() string
	accessToken() string
	apiHost() string
	apiPort() string
}

const ROOT = "api.clarifai.com"
const TAG_PATH = "/v1/token"
const FEEDBACK_PATH = "/v1/feedback"

type ClarifaiClient struct {
	clientID     string
	clientSecret string
	accessToken  string
}

// Initialize a new client object
func InitClient(clientID, clientSecret string) *ClarifaiClient {
	return &ClarifaiClient{clientID, clientSecret, "unasigned"}
}

// clientID getter
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
