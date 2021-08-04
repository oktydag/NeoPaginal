var path = require('path'),
    rootPath = path.normalize(__dirname + '/..'),
    env = process.env.NODE_ENV || 'qa';

var config = require('../config.json');

config[env].root = rootPath;

module.exports = config[env];
