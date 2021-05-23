export default (html: string) => {
  let div = document.createElement("div");
  div.innerHTML = html;
  return div.textContent || "";
}