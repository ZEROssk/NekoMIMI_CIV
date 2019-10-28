'usestrict';

let path = location.pathname + location.search
var searchId

document.addEventListener("DOMContentLoaded", function() {
	if (location.pathname == "/original") {
		addOriginalImg(path);
	} else if (path == "/") {
		addHome()
	} else {
		addThumbnailImg(path);
	}
});

function requestAjax(endpoint, callback) {
	var xhr = new XMLHttpRequest();
	xhr.onreadystatechange = function() {
		if (this.readyState==4 && this.status==200) {
			callback(this.response);
		}
	};
	xhr.responseType = 'json';
	xhr.open('GET',endpoint,true);
	xhr.send();
};

function addHome() {
	let media = document.getElementById('media');
	media.textContent = null;

	let home = document.createElement('div');
	home.setAttribute('id', 'home');
	media.appendChild(home);

	let homeContent =
		'<p>HOME</p>'+
		'<ul>'+
			'<li><a href="/thumbnail?p=1">Thumbnail page</a></li>'+
		'<ul>'
	;

	home.insertAdjacentHTML('beforeend', homeContent);
}

function onSearchButtonClick() {
	if(window.event.keyCode == 13) {
		searchId = document.getElementById("input-keyword").value;
		window.location.href = `/search?tid=${searchId}`;
	}
}

function addOriginalImg(url) {
	let originalImg = document.getElementById('img-container');
	originalImg.id = "original-img-container";

	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg${url}`, function(response){

		let img = response.Image.FileName
		let original =
			'<div class="content-original" target="_blank">'+
				`<img class="original-img" src="../IMAGE/Twitter/${img}"/>`+
			'</div>'
		;

		document.getElementById('original-img-container').insertAdjacentHTML('beforeend', original);
	});
}

function addThumbnailImg(url) {
	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg${url}`, function(response){
	
		for(let i=0; i < response.Thumbnail.length; i++) {
			let img = response.Thumbnail[i].FileName
			let tid = response.Thumbnail[i].TwitterID
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<a href="/original?tid=${tid}&fname=${img}">`+
						`<img class="thumbnail-img" src="../IMAGE/Twitter/${img}"/>`+
					'</a>'+
				'</div>'
			;

			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

