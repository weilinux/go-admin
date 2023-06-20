import cookies from 'js-cookie'

export function _set(name, value, config) {
    cookies.set(name, value, config)
    return value
}

export function _get(name) {
    return cookies.get(name)
}

export function _remove(name, config) {
    cookies.remove(name, config)
}

export function _getJSON(name) {
    // 没有看到有下面这个函数
    // cookies.getJSON(name)
}