const fastify = require('fastify');

class Server {
  constructor () {
    this.server = fastify();
  }

  async start () {
    this.registerRoutes();

    await this.server.ready();

    this.server.listen(5000, '0.0.0.0', () => console.log('Fastify Server listening on port 5000 🚀'));
  }

  registerRoutes () {
    this.server.get('/listen', (req, res) => {
      res.status(200).send({
        message: 'Responding back'
      });
    });

    this.server.post('/postdata', (req, res) => {
      res.status(200).send({
        message: 'Responding back'
      });
    });
  }
}

new Server().start();
