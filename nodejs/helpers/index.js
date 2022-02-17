const sqlite3 = require('sqlite3').verbose();
const fetch = require('node-fetch');

module.exports = {
  DB: function () {
    const db = new sqlite3.Database('nodejs.db');

    return db;
  },
  repoResource: async function () {
    const response = await fetch('https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list');
    const resource = await response.json();

    const filteredResource = resource.filter(item => item.uuid !== null);

    return filteredResource;
  },
  currConvertIDRtoUSD: async function () {
    const url = `${process.env.HOST_CURRCONV}/api/v7/convert?q=IDR_USD&compact=ultra&apiKey=${process.env.CURRCONV_API}`;

    const response = await fetch(url);
    const json = await response.json();

    return json;
  }
};