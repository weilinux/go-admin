import axios from 'axios';

// const API_URL = 'http://localhost:9550/api/';
const API_URL = 'http://localhost:9550/';

class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'login', {
                UserName: user.username,
                Password: user.password
            },{
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                }
            })
            .then(response => {
                if (response.data.Token) {
                    // localStorage.setItem('user', JSON.stringify(response.data));
                    localStorage.setItem('Token', JSON.stringify(response.data));
                }
                return response.data;
            });
    }

    logout() {
        localStorage.removeItem('Token');
    }


    register(user) {
        return axios.post(API_URL + 'signup', {
            UserName: user.username,
            // email: user.email,
            Password: user.password
        }, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        });
    }
}

export default new AuthService();