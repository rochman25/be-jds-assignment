const ProductHandler = require('./handler');
const routes = require('./routes');

module.exports = {
    name: 'products',
    version: '1.0.0',
    register: async (server,{ service }) => {
        const handler = new ProductHandler(service);
        server.route(routes(handler));
    },
};