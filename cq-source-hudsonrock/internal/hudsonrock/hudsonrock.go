package HudsonRock

const defaultURL = "https://cavalier.hudsonrock.com/api/json/v2/search-by-domain"

type HudsonRock struct {
	Date_uploaded string `json:"date_uploaded"`
	Stealer_family string `json:"stealer_family"`
	Antiviruses string `json:"antiviruses"`
	Date_compromised string `json:"date_compromised"`
	Credentials string `json:"credentials"`
	Stealer string `json:"stealer"`
	EmployeeAt string `json:"employeeAt"`
	ClientAt string `json:"clientAt"`
	Ip string `json:"ip"`
	Malware_path string `json:"malware_path"`
}

type Client struct {
	baseURL string
	client *http.Client
}

func newClient () (*Client, error) {
	return &Client{
		baseURL: defaultURL,
		client: http.DefaultClient,
	}, nil
}

func (c *Client) GetHudsonRock(num int) (*HudsonRock, error) {
	resp, err := c.client.Get(url: c.baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	hudsonrock := &HudsonRock{}
	if err := json.NewDecoder(resp.Body).Decode(hudsonrock); err != nil {
		return nil,err
	}
	return hudsonrock, nil
}