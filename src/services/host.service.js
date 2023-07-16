/*
 * 主机管理模块
 */
// import axios from "@/router/carry";
import request from '@/router/carryEns';

class HostService {
    //查询
    getUserHosts(page, limit) {
        return request({
            url: '/api/users/hosts?' + 'page=' + page + '&limit=' + limit,
            method: 'get',
        })
    }
    //删除
    delUserHost(id) {
        return request({
            url: '/api/hosts/' + id,
            method: 'delete',
        })
    }
    //详情
    getHost(id) {
        return request({
            url: '/api/hosts/' + id,
            method: 'get',
        })
    }
    //增加
    addHost(host) {
        return request({
            url: '/api/hosts',
            method: 'post',
            headers: { 'content-type': 'application/x-www-form-urlencoded' },
            data: {
                HostName: host.HostName,
                HostIP: host.HostIP,
                HostPort: host.HostPort
            }
        })
    }
    //编辑
    updateHost(id, host) {
        return request({
            url: '/api/hosts/' + id,
            method: 'put',
            headers: { 'content-type': 'application/x-www-form-urlencoded' },
            data: {
                HostName: host.HostName,
                HostIP: host.HostIP,
                HostPort: host.HostPort
            }
        })
    }
    //搜索
    searchHosts(host, page, limit) {
        return request({
            url: '/api/users/searchhost?',
            method: 'get',
            params: {
                host: host,
                page: page,
                limit: limit
            }
        })
    }

}

export default new HostService();