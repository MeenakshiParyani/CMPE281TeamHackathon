/**
 * http://usejsdoc.org/
 */
var ejs = require("ejs");

exports.homePage = function(req, res) {

	ejs.renderFile('./views/home.ejs', function(err, result) {
		// render on success
		if (!err) {
			res.end(result);
		}
		// render or error
		else {
			res.end('An error occurred');
			console.log(err);
		}
	});
};
