

var mongoose = require('mongoose');
var schema = mongoose.Schema;

var orderSchema = new schema({
	customerName : { type : String },
	location : { type : String },
	items: [{
	  	name     : { type : String },
	  	quantity : { type : Number },
	  	milk     : { type : String },
	  	size     : { type : String }
    }],
    status : {
      type    : String,
      enum    : ['Placed', 'Cancelled', 'Paid', 'Preparing', 'Served', 'Collected'],
      default : 'Placed'
    },
    message : { type : String}
});

module.exports = mongoose.model('orders', orderSchema);