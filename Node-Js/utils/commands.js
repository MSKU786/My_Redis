const { store, expireStore } = require('./map');

const commands = {
  SET: (args) => {
    console.log('Inside SET command', args);
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
  GET: (args) => {
    const [key] = args;

    const value = store.get(key);
    const ttl = expireStore.get(key);

    if (ttl && ttl < Date.now()) {
      store.delete(key);
      expireStore.delete(key);
      return null;
    }
    return value || null;
  },
  DEL: (args) => {
    let count = 0;
    args.forEach((key) => {
      if (store.has(key)) {
        store.delete(key);
        expireStore.delete(key);
        count++;
      }
    });
    return count;
  },
  EXPIRE: (args) => {
    const [key, seconds] = args;

    const ttl = seconds * 1000;
    if (store.has(key)) {
      expireStore.set(key, Date.now() + ttl);
      return 1;
    }
    return 0;
  },
  TTL: (args) => {
    const [key] = args;
    if (!store.has(key)) return -2;
    if (!expiry.has(key)) return -1;

    const remaining = Math.ceil((expiry.get(key) - Date.now()) / 1000);
    return remaining > 0 ? remaining : -2;
  },
};

module.exports = { commands };
