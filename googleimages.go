package googleimages

import (
	"encoding/json"
	"io"
	"net/url"
	"os"

	behavior "github.com/ikeikeikeike/gopkg/net/http"
)

const EndPoint = "http://ajax.googleapis.com/ajax/services/search/images"

func tee(r io.Reader, debug bool) io.Reader {
	if !debug {
		return r
	}
	return io.TeeReader(r, os.Stdout)
}

type Client struct {
	*behavior.UserBehavior
	Debug bool
}

func NewClient() *Client {
	return &Client{
		UserBehavior: behavior.NewUserBehavior(),
		Debug:        false,
	}
}

type Result struct {
	GsearchResultClass  string `json:"GsearchResultClass"`
	Width               string `json:"width"`
	Height              string `json:"height"`
	ImageId             string `json:"imageId"`
	TbWidth             string `json:"tbWidth"`
	TbHeight            string `json:"tbHeight"`
	UnescapedUrl        string `json:"unescapedUrl"`
	Url                 string `json:"url"`
	VisibleUrl          string `json:"visibleUrl"`
	Title               string `json:"title"`
	TitleNoFormatting   string `json:"titleNoFormatting"`
	OriginalContextUrl  string `json:"originalContextUrl"`
	Content             string `json:"content"`
	ContentNoFormatting string `json:"contentNoFormatting"`
	TbUrl               string `json:"tbUrl"`
}

type Page struct {
	Start string `json:"start`
	Label int    `json:"label"`
}

type Cursor struct {
	ResultCount          string  `json:"resultCount"`
	Pages                []*Page `json:"pages"`
	EstimatedResultCount string  `json:"estimatedResultCount"`
	CurrentPageIndex     int     `json:"currentPageIndex"`
	MoreResultsUrl       string  `json:"moreResultsUrl"`
	SearchResultTime     string  `json:"searchResultTime"`
}

type ResponseData struct {
	Cursor  *Cursor   `json:"cursor"`
	Results []*Result `json:"results"`
}

type Googleimages struct {
	ResponseData    *ResponseData
	ResponseStatus  int    `json:"responseStatus"`
	ResponseDetails string `json:"responseDetails"`
}

func (c *Client) Fetch(query string) (*Googleimages, error) {
	res, err := c.Get(EndPoint + "?hl=ja&v=1.0&q=" + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	r := new(Googleimages)
	err = json.NewDecoder(tee(res.Body, c.Debug)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Set Cursor object, And return operate iterate object with google image api.
func (c *Client) Do(query string) (err error) { return }
