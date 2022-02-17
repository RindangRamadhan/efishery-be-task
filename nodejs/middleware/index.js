const jwt = require('jsonwebtoken');

module.exports = {
  authenticateToken: function (req, res, next) {
    const authHeader = req.headers['authorization'];
    const token = authHeader && authHeader.split(' ')[1];

    if (token == null) {
      return res.status(401).json({
        "status": "error",
        "message": "Unauthorized",
      });
    }
    
    jwt.verify(token, process.env.SESSION_SECRET, (err, user) => {
      if (err) {
        return res.status(401).json({
          "status": "error",
          "message": "Forbidden",
        });
      }

      req.user = user;

      next();
    });
  }
}