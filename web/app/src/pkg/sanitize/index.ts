export default (html: string) => {
  var div = document.createElement("div");
  div.innerHTML = html;
  return div.textContent || "";
}