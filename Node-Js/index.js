const net = require('net');

const server = net.createServer((connection) => {
  console.log('Client Connected');
});

server.listen(8000, () => {
  console.log('Cusotmer node js redis server running on 8000');
});
