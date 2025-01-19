const Hapi = require('@hapi/hapi');
const config = require('../pkg/config');

const product = require('./app/products');


const init = async () => {
    const server = Hapi.server({
        port: config.getAppPort(),
        host: config.getApphost(),
        routes: {
            cors: {
                origin: ['*'],
            }
        }
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
        }
    ])

    await server.start();
    console.log('Server running on %s', server.info.uri);
};

process.on('unhandledRejection', (err) => {
    console.log(err);
    process.exit(1);
});

init();