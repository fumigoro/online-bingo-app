<template>
  <div>
    <button @click="login">ログイン</button>
  </div>
</template>
<script>
// Firebase App (the core Firebase SDK) is always required and must be listed first
import firebase from "firebase/app";
// If you are using v7 or any earlier version of the JS SDK, you should import firebase using namespace import
// import * as firebase from "firebase/app"

// If you enabled Analytics in your project, add the Firebase SDK for Analytics
import "firebase/analytics";

// Add the Firebase products that you want to use
import "firebase/auth";
import "firebase/firestore";

export default {

  methods: {
    login() {
      const provider = new firebase.auth.GoogleAuthProvider();
      provider.addScope("https://www.googleapis.com/auth/contacts.readonly");
      firebase.auth().signInWithRedirect(provider);
    },
  },
  mounted: async () => {
    const firebaseConfig = {
      apiKey: "AIzaSyCHPyHzZ14QUeRI-5AoBMeTx0FCVsiZjTY",
      authDomain: "online-bingo-52232.firebaseapp.com",
      projectId: "online-bingo-52232",
      storageBucket: "online-bingo-52232.appspot.com",
      messagingSenderId: "49688864208",
      appId: "1:49688864208:web:accd6a54a204b79e5b45d5",
    };
    // Initialize Firebase
    if (firebase.apps.length === 0) {
      firebase.initializeApp(firebaseConfig);
    }
  },
};
</script>