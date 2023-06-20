export default function authHeader() {
    let user = JSON.parse(localStorage.getItem('Token'));

    if (user && user.Token) {
        return { Authorization: 'Bearer ' + user.Token }; // for Spring Boot back-end
        // return { 'x-access-token': user.accessToken };       // for Node.js Express back-end
    } else {
        return {};
    }
}
