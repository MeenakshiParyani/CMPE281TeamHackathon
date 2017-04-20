// modules =================================================
var express        = require('express');
var app            = express();
var bodyParser     = require('body-parser');
var methodOverride = require('method-override');
var mongoose       = require('mongoose');
// configuration ===========================================
    
//mongo db config
var dbConfig = require('./config/db');
mongoose.connect(dbConfig.url, dbConfig.options);
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
db.on('connected', function(){
	console.log('conn ' + db.readyState);
});
db.once('open', function() {
  console.log('we\'re open!');
});
//console.log(db);

var port = 8080;

app.use(bodyParser.json()); 

app.use(bodyParser.urlencoded({ extended: true })); 

app.use(methodOverride('X-HTTP-Method-Override')); 

var routes = require('./app/routes/routes');
app.get('/api/*', routes);

app.listen(port);               
                
console.log('App running on port ' + port);

// expose app           
exports = module.exports = app;                         
