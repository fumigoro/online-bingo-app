<template>
  <div>
    <h1>管理者画面</h1>
    <p>{{ count }}</p>
    <AdminLogin v-on:login-data="count = $event" v-show="displayFlag.login" />
    <div v-show="displayFlag.main">
      <div>
        <button @click="logout">ログアウト</button>
        <h2>数字を登録</h2>
        <div>
          <label for="newNumber">登録する数字</label>
          <input type="text" id="newNumber" v-model="newNumber" />
          <div>{{ message }}</div>
          <button @click="addNumber">登録</button>
          <button @click="initNumber">初期化</button>
        </div>
        <div>
          <h2>出た数字</h2>
          <div v-for="(value, number) in winNumbers" :key="number">
            {{ number }} {{ value }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AdminLogin from "../../components/AdminLogin.vue";

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
  components: { AdminLogin },
  data() {
    return {
      count: "",
      newNumber: "",
      message: "",
      winNumbers: {},
      displayFlag: {
        login: false,
        main: false,
      },
    };
  },
  methods: {

    async addNumber() {
      this.message = "";
      const newNumber = String(Number(this.newNumber));
      if (!/^[0-7]?[0-9]$/.test(this.newNumber)) {
        this.message = "入力内容を確認してください";
        console.log("入力内容を確認してください");
        return;
      }
      if (Number(newNumber) < 0 || 75 < Number(newNumber)) {
        this.message = "数字が範囲外です";
        console.log("数字が範囲外です");
        return;
      }

      if (this.winNumbers[newNumber]) {
        //既出の場合
        console.log("すでに出ています");
        this.message = "すでに出ています";
        return;
      }

      const data = {};
      data["numbers." + String(newNumber)] = true;
      data["modified"] = firebase.firestore.FieldValue.serverTimestamp();
      console.log({ data });

      const db = firebase.firestore();
      const winNumbersDoc = db.collection("games").doc("active");
      let winNumbersData;
      try {
        winNumbersData = await winNumbersDoc.get();
      } catch (error) {
        console.log("Error getting document:", error);
        this.message = "Error getting document:", error;

        return;
      }
      if (!winNumbersData.exists) {
        console.log("No such document!");
        this.message = "No such document!";
        return;
      }
      if(!("created" in winNumbersData.data())){
        console.log("初期化されていません");
        this.message = "初期化されていません";
        return;
      }

      winNumbersDoc
        .update(data)
        .then(() => {
          console.log("Document successfully updated!");
          this.newNumber = "";
        })
        .catch((error) => {
          console.log(error);
        });
    },
    initNumber() {
      const numbers = {};
      numbers["0"] = true;
      for (let i = 0; i < 75; i++) {
        numbers[i + 1] = false;
      }
      const data = {};
      data["numbers"] = numbers;
      data["created"] = firebase.firestore.FieldValue.serverTimestamp();
      data["modified"] = firebase.firestore.FieldValue.serverTimestamp();
      const db = firebase.firestore();
      db.collection("games")
        .doc("active")
        .set(data)
        .then(() => {
          console.log("Document written ");
        })
        .catch((error) => {
          console.error("Error adding document: ", error);
        });
    },
    logout(){
      firebase.auth().signOut().then(()=>{
        console.log("ログアウトしました");
      })
      .catch( (error)=>{
        console.log(`ログアウト時にエラーが発生しました (${error})`);
      });
    },
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
      if (user) {
        // User is signed in.
        me.displayFlag.login = false;
        me.displayFlag.main = true;
        me.count = user.email + "でログイン中";

        db.collection("games")
          .doc("active")
          .onSnapshot((doc) => {
            const data = doc.data();
            console.log("Current data: ", data);
            if (data && "numbers" in data) {
              me.winNumbers = data.numbers;
              me.winNumbers = Object.assign(me.winNumbers, {});
            } else {
              me.winNumbers = { Error: "データがありません" };
            }
          });
      } else {
        me.displayFlag.login = true;
        me.displayFlag.main = false;
      }
    });
  },
};
</script>

