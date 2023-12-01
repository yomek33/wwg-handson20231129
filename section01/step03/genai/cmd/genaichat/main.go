package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	"example.com/genai"
)

var (
	flagName string
)

func init() {
	// TODO: flagNameに-nameで指定した名前が入るようにしてください。
	// デフォルト値はgopherにしてください
	flag.StringVar(&flagName, "name", "gopher", "bot name")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	apikey := os.Getenv("OPENAI_API_KEY")
	bot, err := genai.NewBot(ctx, flagName, apikey)
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
