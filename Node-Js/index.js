const net = require('net');
const Parser = require('redis-parser');

const server = net.createServer((connection) => {
  console.log('Client Connected');
  connection.on('data', (data) => {
    const parser = new Parser({
      returnReply: (reply) => {
        console.log('=>', reply);
      },
      returnError: (err) => {
        console.log('=>', err);
      },
    });
    parser.execute(data);
    console.log('->', data.toString());
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
