'usestrict'
let thumbnail_list = ['Twitter-1174215925614223360-hirame_sa-EEunKGoW4AAJVaB.jpg','Twitter-1174214910642991104-riosi_RRR-EEumO9zXYAAuPc8.jpg','Twitter-1173645323442372608-hiraga_matsuri-EEmgMhrU8AY2G91.jpg','Twitter-1173588320120557568-herumonnnulu-EElruLEU8AIdTnF.jpg','Twitter-1173243158664110086-TYBAyosizawa-EEgybnDU8AEuby2.jpg']

document.addEventListener("DOMContentLoaded", function() {
	let rootContainer = document.getElementById("root-container");

	let content = document.createElement("div");
	content.setAttribute('id', 'content');
	rootContainer.appendChild(content);

	let media = document.createElement("div");
	media.setAttribute('id', 'media');
	content.appendChild(media);

	let imgContainer = document.createElement("div");
	imgContainer.setAttribute('id', 'img-container');
	media.appendChild(imgContainer);

	addContent();
});

function requestAjax(endpoint, callback) {
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

function open_OriginalImg() {
	console.log("click!");
}

function addContent() {
	const pNum = 1;
	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg/page?p=${pNum}`, function(response){
	
		for(let i=0; i<response.Thumbnail.length; i++) {
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					'<img class="thumbnail-img" onclick="open_OriginalImg()" src="../IMAGE/Twitter/'+ response.Thumbnail[i].FileName +'"/>'+
				'</div>'
			;

			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

