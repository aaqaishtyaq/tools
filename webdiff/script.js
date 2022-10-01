require.config({
  paths: {
    vs: "https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.33.0/min/vs/",
  },
});

require(["vs/editor/editor.main"], () => {
  const editor = monaco.editor.createDiffEditor(
    document.getElementById("container"),
    { theme: "vs-dark", originalEditable: true }
  );

  editor.setModel({
    original: monaco.editor.createModel("", "text/plain"),
    modified: monaco.editor.createModel("", "text/plain"),
  });

  document.querySelector(".inline").addEventListener("change", (e) => {
    editor.updateOptions({ renderSideBySide: !e.target.checked });
  });
});
