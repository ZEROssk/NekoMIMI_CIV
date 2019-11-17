'usestrict';

let path = location.pathname + location.search
var searchId
var nAcquired;

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
		window.location.href = `/search?tid=${searchId}&get=${nAcquired}`;
	}
}

function updateNAcquired(n) {
	nAcquired = n;
	console.log(n);
	console.log(nAcquired);
}

function addPagination(p, limit) {
	let pnnContainer =
		'<div id="pnn-container"></div>'
	;
	document.getElementById('media').insertAdjacentHTML('afterend', pnnContainer);

	let pnnContent = document.getElementById('pnn-container');

	if(`${p-1}` != 0) {
		let back =
			`<a id="pnn-back" class="fa pnn-button" href="/thumbnail?p=${p-1}&get=${nAcquired}">&#xf137</a>`
		;
		pnnContent.insertAdjacentHTML('afterbegin', back);
	
	}

	for(let i=4; i > 0; i--) {
		let pNumber = p-i;
		if(pNumber > 0) {
			let pnnNumber =
				`<a id="pnn-number" href="/thumbnail?p=${pNumber}&get=${nAcquired}">${pNumber}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', pnnNumber);
		} else {
			continue;
		}
	}

	for(let i=0; i < 5; i++) {
		let pNumber = p+i;
		if(pNumber == p) {
			let nowP =
				`<a id="pnn-now">${p}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', nowP);
		} else if(pNumber <= limit) {
			let pnnNumber =
				`<a id="pnn-number" href="/thumbnail?p=${pNumber}&get=${nAcquired}">${pNumber}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', pnnNumber);
		} else {
			break;
		}
	}

	if(`${p+1}` <= limit) {
		let next =
			`<a id="pnn-next" class="fa pnn-button" href="/thumbnail?p=${p+1}&get=${nAcquired}">&#xf138</a>`
		;
		pnnContent.insertAdjacentHTML('beforeend', next);
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
			`<li><a href="/thumbnail?p=1&get=${nAcquired}">Thumbnail page</a></li>`+
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

		let change_nAcquired =
			'<div id="nAcquired-bt-container">'+
				`<a href="/thumbnail?p=${page}&get=50">`+
					'<button id="nAcquired-bt" onclick="updateNAcquired(50)">50</button>'+
				'</a>'+
				`<a href="/thumbnail?p=${page}&get=100">`+
					'<button id="nAcquired-bt" onclick="updateNAcquired(100)">100</button>'+
				'</a>'+
				`<a href="/thumbnail?p=${page}&get=150">`+
					'<button id="nAcquired-bt" onclick="updateNAcquired(150)">150</button>'+
				'</a>'+
			'</div>'
		;

		document.getElementById('media').insertAdjacentHTML('beforebegin', change_nAcquired);

		for(let i=0; i < response.Thumbnail.length; i++) {
			let img = response.Thumbnail[i].FileName
			let tid = response.Thumbnail[i].TwitterID
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<a href="/original?tid=${tid}&fname=${img}">`+
						`<img class="thumbnail-img" src="../IMAGE/Twitter/${img}"/>`+
					'</a>'+
					`<a id="twi-id-link" href="/search?tid=${tid}&get=${nAcquired}">`+
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
				`<a href="https://twitter.com/${tid}" target="_blank">@${tid}</a>`+
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

