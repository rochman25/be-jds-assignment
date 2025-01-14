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

module.exports = {
    getEnvVariable,
    getApphost,
    getAppPort,
    getAppName,
    getAppEnv,
    getAppVersion
};