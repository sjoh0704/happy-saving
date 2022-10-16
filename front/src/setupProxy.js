const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = (app) => {
    app.use(
        createProxyMiddleware(['/apis', '/auth'], {
            target: process.env.REACT_APP_BACK_BASE_URL,
            changeOrigin: true
        })
    )
}
