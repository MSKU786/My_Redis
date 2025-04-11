export function parserRESP(data) {
  const str = data.toString();
  const [type, ...rest] = str.split('\r\n');

  switch (type[0]) {
    //Array
    case '*':
      console.log(type, rest);
      return [];
    default:
      throw new Error(`Unknown RESP type: ${type[0]}`);
  }
}
