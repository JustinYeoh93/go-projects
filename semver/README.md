# Introduction
This section is to learn how to do semver in Go. 

The main discovery aimed here is to understand how modules are managed by Go.

# Links
https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories

# Standard
Many of Go's suggestions are to follow a **one repo one module** system. 

After reading through Go's way of managing modules, I've come to understand that you need to put a tag in your VCS 

# Hypothesis
As I'm interested to see if we can do deeply nested modules (n > 2), there is a few hypothesis that needs to first be challenged.

## Nested requires a root module
I've dug through quite a few docs and most successful nested modules have a common root module. This is ever clear when looking at this diagram. https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories.

As it shows, it is first referring to the root module, then trickles down to the nested module.

### Result
This is proven true with https://github.com/JustinYeoh93/go-projects/releases/tag/semver%2Fmod1%2Fv0.0.1. There is also no need for a release on the VCS, only the existance of the tag.

# Conclusion
There was initially another test expected.
- Only a single layer nest is allowed

However, the above test has proven that as long as the mod exist from the root repo onwards, you can nest all ya want.
