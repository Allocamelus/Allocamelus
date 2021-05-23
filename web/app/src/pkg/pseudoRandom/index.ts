export function PseudoRandom(length = 2): string {
  return Math.random().toString(16).substr(2, length)
}
export default PseudoRandom