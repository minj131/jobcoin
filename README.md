# Gemini Jobcoin

Golang implementation of the Gemini coding challenge.

## Go Modules
Make sure Go Modules are on or exported `GO111MODULE=on` as modules are used instead of vendoring. Check by `go env | grep MODULE`.

## To Build
Run `make build` builds executable binary in `bin/`

## To Run
Run `make build && bin/jobcoin` or `bin/jobcoin`

## Note
- Go version >= 1.12
- Make sure your go environment is configured correctly, if errors occur when building there might be an issue with the `GOPATH`. Either reset your `GOPATH` or reinstall Go version >= 1.12 to fix any possible version issues.
- Make sure go modules are on `GO111MODULE=on` or prefix `GO111MODULE=on` to `make build` or `make unit-tests`.

## Tests
- Run `make unit-tests` to run unit tests.

## Limitations
Traditionally, mixers rely on a large pool of different inputs coming from potentially different people to better anonymize the inputs and outputs. As a result, I don't have to deal with the possible race conditions or concurrency management that might occur when multiple people are trying to mix their coins at the same time. This could easily be solved by using a queue of some sort and is already partially fixed due to the functionality of sending all funds to the master account before distribution.

As for the polling function, I chose to check the state of the wallet every 5 seconds and only proceeding when we have a balance greater than 0. Since the generated wallet is a random hash of length 32, collisions are going to be extremely rare. I added a timeout to 60 seconds to prevent stalling.
