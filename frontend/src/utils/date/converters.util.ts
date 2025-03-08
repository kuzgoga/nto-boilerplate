export const getFullTimestamp = (n: number): number => {
    const length = String(n).length
    let str = ''
    while (str.length + length < 13) {
        str += '0'
    }
    return parseInt(`${n}${str}`)
}
export const toDate = (n: number | Date) => {
    if (n instanceof Date) return n
    return new Date(getFullTimestamp(n))
}
export const toTimestamp = (d: Date) => d.getTime()

export const dateToTimestamp = (obj: any) => {
    if (typeof obj == 'object') {
        for (let key in obj) {
            if (obj[key] instanceof Date) {
                obj[key] = toTimestamp(obj[key])
            } else if (typeof obj[key] == 'object') {
                dateToTimestamp(obj[key])
            }
        }
    }
    return obj
}

export const timestampToDate = (obj: any) => {
    if (typeof obj == 'object') {
        for (let key in obj) {
            if (typeof obj[key] == 'number') {
                obj[key] = toDate(obj[key])
            } else if (typeof obj[key] == 'object') {
                timestampToDate(obj[key])
            }
        }
    }
    return obj
}