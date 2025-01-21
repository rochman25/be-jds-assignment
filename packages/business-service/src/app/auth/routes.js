const routes = (handler) => [
    {
      method: 'GET',
      path: '/api/v1/auth/token_debug',
      handler: handler.getDebugToken,
      options: {
        auth: 'business_service',
      },
    }
  ];
  
  module.exports = routes;