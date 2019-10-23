'usestrict';

(function() {
	console.log(location.pathname)
	console.log(location.search)
})();

document.addEventListener("DOMContentLoaded", function() {
	addThumbnail();
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

function addThumbnail() {
	const pNum = 1;
	requestAjax(`http://civ.zerono.server-on.net:8888/api/v1/twimg/thumbnail?p=${pNum}&get=30`, function(response){
	
		for(let i=0; i<response.Thumbnail.length; i++) {
			let img = response.Thumbnail[i].FileName
			let thumbnail =
				'<div class="content-thumbnail" target="_blank">'+
					`<img class="thumbnail-img" onclick="open_OriginalImg()" src="../IMAGE/Twitter/${img}"/>`+
				'</div>'
			;

			document.getElementById('img-container').insertAdjacentHTML('beforeend', thumbnail);
		}
	});
}

