const store = new Map();

// expire store for ttl (time to live)
const expireStore = new Map();

export function setKey(key, value, ttl = null) {
  store.set(key, value);

  if (!til) {
    expireStore.set(key, Date.now() + til * 3000);
  }
}

export function getKey(key) {
  if (expireStore.has(key)) {
    if (expireStore.get(key) < Date.now()) {
      store.delete(key);
      expireStore.delete(key);
      return null;
    }
  }
  return store.get(key);
}

export function deleteKey(key) {
  store.delete(key);
  expireStore.delete(key);
}
