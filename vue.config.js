module.exports = {
    devServer: {
        open: true,
        host: '0.0.0.0',
        port: 8084,
        https: false,
        historyApiFallback: true,
        allowedHosts: [
            'www.wllinux.com',
        ],
        //以上的ip和端口是我们本机的;下面为需要跨域的
        proxy: {//配置跨域
            '/api': {
                target: 'http://www.wllinux.com:9550',
                changeOrigin: true,
                origin: 'http://www.wllinux.com:8084',
                ws: true,
                secure: false,
                logLevel: 'debug',
                pathRewrite: {
                    '^/api': '',
                }
            }

        }
    }
}