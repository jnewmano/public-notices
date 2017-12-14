
var mymap = L.map('mapdiv').setView([40.415, -111.87], 13);


var tileLayerID = '';
var accessToken = '';
var tileSource = 'https://cartodb-basemaps-{s}.global.ssl.fastly.net/light_all/{z}/{x}/{y}.png';

var attribution = 'WikiMedia <a href="http://openstreetmap.org">OpenStreetMap</a> contributors, <a href="http://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>';

L.tileLayer(tileSource, {
    attribution: attribution,
    maxZoom: 18,
    id: tileLayerID,
    accessToken: accessToken
}).addTo(mymap);


var url = './upcomingActions?entity=abc';
alert(url);

fetch(url)
.then(res => res.json())
.then((out) => {
  console.log('Checkout this JSON! ', out);
  drawNotices(out)
})
.catch(err => { throw err });

function drawNotices(notices) {
	notices.Notices.forEach(function(v) {
		var lat = v.Location.Latitude;
		var long = v.Location.Longitude;

		var circle = L.circle([lat, long], {
		    color: 'green',
		    fillColor: 'green',
		    fillOpacity: 0.5,
		    radius: 300
		}).addTo(mymap);
	});
}
