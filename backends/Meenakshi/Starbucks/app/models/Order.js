

var mongoose = require('mongoose');
var schema = mongoose.Schema;

var orderSchema = new schema({
	orderItem: [{
	  	name     : { type : String},
	  	quantity : { type : Number},
	  	milk     : { type : String},
	  	size     : { type : String}
    }],
});


module.exports = mongoose.model('orders', orderSchema);