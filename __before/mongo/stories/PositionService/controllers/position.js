var Position = require("../models/position")

// Create
exports.createPosition = function(req, res, next) {
  var position = new Position(
    {
      account: req.body.account,
      stock: req.body.stock,
      quantity: req.body.quantity,
      price: req.body.price
    }
  )
  position.save(function(err) {
    if (err) {
      return next(err)
    }
    res.send('仓位纪录添加成功')
  })
}

// Read
exports.queryPosition = function(req, res, next) {
  Position.find({
    account: req.params.account
  }, function(err, position) {
    if (err) {
      return next(err)
    }
    res.send(position)
  })
}

//Update
exports.updatePosition = function(req, res, next) {
  Position.findByIdAndUpdate(req.params.id, {
    $set: req.body
  }, function(err) {
    if (err) {
      return next(err)
    }
    res.send('仓位纪录更新成功')
  })
}

//Delete
exports.deletePosition = function(req, res, next) {
  Position.findByIdAndRemove(req.params.id, function(err) {
    if (err) {
      return next(err)
    }
    res.send('仓位纪录删除成功')
  })
}
