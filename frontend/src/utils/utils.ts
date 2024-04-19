
/**
 * 格式化數字，給數字加上千分隔符
 * @param num 要格式化的數字
 * @returns 格式化後的字符串
 */
export function formatNumber(num: number) {
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}
