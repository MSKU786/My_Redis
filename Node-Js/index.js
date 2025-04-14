const net = require('net');
const Parser = require('redis-parser');
const { parserRESP } = require('./utils/parserRESP');

const server = net.createServer((socket) => {
  socket.on('data', (data) => {
    try {
      const args = parserRESP(data);
      console.log(args);
      const command = args[0].toUpperCase();
      const handler = commands[command];
      if (!handler) {
        console.log(error);
        socket.write('-ERR unknown command\r\n');
        return;
      }
      const result = handler(args.slice(1));
      console.log('this is result', result);
    } catch (err) {}
  });

  socket.on('end', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
