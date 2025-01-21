const getVatRates = require('../../client/vat_client');
const fetchAndMapProducts = require('../../client/jds_client');

class ProductService {
    constructor(cacheService) {
        this._cacheService = cacheService;
    }

    async getProductList() {
        try {
            let vatData = await this._cacheService.get(`vatRates`);
            let vatRate = 0;

            if (vatData === null) {
                vatData = await getVatRates();
                await this._cacheService.set(`vatRates`, JSON.stringify(vatData), 86400);
            } else {
                vatData = JSON.parse(vatData);
            }

            if (typeof vatData.rates.IDR != 'undefined') {
                vatRate = vatData.rates.IDR;
            }

            let products = await fetchAndMapProducts();
            if (products.length > 0) {
                products.forEach((product) => {
                    product.price_IDR = product.price * (vatRate);
                });
            }
            return products;
        } catch (error) {
            throw error;
        }
    }

    async getAggregateProductList() {
        try {
            let vatData = await this._cacheService.get(`vatRates`);
            let vatRate = 0;

            if (vatData === null) {
                vatData = await getVatRates();
                await this._cacheService.set(`vatRates`, JSON.stringify(vatData), 86400);
            } else {
                vatData = JSON.parse(vatData);
            }
            
            if (typeof vatData.rates.IDR != 'undefined') {
                vatRate = vatData.rates.IDR;
            }
            let products = await fetchAndMapProducts();
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
        } catch (error) {
            throw error;
        }
    }


}


module.exports = ProductService;