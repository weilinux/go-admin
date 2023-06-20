export function checkLogin(username,password) {
    if (username === '' || username === undefined) {
        return false
    }
    return !(password === '' || password === undefined);
}

export function checkContent(title,categoryNames,content,cover_img) {
    if (title === '' || title === undefined) {
        alert('标题不能空')
        return false
    }
    if (categoryNames.length===0) {
        alert('分类不能空')
        return false
    }
    if (content === '' || content === undefined) {
        alert('内容不能空')
        return false
    }
    if (cover_img === '' || cover_img === undefined) {
        alert('封面不能空')
        return false
    }
    return true
}