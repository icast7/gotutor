package main

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	const GET = "GET"
	const POST = "POST"
	const LOG_FORMAT = "\nUrl: %s\nHttpCode: %d\nBody: %s\n"

	httpCode, body := nginxServer.handleRequest(appStatusURL, GET)
	fmt.Printf(LOG_FORMAT, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, GET)
	fmt.Printf(LOG_FORMAT, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, GET)
	fmt.Printf(LOG_FORMAT, appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, POST)
	fmt.Printf(LOG_FORMAT, createuserURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, GET)
	fmt.Printf(LOG_FORMAT, createuserURL, httpCode, body)
}
