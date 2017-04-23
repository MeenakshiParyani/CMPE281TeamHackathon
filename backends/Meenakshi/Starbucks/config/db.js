 

 module.exports = {
        url : 'mongodb://ec2-52-52-195-15.us-west-1.compute.amazonaws.com:27017/orders',
        options : {
		  db: { native_parser: true },
		  server: { poolSize: 5 },
		  replset: { rs_name: 'starbucks-replica-set' }
		}
 }