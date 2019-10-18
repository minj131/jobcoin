# Gemini Jobcoin

## To Build Binary
Run `make build`

## To Run Application
Run `make build && bin/jobcoin` or `bin/jobcoin`

## Note
- Go version >= 1.12
- Make sure your go environment is configured correctly
- Make sure go modules are on `GO111MODULE=on`

## Limitations
Traditionally, mixers rely on a large pool of different inputs coming from potentially different people to better anonymize the inputs and outputs.
Since I am the only one sending in funds, it's hard to say how the functionality might work if multiple people are trying to mix their coins at the same time.
Though, the idea that the coins are getting sent to a master wallet (which should be constantly refreshed) as the centralized point to which coins get redistributed should address that issue.

As for the polling function, I chose to check the state of the wallet every 5 seconds and only continuing when we have a balance. Since the generated wallet is a random hash of length 32, collisions are going to be extremely rare. To combat nefarious attempts, I added a timeout to 60 seconds.
