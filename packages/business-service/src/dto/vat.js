class VatRates {
    constructor(date, base, rates) {
        this.date = date;
        this.base = base;
        this.rates = rates;
    }

    static fromJSON(json) {
        return new VatRates(json.date, json.base, json.rates);
    }

    toJSON() {
        return {
            date: this.date,
            base: this.base,
            rates: this.rates
        };
    }
}

module.exports = VatRates;