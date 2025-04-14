function parserRESP(data) {
  const str = data.toString();
  const [type, ...rest] = str.split('\r\n');

  switch (type[0]) {
    //Array
    case '*':
      console.log('this is parserReSP', type);
      console.log('this is rest', rest);
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

module.exports = {
  parserRESP,
};
