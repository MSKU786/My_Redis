const net = require('net');
const Parser = require('redis-parser');
const { parserRESP } = require('./utils/parserRESP');

const server = net.createServer((socket) => {
  socket.on('data', (data) => {
    try {
      const agrs = parserRESP(data);
      const command = args[0].toUpperCase();
      const handler = commands[command];
      if (!handler) {
        socket.write('-ERR unknown command\r\n');
        return;
      }
      const result = handler(args.slice(1));
      console.log(result);
    } catch (err) {}
  });

  socket.on('end', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
