# Cross-Site Request Forgery (CSRF) attack demo

Cross-Site Request Forgery (CSRF) attack demo.

Start the real server:

```sh
cd server
npm install
node app.js
```

Start the malicious server:

```sh
cd evilserver
npm install
node evil.js
```

Login to the real server: http://localhost:3000/login. This will generate the login token called `SESSION_TOKEN`.

Now an attaker able to spoof the DNS could make the user navigate to a fake website to extract the sensitive data.

Example: http://localhost:3666/cookies

After extracting the cookie, the attacker could then use the session identification to invoke the real server prentending to be the real user, such as http://localhost:3000/withdraw=10000.
