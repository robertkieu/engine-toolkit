'use strict'

const express = require('express');
// Create Express app
const app = express();
var multiparty = require('multiparty');

const process = (req, res) => {
	var form = new multiparty.Form();

    form.parse(req, function(err, fields, files) {
		console.log('fields', JSON.stringify(fields));
    });
	const event = {
		startOffsetMs: 1000,
		endOffsetMs: 2000,
		chunkIndex: 1,
		tdoId: 123,

	}
	var series = [];
	const boundingPoly = [
		{x: 0.1,
			y: 0.2},
		{x: 0.1,
			y: 0.2},
		{x: 0.1,
			y: 0.2},
		{x: 0.1,
			y: 0.2}
	];
	const face = {
		startTimeMs: event.startOffsetMs,
		stopTimeMs: event.endOffsetMs,
		entityId: '',
		libraryId: '',
		object: {boundingPoly: boundingPoly,
			confidence: 1,
			label: 'EdgeTestFace' + event.chunkIndex,
			type: 'face',
			uri: '/media-streamer/image/' + event.tdoId + '/2017-12-28T15:53:00?x1=0.78&y1=0.12&x2=0.99&y2=0.12&x3=0.99&y3=0.44&x4=0.78&y4=0.44'
		}
	};

	series.push(face);
	res.json({series: series});
};

// 
app.get('/readyz', (req, res) => res.send('Engine ready!'));
app.post('/process', process);

// Start the Express server
app.listen(8080, '0.0.0.0' , () => console.log('Server running on 0.0.0.0:8080!')) 

