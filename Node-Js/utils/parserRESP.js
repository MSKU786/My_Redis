function parserRESP(data) {
  const str = data.toString();
  const [type, ...rest] = str.split('\r\n');

  switch (type[0]) {
    //Array
    case '*':
      //Count number of elements
      const count = parseInt(type.slice(1), 10);
      const elements = [];

      for (let i = 0; i < count * 2; i += 2) {
        // First char is always string followed by actaull coomand
        const element = rest[i + 1];
        elements.push(element);
      }
      return elements;
    default:
      throw new Error(`Unknown RESP type: ${type[0]}`);
  }
}

function serializeRESP(data) {
  if (data === null) return '$-1\r\n';
  if (typeof data === 'string') return `$${data.length}\r\n${data}\r\n`;
  if (typeof data === 'number') return `:${data}\r\n`;
  if (data === 'OK') return '+OK\r\n';

  throw new Error(`Unsupported data type for RESP: ${typeof data}`);
}

module.exports = {
  parserRESP,
  serializeRESP,
};
