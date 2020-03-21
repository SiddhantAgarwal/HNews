package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SiddhantAgarwal/HNews/internal/model"
)

type service interface {
	GetNews(ctx context.Context, numberOfNews int64) []model.News
}

// Service : News Service Object
type Service struct {
	apiKey string
}

// GetNews : method to get news from HN.
func (svc *Service) GetNews(ctx context.Context, numberOfNews int64) []model.News {
	return svc.fetchNews(numberOfNews)
}

// NewNewsService : Returns a news Service instance
func NewNewsService(apiKey string) *Service {
	return &Service{
		apiKey: apiKey,
	}
}

func (svc *Service) fetchNews(numberOfItems int64) []model.News {
	var resp map[string]interface{}
	request, err := http.NewRequest("GET", fmt.Sprintf("http://newsapi.org/v2/top-headlines?sources=hacker-news&apiKey=%s", svc.apiKey), nil)
	if err != nil {
		fmt.Println(fmt.Errorf("[Service.fetchNews] Error while init request %+v", err))
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(fmt.Errorf("[Service.fetchNews] Get request failure %+v", err))
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&resp)
	if err != nil {
		fmt.Println(fmt.Errorf("[Service.fetchNews] Failed to decode response %+v", err))
	}
	news := parseResponse(resp)
	return news
}

func parseResponse(resp map[string]interface{}) []model.News {
	var news []model.News
	var content, title, url string
	var timeObj time.Time
	var err error

	if articles, ok := resp["articles"]; ok {
		articlesTemp := articles.([]interface{})
		for _, item := range articlesTemp {
			article, ok := item.(map[string]interface{})
			if ok {
				if publishedAt, ok := article["publishedAt"]; ok {
					timeObj, err = time.Parse(time.RFC3339, publishedAt.(string))
					if err != nil {
						timeObj = time.Now()
					}
				}

				content, _ = article["content"].(string)
				title, _ = article["title"].(string)
				url, _ = article["url"].(string)

				news = append(news, model.News{
					Content:  content,
					Headline: title,
					Time:     timeObj,
					URL:      url,
				})
			}
		}
	}
	return news
}
