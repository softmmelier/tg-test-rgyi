package articles

import (
	"context"
	"fmt"

	"encore.app/hello"
	"encore.app/other"
)

// Welcome to Encore!
// This is a simple "Hello World" project to get you started.
//
// To run it, execute "encore run" in your favorite shell.

// ==================================================================

// This is a simple REST API that responds with a personalized greeting.
// To call it, run in your terminal:
//
//     curl http://localhost:4000/other/World
//
//encore:api public path=/article/:name
func World(ctx context.Context, name string) (*Response, error) {
	msg := "Articles there, " + name + "!"

	otherSrv, _ := other.World(ctx, name)
	fmt.Println(otherSrv.Message)

	helloSrv, _ := hello.World(ctx, name)
	fmt.Println(helloSrv.Message)

	return &Response{Message: msg}, nil
}

type Response struct {
	Message string `json:"message"`
}

// ==================================================================

// Encore comes with a built-in development dashboard for
// exploring your API, viewing documentation, debugging with
// distributed tracing, and more. Visit your API URL in the browser:
//
//     http://localhost:4000
//

// ==================================================================

// Next steps
//
// 1. Deploy your application to the cloud with a single command:
//
//     git push encore
//
// 2. To continue exploring Encore, check out one of these topics:
//
//    Building a Slack bot:  https://encore.dev/docs/tutorials/slack-bot
//    Building a REST API:   https://encore.dev/docs/tutorials/rest-api
//    Using SQL databases:   https://encore.dev/docs/develop/sql-database
//    Authenticating users:  https://encore.dev/docs/develop/auth
