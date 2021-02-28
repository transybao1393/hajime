// Package services is simple package to handle main flow of the program
package services

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// RequestHandler is very basic function
func RequestHandler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Path()) {
	case "/":
		informationRoute(ctx)
	case "/hello":
		hello(ctx)
	default:
		ctx.Error("Path not found ", fasthttp.StatusNotFound)
	}

}

func hello(ctx *fasthttp.RequestCtx) {
	fmt.Println("This is a hello for you!")
	fmt.Fprintf(ctx, "This is a hello for you!")
	ctx.SetContentType("application/json; charset=utf8")
}

func informationRoute(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "API information\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	//- content type, headers, cookies
	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	// ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
