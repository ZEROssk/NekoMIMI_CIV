'usestrict';

let path = location.pathname + location.search
var searchId

document.addEventListener("DOMContentLoaded", function() {
	if (location.pathname == "/original") {
		addOriginalImg(path);
	} else if (location.pathname == "/search") {
		addSearchThumbnailImg(path);
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

function onSearchButtonClick() {
	if(window.event.keyCode == 13) {
		searchId = document.getElementById("input-keyword").value;
		window.location.href = `/search?tid=${searchId}`;
	}
}

function addPagination(p, limit) {
	let pnnContainer =
		'<div id="pnn-container"></div>'
	;
	document.getElementById('media').insertAdjacentHTML('afterend', pnnContainer);

	if(`${p-1}` != 0) {
		let back =
			`<a id="pnn-back" class="fa pnn-button" href="/thumbnail?p=${p-1}">&#xf137</a>`
		;
		document.getElementById('pnn-container').insertAdjacentHTML('afterbegin', back);
	
		let pnnNumber =
			`<a id="pnn-number">${p}</a>`
		;
		document.getElementById('pnn-back').insertAdjacentHTML('afterend', pnnNumber);
	} else {
		let pnnNumber =
			`<a id="pnn-number">${p}</a>`
		;
		document.getElementById('pnn-container').insertAdjacentHTML('afterbegin', pnnNumber);
	}

	if(`${p+1}` <= limit) {
		let next =
			`<a id="pnn-next" class="fa pnn-button" href="/thumbnail?p=${p+1}">&#xf138</a>`
		;
		document.getElementById('pnn-container').insertAdjacentHTML('beforeend', next);
	}
}

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

function addOriginalImg(v) {
	let orImg = document.getElementById('img-container');
	orImg.id = "original-img-container";

	let hisback =
		'<button id="hisback" class="fa" onclick="history.back()">&#xf137</button>'
	;

	orImg.parentNode.insertAdjacentHTML('beforebegin', hisback)

	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg${v}`, function(response){

		let img = response.Image.FileName
		let original =
			'<div class="content-original" target="_blank">'+
				`<img class="original-img" src="../IMAGE/Twitter/${img}"/>`+
			'</div>'
		;

		orImg.insertAdjacentHTML('beforeend', original);
	});
}

function addThumbnailImg(v) {
	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg${v}`, function(response){
		let page = response.PageNumber
		let limit = response.PageLimit
		addPagination(page, limit);

		for(let i=0; i < response.Thumbnail.length; i++) {
			let img = response.Thumbnail[i].FileName
			let tid = response.Thumbnail[i].TwitterID
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<a href="/original?tid=${tid}&fname=${img}">`+
						`<img class="thumbnail-img" src="../IMAGE/Twitter/${img}"/>`+
					'</a>'+
					`<a id="twi-id-link" href="/search?tid=${tid}">`+
						`<span id="twi-id-hover">${tid}</span>`+
					'</a>'+
				'</div>'
			;

			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

function addSearchThumbnailImg(v) {
	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg${v}`, function(response){
		let page = response.PageLimit
		let limit = response.PageLimit
		let tid = response.TwitterID
		let displayID =
			'<div id="select-ID">'+
				`<a href="https://twitter.com/${tid}" target="_blank">@${tid} ${plimit}</a>`+
			'</div>'
		;

		document.getElementById('media').insertAdjacentHTML('beforebegin', displayID);
		addPagination(page, limit);
	
		for(let i=0; i < response.Thumbnail.length; i++) {
			let img = response.Thumbnail[i].FileName
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

