# Introduction
This section we will go into the Swagger and touch on a bit of OpenAPI.

# Key takeaway
- The main way of interacting with with Swagger UI is by using the binary CLI `swaggo/swag`.
- The CLI parses the comments in the Go and converts into OpenAPI docs.
- The comments for swagger is split into 2 main parts. Server and Route configs.
- Steps required to get docs up and running
    - swag init
    - import the docs package into the main file, `import _ "<package-name>/docs"
    - add swagger routes to the docs route, router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

# Link
https://sntshk.medium.com/how-to-integrate-swagger-ui-in-go-backend-gin-edition-b69546cd1149