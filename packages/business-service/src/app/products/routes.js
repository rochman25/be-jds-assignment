const routes = (handler) => [
  {
    method: 'GET',
    path: '/products',
    handler: handler.getProductList,
    options: {
      auth: 'business_service',
    },
  },
  {
    method: 'GET',
    path: '/products/aggregate',
    handler: handler.getAggregateProductList,
    options: {
      auth: 'business_service',
    },
  }
];

module.exports = routes;