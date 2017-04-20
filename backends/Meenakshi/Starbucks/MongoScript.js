
db.orders.insert([
    {"location" : "Take-Out",
     "orderItems" : { "name" : "Cafe-Latte" , "quantity" : 1 , "milk" : "Whole", "size" : "Medium" } }
]);

db.orders.find();
    
db.orders.findOneAndDelete();
     
db.orders.insert({  
    customerName : "Meenakshi",
    location : "Take-Out",
    items : [{
        name     : "Tea",
	quantity : 1,
	milk     : "Organic",
	size     : "Large"
    }]
});