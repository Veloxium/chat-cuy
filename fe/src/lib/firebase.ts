import { initializeApp } from "firebase/app";

import {
   getAuth,
   GoogleAuthProvider,
   FacebookAuthProvider,
   signInWithPopup,
} from "firebase/auth";

interface ConfigType {
   apiKey: string;
   authDomain: string;
   projectId: string;
   storageBucket: string;
   messagingSenderId: string;
   appId: string;
}

const firebaseConfig = {
   apiKey: "AIzaSyCHqW9mwctn8e-iAFdWM-n_am6KtBwzRpw",
   authDomain: "chat-cuyy.firebaseapp.com",
   projectId: "chat-cuyy",
   storageBucket: "chat-cuyy.firebasestorage.app",
   messagingSenderId: "16425233438",
   appId: "1:16425233438:web:1e313d1f80bc7bf3be920f",
} satisfies ConfigType;

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
const googleProvider = new GoogleAuthProvider();
const facebookProvider = new FacebookAuthProvider();

export { auth, googleProvider, signInWithPopup, facebookProvider };
