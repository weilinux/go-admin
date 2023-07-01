/*
 * 主机分配模块
 */
import request from '@/router/carryEns';

class HostAssignService {
    //增加
    getUnbindHosts(uid) {
        return request({
            url: '/api/users/' + uid + '/hosts',
            method: 'get',
        })
    }
    //编辑
    assignHost(uid, hid) {
        return request({
            url: '/api/hosts/assign',
            method: 'post',
            headers: { 'content-type': 'application/json' },
            data: {
                "user_id": uid,
                "host_id": hid,
            }
        })
    }
}

export default new HostAssignService();