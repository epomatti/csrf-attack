# Cross-Site Request Forgery (CSRF) attack demo

Cross-Site Request Forgery (CSRF) attack demo.

## Setup

### 👍 Real server

Enter the `./realserver` directory and copy the sample environment file:

```sh
cd realserver
cp sample.env .env
```

Start the application in the `./realserver` module:

```sh
go get
go run .
```

### 😈 Evil server

Open a separate terminal session and start the malicious `./evilserver` server:

```sh
cd ./evilserver
go run .
```

## Simulation

A user would login to the real server and get an authentication token as a cookie named `SESSION_TOKEN`:

```
http://localhost:3000/login
```

Now, an attacker with the ability to spoof the DNS could make the user navigate to a fake website and extract sensitive data. Example:

```
http://localhost:3666/cookies
```

After capturing the cookie, the attacker could then use the session identification to invoke the real server pretending to be the real user:

```
curl --cookie "SESSION_COOKIE=<AUTH>" localhost:3000/withdraw=10000
```
