'usestrict'
let thumbnail_list = ['Twitter-1174215925614223360-hirame_sa-EEunKGoW4AAJVaB.jpg','Twitter-1174214910642991104-riosi_RRR-EEumO9zXYAAuPc8.jpg','Twitter-1173645323442372608-hiraga_matsuri-EEmgMhrU8AY2G91.jpg','Twitter-1173588320120557568-herumonnnulu-EElruLEU8AIdTnF.jpg','Twitter-1173243158664110086-TYBAyosizawa-EEgybnDU8AEuby2.jpg']

document.addEventListener("DOMContentLoaded", function() {
	let rootContainer = document.getElementById("root-container");

	addContent(rootContainer);

	const pNum = "1";


	var requestAjax = function(endpoint, callback) {
		var xhr = new XMLHttpRequest();
		xhr.onreadystatechange = function(){
			if (this.readyState==4 && this.status==200) {
				callback(this.response);
			}
		};
		xhr.responseType = 'json';
		xhr.open('GET',endpoint,true);
		xhr.send();
	};

	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg/page?p=${pNum}`, function(response){
		console.log(response);
	});

	// http://civ_qpi:5200/api/v1/twimg/page?p=${pNum}
});

function open_OriginalImg() {
	console.log("click!");
}

function addContent(rootC) {
	let content = document.createElement("div");
	content.setAttribute('id', 'content');
	rootC.appendChild(content);

	let media = document.createElement("div");
	media.setAttribute('id', 'media');
	content.appendChild(media);

	let imgContainer = document.createElement("div");
	imgContainer.setAttribute('id', 'img-container');
	media.appendChild(imgContainer);

	for(let i=0; i<thumbnail_list.length; i++) {
		let thumbnail =
			'<div class="content-thumbnail" target="_blank">'+
				'<img class="thumbnail-img" onclick="open_OriginalImg()" src="../IMAGE/'+ thumbnail_list[i] +'"/>'+
			'</div>'
		;

		document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
	}
}

