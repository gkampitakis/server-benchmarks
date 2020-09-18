const fastify = require('fastify');
const customHealthCheck = require('fastify-custom-healthcheck');

class Server {
  constructor () {
    this.server = fastify();
    this.server.register(customHealthCheck);
  }

  async start () {
    this.registerRoutes();

    await this.server.ready();

    this.server.listen(5000, '0.0.0.0', () => console.log('Server listening on port 5000 🚀'));
  }

  registerRoutes () {
    this.server.get('/listen', (req, res) => {
      console.log(JSON.stringify(req.query));

      res.status(200).send({
        message: 'Responding back',
        data: req.query
      });
    });

    this.server.post('/postdata', (req, res) => {
      const bodyData = JSON.stringify(req.body);
      console.log(bodyData);

      res.status(201).send(JSON.parse(bodyData));
    });
  }
}

new Server().start();
