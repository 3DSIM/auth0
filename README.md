# auth0
Library written in Go for interacting with auth0

## Technical Specifications

### Platforms Supported
MacOS, Windows, and Linux

## Contributing code
Read this article and follow the steps they outline: http://scottchacon.com/2011/08/31/github-flow.html

All PRs should be signed off by a member of the team before merging.

## Installing dependencies
We use dep for dependency management. To install dependencies:
* Install dep (https://github.com/golang/dep)
* run `dep ensure -v` to download dependencies

## Generating Fakes
We use https://github.com/maxbrunsfeld/counterfeiter to generate fakes (aka mocks).  To regenarate a fake after an interface change:
* Install counterfeiter.  See their github repo for install instructions.
* In the root of the project run `go generate`

## Team
* Tim Sublette
* Ryan Walls
* Chad Queen
* Pete Krull
* Alex Drinkwater

## Original release
April 2017
