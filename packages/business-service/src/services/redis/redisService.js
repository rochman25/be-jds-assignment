const redis = require('redis');
const config = require('../../../pkg/config');

class RedisService {
  constructor() {
    this._client = redis.createClient({
      url: `redis://${config.redisHost()}`,
    });
    // Handle connection errors
    this._client.on('error', (err) => {
      console.error('Redis error:', err);
    });

    // Log successful connection
    this._client.on('connect', () => {
      console.log('Connected to Redis successfully!');
    });

    // Connect to Redis
    this._client.connect().catch((err) => {
      console.error('Error connecting to Redis:', err);
    });
  }

  async set(key, value, expiryInSeconds = null) {
    if (expiryInSeconds) {
      await this._client.setEx(key, expiryInSeconds, value);
    } else {
      await this._client.set(key, value);
    }
  }

  async get(key) {
    return await this._client.get(key);
  }

  async delete(key) {
    return await this._client.del(key);
  }
}

module.exports = RedisService;