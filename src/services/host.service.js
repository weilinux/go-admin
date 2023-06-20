import axios from "@/router/carry";

class HostService {
    getUserHosts(page, limit) {
        return axios.get('/api/users/hosts', {params: {page: page, limit: limit}})
    }

    delUserHost(id) {
        return axios.delete("/api/hosts/" + id)
    }
}

export default new HostService();