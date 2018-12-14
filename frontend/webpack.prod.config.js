var webpack = require('webpack');
var path = require('path');
var MiniCssExtractPlugin = require("mini-css-extract-plugin");
var ManifestPlugin = require('webpack-manifest-plugin');

module.exports = {
    entry: {
      'js/index' : './src/scenes/Dashboard/app.js'
    },
    plugins: [
      new MiniCssExtractPlugin({
        filename: "css/[name].[contenthash].css",
        chunkFilename: "[name].[contenthash].css"
      }),
      new ManifestPlugin()
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
        {
          test: /\.(png|jpg|gif)$/,
          use: [
              {
                  loader: 'file-loader',
                  options: {
                      outputPath: "images/",
                      publicPath: "/public/images/"
                  }
              }
          ]
        }
      ]
    },
    output: {
      path: __dirname + '/dist/',
      filename: '[name].[chunkhash].js',
      chunkFilename: '[name].[chunkhash].js'
    }
};