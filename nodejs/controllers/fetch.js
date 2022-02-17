const { repoResource, currConvertIDRtoUSD, DB } = require('../helpers');
const moment = require('moment');
const ss = require('simple-statistics')

module.exports = {
  fetchRoleBiasa: async function (req, res, next) {
    const db = DB();
    const resource = await repoResource();
    let cursIDR_USD = null;
    const result = [];

    for (let i = 0; i < resource.length; i++) {
      const item = resource[i];

      const modified = await new Promise(resolve => {
        if (!item.price) item.price = '0';

        db.get(`SELECT price_usd FROM currency WHERE price_idr = ?`, item.price, async function (err, row) {
          if (err) console.log('-- err', err);

          let priceUSD = row?.price_usd;

          if (!row) {
            if (!cursIDR_USD) {
              console.log('-- call currency api');
              cursIDR_USD = await currConvertIDRtoUSD();
            }

            priceUSD = (parseInt(item.price) * cursIDR_USD.IDR_USD).toString();
            // insert newly converted currency
            db.run(`INSERT INTO currency (price_idr, price_usd) VALUES (?, ?)`, [item.price, priceUSD]);
          }

          resolve({
            ...item,
            price_usd: priceUSD
          });
        });
      });

      result.push(modified);
    }

    res.json({
      "status": "success",
      "message": "Successfully Get Resources",
      "data": result
    });
  },

  fetchRoleAdmin: async function (req, res, next) {
    if (req.user.sub.role !== 'admin') {
      return res.status(403).json({
        "status": "error",
        "message": "Forbidden",
      });
    };

    const resource = await repoResource();
    const result = [];
    const objTmp = {};

    for (let i = 0; i < resource.length; i++) {
      const item = resource[i];

      if (!item.tgl_parsed) item.tgl_parsed = new Date();
      const tglParsed = moment(new Date(item.tgl_parsed)).format('YYYY-MM-DD HH:mm:ss');

      const modifiedItem = {
        ...item,
        tgl_parsed: tglParsed
      };

      const dateOnly = moment(new Date(tglParsed)).format('YYYY-MM-DD');
      const keyObj = `${item.area_provinsi}_${dateOnly}`;

      if (!objTmp[keyObj]) objTmp[keyObj] = [];

      objTmp[keyObj].push(modifiedItem);
    }

    for (const [key, value] of Object.entries(objTmp)) {
      const keyArr = key.split('_');
      const year = parseInt(moment(new Date(keyArr[1])).format('YYYY'));
      const month = parseInt(moment(new Date(keyArr[1])).format('MM'));
      const week = parseInt(moment(new Date(keyArr[1])).format('w'));

      const sizeArr = [];
      const priceArr = [];
      value.forEach(item => {
        sizeArr.push(parseInt(item.size));
        priceArr.push(parseInt(item.price));
      })

      const minSize = ss.min(sizeArr);
      const maxSize = ss.max(sizeArr);
      const medianSize = ss.median(sizeArr);
      const avgSize = ss.mean(sizeArr);

      const minPrice = ss.min(priceArr);
      const maxPrice = ss.max(priceArr);
      const medianPrice = ss.median(priceArr);
      const avgPrice = ss.mean(priceArr);
      

      result.push({
        year,
        month,
        week,
        province_area: keyArr[0],
        size_stat: {
          min: minSize,
          max: maxSize,
          median: medianSize,
          avg: avgSize
        },
        price_stat: {
          min: minPrice,
          max: maxPrice,
          median: medianPrice,
          avg: avgPrice
        }
      });
    }

    res.json({
      "status": "success",
      "message": "Successfully Get Resources",
      "data": result
    });
  },
  
  fetchClaims: async function (req, res, next) {
    res.json({
      "status": "success",
      "message": "Successfully Claim Token",
      "data": req.user
    });
  }
}