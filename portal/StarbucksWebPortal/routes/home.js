/**
 * http://usejsdoc.org/
 */
var ejs = require("ejs");

exports.homePage = function(req, res) {

	res.render('home.ejs');
};
