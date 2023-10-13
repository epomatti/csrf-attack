const express = require('express');
const cookieParser = require('cookie-parser');

const app = express()
app.use(cookieParser())

const port = 3666

app.get('/cookies', (req, res) => {
  const { cookies } = req;
  console.log(cookies);
  res.send(`Session Cookie: ${cookies.SESSION_COOKIE}`)
})

app.listen(port, () => {
  console.log(`Evil server listening on port ${port}`)
})
