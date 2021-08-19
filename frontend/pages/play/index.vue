<template>
  <div id="app">
    <h1>Play</h1>
    <p>Matrix: {{ matrix }}</p>
    <button @click="validBingo">ビンゴ！！</button>
    <p>
      結果：<span>{{ result }}</span>
    </p>
    <p>
      メッセージ：<span>{{ message }}</span>
    </p>
    <p>
      学籍番号：<span>{{ student_id }}</span>
    </p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      student_id_entry: "",
      matrix: "",
      result: "",
      message: "",
      student_id: "",
    };
  },
  methods: {
    init() {
      const cardData = JSON.parse(localStorage.getItem("cardData"));
      this.matrix = cardData.matrix;
    },
    async validBingo() {
      this.message = "確認中...";
      const cardData = JSON.parse(localStorage.getItem("cardData"));
      const student_id_hash = cardData.raw_value;
      this.student_id = cardData.student_id;
      const matrix = cardData.matrix;
      const array = [];
      for (let i = 0; i < matrix.length; i++) {
        for (let j = 0; j < matrix[0].length; j++) {
          array.push(Number(matrix[i][j]));
        }
      }

      fetch(this.$config.API_URL+"/api/v1/validBingoCard", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          student_id_hash: student_id_hash,
          student_id:cardData.student_id,
          array: array,
          display_name:"やまだたろう"
        }),
      })
        .then(async (response) => {
          if (response.status == 200) {
            const body = await response.json();
            console.log(body.result);
            console.log(body.message);
            this.message = body.message;
            this.result = body.result;
          } else {
            throw new Error(response.status);
          }
        })
        .catch((e) => {
          console.log(e);
        });
    },
  },
  mounted: function () {
    this.init();
    console.log(this.$config)
  },
};
</script>