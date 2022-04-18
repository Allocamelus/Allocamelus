import DOMPurify from "dompurify";

export function sanitize(html: string | Node) {
  return DOMPurify.sanitize(html);
}

export function textContent(html: string | Node): string {
  const div = document.createElement("div");
  div.innerHTML = sanitize(html);
  return div.textContent || "";
}

export default sanitize;
