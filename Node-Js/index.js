const net = require('net');
const { parserRESP, serializeRESP } = require('./utils/parserRESP');
const { commands } = require('./utils/commands');

const server = net.createServer((socket) => {
  socket.on('data', (data) => {
    try {
      const args = parserRESP(data);

      const command = args[0].toUpperCase();

      const handler = commands[command];

      if (!handler) {
        console.log(error);
        socket.write('-ERR unknown command\r\n');
        return;
      }
      const result = handler(args.slice(1));
      socket.write(serializeRESP(result));
    } catch (err) {}
  });

  socket.on('end', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
