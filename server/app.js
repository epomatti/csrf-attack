const express = require('express');
const crypto = require('crypto');
const cookieParser = require('cookie-parser');

const app = express()
app.use(cookieParser())

const port = 3000

let authToken = "";

app.get('/login', (req, res) => {
  const sessionCookie = crypto.randomUUID();
  authToken = sessionCookie;
  res.cookie('SESSION_COOKIE', sessionCookie)
  res.send(`Authenticated with cookie: ${sessionCookie}`)
})

app.get('/withdraw', (req, res) => {
  const { SESSION_COOKIE } = req.cookies;
  if (SESSION_COOKIE === authToken) {
    const { amount } = req.query;
    res.send(`${amount} removed from your account`);
  } else {
    res.status(401).send('Unauthorized');
  }

})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
