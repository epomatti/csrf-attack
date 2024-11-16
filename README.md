# Cross-Site Request Forgery (CSRF) attack demo

Cross-Site Request Forgery (CSRF) attack demo.

Definition from the OWASP Cheat Sheet:

> A [Cross-Site Request Forgery (CSRF)][3] attack occurs when a malicious web site, email, blog, instant message, or program tricks an authenticated user's web browser into performing an unwanted action on a trusted site. If a target user is authenticated to the site, unprotected target sites cannot distinguish between legitimate authorized requests and forged authenticated requests.

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

### Login CSRF

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

## Solution/Mitigation

- CSRF Tokens - Requires persistence on server side
- Double cookie

## Analysis

### Secure Cookie - weak confidentiality

A `secure` cookie will be sent over to (potentially) insecure locations, such as a subdomain or different port.

In our simulation, a `secure` cookie from `localhost:3000` was sent to `localhost:3666` over HTTP without TLS.

There are some limitations from the State Management [RFC 6265][2], for example, [weak confidentiality][1]:

    Cookies do not provide integrity guarantees for sibling domains (and
    their subdomains).  For example, consider foo.example.com and
    bar.example.com.  The foo.example.com server can set a cookie with a
    Domain attribute of "example.com" (possibly overwriting an existing
    "example.com" cookie set by bar.example.com), and the user agent will
    include that cookie in HTTP requests to bar.example.com.  In the
    worst case, bar.example.com will be unable to distinguish this cookie
    from a cookie it set itself.  The foo.example.com server might be
    able to leverage this ability to mount an attack against
    bar.example.com.



[1]: https://www.rfc-editor.org/rfc/rfc6265#section-8.6
[2]: https://www.rfc-editor.org/rfc/rfc6265
[3]: https://owasp.org/www-community/attacks/csrf
