const routes = (handler) => [
  {
    method: 'GET',
    path: '/products',
    handler: handler.getProductList,
  },
  {
    method: 'GET',
    path: '/products/aggregate',
    handler: handler.getAggregateProductList,
  }
];

module.exports = routes;