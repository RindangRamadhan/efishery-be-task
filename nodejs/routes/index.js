const express = require('express');
const router = express.Router();
const { DB } = require('../helpers');
const { fetchRoleBebas, fetchRoleAdmin, fetchClaims } = require('../controllers/fetch');
const { authenticateToken } = require('../middleware');

/* GET home page. */
router.get('/', function(req, res, next) {
  res.json({
    "status": "ok",
    "message": "Hello eFishery",
  });
});

router.get('/fetch', function(req, res, next) {
  res.json({
    "status": "ok",
    "message": "Hello eFishery",
  });
});

router.get('/fetch/resource-bebas', authenticateToken, fetchRoleBebas);
router.get('/fetch/resource-admin', authenticateToken, fetchRoleAdmin);
router.get('/fetch/claims', authenticateToken, fetchClaims);

module.exports = router;
