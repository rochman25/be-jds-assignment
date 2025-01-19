const getVatRates = require('../../client/vat_client');
const getProductList = require('../../client/jds_client');

require('../../client/vat_client');

class ProductHandler {
  constructor() {
    this.getProductList = this.getProductList.bind(this);
  }

  async getProductList(request) {
    const vatData = await getVatRates();
    let vatRate = 0;
    if (typeof vatData.rates.IDR != 'undefined') {
      vatRate = vatData.rates.IDR;
    }
    let products = await getProductList();
    if (products.length > 0) {
      products.forEach((product) => {
        product.price_IDR = product.price * (vatRate);
      });
    }

    return products;
  }

  async getAggregateProductList(request) {
    const vatData = await getVatRates();
    let vatRate = 0;
    if (typeof vatData.rates.IDR != 'undefined') {
      vatRate = vatData.rates.IDR;
    }
    let products = await getProductList();
    let data = [];
    if (products.length > 0) {
      // Aggregate by department, product, and price_IDR
      data = products.map((item) => ({
        department: item.department,
        product: item.product,
        price_IDR: item.price * (vatRate),
      }));

      // Sort by price_IDR in ascending order
      data.sort((a, b) => a.price_IDR - b.price_IDR);
    }

    return data;
  }
}

module.exports = ProductHandler;