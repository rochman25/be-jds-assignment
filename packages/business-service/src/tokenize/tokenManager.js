const Jwt = require("@hapi/jwt");
const config = require("../../pkg/config");

const seconds = config.jwtExpirationDuration() * 60;
const TokenManager = {
    generateAccessToken: (payload) => Jwt.token.generate(payload, {
        key: config.jwtSignatureKey(),
        algorithm: "HS256",
    }, {
        ttl: config.jwtExpiration(seconds),
    }),
};

module.exports = TokenManager;