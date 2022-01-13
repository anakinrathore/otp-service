# otp-service
This is an assessment repository

### Run Server(PORT is 8000)
```
brew install ngrok
go run ./server/server.go
ngrok http 8000
```

### Dependency
```
auth-service on port 9000
test.env which is gitignored, will send that as part of the mail

We need to expose a public endpoint so need ngrok so pub sub can push the event to our service on localhost, will require my assistance while testing so i can use the generated ngrok url in google pubsub configuration in my GCP console.
```

### NOTE
```
Currently its hard coded for Indian phone number only but that can be changed by changing the server/server.go file line no. 62 and replacing the country code from +91 to the country in which you want to send messages(Sorry for not making it work with all the countries could have been achieved with a little bit more work by storing country to code configurations)

I made the service to work on normal web server instead of grpc, felt it was adding a little complexity for a small piece of code or maybe because i am not that comfortable with gRPC as i am working on this stuff for the first time so that also played out in mind. If there is requirement to do so in PR review, will gladly do it.
```
