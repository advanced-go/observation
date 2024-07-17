package routing1

type Matcher struct {
	Path      string
	Template  string
	Authority string
	Route     string
}

type Client struct {
	Match   Matcher
	Timeout int
}

type Cloud struct {
	RateLimiting bool
}

type Config struct {
	Client Client
	Cloud  Cloud
}

type Case struct {
	Desc   string
	Client Client
	Cloud  Cloud
}
