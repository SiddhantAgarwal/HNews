package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/SiddhantAgarwal/HNews/internal/model"
	"github.com/SiddhantAgarwal/HNews/internal/service"

	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

// NewHNewsApp : Creates and returns a new HNews app.
func NewHNewsApp() *cli.App {
	app := cli.NewApp()
	app.Name = "HackerNews CLI"
	app.Usage = "Get your news on the terminal"
	app.Flags = []cli.Flag{
		&cli.Int64Flag{
			Name:  "number_of_news",
			Value: 10,
			Usage: "Number of news items to be fetched",
		},
	}
	app.Action = func(c *cli.Context) error {
		numberOfNews := c.Int64("number_of_news")
		apiKey := os.Getenv("NEWS_API_KEY")
		if apiKey == "" {
			fmt.Println("Please set the newsapi key in environment variables")
		} else {
			newsService := service.NewHackerNewsService(apiKey)
			printNews(newsService.GetNews(context.Background(), numberOfNews))
		}
		return nil
	}
	return app
}

func printNews(news []model.News) {
	for _, item := range news {
		fmt.Println("---------------------------------------------------------------------------------")
		fmt.Printf("%s - %s\n\n", Bold(item.Headline), Faint(item.Time.Format(time.RFC1123)))
		fmt.Println(Italic(item.Content))
		fmt.Printf("\n%s - %s\n", Faint("More Here"), Underline(item.URL))
	}
}
