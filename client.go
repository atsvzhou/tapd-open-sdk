package tapd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

type Client struct {
	apiUser       string
	apiPassWord   string
	authorization string
	client        *http.Client
	baseURL       *url.URL
}

func NewClient(user, pwd string) *Client {
	return &Client{
		apiUser:     user,
		apiPassWord: pwd,
		client:      &http.Client{},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "api.tapd.cn",
		},
	}
}

func (c *Client) Get(path string, params map[string]string, v interface{}) (*http.Response, error) {
	c.baseURL.Path = path
	c.baseURL.RawQuery = c.ParamsToSortQuery(params)
	req := c.NewReq(http.MethodGet, c.baseURL.String(), nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("the tapd client received an unhealthy status code: %d, message: %s", resp.StatusCode, resp.Status)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func (c *Client) Post(path string, body map[string]string, v interface{}) (*http.Response, error) {
	c.baseURL.Path = path
	p := url.Values{}
	for k, v := range body {
		p.Set(k, v)
	}
	req := c.NewReq(http.MethodPost, c.baseURL.String(), strings.NewReader(p.Encode()))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("the tapd client received an unhealthy status code: %d, message: %s", resp.StatusCode, resp.Status)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func (c *Client) ParamsToSortQuery(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var resultList []string
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := params[key]
		resultList = append(resultList, fmt.Sprintf("%s=%s", key, value))
	}

	result := strings.Join(resultList, "&")
	return result
}

func (c *Client) HandleParams(params interface{}) (map[string]string, error) {
	var p = map[string]string{}
	paramsByte, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var mi map[string]interface{}
	err = json.Unmarshal(paramsByte, &mi)
	if err != nil {
		return nil, err
	}

	for k, v := range mi {
		vt := reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Map, reflect.Array, reflect.Slice:
			value, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			p[k] = string(value)
		case reflect.Float64:
			value, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			p[k] = string(value)
		default:
			p[k] = v.(string)
		}
	}

	return p, err
}

func (c *Client) SetAuthorization() {
	str := c.apiUser + ":" + c.apiPassWord
	base64.StdEncoding.EncodeToString([]byte(str))
	c.authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(str))
}

func (c *Client) NewReq(method, url string, body io.Reader) *http.Request {
	if c.authorization == "" {
		c.SetAuthorization()
	}
	r, _ := http.NewRequest(method, url, body)
	r.Header.Add("Authorization", c.authorization)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r
}
