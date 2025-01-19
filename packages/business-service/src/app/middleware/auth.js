const validateAuthHeader = (req, res, next) => {
    const authHeader = req.headers['authorization'];

    if (!authHeader) {
        return res.status(401).json({ message: 'Authorization header is missing' });
    }

    // Assuming the token should start with 'Bearer '
    const token = authHeader.split(' ')[1];

    if (!token) {
        return res.status(401).json({ message: 'Token is missing' });
    }

    // Here you can add additional token validation logic if needed

    next();
};

module.exports = validateAuthHeader;