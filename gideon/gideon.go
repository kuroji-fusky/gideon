package gideon

import (
	"fmt"
	"net"
)

type CookieMuncher struct {
	FandomSession *string
	FandomCSRF    *string
}

type InitConfig struct {
	Headers map[string]any
	Cookies CookieMuncher
}

type wikiPageRequest struct {
	PartialWikiUrl    string
	CSSArticleContent string
	Headers           *map[string]any
	Cookies           *CookieMuncher
}

func checkValidHostname(hs string) error {
	addr, err := net.LookupHost(hs)

	fmt.Printf("DEBUG: hostname valid - %s", addr)

	if err != nil {
		return fmt.Errorf("AMBOBO MO")
	}

	return nil
}

const DEFAULT_ARTICLE_SELECTOR = "#content.page-content"

func Initalize(baseUrl string) *wikiPageRequest {
	checkValidHostname(baseUrl)

	return &wikiPageRequest{
		PartialWikiUrl:    baseUrl + "/wiki/",
		CSSArticleContent: DEFAULT_ARTICLE_SELECTOR,
		Headers:           nil,
		Cookies:           nil,
	}
}

func InitalizeWithConfig(baseUrl string, cfg InitConfig) *wikiPageRequest {
	checkValidHostname(baseUrl)

	return &wikiPageRequest{
		PartialWikiUrl:    baseUrl + "/wiki/",
		CSSArticleContent: DEFAULT_ARTICLE_SELECTOR,
		Headers:           &cfg.Headers,
		Cookies:           &cfg.Cookies,
	}
}

type CategoryList struct {
	Title string
	Image *string
	Url   string
}

type wikiTopNavNestedItems struct {
	Text string  `json:"text"`
	Link *string `json:"link,omitempty"`
}

type WikiTopNav struct {
	HeadingText string             `json:"heading_text"`
	HeadingLink *string            `json:"heading_link,omitempty"`
	Items       *[]WikiTopNavItems `json:"items,omitempty"`
}

type WikiTopNavItems struct {
	wikiTopNavNestedItems
	Subitems *[]wikiTopNavNestedItems `json:"subitems,omitempty"`
}

// Retrieves the top navigation of a given Fandom wiki
func (f *wikiPageRequest) GetTopNav(pagename string) (WikiTopNav, error) {
	return WikiTopNav{}, nil
}

// []string is a placeholder type - will be using a different type later for HTML parsing
type ArticleResponse PageResponse[[]string]

func (f *wikiPageRequest) GetArticle(pagename string) (*ArticleResponse, error) {
	return &ArticleResponse{}, nil
}

// Returns a page with a "Category:" prefix
func (f *wikiPageRequest) RequestCategoryPage(pagename string) (*PageResponse[[]CategoryList], error) {
	return &PageResponse[[]CategoryList]{}, nil
}

// Returns a page with a "Special:" prefix
func (f *wikiPageRequest) RequestSpecialPage(pagename string) (*PageResponse[[]any], error) {
	return &PageResponse[[]any]{}, nil
}

// Returns a page with a "Template:" prefix
func (f *wikiPageRequest) RequestTemplate(pagename string) (*PageResponse[[]any], error) {
	return &PageResponse[[]any]{}, nil
}

// A wrapper for "Special:AllPages"
func (f *wikiPageRequest) GetAllPages() (*PageResponse[[]any], error) {
	return f.RequestSpecialPage("AllPages")
}

// From ArticleResponse constructor
type HistoryStore struct {
	RevisionDate string
	RevisionId   string
	TotalBytes   int32
	ByteDiff     uint32
	EditMsg      *string
	Tags         []string
}

func (f *PageResponse[any]) History() ([]HistoryStore, error) {
	return []HistoryStore{}, nil
}

// Retrieves all the categories from a given article page
func (f *PageResponse[any]) GetCategoryPages() ([]string, error) {
	return []string{}, nil
}

type CommentStoreResponse struct {
	Id           string   `json:"comment_id"`
	Username     string   `json:"username"`
	UsernameLink string   `json:"username_link"`
	Date         string   `json:"date"`
	Contents     []string `json:"contents"`
}

func (f *PageResponse[any]) GetComments() ([]CommentStoreResponse, error) {
	return []CommentStoreResponse{}, nil
}
