const net = require('net');
const Parser = require('redis-parser');

const server = net.createServer((socket) => {
  console.log('Client Connected');
  socket.on('data', (data) => {
    const input = data.toString().trim(); // e.g., "SET mykey myvalue"
    const args = input.split(' ');
    const cmd = args[0].toUpperCase();
    const command = data.toString().trim();

    console.log(args, cmd, command);
    socket.write('+OK\r\n'); // Simple Redis protocol response
  });

  socket.on('end', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
