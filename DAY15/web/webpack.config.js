const path = require('path')

module.exports = {
  mode: "production",
  entry: "./web-client/client.js",
  output: {
    path: path.resolve(__dirname, 'web-client/dist'),
  }
}
