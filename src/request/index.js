import axios from 'axios';
import config from './config';


export default function request(options) {
   return new Promise((resolve, reject) => {
       const instance = axios.create({ ...config })
       }
   )
}


// request intercept

// response intercept