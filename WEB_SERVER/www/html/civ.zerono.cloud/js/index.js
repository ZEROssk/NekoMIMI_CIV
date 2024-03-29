'usestrict';

let path = location.pathname + location.search;
let api = "https://dev.zerono.cloud/api/v1/twimg"

document.addEventListener("DOMContentLoaded", function() {
	if (path == "/") {
		addHome();
	} else if (location.pathname == "/original") {
		addOriginalImg(path);
	} else if (location.pathname == "/search") {
		addSearchThumbnailImg(path);
	} else if (location.pathname == "/upload") {
		addUpload();
		uploadPageFunc();
		preview.style = "display: none"
	} else {
		addThumbnailImg(path);
	}
});

function requestAjax(endpoint, callback) {
	let xhr = new XMLHttpRequest();
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
	if (window.event.keyCode == 13) {
		let searchId = document.getElementById("input-keyword").value;
		window.location.href = `/search?tid=${searchId}&get=50`;
	}
}

function changeImgSize(size) {
	let imgCon = document.getElementById(`img-container`);
	switch(size) {
		case 'small':
			imgCon.className = "img-size-small";
			break;
		case 'medium':
			imgCon.className = "img-size-medium";
			break;
		case 'large':
			imgCon.className = "img-size-large";
			break;
	}
}

function addChangeImgSizeButton(id, p, nA) {
	let pa = location.pathname;
	if (id == "") {
		u = `${pa}?p=${p}&get=${nA}&s=`
	} else {
		u = `${pa}?tid=${id}&p=${p}&get=${nA}&s=`
	}

	let change_imgSize =
		'<div id="cis-bt-menu" class="dropdown">'+
			'<div class="menu-root">ImageSize<div id="icon" class="fa">&#xf078</div></div>'+
			'<div id="cis-bt-content" class="dropdown-content">'+
				`<a href="${u}small">Small</a>`+
				`<a href="${u}medium">Medium</a>`+
				`<a id="last-menu" href="${u}large">Large</a>`+
			'</div>'+
		'</div>'
	;
	document.getElementById('change-button-container').insertAdjacentHTML('beforeend', change_imgSize);
}

function addNAcquiredButton(id, p, s) {
	let pa = location.pathname;
	if (id == "") {
		u = `${pa}?p=${p}&s=${s}&get=`
	} else {
		u = `${pa}?tid=${id}&p=${p}&s=${s}&get=`
	}

	let change_nAcquired =
		'<div id="na-bt-menu" class="dropdown">'+
			'<div class="menu-root">NumberAcquired<div id="icon" class="fa">&#xf078</div></div>'+
			'<div id="na-bt-content" class="dropdown-content">'+
				`<a href="${u}50">50</a>`+
				`<a href="${u}100">100</a>`+
				`<a id="last-menu" href="${u}150">150</a>`+
			'</div>'+
		'</div>'
	;
	document.getElementById('change-button-container').insertAdjacentHTML('beforeend', change_nAcquired);
}

