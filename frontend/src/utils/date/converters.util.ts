const getFullTimestamp = (n: number): number => {
    const length = String(n).length
    let str = ''
    while (str.length + length < 13) {
        str += '0'
    }
    return parseInt(`${n}${str}`)
}
export const toDate = (n: number) => new Date(getFullTimestamp(n))
export const toTimestamp = (d: Date) => d.getTime()
