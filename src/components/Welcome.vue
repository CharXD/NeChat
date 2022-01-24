<template>
  <div class="bg">
    <p style="position: absolute; bottom:0; margin: auto;color: white;">©2022 Char</p>
    <div class="maskLayer">
      <div class="card">
        <h2 style="font-size: 80px;">Ne Chat</h2>
        <h3 style="font-size: 40px;">爆炸吧现充!</h3>
        <el-button round size="large" @click="jumpLogin">Login</el-button>
        <el-button round size="large" @click="jumpRegister">Register</el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Welcome",
  methods: {
    jumpLogin() {undefined
      window.location.href = '/#/Login'
    },
    jumpRegister() {undefined
      window.location.href = '/#/Register'
    },
  }
}
</script>

<style scoped>
.bg {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: url(../assets/background_welcome.jpg) no-repeat fixed center center;
  background-size: cover;
}

.maskLayer {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background-color: rgba(0, 0, 0, .5);
}

.card {
  border-radius: 5px;
  width: 400px;
  height: 300px;
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  margin: auto;
}

.card h2, h3 {
  line-height: 15px;
  color: white;
  font-weight: 500;
  letter-spacing: .3rem;
  text-transform: uppercase;
}

.el-button, .el-button:focus {
  width: 120px;
  background: rgba(0, 0, 0, 0);
  color: white;
  font-size: 20px;
  border: 2px solid white;
}

.el-button:hover {
  background: white;
  border-color: white;
  color: black;
  font-size: 20px;
}
</style>