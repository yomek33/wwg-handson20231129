package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	genaichat "example.com/genai"
)

var (
	flagName string
	message  string
)

func init() {
	flag.StringVar(&flagName, "name", "gopher", "name of chat bot")
	flag.StringVar(&message, "message", "がんばれ！", "頑張ろうとしている人に送るメッセージ")
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	apikey := os.Getenv("OPENAI_API_KEY")
	bot, err := genaichat.NewBot(ctx, flagName, message, apikey)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %s\n", bot.Name, bot.FirstMessage)
	fmt.Print("> ")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		resp, err := bot.Send(ctx, s.Text())
		if err != nil {
			return err
		}

		fmt.Printf("%s: %s\n", bot.Name, resp)
		fmt.Print("> ")
	}

	return nil
}
