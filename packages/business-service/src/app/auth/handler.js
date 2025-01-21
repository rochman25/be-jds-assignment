class AuthHandler {
    constructor() {
        this.getDebugToken = this.getDebugToken.bind(this);
    }

    async getDebugToken(request, h) {
        try {
            return request.auth.artifacts.decoded.payload;
        } catch (error) {
            console.error(error);
            const response = h.response({ error_message: ['internal server error'] });
            response.code(500);
            return response;
        }
    }
}

module.exports = AuthHandler;
