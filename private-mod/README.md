# Introduction
Sometime, we have to use private go modules. The operation is quite similar to how we do `go get` but with 2 more prerequisites.
- Setup the `GOPRIVATE` variables to add the private repo suffix, in our case `github.com/JustinYeoh93/private-go-modules`.
- Setup the credentials either via HTTPS + Token or through SSH. See https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project

