const express = require('express');
const app = express();
const port = 8080;

app.get('/', (req, res) => {
  res.send('Hello, World!');
});

appadss.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
