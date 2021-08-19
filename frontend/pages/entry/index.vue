<template>
  <div id="app">
    <h1>Entry</h1>
    <div>
      <label for="student_id_entry">ID</label>
      <input
        type="text"
        name="student_id"
        id="student_id_entry"
        v-model="student_id_entry"
      />
      <button @click="send">送信</button>
      <p>{{ matrix }}</p>
      <p>{{ test }}</p>
    </div>
  </div>
</template>

<script>
export default {

  data() {
    return {
      student_id_entry: "",
      matrix: "",
      test:"YYY"
    }
  },
  methods: {
    send: async function (event) {
      fetch(process.env.API_URL+"/api/v1/getBingoCard", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          student_id: this.student_id_entry,
        }),
      })
        .then(async (response) => {
          if (response.status == 200) {
            const body = await response.json();
            this.matrix = body.matrix;
            console.log(body);
            const cardData = {
              matrix: body.matrix,
              student_id: this.student_id_entry,
              raw_value: body.student_id_hash,
            };
            localStorage.setItem("cardData", JSON.stringify(cardData));
            this.$router.push('/play')
          } else {
            throw new Error(response.status);
          }
        })
        .catch((e) => {
          console.log(e);
        });
    },
  },
};
</script>