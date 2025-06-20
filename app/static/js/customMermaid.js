document.addEventListener("DOMContentLoaded", function () {
  document.querySelectorAll("pre > code.mermaid").forEach(function (codeBlock) {
    const pre = codeBlock.parentElement;
    const mermaidDiv = document.createElement("div");
    mermaidDiv.className = "mermaid";

    const textarea = document.createElement("textarea");
    textarea.innerHTML = codeBlock.innerHTML;
    mermaidDiv.textContent = textarea.value;

    pre.parentNode.replaceChild(mermaidDiv, pre);
  });

  mermaid.initialize({ startOnLoad: false }); // 自動は無効化
  mermaid.init(undefined, document.querySelectorAll('.mermaid')); // 手動で再描画
});
