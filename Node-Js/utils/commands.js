import { expireStore, store } from './map';

const commands = {
  SET: (args) => {
    const [key, value, ...opts] = args;
    store.set(key, value);

    const exIndex = opts.indexOf('EX');
    if (exIndex !== -1) {
      const ex = parseInt(opts[exIndex + 1]);
      const ttl = ex * 1000;
      expireStore.set(key, ttl);
    }

    return 'OK';
  },
  GET: (args) => {},
  DEL: (args) => {},
  EXPIRE: (args) => {},
  TTL: (args) => {},
};

export default commands;
