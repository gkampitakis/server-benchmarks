'use strict';
const nanoexpress = require('nanoexpress');

const server = nanoexpress();

server.get('/listen', (req, res) => {
  res.status(200).send({
    message: 'Responding back'
  });
});

server.post('/postdata', (req, res) => {
  res.status(200).send({
    message: 'Responding back'
  });
});

server.listen(5000, '0.0.0.0');