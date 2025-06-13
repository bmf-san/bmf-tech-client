document.addEventListener("DOMContentLoaded", function() {
document.querySelectorAll('pre > code.mermaid').forEach(function(codeBlock) {
	const pre = codeBlock.parentElement;
	const mermaidDiv = document.createElement('div');
	mermaidDiv.className = 'mermaid';
	mermaidDiv.textContent = codeBlock.textContent;
	pre.parentNode.replaceChild(mermaidDiv, pre);
});
mermaid.initialize({ startOnLoad: true });
});