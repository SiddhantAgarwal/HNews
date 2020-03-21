package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SiddhantAgarwal/HNews/internal/model"
)

type service struct {
	apiKey string
}

// GetNews : method to get news from HN.
func (svc *service) GetNews(ctx context.Context, numberOfNews int64) []model.News {
	return svc.fetchNews(numberOfNews)
}

// NewHackerNewsService : Returns a news Service instance
func NewHackerNewsService(apiKey string) Service {
	return &service{
		apiKey: apiKey,
	}
}

func (svc *service) fetchNews(numberOfItems int64) []model.News {
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

	if articlesInt, ok := resp["articles"]; ok {
		articles := articlesInt.([]interface{})
		for _, articleInt := range articles {
			article, ok := articleInt.(map[string]interface{})
			if ok {
				news = append(news, parseRespMapToNewsModel(article))
			}
		}
	}
	return news
}

func parseRespMapToNewsModel(article map[string]interface{}) model.News {
	var timeObj time.Time
	var err error

	if publishedAt, ok := article["publishedAt"]; ok {
		timeObj, err = time.Parse(time.RFC3339, publishedAt.(string))
		if err != nil {
			timeObj = time.Now()
		}
	}

	content, _ := article["content"].(string)
	title, _ := article["title"].(string)
	url, _ := article["url"].(string)

	return model.News{
		Content:  content,
		Headline: title,
		Time:     timeObj,
		URL:      url,
	}
}
