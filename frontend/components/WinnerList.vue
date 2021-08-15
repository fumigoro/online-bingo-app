<template>
  <div>
    <h2>ビンゴした人</h2>
    <ul>
      <li v-for="winner in winners" :key="winner">
        <p>
          {{ winner.display_name }} {{ winner.studentId }}
          {{ winner.timestamp }}
        </p>
      </li>
    </ul>
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
  data() {
    return { winners: [] };
  },
  mounted: function () {
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
    const db = firebase.firestore();
    const me = this;
    firebase.auth().onAuthStateChanged(function (user) {
      if (!user) {
        me.displayFlag.login = true;
        me.displayFlag.main = false;
      }
      // User is signed in.
      db.collection("games")
        .doc("active")
        .collection("winners")
        .onSnapshot((querySnapshot) => {
          me.winners = [];
          querySnapshot.forEach((doc) => {
            console.log(doc.id);
            me.winners.push(doc.data());
          });
          me.winners.splice();
        });
    });
  },
  methods: {
    // getDiffs(currentArray, newArray) {},
  },
};
</script>