class ProductHandler {
  constructor(service) {
    this._service = service;
    this.getProductList = this.getProductList.bind(this);
    this.getAggregateProductList = this.getAggregateProductList.bind(this);
  }

  async getProductList(request,h) {
    try {
      const products = await this._service.getProductList();
      return products;
    } catch (error) {
      console.error(error);
      const response  = h.response({  error_message: ['internal server error'] });
      response.code(500);
      return response;
    }
  }

  async getAggregateProductList(request,h) {
    try {
      const products = await this._service.getAggregateProductList();
      return products;
    } catch (error) {
      console.error(error);
      const response  = h.response({  error_message: ['internal server error'] });
      response.code(500);
      return response;
    }
  }
}

module.exports = ProductHandler;