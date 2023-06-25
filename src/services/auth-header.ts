// export default function authHeader() {
export default function authHeader() {
  // const user = JSON.parse(localStorage.getItem("Token"));
  const token = localStorage.getItem("Token")
  const user = token ? JSON.parse(token): null;

  if (user && user.Token) {
    return { Authorization: "Bearer " + user.Token }; // for Spring Boot back-end
    // return { 'x-access-token': user.accessToken };       // for Node.js Express back-end
  } else {
    return {};
  }
}
