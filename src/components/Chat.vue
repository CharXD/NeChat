<template>
  <div class="bg">
    <div class="stage">
      <el-input type="textarea" :rows="5" maxlength="5" placeholder="input..." v-model="textarea" class="inputTextarea"></el-input>
      <el-button round>发送</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "chat.vue",
  data() {
    return {
      textarea: ''
    }
  }
}
</script>

<style scoped>
.bg{
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: url(../assets/background_chat.jpg) no-repeat fixed center center;
  background-size:cover;
}
.stage {
  position: absolute;
  top: 10%;
  left: 10%;
  height: 80%;
  width:  80%;
  background-color: #d0c6bd;
  box-shadow: 0 0 10px 5px #cac0b7;
  border-radius: 30px;
}
</style>