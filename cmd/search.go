package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	DockerHubApiV2Endpoint  = "https://hub.docker.com/v2"
	DockerHubApiV1Endpoint  = "https://hub.docker.com/api/content/v1"
	SearchImageEndpoint     = DockerHubApiV1Endpoint + "/products/search"
	SearchVersion           = "v3"
	ImageTagsEndpointFormat = DockerHubApiV2Endpoint + "/repositories/library/%s/tags"
)

type (
	searchArgs []string

	searchImageResponse struct {
		Count     int
		Summaries []struct {
			Name string
			Slug string
		}
	}

	searchImageTagsResponse struct {
		Count     int
		ImageName string
		Tags      searchImageTags `json:"results"`
	}

	searchImageTags []searchImageTag

	searchImageTag struct {
		Name string
	}

	searchClient struct {
		client http.Client
	}
)

func (s searchArgs) ImageName() string {
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

func (s searchArgs) Tag() string {
	if len(s) < 2 {
		return ""
	}
	return s[1]
}

func (s searchImageTags) String() string {
	r := ""
	for i, t := range s {
		if i > 0 {
			r += " "
		}
		r += t.Name
	}
	return r
}

func buildSearchCmd(rootCmd *cobra.Command) {
	var c = &cobra.Command{
		Use:   "search [image tag]",
		Short: "Search images",
		Long:  `Search images by input word. By Default, images is filtered by is-official.`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearch(searchArgs(args))
		},
	}

	rootCmd.AddCommand(c)
}

func runSearch(args searchArgs) (err error) {
	c := searchClient{http.Client{Timeout: time.Second * 3}}
	images, err := c.searchImages(args.ImageName())
	if err != nil {
		return
	}

	g := errgroup.Group{}
	ch := make(chan searchImageTagsResponse, len(images.Summaries))
	for _, image := range images.Summaries {
		imageName := image.Slug
		g.Go(func() error {
			tags, errTags := c.searchTags(imageName, args.Tag())
			if errTags != nil {
				return err
			}
			ch <- tags
			return nil
		})
	}

	if err = g.Wait(); err != nil {
		return
	}
	close(ch)

	// TODO: printer
	fmt.Println("NAME    TAGS")
	for tagsResp := range ch {
		fmt.Printf("%s    %s \n", tagsResp.ImageName, tagsResp.Tags.String())
	}

	return
}

func (c searchClient) searchImages(imageName string) (searchImageResponse, error) {
	name, _ := validate(imageName)

	reqUrl, err := url.Parse(SearchImageEndpoint)
	if err != nil {
		return searchImageResponse{}, err
	}
	query := reqUrl.Query()
	query.Set("q", name)
	query.Set("image_filter", "official")
	query.Set("type", "image")
	query.Set("page_size", "25")
	reqUrl.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return searchImageResponse{}, err
	}
	req.Header.Set("Search-Version", SearchVersion)
	resp, err := c.client.Do(req)
	if err != nil {
		return searchImageResponse{}, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchImageResponse{}, err
	}
	var r searchImageResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return searchImageResponse{}, err
	}

	return r, nil
}

func validate(word string) (string, error) {
	return word, nil
}

func (c searchClient) searchTags(imageName string, tag string) (resp searchImageTagsResponse, err error) {
	resp.ImageName = imageName

	reqUrl, err := url.Parse(fmt.Sprintf(ImageTagsEndpointFormat, imageName))
	if err != nil {
		return
	}
	query := reqUrl.Query()
	query.Set("name", tag)
	query.Set("page_size", "15")
	reqUrl.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return
	}
	rawresp, err := c.client.Do(req)
	if err != nil {
		return
	}

	if rawresp.Body != nil {
		defer rawresp.Body.Close()
	}
	body, err := ioutil.ReadAll(rawresp.Body)
	if err != nil {
		return
	}
	if rawresp.StatusCode >= 400 {
		err = errors.New(string(body))
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	return
}
