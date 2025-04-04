const net = require('net');
const Parser = require('redis-parser');

const server = net.createServer((socket) => {
  console.log('Client Connected');
  socket.on('data', (data) => {
    const input = data.toString().trim(); // e.g., "SET mykey myvalue"
    const args = input.split(' ');
    const cmd = args[0].toUpperCase();

    switch (cmd) {
      case 'SET':
        store.set(args[1], args[2]);
        socket.write('OK\n'); // Reply with simple response
        break;
      case 'GET':
        const value = store.get(args[1]) || '(nil)';
        socket.write(`${value}\n`);
        break;
      default:
        socket.write('ERROR: Unknown command\n');
    }
  });

  socket.on('end', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
