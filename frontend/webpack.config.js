var webpack = require('webpack');
var path = require('path');
var MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
    entry: {
      'js/index' : './src/scenes/Dashboard/app.js'
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: "css/[name].css"
        })
    ],
    module: {
      rules: [
        {
          test: /\.(js|jsx)$/,
          exclude: /node_modules/,
          use: ['babel-loader']
        },
        {
          test: /\.s?css$/,
          use: [
              MiniCssExtractPlugin.loader,
              "css-loader",
              "sass-loader"
          ]
      },
      ]
    },
    resolve: {
      extensions: ['*', '.js', '.jsx']
    },
    output: {
      path: __dirname + '/dist/',
      filename: '[name].js'
    },
    devServer: {
      hot: true,
      publicPath: "/public/",
      port: 8081,
      https: true,
    },
    devtool: 'cheap-module-source-map',
  };