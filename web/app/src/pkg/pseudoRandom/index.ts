export function PseudoRandom(length:number = 2) {
  return Math.random().toString(16).substr(2, length)
}
export default PseudoRandom