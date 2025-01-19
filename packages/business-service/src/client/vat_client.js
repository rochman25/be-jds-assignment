const axios = require('axios');
const VatRates = require('../dto/vat');
const config = require('../../pkg/config');

async function getVatRates() {
    try {
        const response = await axios.get(config.vatRatesUrl());

        // Create an instance of VatRates from the API response
        const vatRates = VatRates.fromJSON(response.data);

        return vatRates;
    } catch (error) {
        console.error('Error fetching VAT rates:', error);
        throw error;
    }
}

module.exports = getVatRates;