function addPagination(p, limit, numA, s) {
	let pa = location.pathname;
	let pnnContainer =
		'<div id="pnn-container"></div>'
	;
	document.getElementById('media').insertAdjacentHTML('afterend', pnnContainer);

	let pnnContent = document.getElementById('pnn-container');

	if (`${p-1}` != 0) {
		let back =
			`<a id="pnn-back" class="fa pnn-button" href="${pa}?p=${p-1}&get=${numA}&s=${s}">&#xf137</a>`
		;
		pnnContent.insertAdjacentHTML('afterbegin', back);
	
	}

	for(let i=4; i > 0; i--) {
		let pNumber = p-i;
		if (pNumber > 0) {
			let pnnNumber =
				`<a id="pnn-number" href="${pa}?p=${pNumber}&get=${numA}&s=${s}">${pNumber}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', pnnNumber);
		} else {
			continue;
		}
	}

	for(let i=0; i < 5; i++) {
		let pNumber = p+i;
		if (pNumber == p) {
			let nowP =
				`<a id="pnn-now">${p}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', nowP);
		} else if (pNumber <= limit) {
			let pnnNumber =
				`<a id="pnn-number" href="${pa}?p=${pNumber}&get=${numA}&s=${s}">${pNumber}</a>`
			;
			pnnContent.insertAdjacentHTML('beforeend', pnnNumber);
		} else {
			break;
		}
	}

	if (`${p+1}` <= limit) {
		let next =
			`<a id="pnn-next" class="fa pnn-button" href="${pa}?p=${p+1}&get=${numA}&s=${s}">&#xf138</a>`
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
			'<li><a href="/thumbnail?p=1&get=50">Thumbnail page</a></li>'+
			'<li><a href="/upload">Upload page</a></li>'+
		'<ul>'
	;

	home.insertAdjacentHTML('beforeend', homeContent);
}

function addUpload() {
	let media = document.getElementById('media');
	media.textContent = null;

	let home = document.createElement('div');
	home.setAttribute('id', 'home');
	media.appendChild(home);

	let homeContent =
		'<p>Upload</p>'+
		'<input id="upload-input" type="file" multiple="multiple" accept="image/*">'+
		'<div id="upload-area">'+
			'<p>Drop or Click to Select Upload Image.</p>'+
			'<p>Maximum Upload Size is 5MB</p>'+
			'<div id="preview"></div>'+
		'</div>'+
		'<div id="file-data"><div id="file-info"></div></div>'+
		'<div id="upload-button">'+
			'<button id="post" class="fa" type="button">&#xf382</button>'+
		'</div>'
	;

	home.insertAdjacentHTML('beforeend', homeContent);
}

function fileSizeFormat(size) {
	let i = -1;
	let byteUnits = [' KB', ' MB', ' GB', ' TB', 'PB', 'EB', 'ZB', 'YB'];

	if (size < 1024) {
		return size + ' B';
	}

	do {
		size = size / 1024;
		i++;
	} while (size > 1024);

	let fs = Math.max(size, 0.1).toFixed(1);

	return fs + byteUnits[i];
}

function uploadPageFunc() {
	let uploadImgs = new FormData();
	let inputData = document.getElementById("upload-input");
	let uploadArea = document.getElementById("upload-area");
	let preview = document.getElementById("preview");
	let fileInfo = document.getElementById("file-info");
	let fileSize = 0;
	let fl = 0;

	uploadArea.addEventListener("dragover", function(e) {
		e.stopPropagation();
		e.preventDefault();

		this.style.background = "#d3d3d3";
	}, false);

	uploadArea.addEventListener("dragleave", function(e) {
		e.stopPropagation();
		e.preventDefault();

		this.style.background = "#ffffff";
	}, false);

	uploadArea.addEventListener("click", function(e) {
		e.stopPropagation();
		e.preventDefault();

		inputData.click();
		inputData.onchange = function() {
			let files = inputData.files;
			fl += files.length;
			for (let i = 0; i < files.length; i++) {
				(function() {
					let fr = new FileReader();
					fr.onload = function() {
						uploadArea.getElementsByTagName('p')[0].style = "display: none";
						uploadArea.getElementsByTagName('p')[1].style = "display: none";
						preview.style = "display: true"

						let div = document.createElement('div');
						div.setAttribute('class', "content-preview");

						let img = document.createElement('img');
						img.setAttribute('src', fr.result);
						div.appendChild(img);

						preview.appendChild(div);
					};
					fr.readAsDataURL(files[i]);
				})();
				fileInfo.textContent = null;

				uploadImgs.append("file", files[i]);
				fileSize += files[i].size;

				let fs = fileSizeFormat(fileSize);
				let fi = document.createTextNode(`${fs},  ${fl} File`);
				fileInfo.appendChild(fi);
			}
		};
	}, false);

	uploadArea.addEventListener("drop", function(e) {
		e.stopPropagation();
		e.preventDefault();

		this.style.background = "#ffffff";

		let files = e.dataTransfer.files;
		fl += files.length;
		for (let i = 0; i < files.length; i++) {
			(function() {
				let fr = new FileReader();
				fr.onload = function() {
					uploadArea.getElementsByTagName('p')[0].style = "display: none";
					uploadArea.getElementsByTagName('p')[1].style = "display: none";
					preview.style = "display: true"

					let div = document.createElement('div');
					div.setAttribute('class', "content-preview");

					let img = document.createElement('img');
					img.setAttribute('src', fr.result);
					div.appendChild(img);

					preview.appendChild(div);
				};
				fr.readAsDataURL(files[i]);
			})();
			fileInfo.textContent = null;

			uploadImgs.append("file", files[i]);
			fileSize += files[i].size;

			let fs = fileSizeFormat(fileSize);
			let fi = document.createTextNode(`${fs},  ${fl} File`);
			fileInfo.appendChild(fi);
		}
	}, false);

	let postButton = document.getElementById("post");
	postButton.addEventListener("click", function() {
		let request = new XMLHttpRequest();
		request.open("POST", `${api}/upload`);
		request.send(uploadImgs);
		preview.textContent = null;
		preview.style = "display: none"
		uploadArea.getElementsByTagName('p')[0].style = "display: true";
		uploadArea.getElementsByTagName('p')[1].style = "display: true";
		uploadImgs = new FormData();
		fileSize = 0;
		fl = 0;
	});
}

function addOriginalImg(v) {
	let orImg = document.getElementById('img-container');
	orImg.id = "original-img-container";

	let hisback =
		'<button id="hisback" class="fa" onclick="history.back()">&#xf137</button>'
	;

	orImg.parentNode.insertAdjacentHTML('beforebegin', hisback)

	requestAjax(`${api}${v}`, function(response){

		let img = response.Image.FileName
		let original =
			'<div class="content-original" target="_blank">'+
				`<img class="original-img" src="../IMAGE/ORIGIN/${img}"/>`+
			'</div>'
		;
		orImg.insertAdjacentHTML('beforeend', original);
	});
}

function addThumbnailImg(v) {
	requestAjax(`${api}${v}`, function(response){
		let page = response.PageNumber;
		let limit = response.PageLimit;
		let nA = response.NumberAcquired;
		let imgList = response.Thumbnail;
		let imgSize = response.ImgSize;
		let tid = "";

		changeImgSize(imgSize);
		addChangeImgSizeButton(tid, page, nA);
		addNAcquiredButton(tid, page, imgSize);
		addPagination(page, limit, nA, imgSize);

		for(let i=0; i < imgList.length; i++) {
			let img = imgList[i].FileName
			let tid = imgList[i].TwitterID
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<a href="/original?tid=${tid}&fname=${img}">`+
						`<img class="thumbnail-img" src="../IMAGE/THUMBNAIL/${img}"/>`+
					'</a>'+
					`<a id="twi-id-link" href="/search?tid=${tid}&p=1&get=${nA}">`+
						`<span id="twi-id-hover">${tid}</span>`+
					'</a>'+
				'</div>'
			;
			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

function addSearchThumbnailImg(v) {
	requestAjax(`${api}${v}`, function(response){
		let page = response.PageNumber;
		let limit = response.PageLimit;
		let nA = response.NumberAcquired;
		let imgList = response.Thumbnail;
		let imgSize = response.ImgSize;
		let tid = response.TwitterID;

		let hisback =
			'<button id="hisback" class="fa" onclick="history.back()">&#xf137</button>'
		;
		document.getElementById('media').insertAdjacentHTML('beforebegin', hisback);

		let displayID =
			'<div id="select-ID">'+
				`<a href="https://twitter.com/${tid}" target="_blank">@${tid}</a>`+
			'</div>'
		;
		document.getElementById('menu-container').insertAdjacentHTML('afterbegin', displayID);

		changeImgSize(imgSize);
		addChangeImgSizeButton(tid, page, nA);
		addNAcquiredButton(tid, page, imgSize);
		addPagination(page, limit, nA, imgSize);
	
		for(let i=0; i < imgList.length; i++) {
			let img = imgList[i].FileName
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<a href="/original?tid=${tid}&fname=${img}">`+
						`<img class="thumbnail-img" src="../IMAGE/THUMBNAIL/${img}"/>`+
					'</a>'+
				'</div>'
			;
			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

