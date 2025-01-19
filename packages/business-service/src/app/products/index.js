const ProductHandler = require('./handler');
const routes = require('./routes');

module.exports = {
    name: 'products',
    version: '1.0.0',
    register: async (server) => {
        const handler = new ProductHandler();
        server.route(routes(handler));
    },
};