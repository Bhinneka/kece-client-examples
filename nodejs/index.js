'use strict';

const net = require('net');

var client = new net.Socket();
client.connect(8000, '127.0.0.1', () => {
	console.log('Connected');
	client.write('GET 1\r\n');
});

client.on('data', (data) => {
	console.log('Received: ' + data);
	client.destroy(); // kill client after server's response
});

client.on('close', () => {
	console.log('Connection closed');
});
