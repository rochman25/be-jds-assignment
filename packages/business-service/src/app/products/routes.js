const routes = (handler) => [
  {
    method: 'GET',
    path: '/api/v1/products',
    handler: handler.getProductList,
    options: {
      auth: 'admin_business_service',
    },
  },
  {
    method: 'GET',
    path: '/api/v1/products/aggregate',
    handler: handler.getAggregateProductList,
    options: {
      auth: 'admin_business_service',
    },
  }
];

module.exports = routes;