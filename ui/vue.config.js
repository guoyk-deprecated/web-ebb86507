module.exports = {
    runtimeCompiler: true,
    transpileDependencies: [
        /\bvue-awesome\b/
    ],
    devServer: {
        proxy: {
            '^/api': {
                target: 'http://127.0.0.1:3000',
            }
        }
    }
}