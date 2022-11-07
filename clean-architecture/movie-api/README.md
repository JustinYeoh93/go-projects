# Introduction
This is a Clean Architecture GoLang example from the link https://betterprogramming.pub/lets-build-a-movie-api-with-clean-architecture-ef1f555b563d.

# Key takeaways
A few key takeaways from this blog. 

## Handlers, Services, and Repositories
The difference of handlers, services and repositories explained here are bounded within the context whereby the domain is a HTTP API.

**Handlers**: It is a layer which gets http request and returns http response to the client. (In most of my code, I refer to them within the controller.go)

**Services**: It is a layer where business logic lives in.

**Repositories**: It is the data access layer.