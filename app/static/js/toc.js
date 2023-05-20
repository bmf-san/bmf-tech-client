const hTags = ['h1', 'h2', 'h3', 'h4', 'h5', 'h6'];
const tmpList = [];
let tocElement = document.getElementById('toc');
document.querySelector('.container-readable').querySelectorAll(hTags.join()).forEach((element, index) => {
	const headingIndex = hTags.indexOf(element.localName);
	if (headingIndex !== -1) {
		element.id = index + '-' + element.innerText;
		let li = document.createElement('li');
		let a = document.createElement('a');
		let ul = document.createElement(tocElement.tagName);
		a.href = '#' + element.id;
		a.title = element.innerHTML;
		a.innerHTML = element.innerHTML;
		li.appendChild(a);
		li.appendChild(ul);
		tmpList[index] = [headingIndex, li, ul];
	}
});

if (tmpList[0][0] !== 0) {
	let li = document.createElement('li');
	let ul = document.createElement(tocElement.tagName);
	li.appendChild(ul);
	tmpList.unshift([0, li, ul]);
}

for (let i = 1; i < tmpList.length; i++) {
	for (let j = i - 1; j >= 0; j--) {
	if (tmpList[i][0] > tmpList[j][0]) {
		let dif = tmpList[i][0] - tmpList[j][0];
		if (dif > 1) {
			for (let k = dif - 1; k > 0; k--) {
				let li = document.createElement('li');
				let ul = document.createElement(tocElement.tagName);
				ul.appendChild(tmpList[i][1]);
				li.appendChild(ul);
				tmpList[i][1] = li;
			}
		}
		tmpList[j][2].appendChild(tmpList[i][1]);
		break;
	}
	}
}

tmpList.forEach(function(data) {
	if (data[0] === 0) {
		tocElement.appendChild(data[1]);
	}
});

let tagList = tocElement.getElementsByTagName(tocElement.tagName);

for (let i = tagList.length - 1; i >= 0; i--) {
	if (tagList[i].innerHTML === '') {
		tagList[i].remove();
	}
}
