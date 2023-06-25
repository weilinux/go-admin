import axios from "@/router/carry";

class HostService {
  getUserHosts(page: number, limit: number) {
    return axios.get("/api/users/hosts", {
      params: { page: page, limit: limit },
    });
  }

  delUserHost(id: number) {
    return axios.delete("/api/hosts/" + id);
  }
}

export default new HostService();
