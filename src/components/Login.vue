<template>
  <div class="bg">
    <div class="Login">
      <div class="card">
        <div class="cardHead">
          <img src="../assets/logo.png" alt="">
          <p>NeChat - Login</p>
        </div>

        <div class="cardBody">
          <el-input v-model="username" placeholder="Your nick name..."></el-input>
          <el-input v-model="password" show-password placeholder="Enter your password..."></el-input>
          <el-button @click="login">Login</el-button>
          <el-link href="#/Register" type="info">Register</el-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios"
import {Message} from 'element-ui';

export default {
  name: 'Login',
  data() {
    return {
      username: null,
      password: null
    }
  },
  methods: {
    login() {
      if (this.username === null || this.password === null) {
        Message({
          message: '用户名或密码不能为空',
          type: 'warning',
          duration: 1500
        });
        return;
      }

      axios.post('http://127.0.0.1:25566/login', {
        UserName: this.username,
        Password: this.password
      })
          .then(function (response) {
            const data = response.data;
            if (data.code === 0) {
              Message({
                message: data.msg, type: 'success', duration: 1500
              });
            } else {
              Message({
                message: data.msg, type: 'error', duration: 1500
              });
            }
          })
          .catch(function (error) {
            Message({
              message: error, type: 'error', duration: 1500
            });
          })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.bg{
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: url(../assets/background.png) no-repeat fixed center center;
  background-size:cover;
}
.card {
  border-radius: 5px;
  width: 390px;
  height: 464px;
  background-color: #ffffff;
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  margin: auto 15% auto auto;
  box-shadow: 0 0 10px 5px #eeeeee;
}

.cardHead {
  position: absolute;
  left: 42px;
  width: 316px;
  height: 131px;
}

.cardHead img {
  position: absolute;
  left: 0;
  right: 0;
  margin: 15% auto;
}

.cardHead p {
  position: absolute;
  margin: 45% 81px;
  font-size: 20px;
}

.cardBody {
  position: absolute;
  left: 42px;
  bottom: 0;
  width: 316px;
  height: 290px;
}

.el-input {
  margin-top: 30px;
}

.el-button, .el-button:focus {
  margin-top: 30px;
  width: 316px;
  background: #374b4f;
  border-color: #374b4f;
  color: #ffffff;
}

.el-button:hover {
  background: #465f65;
  border-color: #465f65;
  color: #ffffff;
}

.el-link {
  margin-top: 15px;
}
</style>
