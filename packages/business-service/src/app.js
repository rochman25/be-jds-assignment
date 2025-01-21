const Hapi = require('@hapi/hapi');
const config = require('../pkg/config');
const Jwt = require('@hapi/jwt');

const product = require('./app/products');

const ProductService = require('./services/product/productService');

// redis
const RedisService = require('./services/redis/redisService');


const init = async () => {
    const redisService = new RedisService();
    const productService = new ProductService(redisService);
    const server = Hapi.server({
        port: config.getAppPort(),
        host: config.getApphost(),
        routes: {
            cors: {
                origin: ['*'],
            }
        }
    });

    await server.register({
        plugin: Jwt,
    });

    server.auth.strategy('business_service', 'jwt', {
        keys: config.jwtSignatureKey(),
        verify: {
            aud: false,
            iss: false,
            sub: false,
            maxAgeSec: (config.jwtExpirationDuration() * 60),
        },
        validate: (artifacts) => ({
            isValid: true,
            credentials: {
                id: artifacts.decoded.payload.user_id,
            },
        }),
    });

    // Add a global error handler
    server.ext('onPreResponse', (request, h) => {
        const response = request.response;

        // Check if the response is a 401 Unauthorized
        if (response.isBoom && response.output.statusCode === 401) {
            switch (response.output.payload.message) {
                case 'Missing authentication':
                    response.output.payload = {
                        error_message: ['Authentication is required to access this resource.'],
                    }
                    break;
                case 'Invalid token':
                    response.output.payload = {
                        error_message: ['Authentication token is invalid.'],
                    }
                    break;
                case 'Token expired':  
                    response.output.payload = {
                        error_message: ['Authentication token has expired.'],   
                    }
                    break;
                default:
                    response.output.payload = {
                        error_message: [response.output.payload.message],
                    }
            }
        }

        return h.continue;
    });

    server.route({
        method: 'GET',
        path: '/',
        handler: (_, h) => {
            return h.response({
                name: config.getAppName(),
                version: config.getAppVersion(),
            }).code(200);
        }
    });

    await server.register([
        {
            plugin: product,
            options: {
                service: productService,
            }
        }
    ]);

    await server.start();
    console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {
    console.log(err);
    process.exit(1);
});

init();