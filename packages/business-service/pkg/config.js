const fs = require('fs');
const path = require('path');

require('dotenv').config({ path: path.resolve(__dirname, '../../../.env') });

const getEnvVariable = (key) => {
    return process.env[key];
};

const getApphost = () => {
    return getEnvVariable('APP_BUSINESS_HOST');
};

const getAppPort = () => {
    return getEnvVariable('APP_BUSINESS_PORT');
};

const getAppEnv = () => {
    return getEnvVariable('APP_BUSINESS_ENV');
};

const getAppVersion = () => {
    return getEnvVariable('APP_BUSINESS_VERSION');
};

const getAppName = () => {
    return getEnvVariable('APP_BUSINESS_NAME');
};

const getJwtSecret = () => {
    return getEnvVariable('JWT_SIGNATURE_KEY');
};

const getJwtExpiration = () => {
    return getEnvVariable('JWT_EXPIRATION_DURATION');
};

const vatRatesUrl = () => {
    return getEnvVariable('VAT_RATES_URL');
}

module.exports = {
    getEnvVariable,
    getApphost,
    getAppPort,
    getAppName,
    getAppEnv,
    getAppVersion,
    getJwtSecret,
    getJwtExpiration,
    vatRatesUrl,
};