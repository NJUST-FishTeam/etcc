var express = require('express');
var app = express();

app.use('/static', express.static(__dirname + '/static'));

app.get('/', function(req, res){
    res.sendfile(__dirname+'/index.html');
});
var server = app.listen(8083, function() {
    console.log('Listening on port %d', server.address().port);
});