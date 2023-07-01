/*
 * 用户管理模块
 */
import request from '@/router/carryEns';
import axios from '@/router/carry';

class UserService {
    //注册
    signupUser(username, password) {
        return request({
            url: '/api/signup',
            method: 'post',
            headers: { 'content-type': 'application/x-www-form-urlencoded' },
            data: {
                UserName: username,
                Password: password,
            }
        })
    }

    //查询
    getUserList() {
        return request({
            url: '/api/users',
            method: 'get',
        })
    }

    //详情
    showUser(id) {
        return request({
            url: '/api/users/' + id,
            method: 'get',
        })
    }

    //增加
    addUser(id, username, password) {
        return request({
            url: '/api/users',
            method: 'post',
            headers: { 'content-type': 'application/x-www-form-urlencoded' },
            data: {
                ID: id,
                UserName: username,
                Password: password,
            }
        })
    }

    //编辑
    updateUser(id, username, password) {
        return request({
            url: '/api/users/' + id,
            method: 'put',
            headers: { 'content-type': 'application/x-www-form-urlencoded' },
            data: {
                UserName: username,
                Password: password,
            }
        })
    }

    //删除
    //TODO: 注意这里, request封装获取数据后,前台删除用户之后,页面没有实时更新成功
    delUser(id) {
        return axios({
            url: '/api/users/' + id,
            method: 'delete',
        })
    }
}

export default new UserService();