const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join(__dirname, 'home-site')));

// homepage
app.get('/', function (req, res) {
    res.sendFile(path.join(__dirname, 'home-site', 'index.html'));
});


console.log("Listening on port 3001")
app.listen(3001);
